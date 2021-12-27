//Entities are structs that represents an info and its methods, when necessary its constructors too

package domain

type SentEvent struct {
	Name    string
	Message interface{}
}

type SavedEvent struct {
	Name   string
	Topics []string
}
