package cron

import (
	"sync"
	"time"
)

const ANY = -1 // mod by MDR

type cronJob struct {
	year, month, day, weekday int
	hour, minute, second      int
	taskId                    string
	task                      func()
}

func NewCronJobBuilder() *cronJob {
	return &cronJob{}
}

// This function creates a new job that occurs at the given day and the given
// 24hour time. Any of the values may be -1 as an "any" match, so passing in
// a day of -1, the event occurs every day; passing in a second value of -1, the
// event will fire every second that the other parameters match.
func (c *cronJob) At(year, month, day, weekday, hour, minute, second int) *cronJob {
	c.year, c.month, c.day, c.weekday, c.hour, c.minute, c.second = year, month, day, weekday, hour, minute, second

	return c
}

func (c *cronJob) TimeAt(t time.Time) *cronJob {
	return c.At(t.Year(), int(t.Month()), t.Day(), int(t.Weekday()), t.Hour(), t.Minute(), t.Second())
}

// This creates a job that fires monthly at a given time on a given day.
func (c *cronJob) MonthlyAt(day, hour, minute, second int) *cronJob {
	return c.At(ANY, ANY, day, ANY, hour, minute, second)
}

// This creates a job that fires on the given day of the week and time.
func (c *cronJob) WeeklyAt(weekday, hour, minute, second int) *cronJob {
	return c.At(ANY, ANY, ANY, weekday, hour, minute, second)
}

// This creates a job that fires daily at a specified time.
func (c *cronJob) DailyAt(hour, minute, second int) *cronJob {
	return c.At(ANY, ANY, ANY, ANY, hour, minute, second)
}

// Add job on a cronJob
func (c *cronJob) DoJob(id string, task func()) *cronJob {
	c.task, c.taskId = task, id

	return c
}

// Add a cronJob into jobs that type is cronJobs
func (c *cronJob) Build() {
	jobs.addCronJob(c.taskId, c)
}

func (c *cronJob) matches(t time.Time) bool {
	return (c.year == ANY || c.year == t.Year()) &&
		(c.month == ANY || c.month == int(t.Month())) &&
		(c.day == ANY || c.day == t.Day()) &&
		(c.weekday == ANY || c.weekday == int(t.Weekday())) &&
		(c.hour == ANY || c.hour == t.Hour()) &&
		(c.minute == ANY || c.minute == t.Minute()) &&
		(c.second == ANY || c.second == t.Second())
}

type cronJobs struct {
	jobs   map[string]*cronJob
	rwLock sync.RWMutex
}

var jobs = &cronJobs{jobs: map[string]*cronJob{}, rwLock: sync.RWMutex{}}

func (js *cronJobs) addCronJob(id string, c *cronJob) *cronJobs {
	js.rwLock.Lock()
	defer js.rwLock.Unlock()
	js.jobs[id] = c

	return js
}

func (js *cronJobs) removeCronJob(id string) *cronJobs {
	js.rwLock.Lock()
	defer js.rwLock.Unlock()
	delete(js.jobs, id)

	return js
}

func (js *cronJobs) processJobs() {
	js.rwLock.RLock()
	defer js.rwLock.RUnlock()

	now := time.Now()
	for _, j := range js.jobs {
		if j.matches(now) {
			// execute all our cron tasks asynchronously
			go j.task()
		}
	}
}

func RemoveJob(id string) {
	jobs.removeCronJob(id)
}

func Forever(f func(), period time.Duration) {
	for {
		f()
		time.Sleep(period)
	}
}

func init() {
	go Forever(jobs.processJobs, time.Second)
}
