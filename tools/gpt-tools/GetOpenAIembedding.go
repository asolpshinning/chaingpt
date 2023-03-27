package gpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	openAIEmbeddingsEndpoint = "https://api.openai.com/v1/embeddings"
)

type requestBody struct {
	Input string `json:"input"`
	Model string `json:"model"`
}

type responseBody struct {
	Data []float64 `json:"data"`
}

// GetEmbedding generates an embedding for a given text using the OpenAI API.
func GetEmbedding(apiKey, text string) ([]float64, error) {
	data := requestBody{
		Input: text,
		Model: "text-embedding-ada-002",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", openAIEmbeddingsEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result responseBody
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	text := "Your text string goes here"

	embedding, err := GetEmbedding(apiKey, text)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Embedding:", embedding)
	}
}
