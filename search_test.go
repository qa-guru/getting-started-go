package main

import (
	"testing"

	"github.com/mxschmitt/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestSuccessfulSearch(t *testing.T) {
	page := newPage(t)

	_, err := page.Goto("https://github.com/search")
	require.NoError(t, err)

	search := page.Locator("[aria-label='Search GitHub']")
	require.NoError(t, search.Fill("qa.guru"))
	require.NoError(t, search.Press("Enter"))

	assertThat := playwright.NewPlaywrightAssertions()
	require.NoError(t,
		assertThat.Locator(page.Locator("[data-testid='results-list']")).ToContainText("QA.GURU"))
}
