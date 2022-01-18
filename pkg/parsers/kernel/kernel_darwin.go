//go:build darwin
// +build darwin

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

// Package kernel provides helper function to get, parse and compare kernel
// versions for different platforms.

import (
	"fmt"
	"os/exec"
	"strings"
)

// GetKernelVersion gets the current kernel version.
func GetKernelVersion() (*VersionInfo, error) {
	osName, err := getSPSoftwareDataType()
	if err != nil {
		return nil, err
	}
	release, err := getRelease(osName)
	if err != nil {
		return nil, err
	}
	return ParseRelease(release)
}

// getRelease uses `system_profiler SPSoftwareDataType` to get OSX kernel version
func getRelease(osName string) (string, error) {
	var release string
	data := strings.Split(osName, "\n")
	for _, line := range data {
		if !strings.Contains(line, "Kernel Version") {
			continue
		}
		// It has the format like '      Kernel Version: Darwin 14.5.0'
		content := strings.SplitN(line, ":", 2)
		if len(content) != 2 {
			return "", fmt.Errorf("Kernel Version is invalid")
		}

		prettyNames := strings.SplitN(strings.TrimSpace(content[1]), " ", 2)

		if len(prettyNames) != 2 {
			return "", fmt.Errorf("Kernel Version needs to be 'Darwin x.x.x' ")
		}
		release = prettyNames[1]
	}

	return release, nil
}

func getSPSoftwareDataType() (string, error) {
	cmd := exec.Command("system_profiler", "SPSoftwareDataType")
	osName, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(osName), nil
}
