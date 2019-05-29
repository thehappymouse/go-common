package utils

import (
	"fmt"
	"github.com/magiconair/properties/assert"
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

func TestUssNoGenerate(t *testing.T) {
	last := UssNoGenerate("")
	for i := 0; i < 40; i++ {
		fmt.Println(last)
		last = UssNoGenerate(last)
	}

}

func TestBHex2Num(t *testing.T) {
	assert.Equal(t, BHex2Num("1E6K", 36), 65036)
	assert.Equal(t, BHex2Num("FE0C", 16), 65036)

	assert.Equal(t, NumToBHex(65036, 36), "1E6K")
	assert.Equal(t, NumToBHex(65036, 16), "FE0C")

	fmt.Println(BHex2Num("7H", 16))
}
