package main
import (
	"fmt"
)

type SinglyLinkedListNode struct{
	data int
	next *SinglyLinkedListNode
}

type SinglyLinkedList struct{
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

func (ll *SinglyLinkedList) insert_data (data int) {
	n := &SinglyLinkedListNode{data,nil}

	if ll.head!=nil {
		ll.tail.next = n;
	}else{
		ll.head = n;
	}

	ll.tail = n;
}

func (ll *SinglyLinkedList) insert_node (n *SinglyLinkedListNode) {

	if ll.head!=nil {
		ll.tail.next = n;
	}else{
		ll.head = n;
	}

	ll.tail = n;
}

func (ll *SinglyLinkedList) printLL (){

	t := ll.head
	if t!=nil {

		for t!=nil {
			fmt.Println(t.data)
			t = t.next	
		}

	}else{
		fmt.Println("Empty Linked List")
	}
}

func (ll *SinglyLinkedList) deleteLL(){

	if ll.head!=nil{
		for ll.head.next!=nil {
			t := ll.head.next	
			ll.head.next = nil
			ll.head = t
		}
		ll.head=nil
	}
}


func main() {
	var a =  SinglyLinkedListNode{666,nil}

	var ll SinglyLinkedList

	ll.insert_data(1)
	ll.insert_data(2)
	ll.insert_node(&a)
	ll.insert_data(3)
	ll.insert_data(4)
	ll.printLL()
	ll.deleteLL()
	ll.printLL()
	
}