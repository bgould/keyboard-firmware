package keycodes

// Generate keycodes.go from QMK definitions; requires github.com/bgould/qmk-firmware:keycodes-generate-go
// -------------------------------------------------------------------------------------------------------
//go:generate bash -c "rm -f keycodes_qmk.go && qmk generate-keycodes-go -v latest -o keycodes_qmk.go && go fmt keycodes_qmk.go"
