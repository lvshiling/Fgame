package job

type JobManager struct {
	jobRunnerMap map[string]IJobRunner
}

func (m *JobManager) Start() {
	for _, value := range m.jobRunnerMap {
		if value.GetState() == RunnerStateRunning {
			continue
		}
		value.Start()
	}
}

func (m *JobManager) Stop() {
	for _, value := range m.jobRunnerMap {
		if value.GetState() != RunnerStateRunning {
			continue
		}
		value.Stop()
	}
}

func (m *JobManager) AddJob(job IJob) {
	_, exists := m.jobRunnerMap[job.GetId()]
	if exists {
		return
	}
	runner := NewJobRunner(job)
	m.jobRunnerMap[job.GetId()] = runner
}

func NewJobManager() *JobManager {
	rst := &JobManager{}
	rst.jobRunnerMap = make(map[string]IJobRunner)
	return rst
}
