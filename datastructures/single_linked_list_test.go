package datastructures

import (
	"reflect"
	"testing"
)

func TestSingleLinkedList(t *testing.T) {
	sll := &SingleLinkedList[int]{}
	head := &ListNode[int]{
		next:  nil,
		value: 10,
	}

	node1 := &ListNode[int]{
		next:  nil,
		value: 20,
	}

	node2 := &ListNode[int]{
		next:  nil,
		value: 30,
	}
	node3 := &ListNode[int]{
		next:  nil,
		value: 40,
	}

	sll.prepend(head)
	expected_size := 1
	actual_size := sll.getSize()
	if !reflect.DeepEqual(expected_size, actual_size) {
		t.Fatalf("expected list size to be %v, got %v", expected_size, actual_size)
	}

	expected_data := []int{10}
	actual_data := sll.getData()
	if !reflect.DeepEqual(expected_data, actual_data) {
		t.Fatalf("expected data to be %v, got %v", expected_data, actual_data)
	}

	sll.append(node1)
	expected_size = 2
	actual_size = sll.getSize()
	expected_data = []int{10, 20}
	actual_data = sll.getData()
	head_node := sll.head
	tail_node := sll.tail

	if !reflect.DeepEqual(expected_size, actual_size) {
		t.Fatalf("expected lsit size to be %v, got %v", expected_size, actual_size)
	}
	if !reflect.DeepEqual(expected_data, actual_data) {
		t.Fatalf("expected data to be %v, got %v", expected_data, actual_data)
	}

	if !reflect.DeepEqual(head.value, head_node.value) {
		t.Fatalf("expected head node value to be %v, got %v", head.value, head_node.value)
	}
	if !reflect.DeepEqual(node1.value, tail_node.value) {
		t.Fatalf("expected tail node value to be %v, got %v", node1.value, tail_node.value)
	}

	sll.append(node2)
	sll.append(node3)

	expected_size = 4
	actual_size = sll.getSize()
	expected_data = []int{10, 20, 30, 40}
	actual_data = sll.getData()

	if !reflect.DeepEqual(expected_size, actual_size) {
		t.Fatalf("expected lsit size to be %v, got %v", expected_size, actual_size)
	}
	if !reflect.DeepEqual(expected_data, actual_data) {
		t.Fatalf("expected data to be %v, got %v", expected_data, actual_data)
	}

	if reflect.DeepEqual(sll.contains(30), false) {
		t.Fatalf("expected that value %v is contained in the list", 30)
	}

	if reflect.DeepEqual(sll.contains(60), true) {
		t.Fatalf("did not expect value %v to be in list", 60)
	}

	sll.delete(20)
	expected_size = 3
	actual_size = sll.getSize()
	expected_data = []int{10, 30, 40}
	actual_data = sll.getData()

	if !reflect.DeepEqual(expected_size, actual_size) {
		t.Fatalf("expected lsit size to be %v, got %v", expected_size, actual_size)
	}
	if !reflect.DeepEqual(expected_data, actual_data) {
		t.Fatalf("expected data to be %v, got %v", expected_data, actual_data)
	}
}
