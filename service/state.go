package service

type State interface {
	Operation(int, int) int
}

type (
	AddState      struct{}
	SubtractState struct{}
)

func NewAddState() AddState {
	return AddState{}
}

func (AddState) Operation(v1, v2 int) int {
	return v1 + v2
}

func NewSubtractState() SubtractState {
	return SubtractState{}
}

func (SubtractState) Operation(v1, v2 int) int {
	return v1 - v2
}
