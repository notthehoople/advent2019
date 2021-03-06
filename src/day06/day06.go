package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
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
func printMapList(tempMapList map[string]string) {
	for key, value := range tempMapList {
		fmt.Printf("Object %s orbits around object %s\n", key, value)
	}
}

// Part B function. Builts a "distance" map for the supplied orbitList
//		Expectation is that this will have 2 "top level" calls - one for "YOU" and one for "SAN"
func createDistanceMap(orbitor string, orbitList map[string]string, distanceMap map[string]int) {
	// We count 1 each time there is a higher level object to orbit i.e. if the orbitor we're passed is itself orbiting something
	// Most likely we'll stop counting at COM but let's not make that assumption
	thingWeOrbit, moreObjects := orbitList[orbitor]
	if moreObjects {
		// There are more things to count. Take note of the distance of this step, then continue
		distanceMap[thingWeOrbit] = distanceMap[orbitor] + 1
		createDistanceMap(thingWeOrbit, orbitList, distanceMap)
	}
}

func countNumberOrbits(orbitor string, orbitList map[string]string) int {
	// We count 1 each time there is a higher level object to orbit i.e. if the orbitor we're passed is itself orbiting something
	// Most likely we'll stop counting at COM but let's not make that assumption
	thingWeOrbit, moreObjects := orbitList[orbitor]
	if moreObjects {
		// There are more things to count
		return countNumberOrbits(thingWeOrbit, orbitList) + 1
	}

	return 0
}

// func
// Returns:
func orbitCalculation(filename string, debug bool, part byte) int {
	// Read contents of file into a string array
	fileContents, _ := readLines(filename)

	orbitList := make(map[string]string)

	// What do we do with COM to make sure it's counted?

	for i := 0; i < len(fileContents); i++ {
		processObjects := strings.Split(fileContents[i], ")")
		orbitList[processObjects[1]] = processObjects[0]
	}

	if part == 'a' {
		var orbitCount int

		for orbitor := range orbitList {
			// Loop through the orbitors and count the number of orbits they have
			orbitCount += countNumberOrbits(orbitor, orbitList)
		}

		if debug {
			printMapList(orbitList)
		}

		return orbitCount
	}

	// part b

	// Build map of YOU to COM with a distance at each step, counting from YOU to COM (e.g. COM is the LAST step in the list with the HIGHEST distance)
	// Build map of SAN to COM with a distance at each step, counting from SAN to COM (e.g. COM is the LAST step in the list with the HIGHEST distance)

	distanceMapYOU := make(map[string]int)
	distanceMapSAN := make(map[string]int)

	distanceMapYOU["YOU"] = -1
	createDistanceMap("YOU", orbitList, distanceMapYOU)
	distanceMapSAN["SAN"] = -1
	createDistanceMap("SAN", orbitList, distanceMapSAN)

	// Loop through the YOU list
	// 		IF OBJECT is in santa's list AND
	//			distance is YOU distance to OBJECT + santa distance to OBJECT
	//			if this is the lowest distance seen keep it

	orbitTransfers := 999999
	for object, youDistance := range distanceMapYOU {
		santaDistance, gotSanta := distanceMapSAN[object]
		if gotSanta {
			if (santaDistance + youDistance) < orbitTransfers {
				orbitTransfers = santaDistance + youDistance
			}
		}

	}

	return orbitTransfers
}

// Main routine
func main() {
	var debug bool

	filenamePtr := flag.String("file", "input.txt", "Filename containing the program to run")
	execPartPtr := flag.String("part", "a", "Which part of day06 do you want to calc (a or b)")
	flag.BoolVar(&debug, "debug", false, "Turn debug on")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - Number of orbits:", orbitCalculation(*filenamePtr, debug, 'a'))
	case "b":
		fmt.Println("Part b - Not implemented yet", orbitCalculation(*filenamePtr, debug, 'b'))

	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
