module github.com/bgould/keyboard-firmware

go 1.18

// replace tinygo.org/x/drivers => ../../drivers
// replace tinygo.org/x/tinyterm => ../../tinyterm

require (
	tinygo.org/x/drivers v0.23.1-0.20221212024746-491b3438ce72
	tinygo.org/x/tinydraw v0.3.0
	tinygo.org/x/tinyterm v0.1.0
)

require tinygo.org/x/tinyfont v0.3.0 // indirect
