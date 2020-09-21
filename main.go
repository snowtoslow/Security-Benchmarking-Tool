package main

import (
	"Security-Benchmarking-Tool/constants"
	"Security-Benchmarking-Tool/ui"
	"Security-Benchmarking-Tool/utils"
	"fmt"
	"log"
	"os"
)

func init() {
	fmt.Println("This will get called on main initialization")
	HOME, err := utils.GetUserHome()
	if err != nil {
		log.Println(err)
	}
	if _, err := os.Stat(HOME + constants.DESKTOP + constants.AuditDirectory); err != nil {
		if os.IsNotExist(err) {
			if err := utils.CreateAuditsDir(HOME); err != nil {
				log.Println("Error creating my directory")
			}
		} else {
			log.Println("File exists!!!")
		}
	}
}

func main() {

	if err := ui.SetupAboutDialogWindow(); err != nil {
		log.Println("Main err:", err)
	}
}
