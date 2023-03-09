package main

import (
	"fmt"
	"log"

	"github.com/asolpshinning/chaingpt/chains"
	"github.com/asolpshinning/chaingpt/entity"
)

func main() {
	// create a user agent
	userAgent := &entity.Agent{
		Name: "User",
	}

	// create a user prompt
	userPrompt := "I want to know how many new messages I have so far."

	// create user agent's input or prompt to the chain
	userInput := &entity.AgentResponse{
		Input:        "",
		Output:       userPrompt,
		Satisfactory: true,
	}

	// create the tool the chain will use
	tool := &entity.Tool{
		Name:  "QueryDatabase",
		Value: "postgres",
	}

	postgresTools := []*entity.Tool{tool}

	// let the user agent call the SQLChain
	res, err := chains.ChatWithDatabase(userAgent, userInput, postgresTools)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("response: " + res.Response)
}
