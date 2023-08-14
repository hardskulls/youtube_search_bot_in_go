package logging

import (
	"fmt"
	"log"
	"runtime"
)

func LogVar[T any](variable T, msg string) {
	_, file, line, _ := runtime.Caller(1) // Caller(1) for outer function, Caller(0) for this one.
	separator := "------------------------------------------------------"
	emptyLine := "                                                      "

	log.Println(emptyLine)

	log.Println(separator)

	log.Printf(fmt.Sprintf("[ LOG VAR Caller(1)] [ File and line: %v:%v | Message: '%v' ]", file, line, msg))
	log.Printf(fmt.Sprintf("[ VAR ] Type is <| %T |> Value is [ <| %v |> ] ", variable, variable))

	log.Println(separator)

	log.Println(emptyLine)
}

func LogLine() {
	_, file, line, _ := runtime.Caller(1) // Caller(1) for outer function, Caller(0) for this one.
	separator := "------------------------------------------------------"
	emptyLine := "                                                      "

	log.Println(emptyLine)

	log.Println(separator)

	log.Println(fmt.Sprintf("[ LOG LINE ] [ Line <| %v:%v |> successfully executed. ]", file, line))

	log.Println(separator)

	log.Println(emptyLine)
}

func LogError(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)   // Caller(1) for outer function, Caller(0) for this one.
		_, file2, line2, _ := runtime.Caller(4) // Caller(4) for outer function, Caller(0) for this one.
		separator := "------------------------------------------------------"
		emptyLine := "                                                      "

		log.Println(emptyLine)

		log.Println(separator)

		log.Println(fmt.Sprintf("[ ERROR Caller(1) ] : ( error is <| %v |> | file and line: <| %v:%v |> )", err, file, line))
		log.Println(fmt.Sprintf("[ ERROR Caller(4) ] : ( error is <| %v |> | file and line: <| %v:%v |> )", err, file2, line2))

		log.Println(separator)

		log.Println(emptyLine)
	}
}

func LogFuncStart(name string) {
	separator := "------------------------------------------------------"
	emptyLine := "                                                      "

	log.Println(emptyLine)

	log.Println(separator)

	log.Println(fmt.Sprintf("[ STARTED <| '%v' |>. ]", name))

	log.Println(separator)

	log.Println(emptyLine)
}

func LogFuncEnd(name string) {
	separator := "------------------------------------------------------"
	emptyLine := "                                                      "

	log.Println(emptyLine)

	log.Println(separator)

	log.Println(fmt.Sprintf("[ ENDED <| '%v' |>. ]", name))

	log.Println(separator)

	log.Println(emptyLine)
}
