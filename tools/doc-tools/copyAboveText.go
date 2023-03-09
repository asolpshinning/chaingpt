package doctools

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CopyAboveText(fileName string) (string, error) {
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Define the variable to store the lines
	var str string

	// Loop through each line of the file until you reach the line you want to stop at
	for scanner.Scan() {
		line := scanner.Text()

		// Add the line to the string variable
		str += line + "\n"

		// Check if it's the line to start at and empty the string variable
		if strings.Contains(line, "/start") {
			str = ""
			continue
		}

		// Check if it's the line to stop at
		if strings.Contains(line, "/gen") {
			break
		}
	}
	file.Close()
	return str, nil
}
