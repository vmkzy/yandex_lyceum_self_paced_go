package main

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"strings"
	"time"
)

type Ticket struct {
	Ticket string
	User   string
	Status string
	Date   time.Time
}

var validStatuses = map[string]bool{
	"Готово":           true,
	"В работе":         true,
	"Не будет сделано": true,
}

func GetTasks(
	ctx context.Context,
	r io.Reader,
	w io.Writer,
	user *string,
	status *string,
	timeout time.Duration,
) error {
	if ctx == nil {
		ctx = context.Background()
	}
	if timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}
	scanner := bufio.NewScanner(r)
	res := make([]Ticket, 0)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		line := strings.TrimSpace(scanner.Text())
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
		if _, ok := validStatuses[statusT]; !ok {
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
	if err := scanner.Err(); err != nil {
		return err
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		return err
	}
	return nil

}
