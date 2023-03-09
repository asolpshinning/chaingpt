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
	// this autocommits and pushes your changes to github
	go git.GitAutoCommitPush(folderPath, timeInterval)
	prompt, _ := doc.CopyAboveText(fileName)
	chatGPTResponse, _ := gpt.ChatGPT(prompt)
	doc.InsertChatResponse(fileName, chatGPTResponse)
}
