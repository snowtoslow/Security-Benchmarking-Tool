package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//TWO MAIN METHODS FOR NOW WITH PRACTICALLY SAME LOGIC
// here as input need to be path := /home/snowtoslow/Desktop/audit/policy-info/*policy-name*;
// infoToBeParsed result from CreateMapForMultipleElements
// *arrayData:=utils.ParseFile(utils.GenerateFileNames())
//	info:=utils.CreateMapForMultipleItems(arrayData)
//	fmt.Println(info)



func CreateJsonResponse(path string,infoToBeParsed []map[string]string) (err error){
	file, err := json.MarshalIndent(infoToBeParsed, "", " ")
	err = ioutil.WriteFile(path, file, 0644)
	if err!=nil{
		return err
	}
	return nil
}
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*
Here is how fileName are gonna be composed:
HOME, err := utils.GetUserHome()
fileName,err := utils.GenerateSavedFileName(HOME+constants.DESKTOP + constants.AuditDirectory + constants.SavedFileDIRECTORY,constants.AuditFormat,constants.Policy)
*/
func DownloadFileToExpectedLocation(fileName string) (err error){
	/*HOME, err := utils.GetUserHome()
	fileName,err := utils.GenerateSavedFileName(HOME+constants.DESKTOP + constants.AuditDirectory + constants.SavedFileDIRECTORY,constants.AuditFormat,constants.Policy)*/

	if err:=DownloadFile(fileName,"https://www.tenable.com/downloads/api/v1/public/pages/configuration-audit-policies/downloads/11237/download?i_agree_to_tenable_license_agreement=true");err!=nil{
		return err
	}
	defer log.Println("Downloaded")
	return nil
}
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////



func CreateMapForMultipleItems(arrayOfParsedData []string) []map[string]string {
	var arrayOfMaps []map[string]string
	for _,value :=range arrayOfParsedData{
		arrayOfMaps = append(arrayOfMaps,createMapForSingleItem(strings.TrimLeft(strings.TrimRight(value, "</custom_item>"), "<custom_item>")))
	}

	return arrayOfMaps
}


func createMapForSingleItem(myStr string) (mymap map[string]string) {
	mymap = make(map[string]string)
	words := strings.Fields(myStr)
	for i := 0; i <len(words) ; i++ {
		if words[i]==":" {
			mymap[removeQuotes(words[i-1])]=removeQuotes(words[i+1])
		}
	}
	return mymap
}


func removeQuotes(string2 string)string{
	if string2[0] == '"' {
		string2 = string2[1:]
	}
	if i := len(string2)-1; string2[i] == '"' {
		string2 = string2[:i]
	}

	return string2
}

