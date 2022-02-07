package keyboard

import (
	"strconv"
	"time"

	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

type MouseKeysConfig struct {
	MoveDelta     uint16
	MoveMaxSpeed  uint16
	MoveTimeToMax uint16

	WheelDelta     uint16
	WheelMaxSpeed  uint16
	WheelTimeToMax uint16

	Delay    time.Duration
	Interval time.Duration
}

// https://github.com/bgould/costar_tmk_keyboard/blob/master/config.h#L42-L49
func DefaultMouseKeysConfig() MouseKeysConfig {
	return MouseKeysConfig{
		Delay:          0,
		Interval:       20 * time.Millisecond,
		MoveDelta:      3,
		MoveMaxSpeed:   10,
		MoveTimeToMax:  20,
		WheelDelta:     1,
		WheelMaxSpeed:  16,
		WheelTimeToMax: 40,
	}
}

type MouseKeys struct {
	config MouseKeysConfig

	lastTimer time.Time
	repeat    uint16
	accel     uint16

	report struct {
		buttons uint8
		x       int8
		y       int8
		v       int8
		h       int8
	}

	debug bool
}

const (
	mousekeyMaxMove  = 127
	mousekeyMaxWheel = 127
)

func NewMouseKeys(config MouseKeysConfig) *MouseKeys {
	if config.MoveMaxSpeed > mousekeyMaxMove {
		config.MoveMaxSpeed = mousekeyMaxMove
	}
	if config.WheelMaxSpeed > mousekeyMaxWheel {
		config.WheelMaxSpeed = mousekeyMaxWheel
	}
	return &MouseKeys{
		config: config,
	}
}

func (mk *MouseKeys) Task(report *Report) bool {

	since := mk.config.Interval
	if mk.repeat == 0 {
		since = mk.config.Delay
	}
	if time.Since(mk.lastTimer) < since {
		return false
	}

	if mk.report.x == 0 && mk.report.y == 0 &&
		mk.report.v == 0 && mk.report.h == 0 {
		return false
	}

	if mk.repeat < 255 {
		mk.repeat++
	}

	if mk.report.x > 0 {
		mk.report.x = mk.moveUnit()
	}
	if mk.report.x < 0 {
		mk.report.x = int8(mk.moveUnit()) * -1
	}
	if mk.report.y > 0 {
		mk.report.y = mk.moveUnit()
	}
	if mk.report.y < 0 {
		mk.report.y = mk.moveUnit() * -1
	}

	// diagonal move [1/sqrt(2) = 0.7]
	if mk.report.x != 0 && mk.report.y != 0 {
		mk.report.x = int8(0.7 * float32(mk.report.x))
		mk.report.y = int8(0.7 * float32(mk.report.y))
	}

	if mk.report.x > 0 {
		mk.report.x = mk.moveUnit()
	}
	if mk.report.x < 0 {
		mk.report.x = int8(mk.moveUnit()) * -1
	}
	if mk.report.y > 0 {
		mk.report.y = mk.moveUnit()
	}
	if mk.report.y < 0 {
		mk.report.y = mk.moveUnit() * -1
	}

	report.Mouse(MouseButton(mk.report.buttons), mk.report.x, mk.report.y, mk.report.v, mk.report.h)
	return true

	/*
	   if (timer_elapsed(last_timer) < (mousekey_repeat ? mk_interval : mk_delay*10))
	       return;

	   if (mouse_report.x == 0 && mouse_report.y == 0 && mouse_report.v == 0 && mouse_report.h == 0)
	       return;

	   if (mousekey_repeat != UINT8_MAX)
	       mousekey_repeat++;


	   if (mouse_report.x > 0) mouse_report.x = move_unit();
	   if (mouse_report.x < 0) mouse_report.x = move_unit() * -1;
	   if (mouse_report.y > 0) mouse_report.y = move_unit();
	   if (mouse_report.y < 0) mouse_report.y = move_unit() * -1;

	   // diagonal move [1/sqrt(2) = 0.7]
	   if (mouse_report.x && mouse_report.y) {
	       mouse_report.x *= 0.7;
	       mouse_report.y *= 0.7;
	   }

	   if (mouse_report.v > 0) mouse_report.v = wheel_unit();
	   if (mouse_report.v < 0) mouse_report.v = wheel_unit() * -1;
	   if (mouse_report.h > 0) mouse_report.h = wheel_unit();
	   if (mouse_report.h < 0) mouse_report.h = wheel_unit() * -1;

	   mousekey_send();
	*/
}

func (mk *MouseKeys) Make(code keycodes.Keycode) {
	switch code {
	case keycodes.MS_UP:
		mk.report.y = mk.moveUnit() * -1
	case keycodes.MS_DOWN:
		mk.report.y = mk.moveUnit()
	case keycodes.MS_LEFT:
		mk.report.x = mk.moveUnit() * -1
	case keycodes.MS_RIGHT:
		mk.report.x = mk.moveUnit()
	case keycodes.MS_WH_UP:
		mk.report.v = mk.wheelUnit()
	case keycodes.MS_WH_DOWN:
		mk.report.v = mk.wheelUnit() * -1
	case keycodes.MS_WH_LEFT:
		mk.report.h = mk.wheelUnit() * -1
	case keycodes.MS_WH_RIGHT:
		mk.report.h = mk.wheelUnit()
	case keycodes.MS_BTN1:
		mk.report.buttons |= keycodes.BTN1
	case keycodes.MS_BTN2:
		mk.report.buttons |= keycodes.BTN2
	case keycodes.MS_BTN3:
		mk.report.buttons |= keycodes.BTN3
	case keycodes.MS_BTN4:
		mk.report.buttons |= keycodes.BTN4
	case keycodes.MS_BTN5:
		mk.report.buttons |= keycodes.BTN5
	case keycodes.MS_ACCEL0:
		mk.accel |= (1 << 0)
	case keycodes.MS_ACCEL1:
		mk.accel |= (1 << 1)
	case keycodes.MS_ACCEL2:
		mk.accel |= (1 << 2)
	}
}

func (mk *MouseKeys) Break(code keycodes.Keycode) {
	switch code {
	case keycodes.MS_UP:
		if mk.report.y < 0 {
			mk.report.y = 0
		}
	case keycodes.MS_DOWN:
		if mk.report.y > 0 {
			mk.report.y = 0
		}
	case keycodes.MS_LEFT:
		if mk.report.x < 0 {
			mk.report.x = 0
		}
	case keycodes.MS_RIGHT:
		if mk.report.x > 0 {
			mk.report.x = 0
		}
	case keycodes.MS_WH_UP:
		if mk.report.v > 0 {
			mk.report.v = 0
		}
	case keycodes.MS_WH_DOWN:
		if mk.report.v < 0 {
			mk.report.v = 0
		}
	case keycodes.MS_WH_LEFT:
		if mk.report.h < 0 {
			mk.report.h = 0
		}
	case keycodes.MS_WH_RIGHT:
		if mk.report.h > 0 {
			mk.report.h = 0
		}
	case keycodes.MS_BTN1:
		mk.report.buttons &= ^uint8(keycodes.BTN1)
	case keycodes.MS_BTN2:
		mk.report.buttons &= ^uint8(keycodes.BTN2)
	case keycodes.MS_BTN3:
		mk.report.buttons &= ^uint8(keycodes.BTN3)
	case keycodes.MS_BTN4:
		mk.report.buttons &= ^uint8(keycodes.BTN4)
	case keycodes.MS_BTN5:
		mk.report.buttons &= ^uint8(keycodes.BTN5)
	case keycodes.MS_ACCEL0:
		mk.accel &= ^uint16(1 << 0)
	case keycodes.MS_ACCEL1:
		mk.accel &= ^uint16(1 << 1)
	case keycodes.MS_ACCEL2:
		mk.accel &= ^uint16(1 << 2)
	}
	if mk.report.x == 0 && mk.report.y == 0 &&
		mk.report.v == 0 && mk.report.h == 0 {
		mk.repeat = 0
	}
}

func (mk *MouseKeys) moveUnit() int8 {
	var unit uint16
	if mk.accel&(1<<0) > 0 {
		unit = (mk.config.MoveDelta * mk.config.MoveMaxSpeed) / 4
	} else if mk.accel&(1<<1) > 0 {
		unit = (mk.config.MoveDelta * mk.config.MoveMaxSpeed) / 2
	} else if mk.accel&(1<<2) > 0 {
		unit = (mk.config.MoveDelta * mk.config.MoveMaxSpeed)
	} else if mk.repeat == 0 {
		unit = mk.config.MoveDelta
	} else if mk.repeat >= mk.config.MoveTimeToMax {
		unit = mk.config.MoveDelta * mk.config.MoveMaxSpeed
	} else {
		unit = (mk.config.MoveDelta * mk.config.MoveMaxSpeed * mk.repeat) / mk.config.MoveTimeToMax
	}
	if unit > mk.config.MoveMaxSpeed {
		return int8(mk.config.MoveMaxSpeed)
	} else if unit == 0 {
		return 1
	} else {
		return int8(unit)
	}
}

func (mk *MouseKeys) wheelUnit() int8 {
	var unit uint16
	if mk.accel&(1<<0) > 0 {
		unit = (mk.config.WheelDelta * mk.config.WheelMaxSpeed) / 4
	} else if mk.accel&(1<<1) > 0 {
		unit = (mk.config.WheelDelta * mk.config.WheelMaxSpeed) / 2
	} else if mk.accel&(1<<2) > 0 {
		unit = (mk.config.WheelDelta * mk.config.WheelMaxSpeed)
	} else if mk.repeat == 0 {
		unit = mk.config.WheelDelta
	} else if mk.repeat >= mk.config.WheelTimeToMax {
		unit = mk.config.WheelDelta * mk.config.WheelMaxSpeed
	} else {
		unit = (mk.config.WheelDelta * mk.config.WheelMaxSpeed * mk.repeat) / mk.config.WheelTimeToMax
	}
	if unit > mk.config.WheelMaxSpeed {
		return int8(mk.config.WheelMaxSpeed)
	} else if unit == 0 {
		return 1
	} else {
		return int8(unit)
	}
	// return (unit > mk.WheelMaxSpeed ? mk.WheelMaxSpeed : (unit == 0 ? 1 : unit));
}

func (mk *MouseKeys) Debug() string {
	return "mousekey [btn|x y v h](rep/acl): [" +
		hex(mk.report.buttons) + "|" +
		strconv.FormatInt(int64(mk.report.x), 10) + " " +
		strconv.FormatInt(int64(mk.report.y), 10) + " " +
		strconv.FormatInt(int64(mk.report.v), 10) + " " +
		strconv.FormatInt(int64(mk.report.h), 10) + "](" +
		strconv.FormatInt(int64(mk.repeat), 10) + "/" +
		strconv.FormatInt(int64(mk.accel), 10) + ")\n"
	// print_decs(mouse_report.x); print(" ");
	// print_decs(mouse_report.y); print(" ");
	// print_decs(mouse_report.v); print(" ");
	// print_decs(mouse_report.h); print("](");
	// print_dec(mousekey_repeat); print("/");
	// print_dec(mousekey_accel); print(")\n");
}

/*
#define MOUSEKEY_MOVE_DELTA           3
#define mk.config.WheelDelta          1
#define MOUSEKEY_DELAY                0
#define MOUSEKEY_INTERVAL            20
#define MOUSEKEY_MAX_SPEED           10
#define MOUSEKEY_TIME_TO_MAX         20
#define mk.WheelMaxSpeed_SPEED     16
#define MOUSEKEY_WHEEL_TIME_TO_MAX   40

void mousekey_task(void);
void mousekey_on(uint8_t code);
void mousekey_off(uint8_t code);
void mousekey_clear(void);
void mousekey_send(void);
uint8_t mousekey_buttons(void);

static report_mouse_t mouse_report = {};
static uint8_t mousekey_repeat =  0;
static uint8_t mousekey_accel = 0;

static void mousekey_debug(void);


/*
 * Mouse keys  acceleration algorithm
 *  http://en.wikipedia.org/wiki/Mouse_keys
 *
 *  speed = delta * max_speed * (repeat / time_to_max)**((1000+curve)/1000)
// milliseconds between the initial key press and first repeated motion event (0-2550)
uint8_t mk_delay = MOUSEKEY_DELAY/10;
// milliseconds between repeated motion events (0-255)
uint8_t mk_interval = MOUSEKEY_INTERVAL;
// steady speed (in action_delta units) applied each event (0-255)
uint8_t mk.MoveMaxSpeed = MOUSEKEY_MAX_SPEED;
// number of events (count) accelerating to steady speed (0-255)
uint8_t mk_time_to_max = MOUSEKEY_TIME_TO_MAX;
// ramp used to reach maximum pointer speed (NOT SUPPORTED)
//int8_t mk_curve = 0;
// wheel params
uint8_t mk.WheelMaxSpeed = mk.WheelMaxSpeed_SPEED;
uint8_t mk_wheel_time_to_max = MOUSEKEY_WHEEL_TIME_TO_MAX;

716 831 1800 ext 2227
static uint16_t last_timer = 0;


static uint8_t move_unit(void)
{
    uint16_t unit;
    if (mousekey_accel & (1<<0)) {
        unit = (MOUSEKEY_MOVE_DELTA * mk.MoveMaxSpeed)/4;
    } else if (mousekey_accel & (1<<1)) {
        unit = (MOUSEKEY_MOVE_DELTA * mk.MoveMaxSpeed)/2;
    } else if (mousekey_accel & (1<<2)) {
        unit = (MOUSEKEY_MOVE_DELTA * mk.MoveMaxSpeed);
    } else if (mousekey_repeat == 0) {
        unit = MOUSEKEY_MOVE_DELTA;
    } else if (mousekey_repeat >= mk_time_to_max) {
        unit = MOUSEKEY_MOVE_DELTA * mk.MoveMaxSpeed;
    } else {
        unit = (MOUSEKEY_MOVE_DELTA * mk.MoveMaxSpeed * mousekey_repeat) / mk_time_to_max;
    }
    return (unit > MOUSEKEY_MOVE_MAX ? MOUSEKEY_MOVE_MAX : (unit == 0 ? 1 : unit));
}

static uint8_t wheel_unit(void)
{
    uint16_t unit;
    if (mousekey_accel & (1<<0)) {
        unit = (mk.config.WheelDelta * mk.WheelMaxSpeed)/4;
    } else if (mousekey_accel & (1<<1)) {
        unit = (mk.config.WheelDelta * mk.WheelMaxSpeed)/2;
    } else if (mousekey_accel & (1<<2)) {
        unit = (mk.config.WheelDelta * mk.WheelMaxSpeed);
    } else if (mousekey_repeat == 0) {
        unit = mk.config.WheelDelta;
    } else if (mousekey_repeat >= mk_wheel_time_to_max) {
        unit = mk.config.WheelDelta * mk.WheelMaxSpeed;
    } else {
        unit = (mk.config.WheelDelta * mk.WheelMaxSpeed * mousekey_repeat) / mk_wheel_time_to_max;
    }
    return (unit > mk.WheelMaxSpeed ? mk.WheelMaxSpeed : (unit == 0 ? 1 : unit));
}

void mousekey_task(void)
{
    if (timer_elapsed(last_timer) < (mousekey_repeat ? mk_interval : mk_delay*10))
        return;

    if (mouse_report.x == 0 && mouse_report.y == 0 && mouse_report.v == 0 && mouse_report.h == 0)
        return;

    if (mousekey_repeat != UINT8_MAX)
        mousekey_repeat++;


    if (mouse_report.x > 0) mouse_report.x = move_unit();
    if (mouse_report.x < 0) mouse_report.x = move_unit() * -1;
    if (mouse_report.y > 0) mouse_report.y = move_unit();
    if (mouse_report.y < 0) mouse_report.y = move_unit() * -1;

    // diagonal move [1/sqrt(2) = 0.7]
    if (mouse_report.x && mouse_report.y) {
        mouse_report.x *= 0.7;
        mouse_report.y *= 0.7;
    }

    if (mouse_report.v > 0) mouse_report.v = wheel_unit();
    if (mouse_report.v < 0) mouse_report.v = wheel_unit() * -1;
    if (mouse_report.h > 0) mouse_report.h = wheel_unit();
    if (mouse_report.h < 0) mouse_report.h = wheel_unit() * -1;

    mousekey_send();
}

void mousekey_on(uint8_t code)
{
    if      (code == KC_MS_UP)       mouse_report.y = move_unit() * -1;
    else if (code == KC_MS_DOWN)     mouse_report.y = move_unit();
    else if (code == KC_MS_LEFT)     mouse_report.x = move_unit() * -1;
    else if (code == KC_MS_RIGHT)    mouse_report.x = move_unit();
    else if (code == KC_MS_WH_UP)    mouse_report.v = wheel_unit();
    else if (code == KC_MS_WH_DOWN)  mouse_report.v = wheel_unit() * -1;
    else if (code == KC_MS_WH_LEFT)  mouse_report.h = wheel_unit() * -1;
    else if (code == KC_MS_WH_RIGHT) mouse_report.h = wheel_unit();
    else if (code == KC_MS_BTN1)     mouse_report.buttons |= MOUSE_BTN1;
    else if (code == KC_MS_BTN2)     mouse_report.buttons |= MOUSE_BTN2;
    else if (code == KC_MS_BTN3)     mouse_report.buttons |= MOUSE_BTN3;
    else if (code == KC_MS_BTN4)     mouse_report.buttons |= MOUSE_BTN4;
    else if (code == KC_MS_BTN5)     mouse_report.buttons |= MOUSE_BTN5;
    else if (code == KC_MS_ACCEL0)   mousekey_accel |= (1<<0);
    else if (code == KC_MS_ACCEL1)   mousekey_accel |= (1<<1);
    else if (code == KC_MS_ACCEL2)   mousekey_accel |= (1<<2);
}

void mousekey_off(uint8_t code)
{
    if      (code == KC_MS_UP       && mouse_report.y < 0) mouse_report.y = 0;
    else if (code == KC_MS_DOWN     && mouse_report.y > 0) mouse_report.y = 0;
    else if (code == KC_MS_LEFT     && mouse_report.x < 0) mouse_report.x = 0;
    else if (code == KC_MS_RIGHT    && mouse_report.x > 0) mouse_report.x = 0;
    else if (code == KC_MS_WH_UP    && mouse_report.v > 0) mouse_report.v = 0;
    else if (code == KC_MS_WH_DOWN  && mouse_report.v < 0) mouse_report.v = 0;
    else if (code == KC_MS_WH_LEFT  && mouse_report.h < 0) mouse_report.h = 0;
    else if (code == KC_MS_WH_RIGHT && mouse_report.h > 0) mouse_report.h = 0;
    else if (code == KC_MS_BTN1) mouse_report.buttons &= ~MOUSE_BTN1;
    else if (code == KC_MS_BTN2) mouse_report.buttons &= ~MOUSE_BTN2;
    else if (code == KC_MS_BTN3) mouse_report.buttons &= ~MOUSE_BTN3;
    else if (code == KC_MS_BTN4) mouse_report.buttons &= ~MOUSE_BTN4;
    else if (code == KC_MS_BTN5) mouse_report.buttons &= ~MOUSE_BTN5;
    else if (code == KC_MS_ACCEL0) mousekey_accel &= ~(1<<0);
    else if (code == KC_MS_ACCEL1) mousekey_accel &= ~(1<<1);
    else if (code == KC_MS_ACCEL2) mousekey_accel &= ~(1<<2);

    if (mouse_report.x == 0 && mouse_report.y == 0 && mouse_report.v == 0 && mouse_report.h == 0)
        mousekey_repeat = 0;
}

void mousekey_send(void)
{
    mousekey_debug();
    host_mouse_send(&mouse_report);
    last_timer = timer_read();
}

void mousekey_clear(void)
{
    mouse_report = (report_mouse_t){};
    mousekey_repeat = 0;
    mousekey_accel = 0;
}

uint8_t mousekey_buttons(void)
{
    return mouse_report.buttons;
}

static void mousekey_debug(void)
{
    if (!debug_mouse) return;
    print("mousekey [btn|x y v h](rep/acl): [");
    phex(mouse_report.buttons); print("|");
    print_decs(mouse_report.x); print(" ");
    print_decs(mouse_report.y); print(" ");
    print_decs(mouse_report.v); print(" ");
    print_decs(mouse_report.h); print("](");
    print_dec(mousekey_repeat); print("/");
    print_dec(mousekey_accel); print(")\n");
}
*/
