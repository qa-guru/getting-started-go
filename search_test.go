package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tebeka/selenium"
)

func TestSuccessfulSearch(t *testing.T) {
	service, err := selenium.NewChromeDriverService("chromedriver", 9515)
	require.NoError(t, err, "failed to start chromedriver")
	defer service.Stop()

	wd, err := selenium.NewRemote(
		selenium.Capabilities{"browserName": "chrome"},
		"http://localhost:9515/wd/hub",
	)
	require.NoError(t, err, "failed to connect to chromedriver")
	defer wd.Quit()

	require.NoError(t, wd.SetImplicitWaitTimeout(10*time.Second))
	require.NoError(t, wd.Get("https://github.com/search"))

	search, err := wd.FindElement(selenium.ByCSSSelector, "[aria-label='Search GitHub']")
	require.NoError(t, err, "search input not found")
	require.NoError(t, search.SendKeys("qa.guru"+selenium.EnterKey))

	results, err := wd.FindElement(selenium.ByCSSSelector, "[data-testid='results-list']")
	require.NoError(t, err, "results list not found")
	text, err := results.Text()
	require.NoError(t, err)
	require.Contains(t, text, "QA.GURU")
}
