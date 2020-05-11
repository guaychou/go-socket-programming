package models

type Data struct {
	Id int `json:"id"`
	Nama string `json:"nama"`
	Nim string `json:"nim"`
}
type MessageArray struct {
	Message []Data `json:"message"`
}

func (messageArray *MessageArray) AddMessage(data Data) []Data{
	messageArray.Message=append(messageArray.Message,data)
	return messageArray.Message
}