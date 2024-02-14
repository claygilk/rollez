//go:build linux
// +build linux

package keys

import (
	"golang.org/x/sys/unix"
)

const (
	ioctlReadTermios  = unix.TCGETS
	ioctlWriteTermios = unix.TCSETS
)
