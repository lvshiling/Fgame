package tcp

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"sync"
	"time"
)

type TCPConnection struct {
	conn net.Conn
	rio  sync.Mutex
	br   *bufio.Reader
	wio  sync.Mutex
	bw   *bufio.Writer
}

func (tconn *TCPConnection) Connection() net.Conn {
	return tconn.conn
}

func (tconn *TCPConnection) Close() error {
	err := tconn.conn.Close()
	if err != nil {
		return err
	}
	return nil
}

func NewTCPConnection(conn net.Conn) *TCPConnection {
	tconn := &TCPConnection{}
	tconn.conn = conn
	br := bufio.NewReader(conn)
	bw := bufio.NewWriter(conn)
	tconn.br = br
	tconn.bw = bw
	return tconn
}

const (
	bytesOfLength        = 2
	maxBytes      uint16 = 65535
)

var (
	ErrNotSupported    = errors.New("not supported")
	ErrPayloadTooLarge = errors.New("payload too large")
)

func Receive(tconn *TCPConnection, v interface{}) (err error) {

	tconn.rio.Lock()
	defer tconn.rio.Unlock()
	//读取消息长度
	var len uint16
	if err = binary.Read(tconn.br, binary.LittleEndian, &len); err != nil {
		return
	}

	//读取全部payload
	msg := make([]byte, len)
	_, err = io.ReadFull(tconn.br, msg)
	if err != nil {
		return
	}
	switch data := v.(type) {
	case *[]byte:
		*data = msg
		return nil
	}
	return ErrNotSupported
}

func Send(tconn *TCPConnection, v interface{}) (err error) {
	tconn.wio.Lock()
	defer tconn.wio.Unlock()
	var msg []byte
	switch data := v.(type) {
	case []byte:
		msg = data
	default:
		return ErrNotSupported
	}
	len := uint16(len(msg))
	if len > maxBytes {
		return ErrPayloadTooLarge
	}
	var outputBuffer bytes.Buffer
	if err = binary.Write(&outputBuffer, binary.LittleEndian, len); err != nil {
		return
	}
	if err = binary.Write(&outputBuffer, binary.LittleEndian, msg); err != nil {
		return
	}
	tconn.Connection().SetWriteDeadline(time.Now().Add(time.Second * 5))
	if _, err := tconn.Connection().Write(outputBuffer.Bytes()); err != nil {
		return err
	}

	return
}

type Handler interface {
	Handle(conn *TCPConnection)
}

type HandlerFunc func(*TCPConnection)

func (hf HandlerFunc) Handle(conn *TCPConnection) {
	hf(conn)
}
