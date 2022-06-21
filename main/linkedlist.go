/* Date:2022/5/2-2022
 * Author:zhaoyufan
 */
package main

type MyLinkedList struct {
	dummy *Node
	size int
}
type Node struct{
	Val int
	Next *Node
}


func New() MyLinkedList {
	return MyLinkedList{dummy: new(Node)}
}


func (this *MyLinkedList) Get(index int) int {
	if index>this.size-1 || index<0 {return -1}
	cur:=this.dummy
	for index>0{
		cur=cur.Next
		index--
	}
	return cur.Next.Val
}


func (this *MyLinkedList) AddAtHead(val int)  {
	head:=&Node{Val:val,Next:this.dummy.Next}
	this.dummy.Next=head
	this.size++
}


func (this *MyLinkedList) AddAtTail(val int)  {
	cur:=this.dummy
	for cur.Next!=nil{
		cur=cur.Next
	}
	cur.Next = &Node{Val:val}
	this.size++
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index<0 {this.AddAtHead(val)}else if index == this.size{this.AddAtTail(val)}else if index>this.size{return}else{
		cur:=this.dummy
		for index>0{
			cur=cur.Next
			index--
		}
		newNode:=&Node{Val:val,Next:cur.Next}
		cur.Next = newNode
		this.size++
	}
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index<0 || index>this.size-1{return}
	cur:=this.dummy
	for index>0{
		cur=cur.Next
		index--
	}
	if cur.Next!=nil{
		cur.Next=cur.Next.Next
	}
	this.size--
}

