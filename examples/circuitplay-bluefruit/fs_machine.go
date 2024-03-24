//go:build circuitplay_bluefruit && !fs.qspi

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
	blockdev := machine.Flash
	lfs := littlefs.New(blockdev)
	lfs.Configure(&littlefs.Config{
		CacheSize:     512,
		LookaheadSize: 512,
		BlockCycles:   100,
	})
	return lfs
}
