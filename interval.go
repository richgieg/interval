// Package interval provides the Interval type which allows repeated execution
// at a fixed time interval.
package interval

import (
	"sync"
	"time"
)

// An Interval repeatedly calls a user-defined function at a fixed time
// interval.
type Interval struct {
	done     chan bool
	duration time.Duration
	function func()
	mutex    sync.Mutex
	running  bool
	ticker   *time.Ticker
}

// NewInterval returns a new Interval which, once started, repeatedly calls the
// function argument at the time interval specified by the duration arument.
func NewInterval(d time.Duration, f func()) *Interval {
	i := Interval{}
	i.duration = d
	i.function = f
	i.mutex = sync.Mutex{}
	return &i
}

// Start activates the Interval.
func (i *Interval) Start() {
	i.mutex.Lock()
	if i.running {
		i.mutex.Unlock()
		return
	}
	i.done = make(chan bool)
	i.ticker = time.NewTicker(i.duration)
	go func() {
		for {
			select {
			case <-i.ticker.C:
				i.function()
			case <-i.done:
				return
			}
		}
	}()
	i.running = true
	i.mutex.Unlock()
}

// Stop deactivates the Interval.
func (i *Interval) Stop() {
	i.mutex.Lock()
	if !i.running {
		i.mutex.Unlock()
		return
	}
	i.ticker.Stop()
	close(i.done)
	i.running = false
	i.mutex.Unlock()
}
