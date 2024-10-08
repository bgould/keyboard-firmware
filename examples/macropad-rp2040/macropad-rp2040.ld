MEMORY
{
    /* Reserve exactly 256 bytes at start of flash for second stage bootloader */
    BOOT2_TEXT (rx) : ORIGIN = 0x10000000, LENGTH = 256
    FLASH_TEXT (rx) : ORIGIN = 0x10000000 + 256, LENGTH = __flash_size - 256
    /* Application data */
    FLASH_APP  (rw) : ORIGIN = ORIGIN(FLASH_TEXT) + 0x200000, LENGTH = __flash_size - 256 - 0x200000
    RAM (rwx)       : ORIGIN = 0x20000000, LENGTH = 256k
}

_stack_size = 4K;

SECTIONS
{
    /* Second stage bootloader is prepended to the image. It must be 256 bytes
       and checksummed.  The gap to the checksum is zero-padded.
    */
    .boot2 : {
        __boot2_start__ = .;
        KEEP (*(.boot2));

        /* Explicitly allocate space for CRC32 checksum at end of second stage
           bootloader
        */
        . = __boot2_start__ + 256 - 4;
        LONG(0)
    } > BOOT2_TEXT = 0x0

    /* The second stage will always enter the image at the start of .text.
       The debugger will use the ELF entry point, which is the _entry_point
       symbol if present, otherwise defaults to start of .text.
       This can be used to transfer control back to the bootrom on debugger
       launches only, to perform proper flash setup.
    */
}


INCLUDE "targets/arm.ld"

/* For the flash API */
__flash_data_start = ORIGIN(FLASH_APP);
__flash_data_end = ORIGIN(FLASH_APP) + LENGTH(FLASH_APP);

/*

$ tinygo flash -target tinygo.target.json -tags keymap.custom -size=short -monitor examples/flash

Connected to COM9. Press Ctrl-C to exit.
Flash data start:       0x10201000
Flash data end:         0x10800000
Flash data size, bytes: 6287360
Flash write block size: 256
Flash erase block size: 4096

*/