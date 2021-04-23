// This program will calculate how many logs can fit onto a truck.
// Created By: Marlon Poddalgoda
// Created on: 2021-04-22

package main

import "fmt"

// constants
const carryCapacity, weightPerMeter float64 = 1100, 20
const shortLog, mediumLog, longLog float64 = 0.25, 0.5, 1

func main() {
	// this function takes in user input and calculates number of logs

	fmt.Println("This program calculates how many logs of a certain length can fit onto a truck")
	fmt.Println("Enter the length of the logs in meters.")
	fmt.Println("")

	// variables
	var logLength float64 = 0
	var numOfLogs float64 = 0
	var intOfLogs int = 0

	// for loop
	for {
	    // Ask for user input
	    fmt.Print("Valid choices are -> 0.25, 0.5, 1: ")

	    // Check if user inputted number
        if _, err := fmt.Scan(&logLength); err == nil {
            // verify if user picked from options
            if logLength == shortLog || logLength == mediumLog || logLength == longLog {
                // calculate number of logs
                numOfLogs = carryCapacity / (logLength * weightPerMeter)
                intOfLogs = int(numOfLogs)
                logLengthStr := fmt.Sprintf("%.2f", logLength)
                // output
                fmt.Println("")
                fmt.Printf("The truck can carry %d logs that are %s meter(s) long.", intOfLogs, logLengthStr)

                break
            } else {
                fmt.Println("That's not an option!")
                logLength = 0
            }
        } else {
            fmt.Println("That's not a number!")
            err = nil
            logLength = 0
        }
	}
}