# What Are Channels in Go?
- Channels are pipes that connect concurrent goroutines. You can directly send message via one go routine and then receive in another go routine
- Think of it like a thread-safe queue, but baked into the language.

# Primary Purpose of Channels
- Communicate between go routines
- Syncronize executions (wait for the go routine to complete)
    - when the channel is waiting for the message, it will block that operation until some values are not received
- Coordinate work (pass message, results, signals)
- Share data without explicit locks (avoiding  manual use of sync.Mutex in many cases)