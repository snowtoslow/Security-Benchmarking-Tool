package main

import (
	"Security-Benchmarking-Tool/constants"
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
	log.Println(HOME)
	if _, err := os.Stat(HOME + constants.DESKTOP + constants.AuditDirectory); err != nil {
		if os.IsNotExist(err) {
			if err:=utils.CreateAuditsDir(HOME);err!=nil{
				log.Println("Error creating my directory")
			}
		} else {
			log.Println("File exists!!!")
		}
	}
}



func main() {
	HOME, err := utils.GetUserHome()
	policyFileName,err := utils.GenerateSavedFileName(HOME + constants.DESKTOP + constants.AuditDirectory + constants.SavedFileDIRECTORY,constants.AuditFormat,constants.Policy)
	if err!=nil {
		log.Println(err)
	}
	if err:=utils.DownloadFileToExpectedLocation(policyFileName);err!=nil{
		log.Println("ERROR IN DOWNLOADING: ",err)
	}

	arrayData:=utils.ParseFile(policyFileName)
	info:=utils.CreateMapForMultipleItems(arrayData)

	jsonFileName, err:= utils.GenerateSavedFileName(HOME + constants.DESKTOP + constants.AuditDirectory + constants.ParsedDataDirectory,constants.ParsedFileFormat,constants.ParsedPolicy)
	if err:= utils.CreateJsonResponse(jsonFileName,info);err!=nil{
		log.Println("ERROR CREATING JSON",err)
	}

}
