package codeAgents

import (
	"github.com/asolpshinning/chaingpt/entity"
)

func TextToSQL(text string, tool *entity.Tool) (*entity.AgentResponse, error) {
	// put the code the sends the text to GPT-3 and returns the SQL code or a message that says no
	/* var result string
	var err error
	if tool.Value == "chatGPT" {
		result, err = gpt.ChatGPT(text)
	} else if tool.Value == "flanT5" {
		result, err = gpt.FlanT5(text)
	}
	if err != nil {
		return nil, err
	} */
	//just a sample result for test purposes. uncomment above when ready
	result := "SELECT * FROM table WHERE id = 1"
	agentResponse := &entity.AgentResponse{
		Input:        text,
		Output:       result,
		Satisfactory: true,
	}
	return agentResponse, nil
}
