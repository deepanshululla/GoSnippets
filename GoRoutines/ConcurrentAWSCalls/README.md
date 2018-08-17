The purpose of this is to create an implementation that does concurrent AWS
calls to create a mapping between different AWS entities.

In this case we are creating mapping for different AWS entities.

However AWS imposes a limit on the number of concurrent requests. Hence
we need to synchonize the go routines to run one at a time using channels


```text
Run using

go run *.go

```
