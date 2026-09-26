package main

import (
	"strings"
	"testing"
	"time"

	"github.com/tebeka/selenium"
)

func TestSuccessfulSearch(t *testing.T) {
	service, err := selenium.NewChromeDriverService("chromedriver", 9515)
	if err != nil {
		t.Fatalf("failed to start chromedriver: %v", err)
	}
	defer service.Stop()

	wd, err := selenium.NewRemote(
		selenium.Capabilities{"browserName": "chrome"},
		"http://localhost:9515/wd/hub",
	)
	if err != nil {
		t.Fatalf("failed to connect to chromedriver: %v", err)
	}
	defer wd.Quit()

	if err := wd.SetImplicitWaitTimeout(10 * time.Second); err != nil {
		t.Fatalf("failed to set implicit wait: %v", err)
	}
	if err := wd.Get("https://github.com/search"); err != nil {
		t.Fatalf("failed to open github search: %v", err)
	}

	search, err := wd.FindElement(selenium.ByCSSSelector, "[aria-label='Search GitHub']")
	if err != nil {
		t.Fatalf("search input not found: %v", err)
	}
	if err := search.SendKeys("qa.guru" + selenium.EnterKey); err != nil {
		t.Fatalf("failed to submit search: %v", err)
	}

	results, err := wd.FindElement(selenium.ByCSSSelector, "[data-testid='results-list']")
	if err != nil {
		t.Fatalf("results list not found: %v", err)
	}
	text, err := results.Text()
	if err != nil {
		t.Fatalf("failed to read results: %v", err)
	}
	if !strings.Contains(text, "QA.GURU") {
		t.Fatalf("expected search results to contain 'QA.GURU', got: %s", text)
	}
}
