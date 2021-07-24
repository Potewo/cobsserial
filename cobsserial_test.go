package cobsserial

import (
	"reflect"
	"testing"
	"time"

	"github.com/tarm/serial"
)

func TestSuccess(t *testing.T) {
	message := []byte("Hello World")
	c := &serial.Config{
		Name:        "/dev/ttyACM0",
		Baud:        9600,
		ReadTimeout: time.Second * 2,
	}

	s, err := serial.OpenPort(c)
	S = s
	if err != nil {
		t.Fatal(err)
	}

	Write(message)
	b, err := Read()
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Read %#v", b)
	}
	if reflect.DeepEqual(b, message) {
		t.Logf("Success!")
	} else {
		t.Fatalf("Not same\nexpected:  %#v\nreceived: %#v", message, b)
	}
}
