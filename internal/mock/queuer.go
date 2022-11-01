package mock

import "queue-mgr/internal/jobs"

type QueuerMock struct{}

func (q *QueuerMock) Enqueue(job *jobs.Job) {
}

func (q *QueuerMock) Dequeue() {
}

func (q *QueuerMock) RetrieveQueue() []*jobs.Job {

	job1 := &jobs.Job{
		Name: "JobName1",
		ID:   1,
		Data: []int{1, 2, 3},
	}
	job1.SetResult("Result OK. Sum = 6")

	job2 := &jobs.Job{
		Name: "JobName2",
		ID:   2,
		Data: []int{3, 2, 2},
	}
	job2.SetResult("Result OK. Sum = 7")

	job3 := &jobs.Job{
		Name: "JobName3",
		ID:   3,
		Data: []int{3, 2, 4},
	}
	job3.SetStatusPending()

	return []*jobs.Job{job1, job2, job3}
}

func (q *QueuerMock) RetrievePendingQueue() []*jobs.Job {

	job3 := &jobs.Job{
		Name: "JobName3",
		ID:   3,
		Data: []int{3, 2, 4},
	}
	job3.SetStatusPending()

	job4 := &jobs.Job{
		Name: "JobName4",
		ID:   4,
		Data: []int{3, 2, 5},
	}
	job4.SetStatusPending()

	return []*jobs.Job{job3, job4}
}

func (q *QueuerMock) RetrieveProcessedQueue() []*jobs.Job {

	job1 := &jobs.Job{
		Name: "JobName1",
		ID:   1,
		Data: []int{1, 2, 3},
	}
	job1.SetResult("Result OK. Sum = 6")

	job2 := &jobs.Job{
		Name: "JobName2",
		ID:   2,
		Data: []int{3, 2, 2},
	}
	job2.SetResult("Result OK. Sum = 7")

	return []*jobs.Job{job1, job2}
}
