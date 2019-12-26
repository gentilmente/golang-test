package main

import "fmt"

func main() {
	grades := make(map[string]float32)

	grades["Sebas"] = 42
	grades["Juan"] = 93
	grades["Juana"] = 33

	fmt.Println(grades)

	SebasGrade := grades["Sebas"]
	fmt.Println(SebasGrade)

	delete(grades, "Sebas")

	for k, v := range grades {
		fmt.Println(k, ": ", v)
	}
}
