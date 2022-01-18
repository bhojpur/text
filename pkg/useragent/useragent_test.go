package useragent

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

import "testing"

func TestVersionInfo(t *testing.T) {
	vi := VersionInfo{"foo", "bar"}
	if !vi.isValid() {
		t.Fatalf("VersionInfo should be valid")
	}
	vi = VersionInfo{"", "bar"}
	if vi.isValid() {
		t.Fatalf("Expected VersionInfo to be invalid")
	}
	vi = VersionInfo{"foo", ""}
	if vi.isValid() {
		t.Fatalf("Expected VersionInfo to be invalid")
	}
}

func TestAppendVersions(t *testing.T) {
	vis := []VersionInfo{
		{"foo", "1.0"},
		{"bar", "0.1"},
		{"pi", "3.1.4"},
	}
	v := AppendVersions("base", vis...)
	expect := "base foo/1.0 bar/0.1 pi/3.1.4"
	if v != expect {
		t.Fatalf("expected %q, got %q", expect, v)
	}
}
