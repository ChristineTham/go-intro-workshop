package main

import "fmt"

type QuestionAndAnswer struct {
	question string
	answer   string
}

func makeaqa(q string) QuestionAndAnswer {
	return QuestionAndAnswer{
		question: q,
		answer:   "",
	}
}

func main() {
	qaArray := []QuestionAndAnswer{
		makeaqa("What is your last name?"),
		QuestionAndAnswer{question: "What is your street number?"},
		{"What is the name of your street?", ""},
		{"What is the name of your suburb?", ""},
		{"What is the name of your city", ""},
		{"What is your postcode?", ""},
	}

	answer := ""
	for _, qa := range qaArray {
		fmt.Println(qa.question)
		fmt.Scan(&qa.answer)
		answer += qa.answer + ";"
	}

	fmt.Println("Answer ", answer)
}
