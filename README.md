# Go Fuzz Example

This repository includes a fuzz test example written in Go.

To run the fuzz test

```shell
go test -fuzz=FuzzEqual ./...
```

There is currently a failing case in the `/testdate` directory. That failing case will be part of
the seed corpus for any subsequent calls. You can remove that failing case file in order to remove
it from the seed corpus. 
