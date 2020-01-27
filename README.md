# WorkerTest

Just a fun little experiment to get comfortable with worker pools, executing other processes, and collecting their output in Go.

All it does is spawn a worker pool that that call a Python process to hit a public API. It collects the results into a separate channel. 

This isn't an efficient way to solve the problem of hitting an API, but it was a useful learning experience to get a better feel for the language.
