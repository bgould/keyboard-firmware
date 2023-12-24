//go:build !tinygo

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bgould/keyboard-firmware/hosts/usbvial/vial"
	hid "github.com/sstallion/go-hid"
)

func main() {
	b := make([]byte, 65)

	// Initialize the hid package.
	if err := hid.Init(); err != nil {
		log.Fatal(err)
	}

	var found *hid.DeviceInfo
	hid.Enumerate(hid.VendorIDAny, hid.ProductIDAny, hid.EnumFunc(func(info *hid.DeviceInfo) error {
		buf, err := json.Marshal(info)
		if err != nil {
			println("error device:", err.Error())
			return nil
		}
		if strings.HasPrefix(info.SerialNbr, vial.MagicSerialPrefix) &&
			strings.Contains(info.ProductStr, "Advantage2") &&
			info.InterfaceNbr == 3 {
			println("found device:", string(buf))
			found = info
		}
		return nil
	}))

	if found == nil {
		log.Fatal("no matching devices found")
	}

	d, err := hid.OpenPath(found.Path)
	if err != nil {
		log.Fatal(err)
	}

	// Read the Manufacturer String.
	s, err := d.GetMfrStr()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Manufacturer String: %s\n", s)

	// Read the Product String.
	s, err = d.GetProductStr()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Product String: %s\n", s)

	// Read the Serial Number String.
	s, err = d.GetSerialNbr()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Serial Number String: %s\n", s)

	now := time.Now().Unix()
	b[0] = 0xEE
	b[1] = 0x01
	b[2] = byte(now >> 56)
	b[3] = byte(now >> 48)
	b[4] = byte(now >> 40)
	b[5] = byte(now >> 32)
	b[6] = byte(now >> 24)
	b[7] = byte(now >> 16)
	b[8] = byte(now >> 8)
	b[9] = byte(now >> 0)
	if _, err := d.Write(b); err != nil {
		log.Fatal(err)
	}

	// // Request state (cmd 0x81). The first byte is the report number (0x0).
	// b[0] = 0x3
	// b[1] = 0xFE
	// if _, err := d.Write(b); err != nil {
	// 	log.Fatal(err)
	// }
	// // Read Indexed String 1.
	// s, err = d.GetIndexedStr(1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Indexed String 1: %s\n", s)
}
