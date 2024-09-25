package jobmapper

import (
	"encoding/json"
	"errors"

	"github.com/interline-io/transitland-jobs/jobs"
)

///////////

type JobMapper struct {
	jobFns map[string]jobs.JobFn
}

func NewJobMapper() *JobMapper {
	return &JobMapper{jobFns: map[string]jobs.JobFn{}}
}

func (j *JobMapper) AddJobType(jobFn jobs.JobFn) error {
	jw := jobFn()
	j.jobFns[jw.Kind()] = jobFn
	return nil
}

func (j *JobMapper) GetRunner(jobType string, jobArgs jobs.JobArgs) (jobs.JobWorker, error) {
	jobFn, ok := j.jobFns[jobType]
	if !ok {
		return nil, errors.New("unknown job type")
	}
	runner := jobFn()
	jw, err := json.Marshal(jobArgs)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(jw, runner); err != nil {
		return nil, err
	}
	return runner, nil
}
