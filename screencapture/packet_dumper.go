package screencapture

import (
	"fmt"
	"os"
)

type PacketDumper struct {
	writer   UsbWriter
	receiver UsbDataReceiver
	file     *os.File
}

func NewPacketDumper(writer UsbWriter, path string) (PacketDumper, error) {
	dumpFile, err := os.Create(path)
	if err != nil {
		return PacketDumper{}, err
	}
	return PacketDumper{writer: writer, file: dumpFile}, nil
}

func (pd *PacketDumper) SetReceiver(receiver UsbDataReceiver) {
	pd.receiver = receiver
}

func (pd PacketDumper) ReceiveData(data []byte) {
	pd.file.WriteString(fmt.Sprintf("rcv:%x\n", data))
	pd.receiver.ReceiveData(data)
}

func (pd PacketDumper) CloseSession() {
	pd.file.Close()
	pd.receiver.CloseSession()
}

func (pd PacketDumper) WriteDataToUsb(data []byte) {
	pd.file.WriteString(fmt.Sprintf("send:%x\n", data))
	pd.writer.WriteDataToUsb(data)
}
