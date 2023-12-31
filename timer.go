package main

import (
	"context"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Timer struct {
	ctx context.Context
}

func NewTimer() *Timer {
	return &Timer{}
}

func (a *Timer) timerStartup(ctx context.Context) {
	a.ctx = ctx

	err := beeep.Alert("Theia", "Timer has been started", "")
	if err != nil {
    	println("Error:", err.Error())
	}
}

func (a *Timer) Time() {
	err := beeep.Alert("Timer", "It's time to look away from the screen mate! Look at an object far away for 20s", "")
	if err != nil {
    	println("Error:", err.Error())
	}

	runtime.WindowShow(a.ctx)

	time.AfterFunc(20 * time.Second, func ()  {
		runtime.WindowHide(a.ctx)

		err := beeep.Alert("Timer Over", "You can look back at the screen now", "")
		if err != nil {
    		println("Error:", err.Error())
		}

		runtime.WindowReloadApp(a.ctx)
	})
}