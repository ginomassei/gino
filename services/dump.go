package services

import (
	"fmt"
	"os/exec"
	"time"
)

type DumpConfig struct {
	DbUri     string
	BackupDir string
}

type DumpService interface {
	Dump() error
}

type dumpService struct {
	config DumpConfig
}

func NewDumpService(config DumpConfig) *dumpService {
	return &dumpService{
		config: config,
	}
}

func (s *dumpService) Dump() (string, error) {
	// Run mongodump command
	cmd := exec.Command("mongodump", "--uri", s.config.DbUri, "--out", s.config.BackupDir)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running mongodump:", err)
		return "", err
	}

	// Get the current date and time
	currentTime := time.Now()
	// Format the date and time
	dumpFileName := currentTime.Format("2006-01-02_15:04")
	dumpFileName = "dump_" + dumpFileName + ".tar.gz"

	// Create a tar.gz file with the MongoDB dump
	cmd = exec.Command("pwd")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return "", err
	}

	cmd = exec.Command("tar", "-czvf", dumpFileName, "-C", s.config.BackupDir, ".")
	cmd.Dir = s.config.BackupDir

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error creating tar.gz file:", err)
		return "", err
	}

	return dumpFileName, nil
}
