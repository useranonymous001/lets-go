// package main

// import (
// 	"sync"
// )

// // build the reader-write structure

// type ReadWriteMutex struct {
// 	readersCounter int        // counts the number of reader
// 	readersLock    sync.Mutex // indicates the readers lock
// 	globalLock     sync.Mutex // indicates the global lock(write)
// }

// func (rw *ReadWriteMutex) ReadLock() {
// 	// lock the critical section
// 	rw.readersLock.Lock()
// 	// increment the readers count
// 	rw.readersCounter++

// 	// check whether the counter is 1st or not
// 	// is so lock the global-lock (denying access to writers lock
// 	if rw.readersCounter == 1 {
// 		rw.globalLock.Lock() // locks the writers mutex
// 	}

// 	// unlock the readers lock, so other readers would get chance
// 	rw.readersLock.Unlock()
// }

// func (rw *ReadWriteMutex) ReadUnlock() {
// 	// first lock the readers lock to decrement the readers count
// 	rw.readersLock.Lock()
// 	rw.readersCounter--

// 	if readersCounter == 0 {
// 		rw.globalLock.Unlock()
// 	}

// 	rw.readersLock.Unlock()
// }

// func (rw *ReadWriteMutex) WriteLock() {
// 	// just lock the global lock
// 	rw.globalLock.Lock()
// }

// func (rw *ReadWriteMutex) WriteUnlock() {
// 	// unlock the global lock
// 	rw.globalLock.Unlock()
// }

package main
