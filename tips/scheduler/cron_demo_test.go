package scheduler

import "testing"

func TestCronDemo(t *testing.T) {
	CronDemo()
}

func TestMultipleJobs(t *testing.T) {
	r := NewJobRunner("*/1 * * * * ?")
	r.RegisterFunc(Job1)
	r.RegisterFunc(Job2)
	r.RegisterFuncWithSpec(Job3,"*/5 * * * * ?")
	r.Run()
}