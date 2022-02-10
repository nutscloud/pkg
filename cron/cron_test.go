// cron_test.go

package cron

import (
	"fmt"
	"testing"
	"time"
)

var times_run int

func dailyTask() {
	fmt.Println("time = %s.", time.Now().String())
	times_run++
}

func Test_DailyCronJob(t *testing.T) {
	fmt.Println("Test_DailyCronJob.")

	NewCronJobBuilder().DailyAt(ANY, ANY, 5).DoJob("Test_DailyCronJob", dailyTask).Build()
}
