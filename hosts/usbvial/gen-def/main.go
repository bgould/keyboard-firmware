package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bgould/keyboard-firmware/hosts/usbvial/vial"
	"github.com/itchio/lzma"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("this program needs vial.json's path")
	}
	// fmt.Println(os.Args[1])

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	r, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	def := &vial.DeviceDefinition{}
	if err := json.Unmarshal(r, def); err != nil {
		log.Fatal(err)
	}

	// compress
	var tBuf bytes.Buffer
	w := lzma.NewWriter(&tBuf)
	_, err = w.Write(r)
	if err != nil {
		fmt.Printf("lzma.Write() failed: %s", err.Error())
	}
	w.Close()

	var oBuf strings.Builder
	oBuf.WriteString("// This file is auto-generated by github.com/bgould/keyboard-firmware/hosts/usbvial/gen-def, do not edit.")
	oBuf.WriteString("\n")
	oBuf.WriteString("package main\n")
	oBuf.WriteString("\n")
	oBuf.WriteString("import \"github.com/bgould/keyboard-firmware/hosts/usbvial/vial\"\n")
	oBuf.WriteString("\n")
	oBuf.WriteString("var VialDeviceDefinition = vial.DeviceDefinition{\n")
	oBuf.WriteString("	Name: \"" + template.JSEscapeString(def.Name) + "\",\n")
	oBuf.WriteString("	VendorID: \"" + template.JSEscapeString(def.VendorID) + "\",\n")
	oBuf.WriteString("	ProductID: \"" + template.JSEscapeString(def.ProductID) + "\",\n")
	oBuf.WriteString("	Matrix: vial.DeviceMatrix{\n")
	oBuf.WriteString(fmt.Sprintf("		Rows: %d,\n", def.Matrix.Rows))
	oBuf.WriteString(fmt.Sprintf("		Cols: %d,\n", def.Matrix.Cols))
	oBuf.WriteString("  },\n")
	oBuf.WriteString("\tLzmaDefLength: uint16(len(vialLzmaDefinition)),\n")
	oBuf.WriteString("\tLzmaDefWriter: vial.LzmaDefPageWriterFunc(vialWriteLzmaDefPage),\n")
	// oBuf.WriteString("\tLzmaDef: []byte{\n")
	// oBuf.WriteString("\t\t")
	// for i, b := range tBuf.Bytes() {
	// 	if i == (len(tBuf.Bytes()) - 1) {
	// 		oBuf.WriteString(fmt.Sprintf("0x%02X,", b))
	// 	} else {
	// 		oBuf.WriteString(fmt.Sprintf("0x%02X, ", b))
	// 	}
	// }
	// oBuf.WriteString("\n")
	// oBuf.WriteString("\t},\n")
	oBuf.WriteString("}\n")
	oBuf.WriteString("\n")
	oBuf.WriteString("var vialLzmaDefinition = []byte{\n")
	oBuf.WriteString("\t")
	for i, b := range tBuf.Bytes() {
		if i == (len(tBuf.Bytes()) - 1) {
			oBuf.WriteString(fmt.Sprintf("0x%02X,", b))
		} else {
			oBuf.WriteString(fmt.Sprintf("0x%02X, ", b))
		}
	}
	oBuf.WriteString("\n")
	oBuf.WriteString("}\n")
	oBuf.WriteString("\n")

	oBuf.WriteString(`
func vialWriteLzmaDefPage(tx []byte, page uint16) bool {
	start := page * 32
	end := start + 32
	len := uint16(len(vialLzmaDefinition))
	if end < start || start >= len {
		return false
	}
	if end > len {
		end = len
	}
	copy(tx[:32], vialLzmaDefinition[start:end])
	return true
}
	`)
	oBuf.WriteString("\n")

	formatted, err := format.Source([]byte(oBuf.String()))
	if err != nil {
		log.Fatal(err)
	}

	outPath := filepath.Join(filepath.Dir(os.Args[1]), `vial-def.go`)
	// fmt.Println(outPath)
	err = os.WriteFile(outPath, formatted, 0666)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(tBuf.Bytes())
}
