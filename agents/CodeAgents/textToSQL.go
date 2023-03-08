package codeAgents

import "github.com/asolpshinning/chaingpt/entity"

func TextToSQL(text string, tool *entity.Tool) (*entity.AgentResponse, error) {
	// put the code the sends the text to GPT-3 and returns the SQL code or a message that says no

	result := "SELECT * FROM table WHERE id = 1"
	agentResponse := &entity.AgentResponse{
		Input:        text,
		Output:       result,
		Satisfactory: true,
	}
	return agentResponse, nil
}
