package main

import (
	"github.com/esonhugh/openai-platform-api/cmd/subcmd"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := subcmd.RootCmd.Execute()
	if err != nil {
		log.Error(err)
		return
	}
}
