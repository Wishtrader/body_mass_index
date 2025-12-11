package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	for {
		fmt.Println("_____ Body Mass Index _____")
		userHeight, userWeight := getUserInput()
		BMI, err := calculateBMI(userHeight, userWeight)
		if err != nil {
			fmt.Println(err)
			continue
		}
		outputResult(BMI)
		isRepeatCalculation := checkRepeatCalculation()
		if !isRepeatCalculation {
			break
		}
	}
}

func outputResult(bmi float64) {
	result := fmt.Sprintf("Your body mass index: %.0f", bmi)
	fmt.Println(result)
	switch {
		case bmi < 16:
			fmt.Println("You are severely underweight.")
		case bmi < 18.5:
			fmt.Println("You are underweight.")
		case bmi < 25:
			fmt.Println("You have a normal body weight.")
		case bmi < 30:
			fmt.Println("You are overweight.")
		default:
			fmt.Println("You have a degree of obesity.")
	}
}

func calculateBMI(userHeight, userWeight float64) (float64, error) {
	if userHeight <= 0 || userWeight <= 0 {
		return 0, errors.New("You enter wrong data!")
	}
	const IMTPower = 2
	BMI := userWeight / math.Pow(userHeight / 100, IMTPower)
	return BMI, nil
}

func getUserInput() (float64, float64) {
	var userHeight, userWeight float64
	fmt.Print("Enter your height, cm: ")
	fmt.Scan(&userHeight)
	fmt.Print("Enter your weight, kg: ")
	fmt.Scan(&userWeight)
	return userHeight, userWeight
}

func checkRepeatCalculation() bool {
	var userChoice string
	fmt.Print("Do you want to do calculation BMI again? (y/n): ")
	fmt.Scan(&userChoice)
	if userChoice == "y" || userChoice == "Y" {
		return true
	}
	return false
}