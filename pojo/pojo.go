package pojo

type TextInput struct {
	Text string `json:"text"`
}

type WordFrequency struct {
	Key   string `json:"word"`
	Value int    `json:"count"`
}