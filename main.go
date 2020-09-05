package main

import (
	"github.com/DragonSSS/jwt-generator-cli/cmd"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	cmd.Execute()
}
