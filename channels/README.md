# What Are Channels in Go?
- Channels are pipes that connect concurrent goroutines. You can directly send message via one go routine and then receive in another go routine
- Think of it like a thread-safe queue, but baked into the language.

# Primary Purpose of Channels
- Communicate between go routines
- Syncronize executions (wait for the go routine to complete)
    - when the channel is waiting for the message, it will block that operation until some values are not received
- Coordinate work (pass message, results, signals)
- Share data without explicit locks (avoiding  manual use of sync.Mutex in many cases)




GOROUTINES
    - user level threads that is capable of executing independently

CHANNELS - MESSAGE PASSING THROUGH CHANNELS
    - methods used to pass message between goroutine without sharing memory 
    - Buffered: 
        - fixed size and asynchronous (does not block the gorutine execution)
        - pass data and forget about it 

    - UnBuffered: 
        - unlimited size but synchronous (blocks the goroutine execution)
        - pass data and wait until some goroutine receives it [more reliable] 

SELECT CHANNELS
    - switch statement type to select channel based on the data received

FOR-SELECT-LOOP 
    - loop that is used to automate the select statement
    - especially for buffered channels
