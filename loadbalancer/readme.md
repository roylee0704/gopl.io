#Load Balancer

Imagine you have many processes requesting actions and a few workers that share the load. Assume workers are more efficient if they batch many requests.

We want a load balancer that distributes the load fairly across the workers according to their current workload.

In real life, we'd distribute work across many machines, but for simplicity we'll just focus on a local load balancer.

This is simplistic but representative of the core of a realistic problem.

##Specs

### Life of a request

Requesters make a request to the load balancer.

Load balancer immediately sends the request to the most lightly loaded worker.

When it completes the task, the worker replies directly to the requester, and at the same time signal balancer that it has completed the task.

Balancer adjusts its measure of the workload.


The requester sends Requests to the balancer
```go
type Request struct {
  fn func() int // The operation to perform
  c chan int    // The channel on which to return the result
}
```

An artificial but illustrative simulation of a requester.
```go
func requester(work chan Request) {
  c := make(chan int)
  for {
    time.Sleep(rand.Int63n(nWoeker * 2e9))  // spend time
    work <- Request{workFn, c}              // send request
    result := <-c                           // wait for answer
    furtherProcess(result)
  }
}
```
