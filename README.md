# Getting started with Go UI tests

UI test demo with [go-rod](https://github.com/go-rod/rod): opens GitHub search, searches for `qa.guru`, asserts results contain `QA.GURU`.

## Run

```bash
go mod download
go test -v
```

Rod downloads a browser automatically on first run.
