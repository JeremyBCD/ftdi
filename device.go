package ftdi

import (
)

// Type is numeric type id of FTDI device.
type Type uint32

const (
	TypeAM Type = iota
	TypeBM
	Type2232C
	TypeR
	Type2232H
	Type4232H
	Type232H
	Type230x
)

var types = []string{"AM", "BM", "2232C", "R", "2232H", "4232H", "232H", "230X"}

// String returns text name that describes type id.
func (t Type) String() string {
	if t >= Type(len(types)) {
		return "unknown"
	}
	return types[t]
}

// Channel represents channel (interface) of FTDI device. Some devices have more
// than one channel (eg. FT2232 has 2 channels, FT4232 has 4 channels).
type Channel uint32

const (
	ChannelAny Channel = iota
	ChannelA
	ChannelB
	ChannelC
	ChannelD
)

// Mode represents operation mode that FTDI device can work.
type Mode byte

const (
	ModeReset Mode = (1 << iota) >> 1
	ModeBitbang
	ModeMPSSE
	ModeSyncBB
	ModeMCU
	ModeOpto
	ModeCBUS
	ModeSyncFF
	ModeFT1284
)

// MPSSEDivValue calculates the two bytes that are required to be supplied after
// MPSSETCKDivisor to get the desired clock speed (in Hz).
// Set the dvi5 flag if MPSSEEnableDiv5 has been sent, to use a 12MHz base clock,
// instead of 60MHz.
func MPSSEDivValue(rate int, div5 bool) int {
	clk := 60_000_000
	if div5 {
		clk /= 5
	}
	if rate <= 0 || rate > clk {
		return 0
	}
	if (clk/rate)-1 > 0xffff {
		return 0xffff
	}
	return clk/rate - 1
}

// Parity represents the parity mode
type Parity int

// DataBits represents the number of data bits to expect
type DataBits int

// StopBits represents the number of stop bits to expect
type StopBits int

// Break represents the break mode
type Break int

type FlowCtrl int

const (
	// FlowCtrlDisable disables all automatic flow control.
	FlowCtrlDisable FlowCtrl = (1 << iota) >> 1
	// FlowCtrlRTSCTS enables RTS CTS flow control.
	FlowCtrlRTSCTS
	// FlowCtrlDTRDSR enables DTR DSR flow control.
	FlowCtrlDTRDSR
	// FlowCtrlXONXOFF enables XON XOF flow control.
	FlowCtrlXONXOFF
)

func (d *Device) EEPROM() EEPROM {
	return EEPROM{d}
}
