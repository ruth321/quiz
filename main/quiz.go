package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type questionAnswer struct {
	Question string
	Answer   string
}

func main() {
	k := true
	for k {
		fmt.Println("1. Start quiz")
		fmt.Println("2. Add questions")
		fmt.Println("3. Delete questions")
		fmt.Println("4. Exit")
		fmt.Print("->")
		var a int
		_, _ = fmt.Scan(&a)
		switch a {
		case 1:
			file, err := ioutil.ReadFile("quizFile.json")
			if err != nil {
				log.Fatal(err)
			}
			var quiz []questionAnswer
			err = json.Unmarshal(file, &quiz)
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
		case 2:
			file, err := ioutil.ReadFile("quizFile.json")
			if err != nil {
				log.Fatal(err)
			}
			resp := "y"
			var quiz []questionAnswer
			err = json.Unmarshal(file, &quiz)
			if err != nil {
				fmt.Println(err)
			}
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
			file, err = json.Marshal(quiz)
			if err != nil {
				fmt.Println(err)
			}
			err = ioutil.WriteFile("quizFile.json", file, 0644)
			if err != nil {
				fmt.Println(err)
			}
		case 3:
			file, err := ioutil.ReadFile("quizFile.json")
			if err != nil {
				log.Fatal(err)
			}
			var quiz []questionAnswer
			err = json.Unmarshal(file, &quiz)
			resp := "y"
			var n int
			for resp == "y" {
				fmt.Println("Choose question:")
				for i := 0; i < len(quiz); i++ {
					fmt.Printf("%2.d. Question: %s\n    Answer: %s\n", i+1, quiz[i].Question, quiz[i].Answer)
				}
				fmt.Print("->")
				_, _ = fmt.Scan(&n)
				for i := n - 1; i < len(quiz)-1; i++ {
					quiz[i] = quiz[i+1]
				}
				quiz = quiz[:len(quiz)-1]
				fmt.Println("Continue?")
				fmt.Print("(y/n)->")
				_, _ = fmt.Scan(&resp)
			}
			file, err = json.Marshal(quiz)
			if err != nil {
				fmt.Println(err)
			}
			err = ioutil.WriteFile("quizFile.json", file, 0644)
			if err != nil {
				fmt.Println(err)
			}
		case 4:
			k = false
		}
	}
}
