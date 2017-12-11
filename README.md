# tson
**tson** is a JSON parser for Go that is able to change the time format flexibly.

## Installation

```sh
$ go get github.com/cia-rana/tson
```

## Usage

The way of use a function `tson.Unmarshal` is the same as a function `json.Unmarshal`; however, you can change the time format to parse JSON into an object with a function `tson.ChangeLayout` before parsing. See [example](https://github.com/cia-rana/tson/blob/master/_example/main.go) for detail.

## License
MIT
