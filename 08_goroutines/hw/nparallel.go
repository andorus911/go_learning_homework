package main

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	issues := make([]func() error, 5)

	for i := range issues {
		issues[i] = issue2
	}

	err := NParallel(issues, 1, 3)
	fmt.Println("Error:", err)
}

func issue1() error {
	fmt.Println("ISSUE 1 - 5 sec")
	time.Sleep(5 * time.Second)
	return nil
}

func issue2() error {
	fmt.Println("ISSUE 2 - 5 sec + Error")
	time.Sleep(5 * time.Second)
	return errors.New("oops")
}

func NParallel(issuesToRun []func() error, routineNumber uint32, errorNumberToStop uint32) error {
	routineController := make(chan struct{}, routineNumber)
	var errorCounter uint32
	var wg sync.WaitGroup

	for _, issue := range issuesToRun {
		routineController <- struct{}{}
		if atomic.LoadUint32(&errorCounter) < errorNumberToStop {
			wg.Add(1)
			go func(issue func() error, routineController <-chan struct{}) {
				defer func(){
					wg.Done()
					<-routineController
				}()
				err := issue()
				if err != nil {
					atomic.AddUint32(&errorCounter, 1)
				}
			}(issue, routineController)
		} else {
			break
		}
	}
	wg.Wait()
	if errorCounter < errorNumberToStop {
		return nil
	} else {
		return errors.New("error counter is exceeded")
	}
}