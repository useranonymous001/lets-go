package main

import "sync"

/*
Building your own reader writer mutex, to understand it more well.

RLock => ReadLock
WLock => WriteLock
RUnlock => ReadUnlock
WUnlock => WriteUnlock

*/

// readersLock  ==> simulating the readers lock
// writersLock  ==> simulating the writers lock
// readersCounter ==> keeping track of number of readers inside the critical section

// structure for reader writer mutex

type ReaderWriterMutex struct {
	readersCounter int
	readersLock    sync.Mutex
	globalLock     sync.Mutex
}

func (rw *ReaderWriterMutex) ReadLock() {

	rw.readersLock.Lock() // lock the readers so that no other reader interfere
	rw.readersCounter++   // now increment the readers count

	// now if the reader is first one to enter
	// block the writer from entering the critical section

	if rw.readersCounter == 1 {
		rw.globalLock.Lock()
	}
	rw.readersLock.Unlock() // release the reader lock so that other readers can also enter the critical section
}

func (rw *ReaderWriterMutex) WriteLock() {
	// just lock the critical section, blocking all other go routine
	rw.globalLock.Lock()
}

func (rw *ReaderWriterMutex) ReadUnlock() {
	// unlock the global lock

	// lock the readers so that only one reader can exit at a time
	rw.readersLock.Lock()
	rw.readersCounter--

	// check whether the reader exiting is the last reader or not
	// if so, release the global lock, so that write can now access the critical section
	if rw.readersCounter == 0 {
		rw.globalLock.Unlock()
	}

	// now release the final reader
	rw.readersLock.Unlock()
}

func (rw *ReaderWriterMutex) WriteUnlock() {
	// unlock the global locker again free for other reader
	rw.globalLock.Unlock()
}
