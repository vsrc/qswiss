package util

import "log"

// PanicIf is a small util function to panic on errors
// TODO: Change panic behaviour to some kind of silent log
func PanicIf(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
