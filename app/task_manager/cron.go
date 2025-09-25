package task_manager

import "time"

func AddCronTask(id, name, spec string) {
	TaskManager.lock.Lock()
	defer TaskManager.lock.Unlock()
	// 从持久化数据中加载更新时间
	var ut int64
	pt := GetPersist().GetBackgroundJob(name)
	if pt > 0 {
		ut = pt
	} else {
		ut = time.Now().Unix()
	}
	if _, ok := TaskManager.CronJobs[id]; !ok {
		TaskManager.CronJobs[id] = task{
			TaskID:     id,
			TaskName:   name,
			Spec:       spec,
			CreateTime: time.Now().Unix(),
			UpdateTime: ut,
		}
	}
}
