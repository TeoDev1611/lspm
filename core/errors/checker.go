package errors

import "log"

func ErrorChecker(err error) {
	if err != nil {
		log.Fatalf("ERROR( LSP ): %s", err)
	}
}
