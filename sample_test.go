package main

import (
	"github.com/prestonchoate/go-katas/datastructures"
	"testing"
)

var sample = datastructures.SampleDS[int]{}

func TestSampleDSAddItem(t *testing.T) {
	sample.Clear()
	sample.AddItem(5)
	got := sample.Size()
	want := uint32(1)
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
	t.Log("Add Item success")
}

func TestSampleDSPrintItems(t *testing.T) {
	sample.Clear()
	sample.AddItem(4)
	got := sample.PrintItems()
	want := "4"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
	t.Log("Print Items success")
}
