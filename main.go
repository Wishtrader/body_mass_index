package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("_____ Body Mass Index _____")
	userHeight, userWeight := getUserInput()
	BMI := calculateBMI(userHeight, userWeight)
	outputResult(BMI)
}

func outputResult(bmi float64) {
	result := fmt.Sprintf("Your body mass index: %.0f", bmi)
	fmt.Print(result)
}

func calculateBMI(userHeight, userWeight float64) float64 {
	const IMTPower = 2
	BMI := userWeight / math.Pow(userHeight / 100, IMTPower)
	return BMI
}

func getUserInput() (float64, float64) {
	var userHeight, userWeight float64
	fmt.Print("Enter your height, cm: ")
	fmt.Scan(&userHeight)
	fmt.Print("Enter your weight, kg: ")
	fmt.Scan(&userWeight)
	return userHeight, userWeight
}