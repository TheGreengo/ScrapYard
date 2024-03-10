# Notes for Project 3

# What we need to do
- data channel: transfers file content, implemented using shared memory
- command channel: transfers transfer command, implemented using other IPC mechanisms, such as message queue
- relaying the contents of the cache to the proxy process by way of shared memory
- can use either System V or the POSIX API
- a pair of proxy/cache threads may only use 1 shared memory segment for the current connection
- the proxy should query the cache to see if the file is available in local memory already, if the file is available there, it should send the cached version
- **you need to only check the cache**. In other words, you only need to implement client, proxy, cache, and the communication protocols
-  the proxy is responsible for creating and destroying the shared memory (or message queues)
-  The cache must set up some communication mechanism (socket, message queue, shared memory) by which the proxy can communicate
-  Neither the cache daemon nor the proxy should crash if the other process is not started already
-  he proxy process (and perhaps the cache as well depending on your implementation) must use a signal handler for both SIGTERM (signal for kill) and SIGINT (signal for Ctrl-C) to first remove all existing IPC objects
-  

# Plans for how to do it
- test the signal handling
- have some sort of IPC process working
  - make a shared memory IPC (shared memory, mapped files)
  - make message passing IPC (socket, pipes, queues)
  - interprocess reader/writer