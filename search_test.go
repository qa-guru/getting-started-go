package main

import (
	"testing"

	"github.com/mxschmitt/playwright-go"
)

func TestSuccessfulSearch(t *testing.T) {
	pw, err := playwright.Run()
	if err != nil {
		t.Fatalf("could not start playwright: %v", err)
	}
	defer pw.Stop()

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	if err != nil {
		t.Fatalf("could not launch browser: %v", err)
	}
	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		t.Fatalf("could not create page: %v", err)
	}
	if _, err := page.Goto("https://github.com/search"); err != nil {
		t.Fatalf("could not open github search: %v", err)
	}

	search := page.Locator("[aria-label='Search GitHub']")
	if err := search.Fill("qa.guru"); err != nil {
		t.Fatalf("could not fill search input: %v", err)
	}
	if err := search.Press("Enter"); err != nil {
		t.Fatalf("could not submit search: %v", err)
	}

	assertThat := playwright.NewPlaywrightAssertions()
	if err := assertThat.Locator(page.Locator("[data-testid='results-list']")).
		ToContainText("QA.GURU"); err != nil {
		t.Fatalf("expected search results to contain 'QA.GURU': %v", err)
	}
}
