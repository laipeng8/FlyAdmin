package listener

import (
	"fmt"
	event2 "server/app/event"
	"server/pkg/event"
)

type TestListener struct {
}

func NewTestListener() *TestListener {
	return &TestListener{}
}

func (t *TestListener) Process(event event.Event) {
	switch ev := event.(type) {
	case *event2.LoginEvent:
		fmt.Println(ev.Name)
	}
}
