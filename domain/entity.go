//Entities are structs that represents an info and its methods, when necessary its constructors too

package domain

type SentEvent struct {
	Name    string      `json:"name,omitempty"`
	Message interface{} `json:"message,omitempty"`
}

type SavedEvent struct {
	Name   string   `json:"name,omitempty"`
	Topics []string `json:"topics,omitempty"`
}
