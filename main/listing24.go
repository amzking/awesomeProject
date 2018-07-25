package main

import (
	"os"
	"time"
	"errors"
	"os/signal"
)

type Runner1 struct {
	interrupt chan os.Signal

	complete chan error

	// <-chan  表示只能从chan中接收值，不能写值
	timeout <-chan time.Time

	tasks []func(int)
}

var ErrTimeout = errors.New("recived interrupt")

var ErrInterrupt = errors.New("received interrupt")

func New(d time.Duration) *Runner1 {
	return &Runner1 {
		interrupt: make(chan os.Signal, 1),
		complete : make(chan error),
		timeout: time.After(d),
	}
}

func (r *Runner1) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner1) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <- r.timeout :
		return ErrTimeout
	}
}

func (r *Runner1) run() error {
	for id, task := range r.tasks {
		if r.gotInerrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

func (r * Runner1) gotInerrupt() bool {
	select {
	case <- r.interrupt :
		signal.Stop(r.interrupt)
		return true

	default:
		return false
	}

	return nil;
}