# Tests Research

This repository researches test systems in programming languages.

The code here includes tests written in several programming languages. Some tests include fixtures and paraterized tests.

## Go

```shell
cd go/math
go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega/...
go test
```

## Rust

```shell
cd rust
cargo test
```

## Dlang

```shell
cd dlang
rdmd -unittest Math.d
```

## Swift

```shell
cd swift
swift test
```