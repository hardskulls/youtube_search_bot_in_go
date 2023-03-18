package commands

type Command string

const (
	Start  Command = "/start"
	Info   Command = "/info"
	Search Command = "/search"
	List   Command = "/list"
	LogOut Command = "/logout"
)
