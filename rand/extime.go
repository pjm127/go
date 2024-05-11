package main

import (
	"fmt"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("Asia/Seoul")
	const longForm = "Jan 2, 2006 at 3:04pm"
	t1, _ := time.ParseInLocation(longForm, "Jun 14, 2021 at 10:00pm", loc) //1번쨰 형식으로 23
	fmt.Println(t1, t1.Location(), t1.UTC())
	const shrotForm = "2006-Jan-02"
	t2, _ := time.Parse(shrotForm, "2021-Jun-14") // 타임존 없으면 utc기준 0시
	fmt.Println(t2, t2.Location())

	t3, err := time.Parse("2021-03-23 15:20:21", "2021-05,14 12:32:32")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t3, t3.Location())
	d := t2.Sub(t1)
	fmt.Println(d)
}
