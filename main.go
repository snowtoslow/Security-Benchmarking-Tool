package main

import (
	"Security-Benchmarking-Tool/utils"
	"fmt"
	"log"
)

func main() {

	if err:=utils.DownloadFile(utils.GenerateFileNames(),"https://www.tenable.com/downloads/api/v1/public/pages/configuration-audit-policies/downloads/11237/download?i_agree_to_tenable_license_agreement=true");err!=nil{
		log.Println(err)
	} else{
		log.Println("DOWNLOADED")
	}

	arrayData:=utils.ParseFile(utils.GenerateFileNames())
	fmt.Println(arrayData[0])

}
