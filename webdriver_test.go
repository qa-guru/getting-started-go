package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tebeka/selenium"
)

func startWebDriver(t *testing.T) selenium.WebDriver {
	t.Helper()

	service, err := selenium.NewChromeDriverService("chromedriver", 9515)
	require.NoError(t, err, "failed to start chromedriver")
	t.Cleanup(func() { service.Stop() })

	wd, err := selenium.NewRemote(
		selenium.Capabilities{"browserName": "chrome"},
		"http://localhost:9515/wd/hub",
	)
	require.NoError(t, err, "failed to connect to chromedriver")
	t.Cleanup(func() { wd.Quit() })

	require.NoError(t, wd.SetImplicitWaitTimeout(10*time.Second))
	return wd
}
