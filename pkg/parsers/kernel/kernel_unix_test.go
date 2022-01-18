//go:build !windows
// +build !windows

package kernel

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"fmt"
	"testing"
)

func assertParseRelease(t *testing.T, release string, b *VersionInfo, result int) {
	var (
		a *VersionInfo
	)
	a, _ = ParseRelease(release)

	if r := CompareKernelVersion(*a, *b); r != result {
		t.Fatalf("Unexpected kernel version comparison result for (%v,%v). Found %d, expected %d", release, b, r, result)
	}
	if a.Flavor != b.Flavor {
		t.Fatalf("Unexpected parsed kernel flavor.  Found %s, expected %s", a.Flavor, b.Flavor)
	}
}

// TestParseRelease tests the ParseRelease() function
func TestParseRelease(t *testing.T) {
	assertParseRelease(t, "3.8.0", &VersionInfo{Kernel: 3, Major: 8, Minor: 0}, 0)
	assertParseRelease(t, "3.4.54.longterm-1", &VersionInfo{Kernel: 3, Major: 4, Minor: 54, Flavor: ".longterm-1"}, 0)
	assertParseRelease(t, "3.4.54.longterm-1", &VersionInfo{Kernel: 3, Major: 4, Minor: 54, Flavor: ".longterm-1"}, 0)
	assertParseRelease(t, "3.8.0-19-generic", &VersionInfo{Kernel: 3, Major: 8, Minor: 0, Flavor: "-19-generic"}, 0)
	assertParseRelease(t, "3.10.0-862.2.3.el7.x86_64", &VersionInfo{Kernel: 3, Major: 10, Minor: 0, Flavor: "-862.2.3.el7.x86_64"}, 0)
	assertParseRelease(t, "3.12.8tag", &VersionInfo{Kernel: 3, Major: 12, Minor: 8, Flavor: "tag"}, 0)
	assertParseRelease(t, "3.12-1-amd64", &VersionInfo{Kernel: 3, Major: 12, Minor: 0, Flavor: "-1-amd64"}, 0)
	assertParseRelease(t, "3.8.0", &VersionInfo{Kernel: 4, Major: 8, Minor: 0}, -1)
	// Errors
	invalids := []string{
		"3",
		"a",
		"a.a",
		"a.a.a-a",
	}
	for _, invalid := range invalids {
		expectedMessage := fmt.Sprintf("Can't parse kernel version %v", invalid)
		if _, err := ParseRelease(invalid); err == nil || err.Error() != expectedMessage {
			if err == nil {
				t.Fatalf("Expected %q, got nil", expectedMessage)
			}
			t.Fatalf("Expected %q, got %q", expectedMessage, err.Error())
		}
	}
}

func assertKernelVersion(t *testing.T, a, b VersionInfo, result int) {
	if r := CompareKernelVersion(a, b); r != result {
		t.Fatalf("Unexpected kernel version comparison result. Found %d, expected %d", r, result)
	}
}

// TestCompareKernelVersion tests the CompareKernelVersion() function
func TestCompareKernelVersion(t *testing.T) {
	assertKernelVersion(t,
		VersionInfo{Kernel: 3, Major: 8, Minor: 0},
		VersionInfo{Kernel: 3, Major: 8, Minor: 0},
		0)
	assertKernelVersion(t,
		VersionInfo{Kernel: 2, Major: 6, Minor: 0},
		VersionInfo{Kernel: 3, Major: 8, Minor: 0},
		-1)
	assertKernelVersion(t,
		VersionInfo{Kernel: 3, Major: 8, Minor: 0},
		VersionInfo{Kernel: 2, Major: 6, Minor: 0},
		1)
	assertKernelVersion(t,
		VersionInfo{Kernel: 3, Major: 8, Minor: 0},
		VersionInfo{Kernel: 3, Major: 8, Minor: 0},
		0)
	assertKernelVersion(t,
		VersionInfo{Kernel: 3, Major: 8, Minor: 5},
		VersionInfo{Kernel: 3, Major: 8, Minor: 0},
		1)
	assertKernelVersion(t,
		VersionInfo{Kernel: 3, Major: 0, Minor: 20},
		VersionInfo{Kernel: 3, Major: 8, Minor: 0},
		-1)
	assertKernelVersion(t,
		VersionInfo{Kernel: 3, Major: 7, Minor: 20},
		VersionInfo{Kernel: 3, Major: 8, Minor: 0},
		-1)
	assertKernelVersion(t,
		VersionInfo{Kernel: 3, Major: 8, Minor: 20},
		VersionInfo{Kernel: 3, Major: 7, Minor: 0},
		1)
	assertKernelVersion(t,
		VersionInfo{Kernel: 3, Major: 8, Minor: 20},
		VersionInfo{Kernel: 3, Major: 8, Minor: 0},
		1)
	assertKernelVersion(t,
		VersionInfo{Kernel: 3, Major: 8, Minor: 0},
		VersionInfo{Kernel: 3, Major: 8, Minor: 20},
		-1)
}
