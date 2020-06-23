package screencapture

import (
	"fmt"
	"os"
)

// PacketDumper wraps a UsbWriter and a UsbDataReceiver to dump all data received and sent to a file
type PacketDumper struct {
	writer   UsbWriter
	receiver UsbDataReceiver
	file     *os.File
}

// NewPacketDumper takes a UsbWriter and opens a dumpfile at path
func NewPacketDumper(writer UsbWriter, path string) (PacketDumper, error) {
	dumpFile, err := os.Create(path)
	if err != nil {
		return PacketDumper{}, err
	}
	return PacketDumper{writer: writer, file: dumpFile}, nil
}

// SetReceiver can be used to set a UsbDataReceiver after creating a new PacketDumper
func (pd *PacketDumper) SetReceiver(receiver UsbDataReceiver) {
	pd.receiver = receiver
}

// ReceiveData writes the hexdump of data to the dumpfile and forwards the data to the UsbDataReceiver
func (pd PacketDumper) ReceiveData(data []byte) {
	pd.file.WriteString(fmt.Sprintf("rcv:%x\n", data))
	pd.receiver.ReceiveData(data)
}

// CloseSession closes the dump file and the underlying UsbDataReceiver
func (pd PacketDumper) CloseSession() {
	pd.file.Close()
	pd.receiver.CloseSession()
}

// WriteDataToUsb writes the hexdump of data to the dumpfile and forwards the data to the UsbWriter
func (pd PacketDumper) WriteDataToUsb(data []byte) {
	pd.file.WriteString(fmt.Sprintf("send:%x\n", data))
	pd.writer.WriteDataToUsb(data)
}
