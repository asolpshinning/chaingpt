package codeAgents

import "github.com/asolpshinning/chaingpt/entity"

func ProcessDBqueryResult(queryResult string, tool *entity.Tool) (*entity.AgentResponse, error) {
	// put the code the sends the text to GPT-3 and returns the interpretation of the result

	result := "you have 5 new messages... you need to test this with LLM soon"
	agentResponse := &entity.AgentResponse{
		Input:        queryResult,
		Output:       result,
		Satisfactory: true,
	}
	return agentResponse, nil
}
