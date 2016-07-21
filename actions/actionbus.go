package actions

import (
	"log"

	"github.com/tcolar/goed/core"
)

var _ core.ActionDispatcher = (*actionBus)(nil)

type actionBus struct {
	actionChan chan (core.Action)
	quitc      chan (struct{})
}

func NewActionBus() core.ActionDispatcher {
	return actionBus{
		actionChan: make(chan (core.Action), 1000),
		quitc:      make(chan (struct{})),
	}
}

func (a actionBus) Dispatch(action core.Action) {
	a.actionChan <- action
}

func (a actionBus) Start() {
	for {
		select {
		case action := <-a.actionChan:
			if core.Trace {
				log.Printf("> %#v", action)
			}
			action.Run()
			if core.Trace {
				log.Printf("< %#v", action)
			}
		case <-a.quitc:
			break
		}
	}
}

// Flush waits for all actions sent before it to have been processed
func (a actionBus) Flush() {
	c := make(chan (struct{}), 1)
	d(flushAction{c})
	<-c
}

func (a actionBus) Shutdown() {
	a.quitc <- struct{}{}
}

type flushAction struct {
	c chan (struct{})
}

func (a flushAction) Run() {
	a.c <- struct{}{}
}
