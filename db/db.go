package db

import (
	"context"
	"encoding/json"
	"errors"
	"gopkg.in/telebot.v3"
	"net/http"
	"net/url"
	"strings"
	"youtube_search_go_bot/keyboards"
	"youtube_search_go_bot/logging"

	"golang.org/x/oauth2"
	"youtube_search_go_bot/dialogue"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
)

func SaveOAuthToken(userId string, oauthToken *oauth2.Token, databaseURL string) error {
	jsonedToken, err := json.Marshal(oauthToken)
	if err != nil {
		return err
	}
	if jsonedToken == nil {
		return errors.New("jsonedToken is nil")
	}

	dbPool, err := pgxpool.Connect(context.Background(), databaseURL)
	if err != nil {
		return err
	}
	defer dbPool.Close()

	upsertTokens := `INSERT INTO user_oauth_tokens (user_id, jsoned_token) 
	VALUES ($1, $2) 
	ON CONFLICT (user_id) DO UPDATE SET jsoned_token = $2;`
	if _, err = dbPool.Exec(context.Background(), upsertTokens, userId, jsonedToken); err != nil {
		return err
	}

	// 	upsertDialogueData := `INSERT INTO user_dialogue_data (user_id) VALUES ($1) ON CONFLICT DO NOTHING;`
	// 	if _, err = dbPool.Exec(context.Background(), upsertDialogueData, userId); err != nil {
	// 		return err
	// 	}

	logging.LogFuncEnd("SaveOAuthToken")
	return nil
}

func GetOAuthToken(userId, dbUrl string) (*oauth2.Token, error) {
	dbPool, err := pgxpool.Connect(context.Background(), dbUrl)
	if err != nil {
		return nil, err
	}
	defer dbPool.Close()

	var token *oauth2.Token
	var jsonData []byte

	statement := `SELECT jsoned_token FROM user_oauth_tokens WHERE user_id=$1;`
	if err = dbPool.QueryRow(context.Background(), statement, userId).Scan(&jsonData); err != nil {
		return nil, err
	}

	if jsonData == nil {
		err = errors.New("jsonData is nil")
		return nil, err
	}
	if json.Valid(jsonData) == false {
		err = errors.New("jsonData contains invalid json")
		return nil, err
	}
	if err = json.Unmarshal(jsonData, &token); err != nil {
		return nil, err
	}

	logging.LogFuncEnd("GetOAuthToken")
	return token, nil
}

func SaveDialogueData(userId string, dialogueData dialogue.DialogueData, dbURL string) error {
	jsonDialogueData, err := json.Marshal(dialogueData)
	if err != nil {
		return err
	}
	if jsonDialogueData == nil {
		return errors.New("jsonDialogueData is nil")
	}

	dbPool, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		return err
	}
	defer dbPool.Close()

	sqlStr := `INSERT INTO user_dialogue_data (user_id, dialogue_state)
	VALUES ($1, $2)
	ON CONFLICT (user_id)
	DO UPDATE SET dialogue_state = $2;`
	if _, err = dbPool.Exec(context.Background(), sqlStr, userId, dialogueData); err != nil {
		return err
	}

	logging.LogFuncEnd("SaveDialogueData")
	return nil
}

func GetDialogueData(userId, dbURL string) (dialogueState dialogue.DialogueData, err error) {
	dbPool, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		return dialogue.DialogueData{}, err
	}
	defer dbPool.Close()

	var dialogueData dialogue.DialogueData
	var jsonData pgtype.JSON

	statement := `SELECT dialogue_state FROM user_dialogue_data WHERE user_id=$1;`
	if err = dbPool.QueryRow(context.Background(), statement, userId).Scan(&jsonData); err != nil {
		return dialogue.DialogueData{}, err
	}
	if err = json.Unmarshal(jsonData.Bytes, &dialogueData); err != nil {
		return dialogue.DialogueData{}, err
	}

	logging.LogFuncEnd("GetDialogueData")
	return dialogueData, nil
}

// SaveChatId TODO: Add saving and getting user's chatId to notify him after successfuly aquiring oauth token.
func SaveChatId(user string, chatId string) error {
	return nil
}

func CreateTables(dbUrl string) error {
	dbPool, err := pgxpool.Connect(context.Background(), dbUrl)
	if err != nil {
		return err
	}
	defer dbPool.Close()

	dropTable1 := `DROP TABLE IF EXISTS user_dialogue_data;`
	if _, err = dbPool.Exec(context.Background(), dropTable1); err != nil {
		return err
	}

	dropTable2 := `DROP TABLE IF EXISTS user_oauth_tokens;`
	if _, err = dbPool.Exec(context.Background(), dropTable2); err != nil {
		return err
	}

	createUserOAuthTokens := `CREATE TABLE user_oauth_tokens(user_id TEXT PRIMARY KEY, jsoned_token JSON NOT NULL);`
	if _, err = dbPool.Exec(context.Background(), createUserOAuthTokens); err != nil {
		return err
	}

	createUserDialogueData := `CREATE TABLE user_dialogue_data (
		user_id TEXT PRIMARY KEY, 
		dialogue_state JSON NOT NULL
	);`
	if _, err = dbPool.Exec(context.Background(), createUserDialogueData); err != nil {
		return err
	}

	logging.LogFuncEnd("CreateTables")
	return nil
}

func SaveTarget(userId, target, dbUrl string) error {
	var t dialogue.Target
	if strings.Contains(target, "Subscription") {
		t = dialogue.TargetSubscription
	} else if strings.Contains(target, "Playlist") {
		t = dialogue.TargetPlaylist
	} else {
		return telebot.Err("target is wrong")
	}

	dialogueData, err := GetDialogueData(userId, dbUrl)
	if err != nil {
		return err
	}
	dialogueData.Target = t
	err = SaveDialogueData(userId, dialogueData, dbUrl)
	if err != nil {
		return err
	}
	return nil
}

func SaveSearchIn(userId string, searchIn keyboards.SearchIn, dbUrl string) error {
	dialogueData, err := GetDialogueData(userId, dbUrl)
	if err != nil {
		return err
	}
	dialogueData.SearchIn = searchIn
	err = SaveDialogueData(userId, dialogueData, dbUrl)
	if err != nil {
		return err
	}
	return nil
}

func SaveSorting(userId string, sorting keyboards.Sorting, dbUrl string) error {
	dialogueData, err := GetDialogueData(userId, dbUrl)
	if err != nil {
		return err
	}
	dialogueData.Sorting = sorting
	err = SaveDialogueData(userId, dialogueData, dbUrl)
	if err != nil {
		return err
	}
	return nil
}

// Revoke YouTube access token, and if request was successful, deletes token from db.
func LogOut(userId, dbUrl string) error {
	dbPool, err := pgxpool.Connect(context.Background(), dbUrl)
	if err != nil {
		return err
	}
	defer dbPool.Close()

	token, _ := GetOAuthToken(userId, dbUrl)
	t := token.RefreshToken

	resp, err := http.PostForm("https://oauth2.googleapis.com/revoke", url.Values{"token": {t}})
	if err != nil || resp.StatusCode > 299 || resp.StatusCode < 200 {
		return err
	}
	statement := `DELETE FROM user_oauth_tokens WHERE user_id=$1;`
	if _, err = dbPool.Exec(context.Background(), statement, userId); err != nil {
		return err
	}

	logging.LogFuncEnd("LogOut")
	return nil
}
