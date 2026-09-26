package main

import (
	"testing"

	"github.com/go-rod/rod/lib/input"
	"github.com/stretchr/testify/require"
)

func TestSuccessfulSearch(t *testing.T) {
	page := openPage(t, "https://github.com/search")

	page.MustElement("[aria-label='Search GitHub']").
		MustInput("qa.guru").
		MustType(input.Enter)
	page.MustWaitNavigation()

	results := page.MustElement("[data-testid='results-list']").MustText()
	require.Contains(t, results, "QA.GURU")
}
