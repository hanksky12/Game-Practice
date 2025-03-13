package log

import log "github.com/sirupsen/logrus"

func PrintCron(err error) {
	if err != nil {
		log.Println("Error scheduling job:", err)
	}
}
