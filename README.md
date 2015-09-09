# GLLRB
Left Leaning Red Black Tree written in Go

[![Build Status](https://travis-ci.org/levigross/gllrb.svg?branch=master)](https://travis-ci.org/levigross/gllrb) [![GoDoc](https://godoc.org/github.com/levigross/gllrb?status.svg)](https://godoc.org/github.com/levigross/gllrb) [![Coverage Status](https://coveralls.io/repos/levigross/gllrb/badge.svg?branch=master&service=github)](https://coveralls.io/github/levigross/gllrb?branch=master)

License
======

GLLRB is licensed under the Apache License, Version 2.0. See [LICENSE](LICENSE) for the full license text


Example
=======

```go
llrb := gllrb.NewLLRB()
llrb.Put(gllrb.Bytes(word))
llrb.Delete(gllrb.Bytes(word))
```
