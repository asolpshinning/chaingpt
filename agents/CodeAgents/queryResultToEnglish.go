package codeAgents


import (
	"errors"

	"github.com/asolpshinning/chaingpt/entity"
	gpt "github.com/asolpshinning/chaingpt/tools/gpt-tools"
)

func QueryResultToEnglish(queryResult string, tool *entity.Tool) (*entity.AgentResponse, error) {
	if tool.Type != "gpt" {
		return nil, errors.New("the tool provided is not a gpt tool")
	}
	var result string
	var err error
	if tool.Value == "chatGPT" {
		result, err = gpt.ChatGPT(queryResult)
	} else if tool.Value =="flanT5" {
		result, err = gpt.FlanT5(queryResult)
	}
	if err != nil {
		return nil, err
	}
	agentResponse := &entity.AgentResponse{
		Input:        queryResult,
		Output:       result,
		Satisfactory: true,
	}
	return agentResponse, nil
}

func YourCustomizedT5Function(prompt string) (string, error) {
	// your code here
	return "", nil
}