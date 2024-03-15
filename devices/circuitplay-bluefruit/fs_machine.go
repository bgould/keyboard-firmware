//go:build tinygo && !fs.qspi

package main

import (
	"machine"

	"tinygo.org/x/tinyfs/littlefs"
)

func initFilesystem() {
	blockdev = machine.Flash
	lfs := littlefs.New(blockdev)
	lfs.Configure(&littlefs.Config{
		CacheSize:     512,
		LookaheadSize: 512,
		BlockCycles:   100,
	})
	filesystem = lfs
}
