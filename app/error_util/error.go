package error_util

import "log"

func CheckError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
