package utils

import (
	"Security-Benchmarking-Tool/constants"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func GenerateSavedFileName(path string, fileFormat string, fileType string) (savedFileName string, err error) {
	counter, err := fileCount(path)
	if err != nil {
		return "", err
	}
	dt := time.Now()
	date := fmt.Sprint(dt.Format("01022006"))

	savedFileName = path + fileType + date + counter + fileFormat // here fileType if policy or parsedData file format is .json or .audit
	return
}

func fileCount(path string) (numberOfFiles string, err error) {
	i := 0
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return "", err
	}
	for _, file := range files {
		if !file.IsDir() {
			i++
		}
	}
	numberOfFiles = strconv.Itoa(i)
	return
}

func CreateAuditsDir(home string) (err error) {
	home, err = GetUserHome()
	if err != nil {
		return
	}
	err = os.MkdirAll(home+constants.DESKTOP+constants.AuditDirectory+constants.SavedFileDIRECTORY, 0755)
	if err != nil {
		return err
	}

	err = os.MkdirAll(home+constants.DESKTOP+constants.AuditDirectory+constants.ParsedDataDirectory, 0755)
	if err != nil {
		return err
	}

	err = os.MkdirAll(home+constants.DESKTOP+constants.AuditDirectory+constants.CustomAuditDirectory, 0755)
	if err != nil {
		return
	}

	return nil
}

func GetUserHome() (home string, err error) {

	home, err = os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return
}

func RemoveQuotes(string2 string) string {
	if string2[0] == '"' {
		string2 = string2[1:]
	}
	if i := len(string2) - 1; string2[i] == '"' {
		string2 = string2[:i]
	}

	return string2
}
