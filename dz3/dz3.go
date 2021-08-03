// Двусвязный список
package main

import "fmt"

type Item struct {
	next, prev *Item
	list       *List
	Value      interface{}
}

// Получить следующий элемент следующий за текущим
func (el *Item) Next() *Item {
	if p := el.next; el.list != nil && p != &el.list.root {
		return p
	}
	return nil
}

// Получить предыдущий элемент от текущего
func (el *Item) Prev() *Item {
	if p := el.prev; el.list != nil && p != &el.list.root {
		return p
	}
	return nil
}

type List struct {
	root Item
	len  int
}

// Инициализация списка
func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// Создать новый экземпляр списка
func New() *List { return new(List).Init() }

// Получить кол-во элементов в списке
func (l *List) Len() int { return l.len }

// Получить первый элемент
func (l *List) First() *Item {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// Получить последний элемент
func (l *List) Last() *Item {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// Инициализация нулевого списка
func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

func (l *List) insert(e, at *Item) *Item {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

func (l *List) insertValue(v interface{}, at *Item) *Item {
	return l.insert(&Item{Value: v}, at)
}

func (l *List) remove(e *Item) *Item {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil
	e.prev = nil
	e.list = nil
	l.len--
	return e
}

func (l *List) Remove(e *Item) interface{} {
	if e.list == l {
		l.remove(e)
	}
	return e.Value
}

// добавить значение в начало
func (l *List) PushFront(v interface{}) *Item {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

// Добавить в конец
func (l *List) PushBack(v interface{}) *Item {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

func main() {
	list := New()

	list.PushFront(3)
	list.PushFront(2)
	list.PushFront(1)
	list.PushBack(4)

	fmt.Printf("Элементов: %d\n", list.Len())

	fmt.Println("Итерируем по списку с начала:")
	for e := list.First(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	fmt.Println("Итерируем по списку с конца:")
	for e := list.Last(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}
