package main

import (
	"fmt"
	"strings"
	"time"
)

type UserSessionc struct {
	Date     string
	RealTime string
	TimeZone string
	City     string
	Any      string
}

func (p *UserSessionc) met(str []string) {
	p.Date = str[0]
	p.RealTime = str[1]
	p.TimeZone = str[2]
	p.City = str[3]
	p.Any = str[4]
}
func (p UserSessionc) printStruct() {
	fmt.Println(" ", strings.Repeat("-", len(p.RealTime)+12))
	fmt.Println("|Date     = ", p.Date, strings.Repeat(" ", len(p.RealTime)-len(p.Date)-1), "|")
	fmt.Println("|Time     = ", p.RealTime, "|")
	fmt.Println("|TimeZone = ", p.TimeZone, strings.Repeat(" ", len(p.RealTime)-len(p.TimeZone)-1), "|")
	fmt.Println("|City     = ", p.City, strings.Repeat(" ", len(p.RealTime)-len(p.City)-1), "|")
	fmt.Println(" ", strings.Repeat("-", len(p.RealTime)+12))
}
func main() {
	var str []string
	var myStruct [4]UserSessionc
	ch := make(chan UserSessionc, 1)
	go func() {
		for i := 0; i < 4; i++ {
			tmp := time.Now().String()
			str = strings.Split(tmp, " ")
			myStruct[i].met(str)
			ch <- myStruct[i]
			time.Sleep(time.Second)
		}
	}()
	for {
		select {
		case msg1 := <-ch:
			msg1.printStruct()
		case <-time.After(2 * time.Second):
			fmt.Println("time.After timeout happend")
			return
		}
	}
}
