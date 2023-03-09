package doctools

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// this function opens the file and returns the current contents of the file and the file
func OpenFileWrite(fileName string) (*os.File, error) {
	// Open the file for reading and writing
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return file, nil
}

// this function prompts the user for input and returns the input without the delimiter '.' and newline characters
func promptUser() (string, error) {
	// Prompt the user to enter userInput in a loop
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter userInput: ")
	userInput, err := reader.ReadString('.')
	if err != nil {
		if err == io.EOF {
			fmt.Println("EOF")
		} else {
			log.Fatal(err)
			return "", err
		}
	}
	// Remove newline character and delimiter from userInput input
	userInput = strings.TrimRight(userInput, "."+"\n")
	userInput = strings.TrimRight(userInput, "\n"+"\r")
	userInput = strings.TrimLeft(userInput, "\n"+"\r")

	return userInput, err
}

func InsertChatResponse(fileName string, response string) error {

	for {
		userInput, err := promptUser()
		if err != nil {
			log.Fatal(err)
			return err
		}
		// If the user entered "joy", insert the desired userInput
		if userInput == "g" {

			// Open the file for writing
			file, err := OpenFileWrite(fileName)
			if err != nil {
				log.Fatal(err)
			}

			newPos, err := file.WriteString("\n" + response + "\n")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("new position after writing: %v \n", newPos)
			fmt.Println("userInput inserted!")
			file.Close()
		} else {
			fmt.Println("Invalid input. Please try again.")
		}
	}

}
