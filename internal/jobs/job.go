package jobs

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
