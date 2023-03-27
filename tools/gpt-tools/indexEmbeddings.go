package gpt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// ReadMarkdownFiles reads all the markdown files in a specified folder and returns their contents.
func ReadMarkdownFiles(folder string) (map[string]string, error) {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	markdownFiles := make(map[string]string)

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".md") {
			content, err := ioutil.ReadFile(filepath.Join(folder, file.Name()))
			if err != nil {
				return nil, err
			}
			markdownFiles[file.Name()] = string(content)
		}
	}

	return markdownFiles, nil
}

// BuildIndex generates embeddings for the markdown files and stores them in an index.
func BuildIndex(apiKey string, markdownFiles map[string]string) ([]Embedding, error) {
	index := []Embedding{}

	for id, content := range markdownFiles {
		embedding, err := generateEmbedding(apiKey, content)
		if err != nil {
			return nil, err
		}

		embeddingVector, err := parseEmbedding(embedding)
		if err != nil {
			return nil, err
		}

		index = append(index, Embedding{ID: id, Embedding: embeddingVector})
	}

	return index, nil
}

// SaveIndex saves the index to a specified directory as a JSON file.
func SaveIndex(index []Embedding, directory string) error {
	jsonData, err := json.Marshal(index)
	if err != nil {
		return err
	}

	indexFile := filepath.Join(directory, "index.json")
	err = ioutil.WriteFile(indexFile, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

// LoadIndex loads the index from a specified JSON file.
func LoadIndex(directory string) ([]Embedding, error) {
	indexFile := filepath.Join(directory, "index.json")
	data, err := ioutil.ReadFile(indexFile)
	if err != nil {
		return nil, err
	}

	var index []Embedding
	err = json.Unmarshal(data, &index)
	if err != nil {
		return nil, err
	}

	return index, nil
}

// parseEmbedding is a helper function to parse the embedding string into a []float64.
func parseEmbedding(embedding string) ([]float64, error) {
	// Implement your own logic to parse the embedding string
	// into a []float64 based on your chosen format.
	return nil, nil
}

// generateEmbedding is a helper function to generate an embedding for a given text using the OpenAI API.
func generateEmbedding(apiKey, text string) (string, error) {
	// Implement the function to call the OpenAI API as demonstrated earlier.
	return "", nil
}

func TestingIndexEmbedding() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	folder := "path/to/markdown/files"

	// Read markdown files
	markdownFiles, err := ReadMarkdownFiles(folder)
	if err != nil {
		fmt.Println("Error reading markdown files:", err)
		return
	}

	// Build index
	index, err := BuildIndex(apiKey, markdownFiles)
	if err != nil {
		fmt.Println("Error building index:", err)
		return
	}

	// Save index
	err = SaveIndex(index, folder)
	if err != nil {
		fmt.Println("Error saving index:", err)
	}
}
