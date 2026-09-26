package main

import (
	"testing"

	"github.com/mxschmitt/playwright-go"
	"github.com/stretchr/testify/require"
)

func newPage(t *testing.T) playwright.Page {
	t.Helper()

	require.NoError(t, playwright.Install(&playwright.RunOptions{
		Browsers: []string{"chromium"},
	}))

	pw, err := playwright.Run()
	require.NoError(t, err, "could not start playwright")
	t.Cleanup(func() { pw.Stop() })

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	require.NoError(t, err, "could not launch browser")
	t.Cleanup(func() { browser.Close() })

	page, err := browser.NewPage()
	require.NoError(t, err, "could not create page")
	return page
}
