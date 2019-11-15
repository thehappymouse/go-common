package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestWorkDayAdd(t *testing.T) {
	startDate, _ := time.Parse("2006-01-02", "2019-05-15")

	right := map[int]string{
		0:  "2019-05-15",
		1:  "2019-05-16",
		2:  "2019-05-17",
		3:  "2019-05-20",
		4:  "2019-05-21",
		7:  "2019-05-24",
		8:  "2019-05-27",
		11: "2019-05-30",
		12: "2019-05-31",
		13: "2019-06-03",
	}

	for i, v := range right {
		target := WorkDayAdd(i, startDate)
		if v != target.Format("2006-01-02") {
			t.Errorf("%s 开始，第 %d 个工作日，应该是:%s，然后确得到:%s", startDate, i, v, target)
		}
	}

	fmt.Println("当前日期:", startDate)
	for i := 1; i < 13; i++ {
		target := WorkDayAdd(i, startDate)
		fmt.Printf("%d 个工作日后: %s \n", i, target.Format("2006-01-02 Monday"))
	}
}
