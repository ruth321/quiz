package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type questionAnswer struct {
	Question string
	Answer   string
}

func main() {
	var fileName string
	k := true
	var resp string
	fmt.Print("Enter file name: ")
	_, _ = fmt.Scan(&fileName)
	file, err := ioutil.ReadFile(fileName)
	var quiz []questionAnswer
	if err != nil {
		fmt.Println("File does not exist. Create file with this name?")
		fmt.Print("(y/n)->")
		_, _ = fmt.Scan(&resp)
		if resp == "y" {
			_, err = os.Create(fileName)
			if err != nil {
				fmt.Println(err)
			}
			file, err = ioutil.ReadFile(fileName)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("File created.")
			err = json.Unmarshal(file, &quiz)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			return
		}
	} else {
		err = json.Unmarshal(file, &quiz)
		if err != nil {
			fmt.Println(err)
		}
	}
	for k {
		fmt.Println("1. Start quiz")
		fmt.Println("2. Add questions")
		fmt.Println("3. Delete questions")
		fmt.Println("4. Exit")
		fmt.Println("Choose action")
		fmt.Print("->")
		var a int
		_, _ = fmt.Scan(&a)
		for a < 1 || a > 4 {
			fmt.Println("Wrong number")
			fmt.Print("->")
			_, _ = fmt.Scan(&a)
		}
		switch a {
		case 1:
			startQuiz(quiz)
		case 2:
			quiz = addQuestions(quiz)
		case 3:
			quiz = delQuestions(quiz)
		case 4:
			k = false
		}
	}
	file, err = json.Marshal(quiz)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(fileName, file, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func startQuiz(quiz []questionAnswer) {
	resp := "y"
	for resp == "y" {
		var answer string
		var count int
		for i := 0; i < len(quiz); i++ {
			fmt.Print("Question: ")
			fmt.Println(quiz[i].Question)
			fmt.Print("Answer: ")
			_, _ = fmt.Scan(&answer)
			if answer == quiz[i].Answer {
				count++
			}
		}
		fmt.Printf("Right answers: %d out of %d\n", count, len(quiz))
		fmt.Println("Restart?")
		fmt.Print("(y/n)->")
		_, _ = fmt.Scan(&resp)
	}
}

func addQuestions(quiz []questionAnswer) []questionAnswer {
	resp := "y"
	var q questionAnswer
	for resp == "y" {
		fmt.Print("Question: ")
		_, _ = fmt.Scan(&q.Question)
		fmt.Print("Answer: ")
		_, _ = fmt.Scan(&q.Answer)
		quiz = append(quiz, q)
		fmt.Println("Continue?")
		fmt.Print("(y/n)->")
		_, _ = fmt.Scan(&resp)
	}
	return quiz
}

func delQuestions(quiz []questionAnswer) []questionAnswer {
	if len(quiz) == 0 {
		fmt.Println("File is empty")
		return quiz
	}
	var n int
	resp := "y"
	for resp == "y" {
		fmt.Println("Questions:")
		for i := 0; i < len(quiz); i++ {
			fmt.Printf("%2.d. Question: %s\n    Answer: %s\n", i+1, quiz[i].Question, quiz[i].Answer)
		}
		fmt.Println("Choose number")
		fmt.Print("->")
		_, _ = fmt.Scan(&n)
		for n-1 < 0 || n > len(quiz) {
			fmt.Println("Wrong number")
			fmt.Print("->")
			_, _ = fmt.Scan(&n)
		}
		for i := n - 1; i < len(quiz)-1; i++ {
			quiz[i] = quiz[i+1]
		}
		quiz = quiz[:len(quiz)-1]
		fmt.Println("Question deleted. Continue?")
		fmt.Print("(y/n)->")
		_, _ = fmt.Scan(&resp)
	}
	return quiz
}
