package main

import (
	"strings"
	"time"
)

type Ticket struct {
	Ticket string
	User   string
	Status string
	Date   time.Time
}

var Rstatus = map[string]bool{
	"Готово":           true,
	"В работе":         true,
	"Не будет сделано": true,
}

func GetTasks(text string, user *string, status *string) []Ticket {
	lines := strings.Split(text, "\n")
	var res []Ticket
	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}
		if !strings.HasPrefix(line, "TICKET-") {
			continue
		}
		parts := strings.SplitN(line, "_", 4)

		if len(parts) != 4 {
			continue
		}
		numT := parts[0]
		userT := parts[1]
		statusT := parts[2]
		dateT, err := time.Parse("2006-01-02", parts[3])

		if err != nil {
			continue
		}
		if userT == "" {
			continue
		}
		if _, ok := Rstatus[statusT]; !ok {
			continue
		}
		if user != nil && *user != userT {
			continue
		}
		if status != nil && *status != statusT {
			continue
		}

		res = append(res, Ticket{
			Ticket: numT,
			User:   userT,
			Status: statusT,
			Date:   dateT,
		})
	}
	return res
}

/*
func main() {
	requests := "TICKET-12345_Паша Попов_Готово_2024-01-01\nTICKET-12346_Иван Иванов_В работе_2024-01-02\nTICKET-12347_Анна Смирнова_Не будет сделано_2024-01-03\nTICKET-12348_Паша Попов_В работе_2024-01-04\n"
	user := "Паша Попов"
	result := GetTasks(requests, &user, nil)
	for _, val := range result {
		fmt.Printf("%s-%s_%s\n", val.Ticket, val.User, val.Status)
	}

}
*/
