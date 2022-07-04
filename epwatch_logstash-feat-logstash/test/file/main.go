package main

import (
	"fmt"

	"github.com/CMKL-PTTEP/epwatch_logstash/internal/pkg/repository"
)

func main() {
	fileRepo := repository.NewLogstashFileRepository("/mnt/shared-log/qradar/udp")
	logChan := make(chan string)
	quitChan := make(chan bool)
	defer close(logChan)
	defer close(quitChan)

	go fileRepo.GetLog("2022-06-27.log", 0, logChan, quitChan)
	// conn, _ := net.Dial("tcp", "20.157.117.5:6514")
	line := 0

L:
	for {
		select {
		case <-quitChan:
			break L
		case <-logChan:
			line++
		}
	}
	fmt.Println(line)
}
