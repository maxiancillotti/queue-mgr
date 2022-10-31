package jobs

type jobStatus string

const (
	statusPending   jobStatus = "pending"
	statusProcessed jobStatus = "processed"
)
