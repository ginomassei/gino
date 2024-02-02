package services

import (
	"log"
	"os"
)

type OsService interface {
	CreateTempDir(dirName string) error
	CleanUp(dirName string) error
}

type osService struct{}

func NewOsService() *osService {
	return &osService{}
}

func (s *osService) CreateTempDir(dirName string) error {
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		log.Fatal("Error creating temporary directory:", err)
	}
	return nil
}

func (s *osService) CleanUp(dirName string) error {
	err := os.RemoveAll(dirName)
	if err != nil {
		log.Fatal("Error removing temporary directory:", err)
	}
	return nil
}
