// Copyright 2017 syzkaller project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

//go:build freebsd || netbsd || openbsd

package osutil

import (
	"os"
	"os/exec"
	"time"
)

func creationTime(fi os.FileInfo) time.Time {
	return time.Time{}
}

func RemoveAll(dir string) error {
	return os.RemoveAll(dir)
}

func SystemMemorySize() uint64 {
	return 0
}

func prolongPipe(r, w *os.File) {
}

func Sandbox(cmd *exec.Cmd, user, net bool) error {
	return nil
}

func SandboxChown(file string) error {
	return nil
}

func setPdeathsig(cmd *exec.Cmd, hardKill bool) {
}

func killPgroup(cmd *exec.Cmd) {
}
