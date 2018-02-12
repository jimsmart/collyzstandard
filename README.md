# collyzstandard

collyzstandard is a [Go](https://golang.org) package providing an experimental work-in-progress prototype of a cache compressor for [Colly](https://github.com/gocolly/colly).

## Installation
```bash
$ go get github.com/jimsmart/collyzstandard
```

```go
import "github.com/jimsmart/collyzstandard"
```

### Dependencies

- [Zstd (github.com/DataDog/zstd)](https://github.com/DataDog/zstd) — Go wrapper for Facebook's fast real-time compression algorithm.
- Standard library.
- [Ginkgo](https://onsi.github.io/ginkgo/) and [Gomega](https://onsi.github.io/gomega/) if you wish to run the tests.

## Example

```go
import "github.com/jimsmart/collysqlite"
import "github.com/jimsmart/collyzstandard"

cache := collysqlite.NewCache("./cache")
compcache := collyzstandard.NewCompressor(cache)

// TODO SetCache is a proposed method on colly.Collector that is currently unimplemented.
// c := colly.NewCollector()
// c.SetCache(compcache)
// ...

```

## Documentation

GoDocs [https://godoc.org/github.com/jimsmart/collyzstandard](https://godoc.org/github.com/jimsmart/collyzstandard)

## Testing

To run the tests execute `go test` inside the project folder.

For a full coverage report, try:

```bash
$ go test -coverprofile=coverage.out && go tool cover -html=coverage.out
```

## License

Package collyzstandard is copyright 2018 by Jim Smart and released under the [MIT License](LICENSE.md)
