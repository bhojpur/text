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

	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
)

// VersionInfo holds information about the kernel.
type VersionInfo struct {
	kvi   string // Version of the kernel (e.g. 6.1.7601.17592 -> 6)
	major int    // Major part of the kernel version (e.g. 6.1.7601.17592 -> 1)
	minor int    // Minor part of the kernel version (e.g. 6.1.7601.17592 -> 7601)
	build int    // Build number of the kernel version (e.g. 6.1.7601.17592 -> 17592)
}

func (k *VersionInfo) String() string {
	return fmt.Sprintf("%d.%d %d (%s)", k.major, k.minor, k.build, k.kvi)
}

// GetKernelVersion gets the current kernel version.
func GetKernelVersion() (*VersionInfo, error) {

	KVI := &VersionInfo{"Unknown", 0, 0, 0}

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		return KVI, err
	}
	defer k.Close()

	blex, _, err := k.GetStringValue("BuildLabEx")
	if err != nil {
		return KVI, err
	}
	KVI.kvi = blex

	dwVersion, err := windows.GetVersion()
	if err != nil {
		return KVI, err
	}

	KVI.major = int(dwVersion & 0xFF)
	KVI.minor = int((dwVersion & 0xFF00) >> 8)
	KVI.build = int((dwVersion & 0xFFFF0000) >> 16)

	return KVI, nil
}
