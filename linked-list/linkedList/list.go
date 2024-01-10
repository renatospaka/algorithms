package linkedlist

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
}

type List struct {
	Head *Node
	Tail *Node
}

func (l *List) Append(p Pessoa) {
	node := &Node{Value: p}
	if l.Head == nil {
		l.Head = node
	}

	if l.Tail != nil {
		l.Tail = node
	}
	l.Tail = node
}

func (l *List) Display() {

}
