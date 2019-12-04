# str [![GoDoc](https://godoc.org/github.com/ddspog/str?status.svg)](https://godoc.org/github.com/ddspog/str) [![Go Report Card](https://goreportcard.com/badge/github.com/ddspog/str)](https://goreportcard.com/report/github.com/ddspog/str) [![Commitizen friendly](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg)](http://commitizen.github.io/cz-cli/) [![Travis CI](https://travis-ci.org/ddspog/str.svg?branch=master)](https://travis-ci.org/ddspog/str)

by [ddspog](http://github.com/ddspog)

Package **str** represents a string production line.

## License

You are free to copy, modify and distribute **str** package with attribution under the terms of the MIT license. See the [LICENSE](https://github.com/ddspog/str/blob/master/LICENSE) file for details.

## Installation

Install **str** package with:

```shell
go get github.com/ddspog/str
```

## How to use

This package represents a string production line.

This is made with the interfaces in this package Chainer, Splitter and Collecter. Each are created using srt.New as starting point, this function will convert via fmt.Sprintf, a string to Chainer. And then various operations can be performed on it.

Each object have the ability to put its output as a string, as an error and printed onto desired io.Writer.

The package can be used like this:

```go
feat := "The_fire_Blaze"
s := str.New(feat).Split("_").String() // "ThefireBlaze"
```

## Performance

This package is slower than usual solutions. Because it makes lots of conversion, and the use of fmt package itself, that have some
performance problems.

This is due to the objective of this package: to be a clearer tool to operates on strings.

## Testing

This package has tests covering all code on it. Further additions to
code should try to follow this.

## Contribution

This package has some objectives from now:

* Incorporate all operations on strings package.
* Incorporate any new ideas about possible improvements.

Any interest in help is much appreciated.