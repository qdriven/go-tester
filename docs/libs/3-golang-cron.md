# `Golang-cron-如何运行定时任务？`

有时需要执行定时任务，那么在Golang中有什么简单的办法实现吗? 最简单的办法是使用golang的cron相关的包.
所谓定时任务，就是到某个指定时间，系统运行已经给定的任务

## 安装

```shell
go get github.com/robfig/cron/v3@v3.0.0
```

## golang如何运行定时任务

关于定时任务，基本上可以理解为有三个不同的小工具合起来解决定时运行的问题:

- 调度器: 负责记录时间，同时到达时间之后出发需要运行的任务
- 时间表达式 cron-expression
- 任务: 具体需要执行的任务

理解了以上三点，那么cron定时任务用起来就比较方便了. 看一下代码:

```golang

func CronDemo() {
	i := 0
	c := cron.New()
	timeSpec := "*/5 * * * * ?"
	entryId, _ := c.AddFunc(timeSpec, func() {
		i++
		log.Println("this loop is ", i)
	})
	fmt.Println(entryId)
	c.Start()
	defer c.Stop()
	select {}
}
```

- timeSpec: 就是定义了一个Cron的表达式，用来表示什么时候会运行任务
- cron.New：实际上启动了一个定时任务调度器，用来管理什么以后运行什么样的任务
- 把timeSpec和对应需要运行的任务注册给调度器,这里面的匿名方法(**func()**)实际上就是任务
- 调度器知道了注册的时间和对应的任务，那么调度器中的计数开始，到指定的时间，调度器出发运行这册在这个时间的任务运行

```shell
c.AddFunc(timeSpec, func() {
		i++
		log.Println("this loop is ", i)
	}
```

综合以上，其实很好理解,就是把运行的时间和运行的任务注册给调度器，你就完成了一个最简单的定时任务的代码.

## Cron的表达式

Cron表达式规范还是不少，所以附上cron表达式的cheatsheet:

Field name   | Mandatory? | Allowed values  | Allowed special characters
----------   | ---------- | --------------  | --------------------------
Seconds      | Yes        | 0-59            | * / , -
Minutes      | Yes        | 0-59            | * / , -
Hours        | Yes        | 0-23            | * / , -
Day of month | Yes        | 1-31            | * / , - ?
Month        | Yes        | 1-12 or JAN-DEC | * / , -
Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?


1。星号(*)
表示 cron 表达式能匹配该字段的所有值。如在第5个字段使用星号(month)，表示每个月
2。 斜线(/)
表示增长间隔，如第1个字段(minutes) 值是 3-59/15，表示每小时的第3分钟开始执行一次，之后每隔 15 分钟执行一次（即 3、18、33、48 这些时间点执行），这里也可以表示为：3/15
3 逗号(,)用于枚举值，如第6个字段值是 MON,WED,FRI，表示 星期一、三、五 执行
4.连字号(-)表示一个范围，如第3个字段的值为 9-17 表示 9am 到 5pm 直接每个小时（包括9和17）
5.问号(?)


