package main

import (
	"bufio"
	"flag"
	"io/ioutil"
	"os"
	"strconv"
)

var posFile = flag.String("pos", "last.pos", "Position file")

func main() {
	flag.Parse()

	lastPos := resumeLastPosition()
	read := pipe(lastPos)
	content := []byte(strconv.Itoa(read))
	ioutil.WriteFile(*posFile, content, os.ModePerm)
}

func pipe(lastPos int) int {
	scanner := bufio.NewScanner(os.Stdin)
	read := 0
	for scanner.Scan() {
		read++
		if read > lastPos {
			println(scanner.Text())
		}
	}
	return read
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func resumeLastPosition() int {
	lastPosBytes, err := ioutil.ReadFile(*posFile)
	if err != nil {
		return 0
	}

	lastPosStr := string(lastPosBytes)
	if lastPosStr == "" {
		return 0
	}

	lastPos, err := strconv.ParseInt(lastPosStr, 10, 32)
	if err != nil {
		return 0
	}

	return int(lastPos)
}
