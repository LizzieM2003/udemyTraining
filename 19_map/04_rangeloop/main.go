package main

import "fmt"

func main() {
	mySkillSet := map[int]string{
		0: "HTML",
		1: "CSS",
		2: "JavaScript",
		3: "React",
		4: "Responsive Design",
		5: "Go",
		6: "SQL",
	}

	for key, val := range mySkillSet {
		fmt.Println(key, " - ", val)
	}
}
