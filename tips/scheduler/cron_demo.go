package scheduler

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
)

/**
字段名	是否必须	允许的值	允许的特定字符
秒(Seconds)	是	0-59	* / , -
分(Minutes)	是	0-59	* / , -
时(Hours)	是	0-23	* / , -
日(Day of month)	是	1-31	* / , – ?
月(Month)	是	1-12 or JAN-DEC	* / , -
星期(Day of week)	否	0-6 or SUM-SAT	* / , – ?
1）星号(*)
表示 cron 表达式能匹配该字段的所有值。如在第5个字段使用星号(month)，表示每个月
2）斜线(/)
表示增长间隔，如第1个字段(minutes) 值是 3-59/15，表示每小时的第3分钟开始执行一次，之后每隔 15 分钟执行一次（即 3、18、33、48 这些时间点执行），这里也可以表示为：3/15
3）逗号(,)
用于枚举值，如第6个字段值是 MON,WED,FRI，表示 星期一、三、五 执行
4）连字号(-)
表示一个范围，如第3个字段的值为 9-17 表示 9am 到 5pm 直接每个小时（包括9和17）
5）问号(?)
只用于日(Day of month)和星期(Day of week)，\表示不指定值，可以用于代替 *
*/
func CronDemo() {
	i := 0
	c := cron.New()
	timeSpec := "*/5 * * * * ?"
	c.AddFunc(timeSpec, func() {
		i++
		log.Println("this loop is ", i)
	})
	c.Start()
	defer c.Stop()
	select {}
}

type JobRunner struct {
	DefaultSpec string
	Funcs       []func()
	C           *cron.Cron
}

func NewJobRunner(spec string) *JobRunner {
	runner := &JobRunner{
		DefaultSpec: spec,
		Funcs:       make([]func(), 0),
		C:           cron.New(),
	}
	return runner
}

func (job *JobRunner) RegisterFunc(fun func()) *JobRunner {
	job.Funcs = append(job.Funcs, fun)
	return job
}

func (job *JobRunner) RegisterFuncWithSpec(fun func(),Spec string) *JobRunner {
	job.Funcs = append(job.Funcs, fun)
	return job
}

func (job *JobRunner) Run() {
	for _, fun := range job.Funcs {
		_ = job.C.AddFunc(job.DefaultSpec, fun)
	}
	job.C.Start()
	select {} //event case is a channel operation
}

func Job1() {
	fmt.Println("this is job1")
}

func Job2() {
	fmt.Println("this is job2")
}

func Job3(){
	fmt.Println("this is job3 with Self Job Spec")
}