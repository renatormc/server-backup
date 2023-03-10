package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var config Config

type Config struct {
	AppDir       string
	Folder       string `json:"folder"`
	BackupFolder string `json:"backup_folder"`
	Docker       struct {
		UsePostgresDocker      bool   `json:"use_postgres_docker"`
		ContainerName          string `json:"container_name"`
		MountPointBackupFolder string `json:"mount_point_backup_folder"`
	} `json:"docker"`
	FilesBackupFolder string
	DBBackupFolder    string
	PgHost            string   `json:"pg_host"`
	PgUser            string   `json:"pg_user"`
	PgPassword        string   `json:"pg_password"`
	PgPort            string   `json:"pg_port"`
	PgDBName          string   `json:"pg_db_name"`
	BackupTimes       []string `json:"backup_times"`
	BackupDB          bool     `json:"backup_db"`
	BackupFiles       bool     `json:"backup_files"`
	DaysBeforeDelete  int64    `json:"days_before_delete"`
	BackupAtStartup   bool     `json:"backup_at_startup"`
}

func GetConfig() Config {
	return config
}

func LoadConfig(name string) Config {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	config.AppDir = filepath.Dir(ex)

	configFolder := filepath.Join(config.AppDir, "config")
	if !DirectoryExists(configFolder) {
		if err := os.MkdirAll(configFolder, os.ModePerm); err != nil {
			panic(err)
		}
	}

	path := filepath.Join(configFolder, fmt.Sprintf("%s.json", name))
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal([]byte(file), &config); err != nil {
		panic(err)
	}

	config.FilesBackupFolder = filepath.Join(config.BackupFolder, "files")
	if !DirectoryExists(config.FilesBackupFolder) {
		err := os.Mkdir(config.FilesBackupFolder, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	config.DBBackupFolder = filepath.Join(config.BackupFolder, "db")
	if !DirectoryExists(config.DBBackupFolder) {
		err := os.Mkdir(config.DBBackupFolder, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	return config
}
