package logging

import (
	"log"
	"runtime"
)

func LogVar[T any](variable T, msg string) {
	_, file, line, _ := runtime.Caller(1) // Caller(1) for outer function, Caller(0) for this one.
	log.Println("------------------------------------------------------------------------------")
	log.Printf("[ LOG VAR ] [ File and line: %v:%v | Message: '%v' ]", file, line, msg)
	log.Printf("[ VAR ] Type is <| %T |> Value is [ <| %v |> ] ", variable, variable)
	log.Println("                                                                              ")
}

func LogLine() {
	_, file, line, _ := runtime.Caller(1) // Caller(1) for outer function, Caller(0) for this one.
	log.Println("------------------------------------------------------------------------------")
	log.Printf("[ LOG LINE ] [ Line <| %v:%v |> successfully executed. ]", file, line)
	log.Println("                                                                              ")
}

func LogError(err error) {
	_, file, line, _ := runtime.Caller(1) // Caller(1) for outer function, Caller(0) for this one.
	log.Printf("[ ERROR ] : ( error is <| %v |> | file and line: <| %v:%v |> )", err, file, line)
}

func LogFuncStart(name string) {
	log.Printf("[ FUNC '%v' <| Started |>. ]", name)
}

func LogFuncEnd(name string) {
	log.Printf("[ FUNC '%v' <| Ended |>. ]", name)
}
