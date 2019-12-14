package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func processSequence(phase string) []int {
	phaseSequence := make([]int, len(phase))

	for i := 0; i < len(phase); i++ {
		phaseSequence[i], _ = strconv.Atoi(string(phase[i]))
	}
	return phaseSequence
}

// Day05 extensions: Parameter Modes
//    0: parameter is interpreteted as a position. if the parameter is 50, its value is the value stored at address 50 in memory.
//		   Until now, ALL parameters have worked in POSITION mode
//    1: immediate mode. Parmeter is taken to be a VALUE. If the parameter is 50, its value is simply 50.
//
// Day05 opcode extension into 5 digits
//    ABCDE: DE - 2 digit opcode (e.g. 01, 02, 03, 04)
//			  C - mode of 1st parameter (if not present then 0)
//			  B - mode of 2nd parameter
//			  A - mode of 3rd parameter
//
//    1: adds together 2 numbers, stores in 3rd. Next 3 numbers are: number A, number B and where to store
//    2: multiplys together 2 numbers, stores in 3rd. Next 3 numbers are: number A, number B and where to store
//    3: takes a single integer as input and saves it to the position given by its only parameter. For example, the instruction 3,50 would take an input value and store it at address 50.
//    4: outputs the value of its only parameter. For example, the instruction 4,50 would output the value at address 50.
//    5: jump-if-true: if first parameter is non-zero, it sets the instruction pointer to the value from the second parameter. Otherwise do nothing
//    6: jump-if-false: if first parameter is zero, it sets instruction pointer to the value from second parameter. Otherwise, do nothing.
//    7: is less than: if first parameter is less than second parameter, it stores 1 in the position given by third parameter. Otherwise, stores 0.
//    8: is equals: if first parameter is equal to second parameter, it stores 1 in the position given by third parameter. Otherwise, stores 0.
//   99: exit
//  any: anything else means things have gone wrong

func intcodeComputer(programArray []int, inputInstruction []int, debug bool, part byte) int {
	var currPos int
	var opcode, firstParamMode, secondParamMode, thirdParamMode int
	var firstValue, secondValue, thirdValue int
	var diagnosticCode int

	if debug {
		fmt.Println(programArray)
	}

	// While something to do
	// Read op code at current position
	// If 99, exit and output our result
	// If 1, use next 3 numbers as positions and work on them
	// If 2, use next 3 numbers as positions and work on them
	// If anything else, output an error and quit

	for {
		// Extra digits from the opcode. ABCDE: DE - 2 digit opcode (e.g. 01), C, D and E are the mode of the 1st, 2nd, 3rd parameters
		opcode = (programArray[currPos]/10%10)*10 + programArray[currPos]%10
		firstParamMode = programArray[currPos] / 100 % 10
		secondParamMode = programArray[currPos] / 1000 % 10
		thirdParamMode = programArray[currPos] / 10000 % 10

		if debug {
			fmt.Printf("[START] opcode: %2d first: %d second: %d third: %d\n", opcode, firstParamMode, secondParamMode, thirdParamMode)
		}

		// TO DO: Move the position vs immediate stuff out here to make the opcode tasks clearer
		//        e.g. for 1, 2, 7, 8 set all 3 *Value settings
		//             for 5, 6 set 2 *Value settings
		//			   for 3, 4 set 1 *Value settings

		switch opcode {
		case 99: // Exit
			return diagnosticCode

		case 1: // Addition
			if firstParamMode == 0 {
				firstValue = programArray[currPos+1]
			} else {
				firstValue = currPos + 1
			}

			if secondParamMode == 0 {
				secondValue = programArray[currPos+2]
			} else {
				secondValue = currPos + 2
			}

			// Write positions are never immediate mode
			thirdValue = programArray[currPos+3]

			if debug {
				fmt.Println("[OP:01] programArray is:", programArray)
				fmt.Printf("[OP:01] opcode %d: adding %d to %d and storing in position %d\n",
					opcode,
					programArray[firstValue],
					programArray[secondValue],
					programArray[thirdValue])
			}

			programArray[thirdValue] = programArray[firstValue] + programArray[secondValue]
			currPos += 4

			if debug {
				fmt.Println("[OP:01] After addition:", programArray)
			}

		case 2: // Multiply
			if firstParamMode == 0 {
				firstValue = programArray[currPos+1]
			} else {
				firstValue = currPos + 1
			}

			if secondParamMode == 0 {
				secondValue = programArray[currPos+2]
			} else {
				secondValue = currPos + 2
			}

			// Write positions are never immediate mode
			thirdValue = programArray[currPos+3]

			if debug {
				fmt.Println("[OP:02] programArray is:", programArray)
				fmt.Printf("[OP:02] opcode %d: multiplying %d to %d and storing in position %d\n",
					opcode,
					programArray[firstValue],
					programArray[secondValue],
					programArray[thirdValue])
			}

			programArray[thirdValue] = programArray[firstValue] * programArray[secondValue]
			currPos += 4

			if debug {
				fmt.Println("[OP:02] After multiply:", programArray)
			}

		case 3:
			// Opcode 3 takes a single integer as input and saves it to the position given by its only parameter. For example,
			// the instruction 3,50 would take an input value and store it at address 50.

			// [TO DO] Should take an inputInstruction and store in the postition given. NOT WHAT I'M CURRENTLY DOING
			//         Could we have an array that we're moving along as the program runs?
			//		   e.g. we use inputArray[0] then set inputArray = inputArray[1:] ?

			// Write to is never in immediate mode
			if debug {
				fmt.Println("[OP:03] programArray is:", programArray)
				fmt.Printf("[OP:03] opcode %d: Takes input %d and stores in position %d\n", opcode, inputInstruction, programArray[firstValue])
			}
			programArray[programArray[currPos+1]] = inputInstruction[0]
			fmt.Println("[OP:03] inputInstruction is:", inputInstruction[0])
			inputInstruction = inputInstruction[1:]
			fmt.Println("[OP:03] inputInstruction is:", inputInstruction[0])

			if debug {
				fmt.Println("[OP:03] After input stored:", programArray)
			}
			currPos += 2

		case 4:
			// 4: outputs the value of its only parameter. For example, the instruction 4,50 would output the value at address 50.
			if debug {
				fmt.Println("[OP:04] programArray is:", programArray)
			}

			if firstParamMode == 0 {
				fmt.Println("[OP:04] Output:", programArray[programArray[currPos+1]])
				diagnosticCode = programArray[programArray[currPos+1]]
			} else {
				fmt.Println("[OP:04] Output:", programArray[currPos+1])
				diagnosticCode = programArray[currPos+1]
			}
			if debug {
				fmt.Println("[OP:04] After output:", programArray)
			}
			currPos += 2

		case 5:
			// 5: jump-if-true: if first parameter is non-zero, it sets the instruction pointer to the value from the second parameter. Otherwise do nothing
			if debug {
				fmt.Println("[OP:05] Before jump if true:", programArray)
				fmt.Println("[OP:05] currPos before jump:", currPos)
			}
			if firstParamMode == 0 {
				firstValue = programArray[currPos+1]
			} else {
				firstValue = currPos + 1
			}
			if programArray[firstValue] != 0 {
				// Set instruction pointer to the value from the second parameter
				if secondParamMode == 0 {
					currPos = programArray[programArray[currPos+2]]
					if debug {
						fmt.Println("[OP:05] currPos after jump:", currPos)
					}
				} else {
					currPos = programArray[currPos+2]
					if debug {
						fmt.Println("[OP:05] currPos after jump:", currPos)
					}
				}
			} else { // Do nothing
				currPos += 3
				if debug {
					fmt.Println("[OP:05] currPos after DO NOTHING:", currPos)
				}
			}

		case 6:
			// 6: jump-if-false: if first parameter is zero, it sets instruction pointer to the value from second parameter. Otherwise, do nothing.
			if debug {
				fmt.Println("[OP:06] Before jump if true:", programArray)
				fmt.Println("[OP:06] currPos before jump:", currPos)
			}
			if firstParamMode == 0 {
				firstValue = programArray[currPos+1]
			} else {
				firstValue = currPos + 1
			}

			if programArray[firstValue] == 0 {
				// Set instruction pointer to the value from the second parameter
				if secondParamMode == 0 {
					currPos = programArray[programArray[currPos+2]]
					if debug {
						fmt.Println("[OP:06] currPos after jump:", currPos)
					}
				} else {
					currPos = programArray[currPos+2]
					if debug {
						fmt.Println("[OP:06] currPos after jump:", currPos)
					}
				}
			} else { // Do nothing
				currPos += 3
				if debug {
					fmt.Println("[OP:06] currPos after DO NOTHING:", currPos)
				}
			}

		case 7:
			// 7: is less than: if first parameter is less than second parameter, it stores 1 in the position given by third parameter. Otherwise, stores 0.
			if firstParamMode == 0 {
				firstValue = programArray[currPos+1]
			} else {
				firstValue = currPos + 1
			}
			if secondParamMode == 0 {
				secondValue = programArray[currPos+2]
			} else {
				secondValue = currPos + 2
			}

			// Write destination is never in immediate mode
			thirdValue = programArray[currPos+3]

			if debug {
				fmt.Println("[OP:07] programArray is:", programArray)
				fmt.Printf("[OP:07] opcode %d: if %d is less than %d then store 1 in position %d\n",
					opcode,
					programArray[firstValue],
					programArray[secondValue],
					programArray[thirdValue])
			}

			if programArray[firstValue] < programArray[secondValue] {
				programArray[thirdValue] = 1
			} else {
				programArray[thirdValue] = 0
			}
			currPos += 4
			if debug {
				fmt.Println("[OP:07] After less than:", programArray)
			}

		case 8:
			// 8: is equals: if first parameter is equal to second parameter, it stores 1 in the position given by third parameter. Otherwise, stores 0.
			if firstParamMode == 0 {
				firstValue = programArray[currPos+1]
			} else {
				firstValue = currPos + 1
			}
			if secondParamMode == 0 {
				secondValue = programArray[currPos+2]
			} else {
				secondValue = currPos + 2
			}

			// Write destination is never in immediate mode
			thirdValue = programArray[currPos+3]

			if debug {
				fmt.Println("[OP:08] programArray is:", programArray)
				fmt.Printf("[OP:08] opcode %d: if %d is equal to %d then store 1 in position %d\n",
					opcode,
					programArray[firstValue],
					programArray[secondValue],
					programArray[thirdValue])
			}

			if programArray[firstValue] == programArray[secondValue] {
				programArray[thirdValue] = 1
			} else {
				programArray[thirdValue] = 0
			}
			currPos += 4
			if debug {
				fmt.Println("[OP:08] After equals:", programArray)
			}

		default: // This shouldn't happen
			fmt.Printf("Code not implemented yet for instruction %d\n", programArray[currPos])
			return 0
		}
	}
}

// day07 Amplification Circuit
// 5 x Amplifiers which need to run the program
// 1st amp output -> 2nd amp input and so on
// Input provided: a current phase setting for EACH amplifier. Number from 0 to 4 which can be used only once
// Find the largest output signal that can be sent to the thrusters by trying every possible combination of phase settings on the amplifiers
// - Start code for Amp A. Provide phase setting for A for input1. For input2, provide 0 (start value)
// - Start code for Amp B. Provide phase setting for B for input1. For input2, provide output from Amp A
// - .....repeat.....
// - Result is the output from Amp E

func intcodeMaxThrusterSignal(filename string, input int, phase string, debug bool, part byte) int {
	var outputSignal int

	csvFile, _ := os.Open(filename)
	reader := csv.NewReader(bufio.NewReader(csvFile))

	// Only reading 1 line from the file and ignoring errors like a bad person
	lineRead, _ := reader.Read()
	// added explicit close in case we need to call the routine many times
	csvFile.Close()

	phaseSequence := processSequence(phase)

	// Create an array the same size as the records we've read from the file, then assign corresponding entries to the array
	programArray := make([]int, len(lineRead))
	for i := 0; i < len(lineRead); i++ {
		programArray[i], _ = strconv.Atoi(lineRead[i])
	}

	inputInstruction := make([]int, 5)

	// Now process the phaseSequence and modify the programArray to use the appropriate phaseSequence and inputInstruction
	// THIS ISN'T RIGHT. RE-READ THE TASK AND CHANGE WHAT I'M DOING HERE
	inputInstruction[0] = phaseSequence[0] // Change 4 to being the phase sequence for the first amp
	inputInstruction[1] = 0

	// Do I pass the inputInstruction and phaseSequence in, or just modify the programArray and send that?
	fmt.Println(programArray)
	fmt.Println(inputInstruction)
	outputSignal = intcodeComputer(programArray, inputInstruction, debug, part)
	fmt.Println(programArray)
	fmt.Println(inputInstruction)

	fmt.Println("Output signal from Amp1 is:", outputSignal)

	//return 0 // REMOVE THIS

	// Second AMP
	programArray2 := make([]int, len(lineRead))
	for i := 0; i < len(lineRead); i++ {
		programArray2[i], _ = strconv.Atoi(lineRead[i])
	}

	inputInstruction[0] = phaseSequence[1] // Change 4 to being the phase sequence for the first amp
	inputInstruction[1] = outputSignal

	fmt.Println(programArray2)
	outputSignal = intcodeComputer(programArray2, inputInstruction, debug, part)
	fmt.Println(programArray2)

	fmt.Println("Output signal from Amp2 is:", outputSignal)

	// Third AMP

	programArray3 := make([]int, len(lineRead))
	for i := 0; i < len(lineRead); i++ {
		programArray3[i], _ = strconv.Atoi(lineRead[i])
	}

	inputInstruction[0] = phaseSequence[2] // Change 4 to being the phase sequence for the first amp
	inputInstruction[1] = outputSignal

	outputSignal = intcodeComputer(programArray3, inputInstruction, debug, part)
	fmt.Println("Output signal from Amp3 is:", outputSignal)

	// Fourth AMP

	programArray4 := make([]int, len(lineRead))
	for i := 0; i < len(lineRead); i++ {
		programArray4[i], _ = strconv.Atoi(lineRead[i])
	}

	inputInstruction[0] = phaseSequence[3] // Change 4 to being the phase sequence for the first amp
	inputInstruction[1] = outputSignal

	outputSignal = intcodeComputer(programArray4, inputInstruction, debug, part)
	fmt.Println("Output signal from Amp4 is:", outputSignal)

	// Fifth AMP

	programArray5 := make([]int, len(lineRead))
	for i := 0; i < len(lineRead); i++ {
		programArray5[i], _ = strconv.Atoi(lineRead[i])
	}

	inputInstruction[0] = phaseSequence[4] // Change 4 to being the phase sequence for the first amp
	inputInstruction[1] = outputSignal

	outputSignal = intcodeComputer(programArray5, inputInstruction, debug, part)
	fmt.Println("Output signal from Amp5 is:", outputSignal)
	return 0
}

// Main routine
func main() {
	var debug bool

	filenamePtr := flag.String("file", "input.txt", "Filename containing the program to run")
	execPartPtr := flag.String("part", "a", "Which part of day18 do you want to calc (a or b)")
	phasePtr := flag.String("phase", "43210", "Phase setting sequence")
	inputPtr := flag.Int("input", 1, "Input instruction for Amp 1 computer")
	flag.BoolVar(&debug, "debug", false, "Turn debug on")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		fmt.Println("Part a - Max thruster signal:", intcodeMaxThrusterSignal(*filenamePtr, *inputPtr, *phasePtr, debug, 'a'))
	case "b":
		fmt.Println("Part b - Not implemented yet")
	default:
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}
