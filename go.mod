module github.com/bgould/keyboard-firmware

go 1.17

// replace tinygo.org/x/drivers => ../../drivers
// replace tinygo.org/x/tinyterm => ../../tinyterm

require (
	tinygo.org/x/drivers v0.19.1-0.20220220191646-34c498b3bbda
	tinygo.org/x/tinydraw v0.0.0-20220125063109-43cae6615eb5
	tinygo.org/x/tinyfont v0.2.1
	tinygo.org/x/tinyterm v0.1.1-0.20220222045623-7f7de77c9c8e
)
