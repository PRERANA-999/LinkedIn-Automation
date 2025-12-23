package stealth

import "github.com/go-rod/rod"

func Apply(page *rod.Page) {
	page.MustEvalOnNewDocument(`
		Object.defineProperty(navigator, 'webdriver', { get: () => undefined });
		Object.defineProperty(navigator, 'platform', { get: () => 'Win32' });
		Object.defineProperty(navigator, 'languages', { get: () => ['en-US', 'en'] });
		Object.defineProperty(navigator, 'hardwareConcurrency', { get: () => 8 });
		Object.defineProperty(navigator, 'deviceMemory', { get: () => 8 });
	`)
}
