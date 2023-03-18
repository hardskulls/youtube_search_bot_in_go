package db

import (
	"runtime"
	"strings"
	"testing"
	"youtube_search_go_bot/dialogue"
	"youtube_search_go_bot/keyboards"

	"golang.org/x/oauth2"
)

const databaseURL = "postgres://postgres:ex_pekt47jkO4y0u@localhost:5432/postgres"

// func TestDatabaseURL(t *testing.T) {
// 	var (
// 		userData string
// 		postgresURL string
// 	)
//
// 	user, password := os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD")
// 	errors.ExitOnEmptyString(user, "!! user env variable is empty !!")
// 	if password != "" {
// 		userData = user + ":" + password
// 	} else {
// 		userData = user
// 	}
// 	postgresURL = "postgres://" + userData + "@localhost:5432/postgres"
//
// 	err :=  CreateTables(postgresURL)
// 	if err != nil {
// 		programCounter, file, line, _ := runtime.Caller(0)
// 		t.Fatalf(" [ ERROR ] : ( Error is: '%v', program counter is %v, file and line is %v, %v ) ", err, programCounter, file, line)
// 	}
// }

func TestSaveOAuthToken(t *testing.T) {
	err := CreateTables(databaseURL)
	if err != nil {
		programCounter, file, line, _ := runtime.Caller(0)
		t.Fatalf(" [ ERROR ] : ( Error is: '%v', program counter is %v, file and line is %v, %v ) ", err, programCounter, file, line)
	}

	token1 := &oauth2.Token{AccessToken: "AccessToken1", TokenType: "TokenType1", RefreshToken: "RefreshToken"}

	err = SaveOAuthToken("45", token1, databaseURL)
	if err != nil {
		programCounter, file, line, _ := runtime.Caller(0)
		t.Fatalf(" [ ERROR ] : ( Error is: '%v', program counter is %v, file and line is %v, %v ) ", err, programCounter, file, line)
	}

	token2 := &oauth2.Token{AccessToken: "AccessToken2", TokenType: "TokenType2", RefreshToken: "RefreshToken2"}

	err = SaveOAuthToken("45", token2, databaseURL)
	if err != nil {
		programCounter, file, line, _ := runtime.Caller(0)
		t.Fatalf(" [ ERROR ] : ( Error is: '%v', program counter is %v, file and line is %v, %v ) ", err, programCounter, file, line)
	}
}

func TestGetOAuthToken(t *testing.T) { // !! Direct comparisson of tokens gives false even if they are equal.
	err := CreateTables(databaseURL)
	if err != nil {
		programCounter, file, line, _ := runtime.Caller(0)
		t.Fatalf(" [ ERROR ] : ( Error is: '%v', program counter is %v, file and line is %v, %v ) ", err, programCounter, file, line)
	}

	token1 := &oauth2.Token{AccessToken: "AccessToken1", TokenType: "TokenType1", RefreshToken: "RefreshToken"}
	err = SaveOAuthToken("45", token1, databaseURL)
	if err != nil {
		programCounter, file, line, _ := runtime.Caller(0)
		t.Fatalf(" [ ERROR ] : ( Error is: '%v', program counter is %v, file and line is %v, %v ) ", err, programCounter, file, line)
	}

	token2, err := GetOAuthToken("45", databaseURL)
	if err != nil || token1.AccessToken != token2.AccessToken || token1.TokenType != token2.TokenType || token1.RefreshToken != token2.RefreshToken {
		programCounter, file, line, _ := runtime.Caller(0)
		t.Fatalf(" [ ERROR ] : ( Error is: '%v', program counter is %v, file and line is %v, %v ) ", err, programCounter, file, line)
	}
}

func TestSaveDialogueState(t *testing.T) {
	err := CreateTables(databaseURL)
	if err != nil {
		programCounter, file, line, _ := runtime.Caller(0)
		t.Fatalf(" [ ERROR ] : ( Error is: '%v', program counter is %v, file and line is %v, %v ) ", err, programCounter, file, line)
	}

	token1 := &oauth2.Token{AccessToken: "AccessToken1", TokenType: "TokenType1", RefreshToken: "RefreshToken"}
	err = SaveOAuthToken("45", token1, databaseURL)
	if err != nil {
		programCounter, file, line, _ := runtime.Caller(0)
		t.Fatalf(" [ ERROR ] : ( Error is: '%v', program counter is %v, file and line is %v, %v ) ", err, programCounter, file, line)
	}

	err = SaveDialogueData("45", dialogue.DialogueData{}, databaseURL)
	if err != nil {
		programCounter, file, line, _ := runtime.Caller(0)
		t.Fatalf(" [ ERROR ] : ( Error is: '%v', program counter is %v, file and line is %v, %v ) ", err, programCounter, file, line)
	}
}

// TODO: Fix this one
func TestGetDialogueState(t *testing.T) {
	err := CreateTables(databaseURL)
	if err != nil {
		programCounter, file, line, _ := runtime.Caller(0)
		t.Fatalf(" [ ERROR ] : ( Error is: '%v', program counter is %v, file and line is %v, %v ) ", err, programCounter, file, line)
	}

	token1 := &oauth2.Token{AccessToken: "AccessToken1", TokenType: "TokenType1", RefreshToken: "RefreshToken"}
	err = SaveOAuthToken("45", token1, databaseURL)
	if err != nil {
		programCounter, file, line, _ := runtime.Caller(0)
		t.Fatalf(" [ ERROR ] : ( Error is: '%v', program counter is %v, file and line is %v, %v ) ", err, programCounter, file, line)
	}

	// userId, activeCommand, maxRes, dialogueState := "45", yt.StartCommand, 0, "Keyboard:Command:Search:Target"
	dialogueState := dialogue.DialogueData{ActiveCmd: dialogue.SearchCommand, SearchIn: keyboards.SearchInTitle, ResultLimit: 5, Target: dialogue.TargetSubscription}
	err = SaveDialogueData("45", dialogueState, databaseURL)
	if err != nil {
		programCounter, file, line, _ := runtime.Caller(0)
		t.Fatalf(" [ ERROR ] : ( Error is: '%v', program counter is %v, file and line is %v, %v ) ", err, programCounter, file, line)
	}

	state, err := GetDialogueData("45", databaseURL)
	if err != nil || state != dialogueState {
		programCounter, file, line, _ := runtime.Caller(0)
		t.Fatalf(" [ ERROR ] : ( Error is: '%v', program counter is %v, file and line is %v, %v ) ", err, programCounter, file, line)
	}
}

func TestStrings(t *testing.T) {
	activeCommand := "/search_subs_by_title"
	index := strings.Index(activeCommand, "/") // var mode yt.CompareMode
	if index != 0 {
		programCounter, file, line, _ := runtime.Caller(0)
		t.Fatalf(" [ ERROR ] : ( command is: '%v', index is: '%v', program counter is %v, file and line is %v, %v ) ",
			activeCommand, index, programCounter, file, line)
	}
}

// func TestForgotToSaveDialogueKey(t *testing.T) {
// 	err := CreateTables(databaseURL)
// 	if err != nil {
// 		programCounter, file, line, _ := runtime.Caller(0)
// 		t.Fatalf(" [ ERROR ] : ( Error is: '%v', program counter is %v, file and line is %v, %v ) ", err, programCounter, file, line)
// 	}
//
// 	token1 := &oauth2.Token{AccessToken: "AccessToken1", TokenType: "TokenType1", RefreshToken: "RefreshToken"}
// 	err = SaveOAuthToken("45", token1, databaseURL)
// 	if err != nil {
// 		programCounter, file, line, _ := runtime.Caller(0)
// 		t.Fatalf(" [ ERROR ] : ( Error is: '%v', program counter is %v, file and line is %v, %v ) ", err, programCounter, file, line)
// 	}
//
// 	comm, res, state, err := GetDialogueData("45", databaseURL)
// 	if err != nil || comm != "" || state != "" || res != int32(1) {
// 		programCounter, file, line, _ := runtime.Caller(0)
// 		t.Fatalf(" [ ERROR ] : ( Error is: '%v', program counter is %v, file and line is %v, %v ) ", err, programCounter, file, line)
// 	}
// }
