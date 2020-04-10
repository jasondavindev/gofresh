package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

type Job struct {
	command string
	task    *exec.Cmd
	mu      sync.Mutex
}

func NewJob(cmd string) *Job {
	return &Job{command: cmd}
}

func (job *Job) Run() {
	if job.command == "" {
		return
	}

	job.mu.Lock()
	task := createTask(job.command)
	job.task = task
	job.mu.Unlock()

	if err := task.Start(); err != nil {
		log.Fatal("start", err)
		return
	}

	if err := task.Wait(); err != nil {
		log.Fatal("wait", err)
		return
	}
}

func createTask(cmd string) *exec.Cmd {
	args := strings.Split(cmd, " ")
	p := exec.Command(args[0], args[1:]...)
	p.Stdout = os.Stdout
	p.Stderr = os.Stderr
	return p
}
