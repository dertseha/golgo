[![Build Status][drone-image]][drone-url]
[![Coverage Status][coveralls-image]][coveralls-url]

# Game Of Life in Go

This project is an implementation of [Conway's Game of Life](http://en.wikipedia.org/wiki/Conways_Game_of_Life) in [Go](http://golang.org/). It uses OpenGL for visualization, which can be run both native and in a web-browser.

This is a pet-project of which the rationale can be found [here](http://manynames.sevensuns.at/technojoy/2014/07/game-of-life-concept-idea.html). Progress can be tracked on the same blog [here](http://manynames.sevensuns.at/technojoy/go-game-of-life/).

## Dependencies
### Native
The native build is using the [go-gl](http://go-gl.github.io/) projects. Dependencies should be downloaded by the ```install``` command (see below).

### Browser
The web-browser build requires [gopher-js](https://github.com/gopherjs/gopherjs) and [web-gl](https://github.com/gopherjs/webgl). Install them on the command line using
```
go get -u github.com/gopherjs/gopherjs
go get -u github.com/gopherjs/webgl
```

## Build
### Native
The native build is compiled, from the root of the project, calling
```
go install .
```

### Browser
The web-browser build is compiled, from the root of the project, calling
```
$GOPATH/bin/gopherjs build -o ./build/golgo-js.js golgo-js/*.go
```
The output files will be under the ```build``` directory, which is also referenced by the ```golgo-js.html``` file. Open this html file in a browser supporting WebGL to run the application.

## Testing
For testing, you need to have the test framework ```gocheck``` installed:
```
go get gopkg.in/check.v1
```

### OpenGL Tests
To test OpenGL and all output it is necessary to create windows and/or run a browser. These tests are in the folder ```test```.

For these tests, another dependency is necessary to be installed:
```
go get bitbucket.org/tebeka/selenium
```
and a [Selenium Server](http://docs.seleniumhq.org/download/) running. (Note: Not every version works as it seems; v2.39.0 works so far)

## License

The project is available under the terms of the **New BSD License** (see LICENSE file).

[drone-url]: https://drone.io/github.com/dertseha/golgo/latest
[drone-image]: https://drone.io/github.com/dertseha/golgo/status.png
[coveralls-url]: https://coveralls.io/r/dertseha/golgo
[coveralls-image]: https://coveralls.io/repos/dertseha/golgo/badge.png
