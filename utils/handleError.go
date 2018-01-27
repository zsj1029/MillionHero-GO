package utils

import (
	"log"
	"os"
)

func HandleError(err error) {
	if err != nil {
		log.Printf("Error %v .", err)
		os.Exit(1)
	}
}
