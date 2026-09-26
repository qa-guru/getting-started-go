# Getting started with Go UI tests

UI test demo with [playwright-go](https://github.com/mxschmitt/playwright-go): opens GitHub search, searches for `qa.guru`, asserts results contain `QA.GURU`.

## Run

```bash
go mod download
go run github.com/mxschmitt/playwright-go/cmd/playwright install chromium
go test -v
```
