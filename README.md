# Gonuts

> A program to make requests to nats.


## Installing

Run `go get github.com/blmayer/gonuts`.


## Using

Run `gonuts` on the terminal, the first argument is the nats' address, the second is
the channel/subject and the last one is the string data to be sent. The data can be
-f <path/to/file> to use a file as data. E.G.:

**Data string:**

`gonuts localhost:4222 test "this is the test data"`

**File as data:**

`gonuts localhost:4222 test -f test.json`



## Meta


### Author

Brian Mayer ([bleemayer@gmail.com](mailto:bleemayer@gmail.com))


### License

See [LICENSE](LICENSE)


### Created

Jun 10 2021
