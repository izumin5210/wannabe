# wannabe
[![GoDoc](https://godoc.org/github.com/izumin5210/wannabe?status.svg)](https://godoc.org/github.com/izumin5210/wannabe)
[![license](https://img.shields.io/github/license/izumin5210/wannabe.svg)](./LICENSE)
[![Go project version](https://badge.fury.io/go/github.com%2Fizumin5210%2Fwannabe.svg)](https://badge.fury.io/go/github.com%2Fizumin5210%2Fwannabe)

Make everything be boolean ([wannabe_bool](https://github.com/prodis/wannabe_bool) porting for golang)


## Usage

```go
wannabe.Bool("1") // => true
wannabe.Bool("0") // => false
wannabe.Bool("y") // => true
wannabe.Bool("N") // => false

// export FOO_ENABLED=1
wannabe.BoolEnv("FOO_ENABLED") // => true
```

See [./bool_test.go](./bool_test.go)


## Author
- Masayuki Izumi ([@izumin5210](https://github.com/izumin5210))


## License
licensed under the MIT License. See [LICENSE](./LICENSE)
