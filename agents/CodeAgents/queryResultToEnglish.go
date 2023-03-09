package codeAgents

import (
	"github.com/asolpshinning/chaingpt/entity"
	"github.com/asolpshinning/chaingpt/tools/gpt"
)

func QueryResultToEnglish(queryResult string, tool *entity.Tool) (*entity.AgentResponse, error) {

	result, err := gpt.ChatGPT(queryResult)
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
