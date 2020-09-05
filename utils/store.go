package utils

import (
	"strings"
)


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
			mymap[words[i-1]]=words[i+1]
		}
	}
	return mymap
}
