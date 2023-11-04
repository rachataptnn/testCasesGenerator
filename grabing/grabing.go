package grabing

import (
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
)

func GrabExamples(url string) (string, string, error) {
	content, err := grabWholePageContent(url)
	if err != nil {
		return "", "", err
	}
	funcName := grabFunctionName(content)
	problemDesc := grabProblemDesc(content)

	parts := strings.Split(problemDesc, "\u00a0")
	if len(parts) > 1 {
		examples := parts[1]
		return examples, funcName, nil
	}

	return "", "", errors.New("parts length incorrect")
}

func grabWholePageContent(url string) (string, error) {
	content, err := fetchURLContents(url)
	if err != nil {
		return "", err
	}
	return content, nil
}

func grabProblemDesc(content string) string {
	descriptionStartIndex := strings.Index(content, "<meta name=\"description\"")

	if descriptionStartIndex >= 0 {
		choppedContent := content[descriptionStartIndex:]
		descriptionEndIndex := strings.Index(choppedContent, "><")
		if descriptionEndIndex >= 0 {
			problemDescription := choppedContent[:descriptionEndIndex]
			return problemDescription
		}
		log.Println("description end index is not found")
		return ""
	}
	log.Println("description start index is not found")
	return ""
}

func fetchURLContents(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	content := string(body)

	return content, nil
}

func grabFunctionName(content string) string {
	// find go {} first
	goIndex := strings.Index(content, "{\"lang\":\"Go\",\"langSlug\":\"golang\"")
	slice1 := content[goIndex:]
	goOnly := strings.Split(slice1, "},{")[0]

	// there are two 'func' so i choose \nfunc
	secondFuncIndex := strings.Index(goOnly, "\\nfunc")
	slice2 := ""
	if secondFuncIndex >= 0 {
		slice2 = goOnly[secondFuncIndex:]
	} else {
		secondFuncIndex := strings.Index(goOnly, "\"func")
		slice2 = goOnly[secondFuncIndex:]
	}
	// want only function name
	startBracketIndex := strings.Index(slice2, "(")
	slice3 := slice2[:startBracketIndex]

	funcName := strings.Split(slice3, " ")[1]

	return funcName
}
