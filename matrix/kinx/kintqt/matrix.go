package kintqt

import (
	"github.com/bgould/keyboard-firmware/keyboard"
	"tinygo.org/x/drivers"
	"tinygo.org/x/drivers/mcp23017"
)

const (
	NumRows = 15
	NumCols = 7
)

func (a *Adapter) NewMatrix() *keyboard.Matrix {
	return keyboard.NewMatrix(NumRows, NumCols, a)
}

func NewAdapter(bus drivers.I2C) *Adapter {
	return &Adapter{bus: bus}
}

// Adapter
type Adapter struct {

	// I2C bus that the port expanders are attached to
	bus drivers.I2C

	// port0 is the device with the column inputs and extra GPIOs
	port0 *mcp23017.Device

	// port1 is the device with outputs for the rows
	port1 *mcp23017.Device

	// we need to keep track of the state of the LEDs
	// because we are using some of the GPIOs from the
	// port expanders to control the LEDs on the kinT
	ledState LEDs
}

// Initialize makes a connection to each of the port expanders,
// returning an error if either is unavailable.
func (m *Adapter) Initialize() (err error) {

	// first attempt to initialize the port expanders at all
	if m.port1, m.port0, err = m.configurePorts(); err != nil {
		return err
	}

	// next configure the inputs/outputs appropriately for each port
	if err = m.configurePins(); err != nil {
		return err
	}

	return nil
}

// ReadRow
func (m *Adapter) ReadRow(rowIndex uint8) (row keyboard.Row) {

	// set all row outputs to high except for rowIndex
	rows := ^(uint16(1) << rowIndex)
	m.port1.SetPins(mcp23017.Pins(rows), port1_rowMask)

	// read input pins to determine which keys are pressed;
	// any inputs with logic low indicate a key press at the
	// given row,column in the matrix
	pins, err := m.port0.GetPins()
	if err != nil {
		pins = 0
	}
	row = keyboard.Row((^pins) & port0_colMask)

	// set all row outputs to high
	m.port1.SetPins(mcp23017.Pins(^uint16(0)), port1_rowMask)

	return row
}

func (m *Adapter) UpdateLEDs(ledState LEDs) {
	m.ledState = ledState
	m.port0.SetPins(ledState.port0state(), port0_ledMask)
	m.port1.SetPins(ledState.port1state(), port1_ledMask)
}

// initialize both MCP23017 expanders
func (m *Adapter) configurePorts() (*mcp23017.Device, *mcp23017.Device, error) {
	ports := []*mcp23017.Device{nil, nil}
	for i := range ports {
		addr := 0x20 + uint8(i)
		dev, err := mcp23017.NewI2C(m.bus, addr)
		if err != nil {
			return nil, nil, err
		}
		ports[i] = dev
	}
	return ports[0], ports[1], nil
}

func (m *Adapter) configurePins() error {
	// Configure port0 with GPA[0-6] as input with pull-ups,
	// and configure all other pins as output
	if err := m.port0.SetModes([]mcp23017.PinMode{
		mcp23017.Pullup,
		mcp23017.Pullup,
		mcp23017.Pullup,
		mcp23017.Pullup,
		mcp23017.Pullup,
		mcp23017.Pullup,
		mcp23017.Pullup,
		mcp23017.Output,
	}); err != nil {
		return err
	}
	// Configure port1 with all pins as output
	if err := m.port1.SetModes([]mcp23017.PinMode{mcp23017.Output}); err != nil {
		return err
	}
	return nil
}

const (

	// mask for column inputs on port0
	port0_colMask = mcp23017.Pins(0b00000000_01111111)

	// mask for row outputs on port1
	port1_rowMask = mcp23017.Pins(0b01111111_11111111)

	// port0_colMask = mcp23017.Pins(0b_01111111_00000000)
	// port0_pinMask = mcp23017.Pins(0b_00000000_00111111)
	// port0_ledMask = mcp23017.Pins(0b_10000000_11000000)

	// port0_pinMask = mcp23017.Pins(0b00111111_00000000)
)
