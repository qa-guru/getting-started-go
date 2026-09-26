package main

import (
	"testing"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/stretchr/testify/require"
)

func TestSuccessfulSearch(t *testing.T) {
	browser := rod.New().
		ControlURL(launcher.New().Headless(false).MustLaunch()).
		MustConnect()
	defer browser.MustClose()

	page := browser.MustPage("https://github.com/search")
	page.Timeout(15 * time.Second).
		MustElement("[aria-label='Search GitHub']").
		MustInput("qa.guru").
		MustType(input.Enter)
	page.MustWaitNavigation()

	results := page.Timeout(15 * time.Second).
		MustElement("[data-testid='results-list']").
		MustText()

	require.Contains(t, results, "QA.GURU")
}
