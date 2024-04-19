package MyQueue

import linkList "github.com/https-whoyan/ADS/tree/main/Assigment2/pkg/Lists/MyLinkedList"

type MyQueue struct {
	data *linkList.MyLinkedList
}

func NewQueue() *MyQueue {
	return &MyQueue{
		data: linkList.NewLinkedList(),
	}
}

func (s *MyQueue) Clear() {
	s.data.Clear()
}

func (s *MyQueue) Contains(el any) bool {
	return s.data.Exists(el)
}

func (s *MyQueue) Push(newEl any) error {
	return s.data.AddFirst(newEl)
}

func (s *MyQueue) Pop() (any, error) {
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

func (s *MyQueue) Peek() (any, error) {
	return s.data.GetLast()
}
