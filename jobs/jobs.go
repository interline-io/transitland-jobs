package jobs

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
)

type JobArgs map[string]any

// Job queue
type JobQueue interface {
	Use(JobMiddleware)
	AddQueue(string, int) error
	AddJobType(JobFn) error
	AddJob(context.Context, Job) error
	AddJobs(context.Context, []Job) error
	RunJob(context.Context, Job) error
	Run(context.Context) error
	Stop(context.Context) error
}

// Job defines a single job
type Job struct {
	Queue       string  `json:"queue"`
	JobType     string  `json:"job_type" river:"unique"`
	JobArgs     JobArgs `json:"job_args" river:"unique"`
	Unique      bool    `json:"unique"`
	JobDeadline int64   `json:"job_deadline"`
	jobId       string  `json:"-"`
}

func (job *Job) HexKey() (string, error) {
	bytes, err := json.Marshal(job.JobArgs)
	if err != nil {
		return "", err
	}
	sum := sha1.Sum(bytes)
	return job.JobType + ":" + hex.EncodeToString(sum[:]), nil
}

// JobWorker defines a job worker
type JobWorker interface {
	Kind() string
	Run(context.Context, Job) error
}

type JobFn func() JobWorker

type JobMiddleware func(JobWorker) JobWorker