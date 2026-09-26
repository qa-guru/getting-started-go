# Getting started with Go UI tests

Three ways to do the same UI test: open GitHub search, search for `qa.guru`, assert results contain `QA.GURU`.

- `selenium_test.go` — [tebeka/selenium](https://github.com/tebeka/selenium) WebDriver client (needs `chromedriver` on PATH, e.g. `brew install chromedriver`)
- `rod_test.go` — [go-rod](https://github.com/go-rod/rod), high-level Chrome DevTools Protocol client (auto-downloads browser)
- `playwright_test.go` — [playwright-go](https://github.com/mxschmitt/playwright-go) (auto-installs browser)

Driver plumbing lives in the `*_driver_test.go` helpers next to each test.

## Run

```bash
go mod download
go test -v                  # all three
go test -v -run TestRodSearch   # one
```
