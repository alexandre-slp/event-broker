//Entities are structs that represents an info and its methods, when necessary its constructors too

package domain

//SentEvent todo
type SentEvent struct {
	Name    string
	Message interface{}
}

//SavedEvent todo
type SavedEvent struct {
	Name   string
	Topics []string
}
