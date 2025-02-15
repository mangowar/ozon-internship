package shorten_test

import (
	"shortener/shorten"
	"testing"
)

func TestUniqueShortLink(t *testing.T) {
	urls := []string{
		"https://example.com/long-url-1",
		"https://example.com/long-url-2",
		"https://example.com/long-url-3",
		"https://example.com/long-url-4",
		"https://example.com/long-url-5",
		"https://example.com/long-url-6",
	}
	results := make(map[string]string)
	for _, url := range urls {
		shortened_link := shorten.ShortenLink(url)
		if value, exists := results[shortened_link]; exists {
			t.Errorf("%s failed: equals %s", url, value)
		} else {
			results[shortened_link] = url
		}
	}
}

func TestIndeptity(t *testing.T) {
	const link = "https://example.com/long-url-1"
	test_subject := shorten.ShortenLink(link)
	for i := 0; i < 100; i++ {
		if result := shorten.ShortenLink(link); result != test_subject {
			t.Errorf("%s failed: on attempt %d got result == %s insted %s", link, i, result, test_subject)
		}
	}
}
