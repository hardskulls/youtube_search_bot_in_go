package dialogue

import (
	"youtube_search_go_bot/internal/keyboards"
)

type KVStruct struct {
	Key   string
	Value string
}

type Command string

const (
	SearchCommand Command = "SearchCommand"
	ListCommand   Command = "ListCommand"
)

type Target string

const (
	TargetSubscription Target = "Target/TargetSubscription"
	TargetPlaylist     Target = "Target/TargetPlaylist"
)

// Stores dialogue context and parameters required for bot commands.
type DialogueData struct {
	ActiveCmd         Command
	LastCallback      string
	MsgWithCallbackId int
	Target            Target
	ResultLimit       uint16
	TextToSearch      string
	Sorting           keyboards.Sorting
	SearchIn          keyboards.SearchIn
}
