package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/akamensky/argparse"
	"github.com/jasonlvhit/gocron"
)

func main() {

	parser := argparse.NewParser("Server backup", "App for making backup of database and files from server")
	configName := parser.String("c", "config", &argparse.Options{Required: true, Help: "Configuration file name without extension"})
	logToFile := parser.Flag("l", "logfile", &argparse.Options{Default: false, Help: "Log to file instead of console"})

	schedulerCmd := parser.NewCommand("scheduler", "Start scheduler")
	backupCmd := parser.NewCommand("backup", "Make one backup")

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	cf := LoadConfig(*configName)
	if *logToFile {
		f, err := os.OpenFile(filepath.Join(cf.AppDir, "logs", fmt.Sprintf("%s.log", *configName)), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		log.SetOutput(f)
	}

	switch {
	case schedulerCmd.Happened():

		for _, t := range cf.BackupTimes {
			err = gocron.Every(1).Day().At(t).Do(func() {
				go DeleteOld()
				BackupAll()
			})
			if err != nil {
				log.Fatal(err.Error())
			}
		}
		log.Println("Starting scheduler")
		<-gocron.Start()
	case backupCmd.Happened():
		BackupAll()
	}
}
