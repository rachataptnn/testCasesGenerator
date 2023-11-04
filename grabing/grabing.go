package grabing

import (
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
)

func GrabExamples(url string) (string, error) {
	content, err := grabWholePageContent(url)
	if err != nil {
		return "", err
	}
	problemDesc := grabProblemDesc(content)

	parts := strings.Split(problemDesc, "\u00a0")
	if len(parts) > 1 {
		examples := parts[1]
		return examples, nil
	}

	return "", errors.New("parts length incorrect")
}

func grabWholePageContent(url string) (string, error) {
	// url = "https://leetcode.com/problems/guess-number-higher-or-lower/description/?fbclid=IwAR1d9rN7CE8dHbgoJKm7rOanvl8jX45NHV7jv7sj1OHXOf6RzCftXwZZBwg"
	content, err := fetchURLContents(url)
	if err != nil {
		return "", err
	}
	return content, nil
}

func grabProblemDesc(content string) string {
	searchString := "<meta name=\"description\""
	descriptionStartIndex := strings.Index(content, searchString)

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
