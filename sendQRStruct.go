package connection

import (
	"encoding/json"
	"net"
)

// SendQRStruct - Отправляет структуру QueryResponse в conn
func SendQRStruct(qr QueryResponse, conn net.Conn) error {
	data, err := json.Marshal(qr)
	if err != nil {
		return err
	}
	data = append(data, '\n')
	err = SendBytes(data, conn)
	if err != nil {
		return err
	}
	return nil
}
