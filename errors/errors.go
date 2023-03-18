package errors

import (
	"log"
	"runtime"
)

func ExitOnError(err error) {
	if err != nil {
		progCounter, file, line, _ := runtime.Caller(1) // Caller(1) for outer function, Caller(0) for this one.
		log.Fatalf("[   ERROR || Error is: '%v' || programm counter - %v | file and line - %v:%v   ]", err, progCounter, file, line)
	}
}

func ExitOnEmptyString(str string, errMsg string) {
	if str == "" {
		progCounter, file, line, _ := runtime.Caller(1) // Caller(1) for outer function, Caller(0) for this one.
		log.Fatalf("[   ERROR || Error: string is empty || "+
			"Error message is: '%v' || "+"programm counter - %v | file and line - %v:%v   ]", errMsg, progCounter, file, line)
	}
}
