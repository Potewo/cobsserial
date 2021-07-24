package cobsserial

import (
	"github.com/dim13/cobs"
	"github.com/tarm/serial"
)

var S *serial.Port
var Buf = make([]byte, 128)

func Write(b []byte) error {
	encodedData := cobs.Encode(b)
	_, err := S.Write(encodedData)
	return err
}

func Read() (b []byte, err error) {
	for {
		n, err := S.Read(Buf)
		if err != nil {
			return nil, err
		}
		b = append(b, Buf[:n]...)
		if contains(Buf[:n], byte(0x0)) {
			break
		}
	}
	b = cobs.Decode(b)
	return b, err
}

func contains(a []byte, d byte) bool {
	for _, v := range a {
		if d == v {
			return true
		}
	}
	return false
}
