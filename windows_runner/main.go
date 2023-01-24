package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/akamensky/argparse"
	"github.com/jasonlvhit/gocron"
)

func main() {

	parser := argparse.NewParser("Server backup windows runner", "Service to run on windows that call wsl service")
	configName := parser.StringPositional(&argparse.Options{Required: true, Help: "Configuration file name without extension"})
	backupTimes := parser.List("t", "time", &argparse.Options{Required: true, Help: "Backup time"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	for _, t := range *backupTimes {
		err = gocron.Every(1).Day().At(t).Do(func() {
			log.Println("Running backup")
			cmd := exec.Command("wsl", "--", "server-backup", "backup", "-c", *configName, "-l")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				log.Println(err)
			}
		})
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	log.Println("Starting scheduler")
	<-gocron.Start()
}
