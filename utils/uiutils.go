package utils

import (
	"strconv"
)

// create array of custom maps from indexes of main array of audits;
func CreateMapOfAuditsFromIndexArray(indexArray []int, auditsMapArray []map[string]string) (customAudits []map[string]string) {
	for _, v := range indexArray {
		customAudits = append(customAudits, auditsMapArray[v])
	}
	return
}

// convert array of indexes(string) to array of indexes on int;
func ConvertArrayToInt(indexesArrayOfString []string) (intArray []int, err error) {
	intArray = make([]int, 0, len(indexesArrayOfString))
	for _, v := range indexesArrayOfString {
		if k, err := strconv.Atoi(v); err != nil {
			return nil, err
		} else {
			intArray = append(intArray, k)
		}
	}
	return
}

// xuio znaet pentru ce no on nujen ea zabil for what
// used in such a way:
func test(map1 map[string]string, mapWithInts map[int]string) (arrayWithLengthEleven []string) {
	arrayWithLengthEleven = make([]string, len(mapWithInts))
	for k, v := range mapWithInts {
		for key, value := range map1 {
			if v == key {
				arrayWithLengthEleven[k-1] = value
			}
		}
	}
	return
}

// create interface to add row
func createInterface(stringArr []string) (myInterface []interface{}) {
	myInterface = make([]interface{}, len(stringArr))
	for i, s := range stringArr {
		myInterface[i] = s
	}

	return
}

// get maps with max len of keys
func getMapsWithMaxNumberOfKey(myMap []map[string]string) (maxMap map[string]string) {
	maxMap = myMap[0]
	for _, v := range myMap {
		if len(maxMap) < len(v) {
			maxMap = v
		}
	}
	return
}
