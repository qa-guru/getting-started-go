package main

import (
	"testing"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func openPage(t *testing.T, url string) *rod.Page {
	t.Helper()

	browser := rod.New().
		ControlURL(launcher.New().Headless(false).MustLaunch()).
		MustConnect()
	t.Cleanup(func() { browser.MustClose() })

	return browser.MustPage(url).Timeout(15 * time.Second)
}
