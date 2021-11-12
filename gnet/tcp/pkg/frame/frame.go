package frame

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/panjf2000/gnet"
)

type Frame []byte

func (cc *Frame) Encode(c gnet.Conn, buf []byte) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	length := uint32(4 + len(buf))
	if err := binary.Write(buffer, binary.BigEndian, length); err != nil {
		return nil, fmt.Errorf("pack length err: %v", err)
	}
	n, err := buffer.Write(buf)
	if err != nil {
		return nil, fmt.Errorf("pack frame payload er: %v", err)
	}
	if n != len(buf) {
		return nil, fmt.Errorf("pack payload length err: %v", err)
	}
	return buffer.Bytes(), nil
}

func (cc *Frame) Decode(c gnet.Conn) ([]byte, error) {
	var frameLength uint32
	if n, header := c.ReadN(4); n == 4 {
		byteBuffer := bytes.NewBuffer(header)
		_ = binary.Read(byteBuffer, binary.BigEndian, &frameLength)
		if frameLength > 100 {
			c.ResetBuffer()
			return nil, errors.New("length value is wrong")
		}
		if n, wholeFrame := c.ReadN(int(frameLength)); n == int(frameLength) {
			c.ShiftN(int(frameLength))
			return wholeFrame[4:], nil
		} else {
			return nil, errors.New("not enough frame payload data")
		}
	}
	return nil, errors.New("not enough frame length data")
}
