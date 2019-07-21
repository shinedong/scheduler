package main

import (
	"fmt"
	cron "github.com/gorhill/cronexpr"
	"time"
)

func main() {
	fmt.Println("debug")
	expression, _ := cron.Parse("5 * * * * * *")
	next := expression.Next(time.Now())
	fmt.Println(next.String())

}
