package cron

import (
	"log"
	"time"
)

type Task struct {
	MonitorID int
	Schedule  Schedule
	Action    func(int) error
}

// Schedule These are the parameters of the time when the task starts.
//
// The IsDate parameter is a parameter which means that the task is executed relative to 00:00 real time. e.g.
//
//	cron.AddTask(cron.Task{
//			MonitorID: 1,
//			Schedule: cron.Schedule{
//				IsDate: true,
//				Day: 0,
//				Hours: 12,
//				Minutes: 0,
//				Seconds: 0,
//			},
//			Action: someFunc,
//		})
//
// IsDate = true this task will run every day at 12:00:00
//
// IsDate = false, then it will run every 12 hours
type Schedule struct {
	IsDate  bool
	Day     int
	Hours   int
	Minutes int
	Seconds int
}

type Cron struct {
	tasks []*Task
}

func NewCron() *Cron {
	return &Cron{}
}

func (c *Cron) AddTask(task Task) {
	c.tasks = append(c.tasks, &task)
}

func (c *Cron) Start() {
	for _, task := range c.tasks {
		go c.runTask(task)
	}
}

func (c *Cron) runTask(task *Task) {
	for {
		nextRun := task.Schedule.getTime()
		now := time.Now()
		duration := nextRun.Sub(now)

		if duration < 0 {
			nextRun = task.Schedule.getTime()
			duration = nextRun.Sub(now)
		}

		time.Sleep(duration)

		err := task.Action(task.MonitorID)
		if err != nil {
			log.Printf("Ping func is failed with error: %s", err)
		}

		nextRun = task.Schedule.getTime()
	}
}

func (s *Schedule) calcTime() time.Duration {
	return time.Hour*24*time.Duration(s.Day) + time.Hour*time.Duration(s.Hours) + time.Minute*time.Duration(s.Minutes) + time.Second*time.Duration(s.Seconds)
}

func (s *Schedule) getTime() time.Time {
	if s.IsDate {
		year, month, day := time.Now().AddDate(0, 0, 1+s.Day).Date()
		loc := time.Now().Location()
		date := time.Date(year, month, day, s.Hours, s.Minutes, s.Seconds, 0, loc)
		return date
	}
	return time.Now().Add(s.calcTime())
}
