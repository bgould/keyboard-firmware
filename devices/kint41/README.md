Building Firmware for Teensy 4.x
================================

USB HID support for Teensy 4.x in TinyGo is in an *alpha* state. To compile firmware with USB HID support, build the version of TinyGo
found at https://github.com/bgould/tinygo/tree/usb-common-keyboard-firmware according to the [instructions][build-from-source]
on the TinyGo website. If not utilizing USB HID support, mainline TinyGo v0.21.0+ can be used.

There is a folder called `devices` containing the main package with hardware-specific glue code for each implemented device.
To build one of the example devices, use a command such as:

    tinygo flash -target=teensy40 -serial=uart ./devices/four_button

For the four_button example, switching between USB HID can be enabled/disabled by using -serial=uart or -serial=usb respectively.
When USB HID is disabled, text representation of HID reports will be output over the serial connection.