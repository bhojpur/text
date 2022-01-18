package operatingsystem

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

	"github.com/Microsoft/hcsshim/osversion"
	"golang.org/x/sys/windows/registry"
)

// GetOperatingSystem gets the name of the current operating system.
func GetOperatingSystem() (string, error) {
	os, err := withCurrentVersionRegistryKey(func(key registry.Key) (os string, err error) {
		if os, _, err = key.GetStringValue("ProductName"); err != nil {
			return "", err
		}

		releaseId, _, err := key.GetStringValue("ReleaseId")
		if err != nil {
			return
		}
		os = fmt.Sprintf("%s Version %s", os, releaseId)

		buildNumber, _, err := key.GetStringValue("CurrentBuildNumber")
		if err != nil {
			return
		}
		ubr, _, err := key.GetIntegerValue("UBR")
		if err != nil {
			return
		}
		os = fmt.Sprintf("%s (OS Build %s.%d)", os, buildNumber, ubr)

		return
	})

	if os == "" {
		// Default return value
		os = "Unknown Operating System"
	}

	return os, err
}

func withCurrentVersionRegistryKey(f func(registry.Key) (string, error)) (string, error) {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer key.Close()
	return f(key)
}

// GetOperatingSystemVersion gets the version of the current operating system, as a string.
func GetOperatingSystemVersion() (string, error) {
	version := osversion.Get()
	return fmt.Sprintf("%d.%d.%d", version.MajorVersion, version.MinorVersion, version.Build), nil
}

// IsContainerized returns true if we are running inside a container.
// No-op on Windows, always returns false.
func IsContainerized() (bool, error) {
	return false, nil
}
