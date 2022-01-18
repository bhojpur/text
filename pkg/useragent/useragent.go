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

import (
	"strings"
)

// VersionInfo is used to model UserAgent versions.
type VersionInfo struct {
	Name    string
	Version string
}

func (vi *VersionInfo) isValid() bool {
	const stopChars = " \t\r\n/"
	name := vi.Name
	vers := vi.Version
	if len(name) == 0 || strings.ContainsAny(name, stopChars) {
		return false
	}
	if len(vers) == 0 || strings.ContainsAny(vers, stopChars) {
		return false
	}
	return true
}

// AppendVersions converts versions to a string and appends the string to the string base.
//
// Each VersionInfo will be converted to a string in the format of
// "product/version", where the "product" is get from the name field, while
// version is get from the version field. Several pieces of version information
// will be concatenated and separated by space.
//
// Example:
// AppendVersions("base", VersionInfo{"foo", "1.0"}, VersionInfo{"bar", "2.0"})
// results in "base foo/1.0 bar/2.0".
func AppendVersions(base string, versions ...VersionInfo) string {
	if len(versions) == 0 {
		return base
	}

	verstrs := make([]string, 0, 1+len(versions))
	if len(base) > 0 {
		verstrs = append(verstrs, base)
	}

	for _, v := range versions {
		if !v.isValid() {
			continue
		}
		verstrs = append(verstrs, v.Name+"/"+v.Version)
	}
	return strings.Join(verstrs, " ")
}
