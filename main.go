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
			utils.CreateAuditsDir(HOME)
		} else {
			log.Println("File exists!!!")
		}
	}
}



func main() {
	HOME, err := utils.GetUserHome()
	fileName,err := utils.GenerateSavedFileName(HOME+constants.DESKTOP + constants.AuditDirectory + constants.SavedFileDIRECTORY,constants.AuditFormat,constants.Policy)

	if err!=nil {
		log.Println(err)
	}


	if err:=utils.DownloadFile(fileName,"https://www.tenable.com/downloads/api/v1/public/pages/configuration-audit-policies/downloads/11237/download?i_agree_to_tenable_license_agreement=true");err!=nil{
		log.Println(err)
	} else{
		log.Println("DOWNLOADED")
	}

	/*arrayData:=utils.ParseFile(utils.GenerateFileNames())
	info:=utils.CreateMapForMultipleItems(arrayData)
	fmt.Println(info)*/
}
