package test

import (
	"github.com/joho/godotenv"
	"log"
)

func Init() error {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatalf("env not read: %v", err)
		return err
	}
	return nil
}
