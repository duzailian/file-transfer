package cmd

import (
	"testing"

	"github.com/Mikubill/transfer/apis/public/wenshushu"
)

func TestParseLinkMatchesWenshushuShortDomain(t *testing.T) {
	backend := ParseLink("https://c.wss.ink/f/k0117xonzz8/")
	if backend != wenshushu.Backend {
		t.Fatalf("ParseLink returned %T, want wenshushu backend", backend)
	}
}

func TestParseLinkDoesNotPanicOnBackendsWithoutMatcher(t *testing.T) {
	backend := ParseLink("https://example.invalid/f/k0117xonzz8/")
	if backend != nil {
		t.Fatalf("ParseLink returned %T, want nil", backend)
	}
}
