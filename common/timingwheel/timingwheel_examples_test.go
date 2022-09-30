package timingwheel

import (
	"fmt"
	"time"
)

func Example_startTimer() {
	tw := NewTimingWheel(time.Millisecond, 20)
	tw.Start()
	defer tw.Stop()

	exitC := make(chan time.Time, 1)
	tw.AfterFunc(time.Second, func() {
		fmt.Println("The timer fires")
		exitC <- time.Now().UTC()
	})

	<-exitC

	// Output:
	// The timer fires
}

func Example_stopTimer() {
	tw := NewTimingWheel(time.Millisecond, 20)
	tw.Start()
	defer tw.Stop()

	t := tw.AfterFunc(time.Second, func() {
		fmt.Println("The timer fires")
	})

	<-time.After(900 * time.Millisecond)
	// Stop the timer before it fires
	t.Stop()

	// Output:
	//
}
