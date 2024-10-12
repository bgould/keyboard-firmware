Go Keyboard Firmware framework
==============================

This is a framework for building USB HID type devices such as keyboards with [TinyGo][tinygo].

Architecture
------------

The architecture of the firmware code is loosely based on concepts from the excellent [TMK Keyboard][tmk] library.
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

Design Considerations
---------------------

 * `import`-able as an idiomatic Go module; no Makefiles etc necessary to use core functionality and most use cases.
 * Usable from both TinyGo and full-sized Go... all TinyGo-specific packages (machine, etc.) should be
   protected with build tags. Conversely ... functionality not available in TinyGo must be protected as well.
 * No/few allocations in core library ... user code should initialize memory when necessary.

Building Firmware
-----------------

In order to build firmware in this repository, please install [Go][golang] 1.22+ and [TinyGo][tinygo] 0.33+.

*Note: for Teensy 4.1 targets, please see README.md in the `devices/kint41` folder.*

License
-----------------------

Copyright © 2019-2023 Benjamin Gould

Licensed GPLv3 or later.

[golang]: https://golang.org/
[tinygo]: https://tinygo.org/
[tmk]: https://github.com/tmk/tmk_keyboard
[tmk-architecture]: https://github.com/tmk/tmk_keyboard/tree/master/tmk_core#architecture
[build-from-source]: https://tinygo.org/docs/guides/build/
