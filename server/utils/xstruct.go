package utils

import (
	"reflect"
	"regexp"
)

const (
	REGEX_COMMENT_VALUE = `comment:'(.*?)'`
	FIELD_JSON_VALUE    = "json"
	FIELD_GORM_VALUE    = "gorm"
)


func GetStructTagName(i interface{}, fieldName string) []string {
	types := reflect.TypeOf(i)
	numField := types.Elem().NumField()
	var tagValues = make([]string, numField)
	for i := 0; i < numField; i++ {
		tagValue := types.Elem().Field(i).Tag.Get(fieldName)
		if tagValue == "" {
			continue
		}
		tagValues[i] = tagValue
	}
	return tagValues
}


func GetStructTagNameRegex(i interface{}, fieldName string, regStr string) []string {
	reg := regexp.MustCompile(regStr)
	types := reflect.TypeOf(i)
	numField := types.Elem().NumField()
	var tagValues = make([]string, numField)
	for i := 0; i < numField; i++ {
		tagValue := types.Elem().Field(i).Tag.Get(fieldName)
		if tagValue == "" {
			continue
		}
		matchArr := reg.FindStringSubmatch(tagValue)
		tagValues[i] = matchArr[len(matchArr)-1]
	}
	return tagValues
}
