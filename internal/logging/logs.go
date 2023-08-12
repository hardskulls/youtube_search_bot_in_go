package logging

import (
	"log"
	"runtime"
)

func LogVar[T any](variable T, msg string) {
	_, file, line, _ := runtime.Caller(1) // Caller(1) for outer function, Caller(0) for this one.
	separator := "------------------------------------------------------"
	log.Printf("\n\n%v", separator)
	log.Printf("\n\n[ LOG VAR Caller(1)] [ File and line: %v:%v | Message: '%v' ]", file, line, msg)
	log.Printf("\n[ VAR ] Type is <| %T |> Value is [ <| %v |> ] ", variable, variable)
	log.Printf("\n\n%v \n\n\n\n", separator)
}

func LogLine() {
	_, file, line, _ := runtime.Caller(1) // Caller(1) for outer function, Caller(0) for this one.
	separator := "------------------------------------------------------"
	log.Printf("\n\n%v", separator)
	log.Printf("\n\n[ LOG LINE ] [ Line <| %v:%v |> successfully executed. ]", file, line)
	log.Printf("\n\n%v \n\n\n\n", separator)
}

func LogError(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1) // Caller(1) for outer function, Caller(0) for this one.
		separator := "------------------------------------------------------"
		log.Printf("\n\n%v", separator)
		log.Printf("\n\n[ ERROR ] : ( error is <| %v |> | file and line: <| %v:%v |> )", err, file, line)
		log.Printf("\n\n%v \n\n\n\n", separator)
	}
}

func LogFuncStart(name string) {
	separator := "------------------------------------------------------"
	log.Printf("\n\n%v", separator)
	log.Printf("\n\n[ STARTED <| '%v' |>. ]", name)
	log.Printf("\n\n%v \n\n\n\n", separator)
}

func LogFuncEnd(name string) {
	separator := "------------------------------------------------------"
	log.Printf("\n\n%v", separator)
	log.Printf("\n\n[ ENDED <| '%v' |>. ]", name)
	log.Printf("\n\n%v \n\n\n\n", separator)
}
