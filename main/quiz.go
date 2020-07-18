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
		fmt.Println("3. Exit")
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
			k = false

		}

	}

}
