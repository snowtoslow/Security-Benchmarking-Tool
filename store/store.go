package store

import (
	"Security-Benchmarking-Tool/constants"
	"Security-Benchmarking-Tool/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//TWO MAIN METHODS FOR NOW WITH PRACTICALLY SAME LOGIC
// here as input need to be path := /home/snowtoslow/Desktop/audit/policy-info/*policy-name*;
// infoToBeParsed result from CreateMapForMultipleElements
// *arrayData:=utils.ParseFile(utils.GenerateFileNames())
//	info:=utils.CreateMapForMultipleItems(arrayData)
//	fmt.Println(info)

func CreateJsonResponse(path string, infoToBeParsed []map[string]string) (err error) {
	file, err := json.MarshalIndent(infoToBeParsed, "", " ")
	err = ioutil.WriteFile(path, file, 0644)
	if err != nil {
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
func DownloadFileToExpectedLocation(fileName string) (err error) {
	/*HOME, err := utils.GetUserHome()
	fileName,err := utils.GenerateSavedFileName(HOME+constants.DESKTOP + constants.AuditDirectory + constants.SavedFileDIRECTORY,constants.AuditFormat,constants.Policy)*/

	if err := utils.DownloadFile(fileName, constants.LinkToDownloadFrom); err != nil {
		return err
	}
	defer log.Println("Downloaded")
	return nil
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func CreateMapForMultipleItems(arrayOfParsedData []string) []map[string]string {
	var arrayOfMaps []map[string]string
	for _, value := range arrayOfParsedData {
		arrayOfMaps = append(arrayOfMaps, createMapForSingleItem(strings.TrimLeft(strings.TrimRight(value, "</custom_item>"), "<custom_item>")))
	}

	return arrayOfMaps
}

func createMapForSingleItem(myStr string) (mymap map[string]string) {
	regex := *regexp.MustCompile(`(.+?)\s*:\s*\\*"(.+?)"`) /*`(.+?)\s*:\s*\\*"(.+?\".+?(\s.+?))*/
	regexType := *regexp.MustCompile(`(.+?)\s*:\s*(.[A-Z_]*)`)
	mymap = make(map[string]string)
	res := regex.FindAllStringSubmatch(myStr, -1)

	for i := range res {
		if strings.Contains(res[i][1], "type") {
			mynewArrWithWords := regexType.FindStringSubmatch(res[i][1])[0]
			res[i][1] = strings.Replace(regexType.Split(res[i][1], -1)[1], " ", "", -1)
			mymap[res[i][1]] = res[i][2]
			mymap[strings.Replace(strings.Split(mynewArrWithWords, ":")[0], " ", "", -1)] = strings.Replace(strings.Split(mynewArrWithWords, ":")[1], " ", "", -1)
		} else {
			mymap[strings.Replace(res[i][1], " ", "", -1)] = res[i][2]
		}
	}
	return
}

func SearchItemsByKey(arrayToSearchIn []map[string]string, searcheableItem string) (mapOfSearchedValues []map[string]string, message string) {
	for _, value := range arrayToSearchIn {
		for k, _ := range value {
			if strings.ToLower(value[k]) == strings.ToLower(searcheableItem) {
				mapOfSearchedValues = append(mapOfSearchedValues, value)
				message = "The configurations was found!"
			} else {
				message = "Not found!"
			}
		}
	}

	return
}

func CreateCustomPolicy(path string, mapOfCharacteristics []map[string]string) (err error) {
	fileWriter, err := os.Create(path)
	if err != nil {
		return err
	}

	defer fileWriter.Close()

	for i := 0; i < len(mapOfCharacteristics); i++ {
		_, err := fileWriter.WriteString(createStringWithCustomItem(mapOfCharacteristics[i]))
		if err != nil {
			return err
		}
	}
	return nil
}

// function to create array of custom strings which will be put in file:
func createStringWithCustomItem(mapOfCharacteristics map[string]string) (customString string) {
	var valueOfCharacteristics string
	for k, v := range mapOfCharacteristics {
		valueOfCharacteristics += fmt.Sprintf("\t%s:%s\n", k, v)
	}
	customString = fmt.Sprintf("%s\n%s%s\n", constants.CustomItemStart, valueOfCharacteristics, constants.CustomItemEnd)
	return
}
