package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

// Problem struct to store question and its corresponding answer
type problem struct {
	q string
	a string
}

func problemPuller(fileName string) ([]problem, error) {

	// Open the CSV quize file
	if fileObj, err := os.Open(fileName); err == nil {

		// Read all the problems from the CSV file
		csvReader := csv.NewReader(fileObj)
		if cLines, err := csvReader.ReadAll(); err == nil {
			return problemParser(cLines), nil
		} else {
			return nil, fmt.Errorf("Error in reading data in csv"+"format from %s file; %s", fileName, err.Error())
		}
	} else {
		return nil, fmt.Errorf("Error in opening the %s file; %s", fileName, err.Error())
	}
}

// To convert each CSV line into a problem struct
func problemParser(lines [][]string) []problem {
	r := make([]problem, len(lines))
	for i := 0; i < len(lines); i++ {
		if len(lines[i]) < 2 {
			continue
		}
		r[i] = problem{q: lines[i][0], a: lines[i][1]}
	}
	return r
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func main() {
	fName := flag.String("f", "quiz.csv", "path of csv file")

	// Parse the CSV file into problems
	prob, err := problemPuller(*fName)

	// Error handling
	if err != nil {
		exit(fmt.Sprintf("Something went wrong: %s", err.Error()))
	}

	// Calculate and set the duration of timer (5 seconds per question)
	timer := flag.Int("t", 5*len(prob), "timer for the quiz")
	flag.Parse()

	fmt.Printf("You have %d seconds to answer %d questions\n", *timer, len(prob))

	// Variable to count the correct answers
	correctAns := 0

	// Initialize the timer
	timeObj := time.NewTimer((time.Duration(*timer) * time.Second))

	// Channel to receive user answers
	ansC := make(chan string)

problemLoop:
	for i, p := range prob {
		var answer string
		fmt.Printf("Problem %d: %s=", i+1, p.q)

		// Run a goroutine to capture user input
		go func() {
			fmt.Scanln(&answer)
			ansC <- answer
		}()

		// Wait for either input or timer expiration
		select {
		case <-timeObj.C:
			fmt.Println()
			break problemLoop
		case iAns := <-ansC:
			if iAns == p.a {
				correctAns++
			}
		}
	}

	// Print the final result
	fmt.Printf("Your result is %d out of %d\n", correctAns, len(prob))
	fmt.Printf("Press enter to exit")
	fmt.Scanln()
}
