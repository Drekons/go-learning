// Параллельное исполнение
package main

import (
	"fmt"
	"math"
)

type process func() error

func processor(processes []process, countOneTime int, countError int) error {
	chErr := make(chan error, countOneTime)
	cErr := 0
	paths := int(math.Ceil(float64(len(processes)) / float64(countOneTime)))

	for i := 0; i < paths; i++ {
		max := i*countOneTime + countOneTime
		if max > len(processes) {
			max = len(processes)
		}

		for _, f := range processes[i*countOneTime : max] {
			go func(f process) {
				chErr <- f()
			}(f)
		}

		for j := i * countOneTime; j < max; j++ {
			err := <-chErr
			if err != nil {
				cErr++
			}

			if cErr >= countError {
				close(chErr)
				return fmt.Errorf("%d process failed", cErr)
			}
		}
	}

	close(chErr)

	return nil
}
