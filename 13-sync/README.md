## Sync Library

### WaitGroup

> A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished.