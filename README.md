## Workz

A project based around worker queues built on a wednesday night to learn about goroutines, channels and clousures.

# What's happening here?
**workers** is a global buffered *chan chan* (a channel that holds channels) of the *Job* type. This is used as a queue.

**Jobs** is a global *chan* of the *Job* type. Also works like a queue.

At startup we use CreateWorkers() to create 10 *worker*s that when instantiated send their respective *chan Job* field to the **workers** *chan chan* and start reading from said field. CreateWorkers() then launches an anonymous goroutine that receives from the **Jobs** *chan*.

Before anything else happens the user is prompted to "Enter text". This string is wrapped in a *Job* and passed to **Jobs** where it is received by the previously mentioned anonymous goroutine. The same goroutine then recives a *chan Job* from **workers** (Belonging to a worker) that it send the previously received *Job* to. As i mentioned earlier the *workers* have goroutines receiving from their respective *chan Job* fields. So the worker now has it's *Job*! It uses the data from the *Job* to create a string containing the *worker*s ID, the *Job*s ID and *Job*s message (The string the user entered). This string is sent to a shared *chan Job* called "Out" that we are receiving from in main.

The users finds out the ID of the *worker* that got the job and thats about all.

But why? Curiosity, also it is a cool "pattern" and it is impressive how much we can do with so little.

![alt gopher](https://blog.golang.org/gopher/gopher.png)