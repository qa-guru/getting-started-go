package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tebeka/selenium"
)

func TestSeleniumSearch(t *testing.T) {
	wd := startWebDriver(t)

	require.NoError(t, wd.Get("https://github.com/search"))

	search, err := wd.FindElement(selenium.ByCSSSelector, "[aria-label='Search GitHub']")
	require.NoError(t, err)
	require.NoError(t, search.SendKeys("qa.guru"+selenium.EnterKey))

	results, err := wd.FindElement(selenium.ByCSSSelector, "[data-testid='results-list']")
	require.NoError(t, err)
	text, err := results.Text()
	require.NoError(t, err)
	require.Contains(t, text, "QA.GURU")
}
