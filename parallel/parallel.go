package parallel

import (
	"github.com/hashicorp/go-multierror"
	"github.com/lithictech/aperitif/mariobros"
	"sync"
)

type empty struct{}
type Processor func(idx int) error

// ForEach processes data in parallel.
// total is the total number of items to process.
// n is the degree of parallelism.
// process is called with the index of the item being processed.
//
// ParallelFor acts as a semaphore over a total WaitGroup fan-out,
// and also coalesces errors into a single error result.
//
// If callers need process to return actual data,
// they should allocate a slice of the data they need,
// and assign to the slice index while processing.
// See ParallelForFiles for an example usage.
func ForEach(total int, n int, process Processor) error {
	semaphore := make(chan empty, n)
	errors := make([]error, total)

	wg := sync.WaitGroup{}
	wg.Add(total)
	for i := 0; i < total; i++ {
		go func(i int) {
			mario := mariobros.Yo("parallel.foreach")
			defer mario()
			semaphore <- empty{}
			errors[i] = process(i)
			<-semaphore
			wg.Done()
		}(i)
	}
	wg.Wait()
	return multierror.Append(nil, errors...).ErrorOrNil()
}
