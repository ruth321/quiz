package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type questionAnswer struct {
	question string
	answer   string
}

func main() {
	k := true
	for k {
		fmt.Println("1. Start quiz")
		fmt.Println("2. Add questions")
		fmt.Println("3. Exit")
		fmt.Print("->")
		var a int
		fmt.Scan(&a)
		switch a {
		case 1:

		case 2:
			file, err := ioutil.ReadFile("guizFile.json")
			err = err
			//			if errors.Is(err, "2") {

			//			}
			var quiz []questionAnswer
			_ = json.Unmarshal(file, &quiz)
			var rep string
			fmt.Scan(&rep)
			if rep == "y" {

			}

		case 3:
			k = false

		}

	}

}
