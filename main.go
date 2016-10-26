package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var timeLayout = flag.String("layout", "2006/01/02 15:04:05", "日期模板，遵循golang规范")
var validTime = flag.Duration("valid_time", 10*time.Minute, "有效时间，当前时间前 valid_time 分钟之内的日志将会输出")
var timeStartPos = flag.Int("offset_start", 0, "时间开始偏移位置")
var timeEndpos = flag.Int("offset_end", 19, "时间截止偏移位置")

func main() {

	flag.Parse()

	location, err := time.LoadLocation("Local")
	if err != nil {
		log.Fatalf("Parse Error: %v", err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		logTime, err := time.ParseInLocation(*timeLayout, line[*timeStartPos:*timeEndpos], location)
		if err != nil {
			continue
		}
		diff := time.Now().Sub(logTime)
		if diff < *validTime {
			fmt.Print(line)
		}
	}
}