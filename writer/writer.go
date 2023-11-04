package writer

import (
	"fmt"
	"os"
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

	testStructPart = `	type testStruct struct {
		%v
	}`

	arrayPart = `	testcases := []testStruct{
		%v
	}`

	loopPart = `	for _, tc := range testcases {
		actual := leetcodefuncName(tc.n, tc.pick)
		if !reflect.DeepEqual(actual, tc.Output) {
			t.Errorf("Expected %v, but got %v", tc.Output, actual)
		}
	}`
)

func MakeTestcases(examples string) string {
	theCoolestArr := removeEmptyStr(examples)
	theText := createTestStruct(theCoolestArr)

	writeStringToFile("main_test.go", theText)

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

func createTestStruct(arr []string) string {
	theText := ""
	for _, v := range arr {
		isExample := strings.Index(v, "Example")
		if isExample > 0 {
			continue
		}
		theText += v + "\n"
	}
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
