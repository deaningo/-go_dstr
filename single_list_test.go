package go_dstr

import "testing"

func TestSingleList_Init(t *testing.T) {
	list := new(SingleList)
	list.Init()
	t.Log("single list init success")
}

func TestSingleListAppend(t *testing.T) {
	list := new(SingleList)
	list.Init()

	b := list.Append(nil)
	if b {
		t.Error("single list append nil failed")
	} else {
		t.Log("single list append nil success")
	}

	b = list.Append(&SingleNode{Data: 1})
	if b {
		t.Log("single list append first success")
	} else {
		t.Error("single list append first failed")
	}

	b = list.Append(&SingleNode{Data: 2})
	if b {
		t.Log("single list append second success")
	} else {
		t.Error("single list append second failed")
	}
}
