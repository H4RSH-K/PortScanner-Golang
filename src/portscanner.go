package main

import (
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func printUsage() {
	log.Println("Usage:")
	log.Println("	go run postscanner.go <host>")
	log.Println("Example:")
	log.Println("	go run portscanner.go 8.8.8.8")
}

func testTCPConnect(host string, port int, doneChannel chan bool) {
	timeoutLength := 5 * time.Second
	conn, err :=net.DialTimeout("tcp", host + ":" + strconv.Itoa(port), timeoutLength)
	if err != nil {
		doneChannel <- false
		return //Could Not Connect
	}
	conn.Close()
	log.Printf("[+] %d connected", port)
	doneChannel <- true
}

func main() {
	if len(os.Args) == 1 {
		log.Println("No Arguments Provided.")
		printUsage()
		os.Exit(1)
	}
	doneChannel := make(chan bool)
	activeThreadCount := 0
	log.Println("Scanning host: " + os.Args[1])
	for portNumber := 1; portNumber <= 65535; portNumber++ {
		activeThreadCount++
		go testTCPConnect(os.Args[1], portNumber, doneChannel)
	}
	for {
		<-doneChannel
		activeThreadCount--;
		log.Printf("Reducing threadCount. Now at %d\n", activeThreadCount)
		if activeThreadCount == 0{
			break
		}
	}
	//until activeThreadCount == 0 Keep Checking
}
