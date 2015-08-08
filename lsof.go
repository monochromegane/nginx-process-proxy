package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type lsof struct{}

func (l lsof) portByPid(pid int) (int, error) {
	files, err := l.openFiles(pid)
	if err != nil {
		return 0, err
	}
	return l.portByOutput(files)
}

func (l lsof) portByOutput(source []byte) (int, error) {
	tcp, err := l.grep(source, "TCP")
	if err != nil {
		return 0, err
	}

	var words []string
	line := strings.Split(tcp[0], " ")
	for _, l := range line {
		if l == "" {
			continue
		}
		words = append(words, l)
	}
	hostAndPort := strings.Split(words[8], ":")
	port, _ := strconv.Atoi(hostAndPort[1])
	return port, nil
}

func (l lsof) openFiles(pid int) ([]byte, error) {
	cmd := exec.Command("lsof", "-n", "-P", "-p", fmt.Sprintf("%d", pid))
	return cmd.Output()
}

func (l lsof) grep(source []byte, pattern string) ([]string, error) {

	scanner := bufio.NewScanner(bytes.NewReader(source))
	var matches []string
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), pattern) {
			matches = append(matches, scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return matches, nil
}
