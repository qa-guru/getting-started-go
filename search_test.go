package main

import (
	"testing"

	"github.com/mxschmitt/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestSuccessfulSearch(t *testing.T) {
	require.NoError(t, playwright.Install(&playwright.RunOptions{
		Browsers: []string{"chromium"},
	}))

	pw, err := playwright.Run()
	require.NoError(t, err, "could not start playwright")
	defer pw.Stop()

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	require.NoError(t, err, "could not launch browser")
	defer browser.Close()

	page, err := browser.NewPage()
	require.NoError(t, err, "could not create page")
	_, err = page.Goto("https://github.com/search")
	require.NoError(t, err, "could not open github search")

	search := page.Locator("[aria-label='Search GitHub']")
	require.NoError(t, search.Fill("qa.guru"))
	require.NoError(t, search.Press("Enter"))

	assertThat := playwright.NewPlaywrightAssertions()
	require.NoError(t,
		assertThat.Locator(page.Locator("[data-testid='results-list']")).ToContainText("QA.GURU"),
		"expected search results to contain 'QA.GURU'")
}
