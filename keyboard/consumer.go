package keyboard

import (
	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

type ConsumerKey uint16

const (

	// USB HID Specifications
	// https://www.usb.org/hid#approved-usage-table-review-requests

	// Consumer Page(0x0C)
	// https://github.com/tmk/tmk_keyboard/issues/370
	consumer_AUDIO_MUTE           = 0x00E2
	consumer_AUDIO_VOL_UP         = 0x00E9
	consumer_AUDIO_VOL_DOWN       = 0x00EA
	consumer_TRANSPORT_NEXT_TRACK = 0x00B5
	consumer_TRANSPORT_PREV_TRACK = 0x00B6
	consumer_TRANSPORT_STOP       = 0x00B7
	consumer_TRANSPORT_STOP_EJECT = 0x00CC
	consumer_TRANSPORT_PLAY_PAUSE = 0x00CD
	// application launch
	consumer_APPLAUNCH_CC_CONFIG     = 0x0183
	consumer_APPLAUNCH_EMAIL         = 0x018A
	consumer_APPLAUNCH_CALCULATOR    = 0x0192
	consumer_APPLAUNCH_LOCAL_BROWSER = 0x0194
	// application control
	consumer_APPCONTROL_SEARCH    = 0x0221
	consumer_APPCONTROL_HOME      = 0x0223
	consumer_APPCONTROL_BACK      = 0x0224
	consumer_APPCONTROL_FORWARD   = 0x0225
	consumer_APPCONTROL_STOP      = 0x0226
	consumer_APPCONTROL_REFRESH   = 0x0227
	consumer_APPCONTROL_BOOKMARKS = 0x022A
	// supplement for Bluegiga iWRAP HID(not supported by Windows?)
	consumer_APPLAUNCH_LOCK         = 0x019E
	consumer_TRANSPORT_RECORD       = 0x00B2
	consumer_TRANSPORT_FAST_FORWARD = 0x00B3
	consumer_TRANSPORT_REWIND       = 0x00B4
	consumer_TRANSPORT_EJECT        = 0x00B8
	consumer_APPCONTROL_MINIMIZE    = 0x0206
	// Display Brightness Controls  https://www.usb.org/sites/default/files/hutrr41_0.pdf */
	consumer_BRIGHTNESS_INCREMENT = 0x006F
	consumer_BRIGHTNESS_DECREMENT = 0x0070
)

func keycode2consumer(keycode keycodes.Keycode) ConsumerKey {
	switch keycode {
	case keycodes.AUDIO_MUTE:
		return consumer_AUDIO_MUTE
	case keycodes.AUDIO_VOL_UP:
		return consumer_AUDIO_VOL_UP
	case keycodes.AUDIO_VOL_DOWN:
		return consumer_AUDIO_VOL_DOWN
	case keycodes.MEDIA_NEXT_TRACK:
		return consumer_TRANSPORT_NEXT_TRACK
	case keycodes.MEDIA_PREV_TRACK:
		return consumer_TRANSPORT_PREV_TRACK
	case keycodes.MEDIA_FAST_FORWARD:
		return consumer_TRANSPORT_FAST_FORWARD
	case keycodes.MEDIA_REWIND:
		return consumer_TRANSPORT_REWIND
	case keycodes.MEDIA_STOP:
		return consumer_TRANSPORT_STOP
	case keycodes.MEDIA_EJECT:
		return consumer_TRANSPORT_STOP_EJECT
	case keycodes.MEDIA_PLAY_PAUSE:
		return consumer_TRANSPORT_PLAY_PAUSE
	case keycodes.MEDIA_SELECT:
		return consumer_APPLAUNCH_CC_CONFIG
	case keycodes.MAIL:
		return consumer_APPLAUNCH_EMAIL
	case keycodes.CALCULATOR:
		return consumer_APPLAUNCH_CALCULATOR
	case keycodes.MY_COMPUTER:
		return consumer_APPLAUNCH_LOCAL_BROWSER
	case keycodes.WWW_SEARCH:
		return consumer_APPCONTROL_SEARCH
	case keycodes.WWW_HOME:
		return consumer_APPCONTROL_HOME
	case keycodes.WWW_BACK:
		return consumer_APPCONTROL_BACK
	case keycodes.WWW_FORWARD:
		return consumer_APPCONTROL_FORWARD
	case keycodes.WWW_STOP:
		return consumer_APPCONTROL_STOP
	case keycodes.WWW_REFRESH:
		return consumer_APPCONTROL_REFRESH
	case keycodes.WWW_FAVORITES:
		return consumer_APPCONTROL_BOOKMARKS
		/* FIXME
		case keycodes.BRIGHTNESS_INCREMENT:
			return consumer_BRIGHTNESS_INCREMENT
		case keycodes.BRIGHTNESS_DEC:
			return consumer_BRIGHTNESS_DECREMENT
		*/
	default:
		return 0
	}
}
