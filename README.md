# Introduction

This code base is a simple library I am developing for AI-powered applications that can make use of multiple agents with different skills or abilities. The agents are simply functions that can execute certain instructions (as specified). The `agents` can use `tools` to carry out their tasks. More interestingly, the concept of `chains` has also been introduced to allow `agents` to work together to achieve a specific goal. A chain will be able to call an agent or another chain to execute a task and then pass the result to another agent to execute another task. It is designed to be super flexible. This allows for a more complex task or REASONING to be achieved by a chain of agents and different tools.


## Tools Built So Far
- `gpt-tools` - This tool makes use of the [chatGPT API](https://chatgpt.com/) to communicate with the OpenAI GPT-3 model. It can be used to generate text from a prompt.
- `doc-tools` - This tool helps to read and write to document files like `markdown` files.
- `git-tools` - This tool currently helps to automate committing and pushing changes to a git repository.

## Agents Built So Far
- `codeAgents.TextToSQL` - This agent takes natural language (english) input from user and then convert it to `sql` query.
- `codeAgents.QueryResultToEnglish` - This agent takes the result of a `sql` query and then convert it to natural language the user can understand.

## Chains Built So Far
- `chain.ChatWithDatabase` - This chain takes natural language (english) input from user and then convert it to `sql` query, run the query on a database and then convert the result to natural language and return it to the user. This chain makes use of agents like `codeAgents.TextToSQL`, and `codeAgents.QueryResultToEnglish` to achieve this. It also makes use of `entityTools.RunPostgresQuery` to run the query on the database, and `gpt.ChatGPT` to communicate with chatGPT API.

***
# Example Usage of the Library
***

## Chat with `ChatGPT` In Your Markdown Files Using `doc-tools`, `gpt-tools` and `git-tools`
Prerequisites:
- [Install Go on your machine](https://golang.org/doc/install)
- Open the repository you want to use and run the following command to install this library
```bash
go get github.com/asolpshinning/chaingpt/tools
```
-  Add `token="your_chatGpt_openAI_api_token"` to your `.env` file.
- Remember to add `/gen` as a text to your markdown file to indicate that any text above `/gen` is your prompt to chatGPT.
- Then put the following example in a file called `main.go` in the root of your repository and run `go run main.go` to start chatting in your markdown file.

<details>
<summary><b><i>See the example</i></b></summary>

```go
package main

import (
	doc "github.com/asolpshinning/chaingpt/tools/doc-tools"
	git "github.com/asolpshinning/chaingpt/tools/git-tools"
	gpt "github.com/asolpshinning/chaingpt/tools/gpt-tools"
)

// your git folder path (example below is for Windows OS)
var folderPath = "C:\\Users\\asolp\\OneDrive\\Documents\\Coding\\chaingpt"

// your markdown file path (where you are doing your chatting)
var fileName = "test/chat.md"

// time interval for git auto commit and push
var timeInterval = 300 //seconds

func main() {
	// this autocommits and pushes your changes to github (comment it out if you do not want this)
	go git.GitAutoCommitPush(folderPath, timeInterval)
	// this gets the prompt from your markdown file
	prompt, _ := doc.CopyAboveText(fileName)
	// this sends the prompt to chatGPT and gets the response
	chatGPTResponse, _ := gpt.ChatGPT(prompt)
	// this inserts the response from chatGPT into your markdown file
	doc.InsertChatResponse(fileName, chatGPTResponse)
}
```

</details>

***

## Using `ChatWithDatabase` Chain
<details>
<summary><b><i>See the example</i></b></summary>

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
</details>

***

## Request for Features
- Feel free to open an issue if you have any feature request or suggestion. I will try to implement it as soon as I can.