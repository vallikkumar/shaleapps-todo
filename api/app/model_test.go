package app

import "testing"

func TestModelUpdate(t *testing.T) {
	td := Todo{
		ID:       1,
		Message:  "Title1",
		Complete: false,
	}

	titleUpdate := "update title"
	u := Update{Message: &titleUpdate}
	td.Update(u)

	if td.Message != titleUpdate {
		t.Errorf("error: expected %s: got %s", titleUpdate, td.Message)
	}
}

func TestModelInput(t *testing.T) {
	i := Input{
		Message:  "InputTitle",
		Complete: false,
	}

	td := i.toTodo()

	if td.Message != i.Message {
		t.Errorf("error: expected %s: got %s", i.Message, td.Message)
	}

	if td.Complete != i.Complete {
		t.Errorf("error: expected %v: got %v", i.Complete, td.Complete)
	}
}
