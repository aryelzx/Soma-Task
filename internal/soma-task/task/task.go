package task

import (
	"bytes"
	"io"
	"time"
)

type Task struct {
	Id          uint64
	Name        string
	Message     *Message
	TimeSeconds uint16
	Repeat      bool
	CreatedAt   time.Time
}

type Message struct {
	From    uint64
	Payload []byte
}

func NewTask(id uint64, name string, message io.Reader, timeSeconds uint16, repeat bool) *Task {
	return &Task{
		Id:          id,
		Name:        name,
		Message:     NewMessage(message, id),
		TimeSeconds: timeSeconds,
		Repeat:      repeat,
		CreatedAt:   time.Now(),
	}
}

func NewMessage(payload io.Reader, from uint64) *Message {
	buf := new(bytes.Buffer)
	buf.ReadFrom(payload)

	return &Message{
		From:    from,
		Payload: buf.Bytes(),
	}
}
