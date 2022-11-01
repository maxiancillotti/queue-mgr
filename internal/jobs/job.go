package jobs

import (
	"errors"
	"strings"
)

type Job struct {
	Name   string    `json:"name"`
	ID     int       `json:"id"`
	Status jobStatus `json:"status"`
	Data   []int     `json:"data"`
	Result jobResult `json:"result"`
}

func (j *Job) SetResult(result string) {
	j.Status = statusProcessed
	j.Result = jobResult(result)
}

func (j *Job) ValidateInput() error {

	if j.Name == "" {
		return errors.New("name cannot be empty")
	}
	if j.Data == nil {
		return errors.New("data cannot be empty")
	}
	return nil
}

func (j *Job) ValidateStatusFilter() error {

	status := strings.ToLower(string(j.Status))

	if status != string(statusPending) && status != string(statusProcessed) {
		return errors.New("status is invalid")
	}
	return nil
}

func (j *Job) IsProcessed() bool {
	return j.Status == statusProcessed
}
