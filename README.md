[![Build Status](https://travis-ci.org/hauke96/go-cape.svg?branch=master)](https://travis-ci.org/hauke96/go-cape)
# go-cape
go-cape (go-cape offers a chained argument parsing engine) is an argument parsing engine for go using chained method calls.

## Install and update `go-cape`
To install or update `go-cape` just execute `go get -u github.com/hauke96/go-cape`.

## Chained calls
Chained method calls are calls like this:
```
a.b().c().d()
```
Every method `b, c, d` returns the reference to `a` so that a new method can be called.

## Using go-cape
Here's an example from the [g0Ch@]() project:
```
p := parser.NewParser()
port := p.RegisterArgument("port", "p", "The port of the g0Ch@ server (usually 44494)").Default("44494").String()
ip := p.RegisterArgument("ipadress", "i", "The IP adress of the server").Required().String()
```
This creates a new parser and registeres the argument `port` with the short variant `p`, the help description `The port of the g0Ch@ server (usually 44494)` and the default value of `44494`. The variable `port` contains a pointer to a string because of the `String()` method call at the end.

The `ipadress` argument is a required on. Not setting this will cause a error in the parse method.

## TODO
* Implement the required condition
* ~~Implement flags (e.g. just setting `-iLb` or `-i` and interpreting this as `true` for the matching arguments)~~
* ~~Remove `PREDEF: ` output~~
