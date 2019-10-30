package main

import "fmt"

func main() {
	questionsSlice := []string{
		"What is your last name?",
		"What is your street number?",
		"What is the name of your street?",
		"What is the name of your suburb?",
		"What is the name of your state?",
		"What is your postcode?",
	}
	var answersSlice []string
	var aSlice = make([]string, len(questionsSlice))

	concatAnswer := ""
	for i, question := range questionsSlice {
		var answer string
		fmt.Println(question)
		fmt.Scan(&answer)
		answersSlice = append(answersSlice, answer)
		aSlice[i] = answer
		concatAnswer += answer + ";"
	}

	fmt.Println("Answer ", concatAnswer)
}
