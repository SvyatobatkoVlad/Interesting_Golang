package main

import "fmt"

// LinkNode определим структуру для односвязного списка
type LinkNode struct {
	next  *LinkNode
	value int
}

// Peverse определим метод обращения списка
func (n *LinkNode) Peverse() *LinkNode {

	//нам понадобятся 2 указателя на предыдущий и текущий элемент
	var cur = n
	var prev *LinkNode

	for cur != nil {
		//сохраним ссылку на следующий элемент что бы заменить ее на предыдущий
		next := cur.next
		cur.next = prev
		prev = cur
		cur = next
	}

	return prev
}

func (n *LinkNode) Print() {

	//выведем элементы списка на экран
	cur := n
	for cur != nil {
		splitter := ""
		if cur != n {
			splitter = " -> "
		}
		fmt.Printf("%s%d", splitter, cur.value)
		cur = cur.next
	}

	fmt.Println()
}

func main() {
	//приведем пример инициализации списка
	n1 := &LinkNode{value: 1}
	n2 := &LinkNode{value: 2}
	n1.next = n2
	n3 := &LinkNode{value: 3}
	n2.next = n3

	n1.Print()

	n1.Peverse().Print()
}
