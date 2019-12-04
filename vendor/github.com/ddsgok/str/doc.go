// Copyright 2018 Dênnis Dantas de Sousa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package str represents a string production line.

This is made with the interfaces in this package Chainer, Splitter and
Collecter. Each are created using srt.New as starting point, this
function will convert via fmt.Sprintf, a string to Chainer. And then
various operations can be performed on it.

Each object have the ability to put its output as a string, as an
error and printed onto desired io.Writer.

The package can be used like this:

	feat := "The_fire_Blaze"
	s := str.New(feat).Split("_").String() // "ThefireBlaze"

Performance

This package is slower than usual solutions. Because it makes lots of
conversion, and the use of fmt package itself, that have some
performance problems.

This is due to the objective of this package: to be a clearer tool to
operates on strings.
*/package str
