package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type jsonTime time.Time

func (t jsonTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(time.RFC3339))
	return []byte(stamp), nil
}

type Message struct {
	CreatedAt jsonTime `json:"createdAt"`
}

func (m Message) MarshalJSON() ([]byte, error) {
	// 取别名，避免对Message.MarchalJSON进行递归调用
	type Alias Message
	return json.Marshal(struct {
		Alias
		CreatedAt string `json:"createdAt"`
	}{
		CreatedAt: time.Time(m.CreatedAt).Format("2006/01/02 15:04:05"),
		Alias:     Alias(m),
	})
}

func main() {
	marshal, err := json.Marshal(Message{
		CreatedAt: jsonTime(time.Now()),
	})
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%#v", string(marshal))
}
