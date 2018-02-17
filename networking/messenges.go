package networking

import "encoding/json"

type Message struct {
	Type string          `json:"Type"`
	Data json.RawMessage `json:"Data,omitempty"`
}

func (m *Message) Serialize() ([]byte, error) {
	return json.Marshal(m)
}

func NewMessage(messageType string, dataStruct interface{}) (*Message, error) {
	b, err := json.Marshal(dataStruct)
	if err != nil {
		return nil, err
	}
	return &Message{messageType, b}, nil
}

func DeserializeToMessage(jsonString []byte) (*Message, error) {
	m := &Message{}
	err := json.Unmarshal(jsonString, m)
	if err != nil {
		return nil, err
	}
	return m, err
}

func DeserializeMessageData(message *Message, dataStruct interface{}) (error) {
	return json.Unmarshal(message.Data, dataStruct)
}
