package MyStack

import linkList "github.com/https-whoyan/ADS/tree/main/Assigment2/pkg/Lists/MyLinkedList"

type MyStack struct {
	data *linkList.MyLinkedList
}

func NewStack() *MyStack {
	return &MyStack{
		data: linkList.NewLinkedList(),
	}
}

func (s *MyStack) Clear() {
	s.data.Clear()
}

func (s *MyStack) Contains(el any) bool {
	return s.data.Exists(el)
}

func (s *MyStack) Push(newEl any) error {
	return s.data.AddLast(newEl)
}

func (s *MyStack) Pop() (any, error) {
	val, err := s.data.GetLast()
	if err != nil {
		return nil, err
	}
	err = s.data.RemoveLast()
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (s *MyStack) Peek() (any, error) {
	return s.data.GetLast()
}
