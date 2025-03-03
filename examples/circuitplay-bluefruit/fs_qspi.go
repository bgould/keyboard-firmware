//go:build circuitplay_bluefruit && !fs.machine

package main

import (
	"machine"

	"tinygo.org/x/drivers/flash"
	"tinygo.org/x/tinyfs"
	"tinygo.org/x/tinyfs/littlefs"
)

func init() {
	kbd.SetFS(initFilesystem())
}

func initFilesystem() tinyfs.Filesystem {
	flashdev := flash.NewSPI(
		&machine.SPI0,
		machine.SPI0_SDO_PIN,
		machine.SPI0_SDI_PIN,
		machine.SPI0_SCK_PIN,
		machine.P0_15,
	)
	flashdev.Configure(&flash.DeviceConfig{})
	blockdev := flashdev
	lfs := littlefs.New(blockdev)
	lfs.Configure(&littlefs.Config{
		CacheSize:     512,
		LookaheadSize: 512,
		BlockCycles:   100,
	})
	return lfs
}
