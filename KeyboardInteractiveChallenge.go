package main

import (
	"log"
)

// Another auth. method.
func keyboardInteractiveChallenge(user, instruction string, questions []string, echos []bool) (answers []string, err error) {

	// Log all the provided data:
	log.Println(`User: ` + user)
	log.Println(`Instruction: ` + instruction)
	log.Println(`Questions:`)
	for q := range questions {
		log.Println(q)
	}

	// How many questions are asked?
	countQuestions := len(questions)

	if countQuestions == 1 {

		// We expect that in this case (only one question is asked), that the server want to know the password ;-)
		answers = make([]string, countQuestions, countQuestions)
		answers[0] = password

	} else if countQuestions > 1 {

		// After logging, this call will exit the whole program:
		log.Fatalln(`The SSH server is asking multiple questions! This program cannot handle this case.`)
	}

	err = nil
	return
}
