package aocutils

import (
	"fmt"
	"os"
	"strings"
)

func ReadInput(fileName string) []string {

	b, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
	}

	fileAsString := string(b)
	return strings.Split(fileAsString, "\n")
}

func ReadInputWithDelimeter(fileName string, delimeter string) []string {

	b, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
	}

	fileAsString := string(b)
	return strings.Split(fileAsString, delimeter)
}
