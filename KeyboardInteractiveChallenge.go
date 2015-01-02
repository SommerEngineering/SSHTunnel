package main

import (
	"log"
)

func keyboardInteractiveChallenge(user, instruction string, questions []string, echos []bool) (answers []string, err error) {

	log.Println(`User: ` + user)
	log.Println(`Instruction: ` + instruction)
	log.Println(`Questions:`)
	for q := range questions {
		log.Println(q)
	}

	countQuestions := len(questions)
	answers = make([]string, countQuestions, countQuestions)

	if countQuestions > 0 {
		answers[0] = password
	}

	err = nil
	return
}
