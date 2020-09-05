package utils

import (
	"strings"
	"unicode"
)


func CreateMapForMultipleItems(arrayOfParsedData []string) []map[string]interface{} {
	var arrayOfMaps []map[string]interface{}
	for _,value :=range arrayOfParsedData{
		arrayOfMaps = append(arrayOfMaps,CreateMapForSingleItem(strings.TrimLeft(strings.TrimRight(value, "</custom_item>"), "<custom_item>")))
	}

	return arrayOfMaps
}


func CreateMapForSingleItem(myStr string) map[string]interface{} {
	m := map[string]interface{}{}

	for k,v := range myStr {
		if v==58 {
			m["First"]=map[string]interface{}{
				myStr[:k]:myStr[k+1:],
			}
		}
	}

	return m
}

func removeSpace(s string) string {
	rr := make([]rune, 0, len(s))
	for _, r := range s {
		if !unicode.IsSpace(r) {
			rr = append(rr, r)
		}
	}
	return string(rr)
}
