package main

import (
	"fmt"
)

func calculateAverage(grades []float64) float64{
	total := 0.0
	for _,gread := range grades {
		total += gread
	}
	return total / float64(len(grades))
}

func main() {
	var studentName string
	var numSubjects int

	fmt.Println("Enter the student name: ")
	fmt.Scanln(&studentName)


	fmt.Println("Enter the number of subjects: ")
	fmt.Scanln(&numSubjects)

	if numSubjects <= 0 {
		fmt.Println("Invalid number of subjects")
		return
	}

	var subjects = make([]string, numSubjects)
	var grades = make([]float64, numSubjects)
    
	for i := 0; i <  numSubjects; i++ {
		fmt.Println("Enter the subject name: ")
		fmt.Scanln(&subjects[i])
		var grade float64

		for{
			fmt.Printf("Enter the grade for  %s (0-100): " , subjects[i])
			fmt.Scanln(&grade)

			if grade >= 0 && grade <= 100{
				grades[i] = grade
				break
			}else{
				fmt.Println("Invalide grade. Please enter a grade between 0 and 100.")

			}

		}

	}

	averageGrade := calculateAverage(grades)

	fmt.Printf("\nStudent Name: %s\n" , studentName)
	fmt.Println("Subjects and Grades:")
	for i ,subject := range subjects{
		fmt.Printf("%s: %.2f\n" , subject , grades[i])
	}
	fmt.Printf("Average Grade: %2f\n" , averageGrade)


}