// WRITE PREFERRED RWMUTEX

package main

import "sync"

// We could block new readers from acquiring the read lock as soon as a writer calls the WriteLock() function.

/*
To design a write-preferred lock, we need a few
properties:

 Readers’ counter—Initially set to 0, this tells us how many reader goroutines are
actively accessing the shared resources.

 Writers’ waiting counter—Initially set to 0, this tells us how many writer goroutines
are suspended waiting to access the shared resource.

 Writer active indicator—Initially set to false, this flag tells us if the resource is cur-
rently being updated by a writer goroutine.

 Condition variable with mutex—This allows us to set various conditions on the
preceding properties, suspending execution when the conditions aren’t met.
*/

// readers counter = int
// writers waiting counter = int
// writers active indicator = bool

/*
Conditions:

-   when nothing is accessing the critical section (when no reader and no writer)
	allow the reader to acceess the critical section and increment the readers counter

-	when readers acquire the lock.
	that means any write trying to access the cs needs to be in waiting section. increment the  writers waiting

-	when some writers are in waiting state
	that means the write some writer are already in waiting state, so no any new reader can access the cs
	they will be blocked until the writers lock is unlocked

*/

type ReadWriteMutex struct {
	readersCounter int
	writersWaiting int
	activeWriter   bool
	cond           *sync.Cond
}

func (rw *ReadWriteMutex) NewReadWriteMutex() *ReadWriteMutex {
	// initializing new read write mutex with new condition variable and associated mutex
	return &ReadWriteMutex{cond: sync.NewCond(&sync.Mutex{})}
}

func (rw *ReadWriteMutex) ReadLock() {

	// apply readers lock
	// acquires mutex
	rw.cond.L.Lock()

	// check whether theres any writers lock waiting or not
	// if there's block other wise apply readers lock
	// Waits on condition variable while writers are waiting or active
	for rw.writersWaiting > 0 || rw.activeWriter {
		rw.cond.Wait()
	}

	// increment the readers counter
	rw.readersCounter++

	// unlock so other reader go routine can access the cs again
	// realeases the mutex
	rw.cond.L.Unlock()
}

// checks for the condition and decrement the readers the counter and then release the acquired  mutex
// and send the broadcast signal to all waiting go routine if its the last remaining reader
func (rw *ReadWriteMutex) ReadUnlock() {

	// acquire the mutex
	rw.cond.L.Lock()

	// decrement the readers count
	rw.readersCounter--

	// assuming this was the last reader
	if rw.readersCounter == 0 {
		rw.cond.Broadcast()
	}

	// release the mutex
	rw.cond.L.Unlock()

}

// check if already existing readers, either wait or acquire the mutex
func (rw *ReadWriteMutex) WriteLock() {
	// acquire the lock
	rw.cond.L.Lock()

	// increment the writers waiting
	rw.writersWaiting++
	for rw.readersCounter > 0 || rw.activeWriter {
		rw.cond.Wait()
	}

	// once wait is over, the waiting writer is not active so decrement the waiting writers
	rw.writersWaiting--

	// set the active writer to true
	rw.activeWriter = true

	// release the mutex
	rw.cond.L.Unlock()
}

// check the number of active writers and if 0 waiting writers, make it inactive again
func (rw *ReadWriteMutex) WriteUnlock() {
	// acquire the lock
	rw.cond.L.Lock()

	// set the writer active to false
	rw.activeWriter = false

	// send broadcase signal to all the goroutine
	rw.cond.Broadcast()

	// release the lock
	rw.cond.L.Unlock()

}
