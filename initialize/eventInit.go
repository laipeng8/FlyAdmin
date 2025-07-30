package initialize

import (
	"server/app/event"
	"server/app/listener"
	event2 "server/pkg/event"
)

func EventInit() event2.EventDispatcher {

	EventDispatcher := event2.NewDispatcher()
	EventDispatcher.Register(event.TestEvent{}.GetEventName(), listener.NewTestListener())
	EventDispatcher.Register(event.LoginEvent{}.GetEventName(), listener.NewTestListener())
	return EventDispatcher

}
