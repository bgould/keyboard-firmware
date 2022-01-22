Go Keyboard Firmware framework
==============================

This is an experimental project that I am using to test out and iterate on the USB HID implementation being developed
for [TinyGo][tinygo]. Once things become more stable I may add features and clean up the codebase to be more
general-purpose and maintain it as a library.

Architecture
------------

The architecture of the firmware code is loosely based on concepts from the excellent [TMK Keyboard][tmk] libary.
The high-level abstractions in the codebase are the same as [defined by TMK][tmk-architecture], including:

 * Device (actually called "Keyboard" in the context of TMK) - often this is a physical keyboard, but also could
   could be something like a protocol converter, logger, re-mapper, etc. It is comprised of a specific Host, Matrix,
   and Keymap implementation. At runtime the Device maintains state such as the status of LEDs, active Layer, etc.
 * Host - represents the hardware and protocol used to communicate with the host to which the device is connected;
   typically this may be something like USB HID, Bluetooth, PS/2, Serial, etc.
 * Matrix - contains the state of the device's switches/inputs, organized in memory as a matrix of rows and columns.
   For a keyboard, this may represent an actual matrix of rows and columns to which physical switches are connected.
   For a converter, the rows/columns are virtual and are used to maintain state based on events from the peripheral.
 * Keymap - user-configurable, ordered set of one or more Layers for translating the state of the matrix into events
   and/or keycodes to be sent to the host. For keymaps with multiple layers, the Device keeps track of the "active"
   layer. When translating a matrix location to a keycode, if a layer does not define a mapping for a matrix location,
   any layers below the active layer are searched in order (from highest to lowest layer) until a mapping is found.
 * Layer - a mapping of matrix locations (row/column) to specific keycodes or framework events. A layer does not need
   to define a mapping for every matrix location; some locations can be "transparent" to delegate to a lower-level
   layer as defined by the keymap.

*Please note that not all protocols and features described above, such as multiple layers, are implemented yet*

Building Firmware
-----------------

USB HID support in TinyGo is in an *alpha* state. To compile firmware with USB HID support, build the version of TinyGo
found at https://github.com/bgould/tinygo/tree/usb-common-keyboard-firmware according to the [instructions][build-from-source]
on the TinyGo website. If not utilizing USB HID support, mainline TinyGo v0.21.0+ can be used.

* The build setup described below is temporary, and will change before this framework is ready for general use *

There is a folder called `devices` containing the main package with hardware-specific glue code for each implemented device.
To build one of the example devices, use a command such as:

    tinygo flash -target=teensy40 -serial=uart ./devices/four_button

For the four_button example, switching between USB HID can be enabled/disabled by using -serial=uart or -serial=usb respectively.
When USB HID is disabled, text representation of HID reports will be output over the serial connection.

License
-----------------------

Copyright © 2019-2022 Benjamin Gould

Licensed GPLv3 or later. May consider re-licensing to 3-clause BSD at a later time.

[tinygo]: https://tinygo.org/
[tmk]: https://github.com/tmk/tmk_keyboard
[tmk-architecture]: https://github.com/tmk/tmk_keyboard/tree/master/tmk_core#architecture
[build-from-source]: https://tinygo.org/docs/guides/build/