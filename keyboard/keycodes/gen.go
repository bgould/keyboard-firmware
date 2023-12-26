package keycodes

//go:generate bash -c "rm -f keycodes_qmk.go && qmk generate-keycodes-go -v latest -o keycodes_qmk.go && go fmt keycodes_qmk.go"
