package config

type ServiceConfig struct {
	TimeBetweenJobProcesses StrTimeDuration `toml:"time_between_job_processes" env:"QUEUE_MGR_TIME_BTW_JOB_PROCESSES" env-required`

	// Could also add:
	//maxWorkers // Default CPU cores
	//buffers // Default 1000. Depends on particular usage and memory
}
