package internal

type Subject struct {
	name string
	//groups
}

func NewSubject(name string) *Subject {
	return &Subject{name: name}
}