[![Build Status](https://travis-ci.org/terryh/goisbn.svg?branch=master)](https://travis-ci.org/terryh/goisbn)

goisbn: a go library for parsing ISBN
=====================================

This small library provodes functionality to validate ISBN and
convert to ISBN 10 and ISBN 13

Example:
--------

    package main

    import (
        "fmt"
        "github.com/terryh/goisbn"
    )

    func main() {
        //var is goisbn.ISBN

	    is, err := goisbn.ToISBN("981-02-4410-X")

        fmt.Println(is.ISBN10(), err)
        fmt.Println(is.ISBN13(), err)

    }
