//go:build macropad_rp2040 && !fs.none

package main

import (
	"machine"

	"tinygo.org/x/tinyfs"
	"tinygo.org/x/tinyfs/littlefs"
)

func init() {
	kbd.SetFS(initFilesystem())
}

func initFilesystem() tinyfs.Filesystem {
	lfs := littlefs.New(machine.Flash)
	lfs.Configure(&littlefs.Config{
		CacheSize:     512,
		LookaheadSize: 512,
		BlockCycles:   100,
	})
	return lfs
}
