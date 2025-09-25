package task_manager

var persistData PersistManager

func GetPersist() *PersistManager {
	return &persistData
}

func (p *PersistManager) GetCronJob(name string) int64 {
	for _, t := range p.CronJobs {
		if t.Name == name {
			return t.UpdateTime
		}
	}
	return 0
}

func (p *PersistManager) SetCronJob(name string, ut int64) {
	for _, t := range p.CronJobs {
		if t.Name == name {
			t.UpdateTime = ut
		}
	}
}

func (p *PersistManager) GetBackgroundJob(name string) int64 {
	for _, t := range p.BackGroundJobs {
		if t.Name == name {
			return t.UpdateTime
		}
	}
	return 0
}

func (p *PersistManager) SetBackgroundJob(name string, ut int64) {
	for _, t := range p.BackGroundJobs {
		if t.Name == name {
			t.UpdateTime = ut
		}
	}
}
