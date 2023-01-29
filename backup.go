package main

import (
	"fmt"

	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func BackupFiles() {
	log.Println("Starting backup files")
	cf := GetConfig()
	cmd := exec.Command("rdiff-backup", cf.Folder, cf.FilesBackupFolder)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()
	// args := []string{"rdiff-backup", cf.Folder, cf.FilesBackupFolder}
	// err := CmdExecConsole(args...)
	err := cmd.Run()
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("Finishing backup files")
}

func BackupDB() {
	log.Println("Starting backup DB")
	cf := GetConfig()
	var cmd *exec.Cmd
	var filename string
	var outfile *os.File
	var err error
	if cf.Docker.UsePostgresDocker {
		cmd = exec.Command("docker", "exec", "-i", "-t", cf.Docker.ContainerName, "pg_dump", "-d", cf.PgDBName, "-U", cf.PgUser, "-p", cf.PgPort, "-h", cf.PgHost, "-O", "-x", "-Ft")
	} else {
		cmd = exec.Command("pg_dump", "-d", cf.PgDBName, "-U", cf.PgUser, "-p", cf.PgPort, "-h", cf.PgHost, "-O", "-x", "-Ft")
	}
	filename = fmt.Sprintf("%d.tar", time.Now().Unix())
	outfile, err = os.Create(filepath.Join(cf.DBBackupFolder, filename))
	if err != nil {
		log.Println(err.Error())
		return
	}

	defer outfile.Close()
	cmd.Stdout = outfile
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, fmt.Sprintf("PGPASSWORD=%s", cf.PgPassword))
	err = cmd.Run()
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("Finishing backup DB")
}

func BackupAll() {
	cf := GetConfig()
	if cf.BackupDB {
		BackupDB()
	}
	if cf.BackupFiles {
		BackupFiles()
	}
}

func DeleteOld() {
	log.Println("Deleting old dbs")
	cf := GetConfig()
	entries, err := os.ReadDir(cf.DBBackupFolder)
	if err != nil {
		log.Println(err)
		return
	}
	delta := cf.DaysBeforeDelete * 24 * 60 * 60
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := strings.TrimSuffix(e.Name(), filepath.Ext(e.Name()))
		i, err := strconv.ParseInt(name, 10, 64)
		if err != nil {
			continue
		}
		if (time.Now().Unix() - delta) > i {
			err = os.Remove(filepath.Join(cf.DBBackupFolder, e.Name()))
			if err != nil {
				log.Println(err)
				continue
			}
		}

	}
}

func SyncFolder(source, dest string) {
	log.Println("Starting sync folders")
	cmd := exec.Command("rsync", "-av", source, dest)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Println(err)
	}
	log.Println("Finish sync folders")
}
