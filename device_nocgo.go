//go:build !cgo

package ftdi

import (
	"errors"
)

var ErrNoCGO = errors.New("no CGO")

type USBDev struct {
	Manufacturer, Description, Serial string
}

type EEPROM struct {
	d *Device
}

type Device struct {
}

type Transfer struct {
}

func (u *USBDev) Close() {
}

func FindAll(vendor, product int) ([]*USBDev, error) {
	return nil, ErrNoCGO
}

// Close closes device
func (d *Device) Close() error {
	return ErrNoCGO
}

func OpenUSBDev(u *USBDev, c Channel) (*Device, error) {
	return nil, ErrNoCGO
}

func OpenFirst(vendor, product int, c Channel) (*Device, error) {
	return nil, ErrNoCGO
}

// Open opens the index-th device with a given vendor id, product id,
// description and serial. Uses specified channel c.
func Open(vendor, product int, description, serial string, index uint,
	c Channel) (*Device, error) {
	return nil, ErrNoCGO
}

// OpenBusAddr opens the device at a given USB bus and device address. Uses
// specified channel c.
func OpenBusAddr(bus, address int, c Channel) (*Device, error) {
	return nil, ErrNoCGO
}

// OpenString opens the ftdi-device described by a description-string. Uses
// specified channel c.
func OpenString(description string, c Channel) (*Device, error) {
	return nil, ErrNoCGO
}

// SetBitmode sets operation mode for device d to mode. iomask bitmask
// configures lines corresponding to its bits as input (bit=0) or output (bit=1).
func (d *Device) SetBitmode(iomask byte, mode Mode) error {
	return ErrNoCGO
}

// Reset resets device d.
func (d *Device) Reset() error {
	return nil
}

// PurgeWriteBuffer clears Rx buffer (buffer for data received from USB?).
func (d *Device) PurgeWriteBuffer() error {
	return ErrNoCGO
}

// PurgeReadBuffer clears Tx buffer (buffer for data that will be sent to USB?).
func (d *Device) PurgeReadBuffer() error {
	return ErrNoCGO
}

// PurgeBuffers clears both (Tx and Rx) buffers.
func (d *Device) PurgeBuffers() error {
	return ErrNoCGO
}

// ReadChunkSize returns current value of read buffer chunk size.
func (d *Device) ReadChunkSize() (int, error) {
	return 0, ErrNoCGO
}

// SetReadChunkSize configure read chunk size for device (default is 4096B) and
// size of software buffer dedicated for reading data from device...
func (d *Device) SetReadChunkSize(cs int) error {
	return ErrNoCGO
}

// WriteChunkSize returns current value of write chunk size.
func (d *Device) WriteChunkSize() (int, error) {
	return 0, ErrNoCGO
}

// SetWriteChunkSize configure write chunk size (default is 4096). If more than
// cs bytes need to be send to device, they will be split to chunks of size cs.
func (d *Device) SetWriteChunkSize(cs int) error {
	return ErrNoCGO
}

// LatencyTimer returns latency timer value [ms].
func (d *Device) LatencyTimer() (int, error) {
	return 0, ErrNoCGO
}

// SetLatencyTimer sets latency timer to lt (value beetwen 1 and 255). If FTDI
// device has fewer data to completely fill one USB packet (<62 B) it waits for
// lt ms before sending data to USB.
func (d *Device) SetLatencyTimer(lt int) error {
	return ErrNoCGO
}

// Read reads data from device to data. It returns number of bytes read.
func (d *Device) Read(data []byte) (int, error) {
	return 0, ErrNoCGO
}

// Write writes data from buf to device. It returns number of bytes written.
func (d *Device) Write(data []byte) (int, error) {
	return 0, ErrNoCGO
}

// WriteString writes bytes from string s to device. It returns the number of bytes written.
func (d *Device) WriteString(s string) (int, error) {
	return 0, ErrNoCGO
}

// ReadByte reads one byte from device.
func (d *Device) ReadByte() (byte, error) {
	return 0, ErrNoCGO
}

// WriteByte writes one byte to device.
func (d *Device) WriteByte(b byte) error {
	return ErrNoCGO
}

// Pins returns current state of pins (circumventing the read buffer).
func (d *Device) Pins() (b byte, err error) {
	return
}

// SetBaudrate sets the rate of data transfer.
func (d *Device) SetBaudrate(br int) error {
	return ErrNoCGO
}

// SetLineProperties sets the uart data bit count, stop bits count, and parity mode
// - values are defined in `device_nocgo.go`
func (d *Device) SetLineProperties(bits DataBits, stopbits StopBits, parity Parity) error {
	return ErrNoCGO
}

// SetLineProperties2 sets the uart data bit count, stop bits count, parity mode,
// and break usage
// - values are defined in `device_nocgo.go`
func (d *Device) SetLineProperties2(bits DataBits, stopbits StopBits, parity Parity, breaks Break) error {
	return ErrNoCGO
}

// SetFlowControl sets the flow control parameter
func (d *Device) SetFlowControl(flowctrl FlowCtrl) error {
	return ErrNoCGO
}

// SetDTRRTS manually sets the DTR and RTS output lines from the
// least significant bit of dtr and rts.
func (d *Device) SetDTRRTS(dtr, rts int) error {
	return ErrNoCGO
}

// SetDTR manually sets the DTR output line from the least significant
// bit of dtr.
func (d *Device) SetDTR(dtr int) error {
	return ErrNoCGO
}

// SetRTS manually sets the RTS output line from the least significant
// bit of rts.
func (d *Device) SetRTS(rts int) error {
	return ErrNoCGO
}

// ChipID reads FTDI Chip-ID (not all devices support this).
func (d *Device) ChipID() (uint32, error) {
	return 0, ErrNoCGO
}

func (d *Device) SubmitRead(data []byte) (*Transfer, error) {
	return nil, ErrNoCGO
}

func (d *Device) SubmitWrite(data []byte) (*Transfer, error) {
	return nil, ErrNoCGO
}

func (t *Transfer) Done() (int, error) {
	return 0, ErrNoCGO
}
