package handler

import (
	"net/http"
	"testcases-gen/grabing"
	"testcases-gen/writer"

	"github.com/labstack/echo"
)

type GenTestcasesRequest struct {
	URL  string `json:"url"`
	IsOk bool   `json:"isOk"`
}

type GenTestcasesResponse struct {
	IsOk      bool   `json:"isOk"`
	TestCases string `json:"testCases"`
}

func GenTestcasesHandler(c echo.Context) error {
	req := new(GenTestcasesRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	response, err := generateTestcases(*req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error generating test cases: "+err.Error())
	}

	return c.JSON(http.StatusOK, response)
}

func generateTestcases(req GenTestcasesRequest) (GenTestcasesResponse, error) {
	exampleArr, err := grabing.GrabExamples(req.URL)
	if err != nil {
		return GenTestcasesResponse{}, err
	}
	script := writer.MakeTestcases(exampleArr)

	response := GenTestcasesResponse{
		IsOk:      req.IsOk,
		TestCases: script,
	}
	return response, nil
}
