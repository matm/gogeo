[![GoDoc](https://godoc.org/github.com/matm/gogeo?status.svg)](https://godoc.org/github.com/matm/gogeo)

## gogeo

A Simple geocoding library for Go.

### What Is Geocoding?

Geocoding is the process of finding associated geographic coordinates
(often expressed as latitude and longitude) from other geographic data, such as street addresses, 
or ZIP codes (postal codes).

### Installation

Get the library with

    go get github.com/matm/gogeo

### Try It!

You can give a try in a terminal by installing the `geocode` binary:

    go get github.com/matm/gogeo/geocode

For example, pass the `-g` option flag to use Google's geocoding API to locate Paris:

    $ geocode -g "Paris, France"

### License

This is free software, released under the GNU GPL.
