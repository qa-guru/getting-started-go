# Getting started with Go UI tests

UI test demo with [tebeka/selenium](https://github.com/tebeka/selenium): opens GitHub search, searches for `qa.guru`, asserts results contain `QA.GURU`.

## Run

```bash
brew install chromedriver   # chromedriver must be on PATH
go mod download
go test -v
```
