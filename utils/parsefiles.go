package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

func DownloadFile(filePath string, url string) error {

	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	out, err := os.Create(filePath)

	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, response.Body)

	return err

}

func ParseFile(path string) (dataArray []string) {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	filteredDataRegex := regexp.MustCompile(`(<custom_item>)(.*?)(</custom_item>)`)

	dataArray = filteredDataRegex.FindAllString(replaceNewLines(data), -1)

	return
}

func replaceNewLines(data []byte) (dataWithOutNewLines string) {
	regexToDeleteNeLines := regexp.MustCompile(`\r?\n`)
	dataWithOutNewLines = regexToDeleteNeLines.ReplaceAllString(string(data), " ")
	return
}
