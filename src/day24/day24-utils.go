package main

import (
	"bufio"
	"fmt"
	"os"
)

// Read the text file passed in by name into a array of strings
// Returns the array as the first return variable
func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// Prints the map list
func printMap(tempMap []string) {
	for i := 0; i < len(tempMap); i++ {
		fmt.Printf("%s\n", tempMap[i])
	}
}

// func: readInitialState
// takes an array of strings and breaks it into a 2D array of bytes
func readInitialState(tempString []string, tempSlice [][]byte) {
	for i := 0; i < len(tempString); i++ {
		for j := 0; j < len(tempString[i]); j++ {
			tempSlice[i][j] = tempString[i][j]
		}
	}
}

func print2DSlice(tempSlice [][]byte) {
	for i := 0; i < len(tempSlice); i++ {
		for j := 0; j < len(tempSlice[i]); j++ {
			fmt.Printf("%c", tempSlice[i][j])
		}
		fmt.Printf("\n")
	}
}

// Assumes that destination slice has enough memory allocated to hold the contents of sourceSlice
func copy2DSlice(sourceSlice [][]byte, destinationSlice [][]byte) {
	for i := 0; i < len(sourceSlice); i++ {
		for j := 0; j < len(sourceSlice[i]); j++ {
			destinationSlice[i][j] = sourceSlice[i][j]
		}
	}
}
