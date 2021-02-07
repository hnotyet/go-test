package main

import (
	"fmt"
	"time"
)

func main() {
	//yyyyMMddHHmmss
	timestamp := time.Now().Format("20060102150405")
	//timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(timestamp)


}