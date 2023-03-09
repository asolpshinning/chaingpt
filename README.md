# Introduction

This code base is a simple library I am developing for AI-powered applications that can make use of multiple agents with different skills or abilities. The agents are simply functions that can execute certain instructions (as specified). The `agents` can use `tools` to carry out their tasks. More interestingly, the concept of `chains` has also been introduced to allow `agents` to work together to achieve a specific goal. A chain will be able to call an agent or another chain to execute a task and then pass the result to another agent to execute another task. It is designed to be super flexible. This allows for a more complex task or REASONING to be achieved by a chain of agents and different tools.


## Tools Built So Far
- 

## Agents Built So Far
- `codeAgents.TextToSQL` - This agent takes natural language (english) input from user and then convert it to `sql` query.
- `codeAgents.QueryResultToEnglish` - This agent takes the result of a `sql` query and then convert it to natural language the user can understand.

## Chains Built So Far
- `chain.ChatWithDatabase` - This chain takes natural language (english) input from user and then convert it to `sql` query, run the query on a database and then convert the result to natural language and return it to the user. This chain makes use of agents like `codeAgents.TextToSQL`, and `codeAgents.QueryResultToEnglish` to achieve this. It also makes use of `entityTools.RunPostgresQuery` to run the query on the database, and `gpt.ChatGPT` to communicate with chatGPT API.


## Example Usage of the Library

<details>
<summary>Using ChatWithDatabase Chain</summary>

```go
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
	userPrompt := "I want to know how many new messages I have on Friday and Sunday."

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

	// let the user agent call the ChatWithDatabase chain
	res, err := chains.ChatWithDatabase(userAgent, userInput, postgresTools)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("response: " + res.Response)
}
```
## .env example to run the example code above
```bash
token="your_chatGpt_openAI_api_token"
```
</detials>

## Request for Features
- Feel free to open an issue if you have any feature request or suggestion. I will try to implement it as soon as I can.