package urlutil

import "testing"

var (
	gitUrls = []string{
		"git://github.com/bhojpur/platform",
		"git@github.com:bhojpur/platform.git",
		"git@bitbucket.org:atlassianlabs/bhojpur-platform.git",
		"https://github.com/bhojpur/platform.git",
		"http://github.com/bhojpur/platform.git",
		"http://github.com/bhojpur/platform.git#branch",
		"http://github.com/bhojpur/platform.git#:dir",
	}
	incompleteGitUrls = []string{
		"github.com/bhojpur/platform",
	}
	invalidGitUrls = []string{
		"http://github.com/bhojpur/platform.git:#branch",
	}
	transportUrls = []string{
		"tcp://example.com",
		"tcp+tls://example.com",
		"udp://example.com",
		"unix:///example",
		"unixgram:///example",
	}
)

func TestIsGIT(t *testing.T) {
	for _, url := range gitUrls {
		if !IsGitURL(url) {
			t.Fatalf("%q should be detected as valid Git url", url)
		}
	}

	for _, url := range incompleteGitUrls {
		if !IsGitURL(url) {
			t.Fatalf("%q should be detected as valid Git url", url)
		}
	}

	for _, url := range invalidGitUrls {
		if IsGitURL(url) {
			t.Fatalf("%q should not be detected as valid Git prefix", url)
		}
	}
}

func TestIsTransport(t *testing.T) {
	for _, url := range transportUrls {
		if !IsTransportURL(url) {
			t.Fatalf("%q should be detected as valid Transport url", url)
		}
	}
}
