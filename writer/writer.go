package writer

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	wholeScript = `package main

import (
	"reflect"
	"testing"
)
	
func TestFunc(t *testing.T) {
	%v
}`

	testStructPart = `type testStruct struct {
		%v
	}`

	testArrayPart = `	testcases := []testStruct{
		%v
	}`

	loopPart = `	for _, tc := range testcases {
		actual := leetcodefuncName(%v)
		if !reflect.DeepEqual(actual, tc.Output) {
			t.Errorf("Expected %v, but got %v", tc.Output, actual)
		}
	}`
	// need to grab leetcodefuncName from web content
)

func MakeTestcases(examples string) string {
	theCoolestArr := removeEmptyStr(examples)
	structText, loopText := createTestStruct(theCoolestArr)
	arrText := createTestArr(theCoolestArr)

	innetScript := structText + "\n\n" + arrText + "\n\n" + loopText
	wholeScript := fmt.Sprintf(wholeScript, innetScript)
	writeStringToFile("result.txt", wholeScript)

	return ""
}

func removeEmptyStr(examples string) []string {
	splitNewLine := strings.Split(examples, "\n")
	noEmptyArr := []string{}

	for _, v := range splitNewLine {
		if len(v) > 0 {
			noEmptyArr = append(noEmptyArr, v)
		}
	}

	return noEmptyArr
}

func createTestStruct(arr []string) (string, string) {
	stuctText := ""
	fieldsForStruct := ""
	limiter := 1
	functionInput := ""

	for _, v := range arr {
		isExample := strings.Index(v, "Example")
		if isExample >= 0 {
			if limiter == 0 {
				break
			}
			continue
		}
		isInput := strings.Index(v, "Input")
		if isInput >= 0 {
			parts := strings.Split(v, ",")
			for _, v := range parts {
				parts := strings.Split(v, " ")
				fieldName := parts[1]
				fieldValue := parts[len(parts)-1]
				fieldType := determineType(fieldValue)
				fieldsForStruct += fieldName + " " + fieldType + "\n\t\t"
				functionInput += "tc." + fieldName + ", "
			}
		}
		isOutput := strings.Index(v, "Output")
		if isOutput >= 0 {
			limiter -= 1
			parts := strings.Split(v, " ")
			fieldName := "Output"
			fieldValue := parts[len(parts)-1]
			fieldType := determineType(fieldValue)
			fieldsForStruct += fieldName + " " + fieldType
		}
	}
	stuctText += fmt.Sprintf(testStructPart, fieldsForStruct)
	loopText := fmt.Sprintf(loopPart, functionInput)
	return stuctText, loopText
}

func createTestArr(arr []string) string {
	theText := ""
	fieldsForArr := ""
	for _, v := range arr {
		isExample := strings.Index(v, "Example")
		if isExample >= 0 {
			fieldsForArr += "{\n\t\t"
			continue
		}
		isInput := strings.Index(v, "Input")
		if isInput >= 0 {
			parts := strings.Split(v, ",")
			for _, v := range parts {
				parts := strings.Split(v, " ")
				fieldName := parts[1]
				fieldValue := parts[len(parts)-1]
				fieldsForArr += "\t" + fieldName + ": " + fieldValue + ",\n\t\t"
			}
		}
		isOutput := strings.Index(v, "Output")
		if isOutput >= 0 {
			parts := strings.Split(v, " ")
			fieldName := "Output"
			fieldValue := parts[len(parts)-1]
			fieldsForArr += "\t" + fieldName + ": " + fieldValue + ",\n\t\t},\n\t\t"
		}
	}
	theText += fmt.Sprintf(testArrayPart, fieldsForArr)

	fmt.Println(theText)
	return theText
}

func writeStringToFile(filename, data string) error {
	// Write the string data to the file
	err := os.WriteFile(filename, []byte(data), 0644) // 0644 for file permissions
	if err != nil {
		return err
	}
	return nil
}

func determineType(value string) string {
	// Attempt to parse the value as a bool
	if _, err := strconv.ParseBool(value); err == nil {
		return "bool"
	}

	// Attempt to parse the value as an int
	if _, err := strconv.Atoi(value); err == nil {
		return "int"
	}

	// Attempt to parse the value as a float64
	if _, err := strconv.ParseFloat(value, 64); err == nil {
		return "float64"
	}

	// If none of the above types can be parsed, consider it a string
	return "string"
}
