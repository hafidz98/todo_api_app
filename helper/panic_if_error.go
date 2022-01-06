package helper

import "log"

func PanicIfError(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
