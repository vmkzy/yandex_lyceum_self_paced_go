package main

import (
	"strings"
	"time"
)

func QuizRunner(questions, answers []string, answerCh chan string) int {
	ans := 0
	for i := 0; i < len(questions); i++ {
		timer := time.NewTimer(time.Second)
		select {
		case userAns := <-answerCh:
			if strings.EqualFold(strings.TrimSpace(userAns), strings.TrimSpace(answers[i])) {
				ans++
			}
			if !timer.Stop() {
				<-timer.C
			}
		case <-timer.C:
		}
	}
	return ans
}
