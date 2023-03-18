package dialogue

import "youtube_search_go_bot/keyboards"

type KVStruct struct {
	Key   string
	Value string
}

type ListSettings struct {
	ResultLimit uint16
	Target      keyboards.SearchTarget
	Sorting     keyboards.Sorting
}

type SearchSettings struct {
	ResultLimit uint16
	Target      keyboards.SearchTarget
	SearchIn    keyboards.SearchIn
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