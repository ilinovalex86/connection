package connection

import (
	"encoding/json"
	"net"
)

// SendEvent - Отправляет структуру Event в conn
func SendEvent(e Event, conn net.Conn) error {
	data, err := json.Marshal(e)
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
