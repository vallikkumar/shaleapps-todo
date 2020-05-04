package app

import "testing"

func setup(t *testing.T) *Service {
	t.Parallel()
	s := NewService()
	return s
}

func TestResolve(t *testing.T) {
	s := setup(t)
	_, err := s.Resolve()
	if err != nil {
		t.Fatal("could not resolve data")
	}
}

func TestResolveByID(t *testing.T) {
	i := Input{
		Message:  "test",
		Complete: true,
	}
	s := setup(t)

	res, err := s.Create(i)
	if err != nil {
		t.Error("could not save data")
	}

	_, err = s.ResolveByID(res.ID)
	if err != nil {
		t.Fatal("could not resolve data")
	}
}

func TestCreate(t *testing.T) {
	i := Input{
		Message:  "test",
		Complete: true,
	}
	s := setup(t)

	_, err := s.Create(i)
	if err != nil {
		t.Error("could not save data")
	}
}

func TestRemove(t *testing.T) {
	i := Input{
		Message:  "test",
		Complete: true,
	}
	s := setup(t)

	res, err := s.Create(i)
	if err != nil {
		t.Error("could not save data")
	}

	err = s.Remove(res.ID)
	if err != nil {
		t.Fatal("could not resolve data")
	}
}

func TestUpdate(t *testing.T) {
	i := Input{
		Message:  "test",
		Complete: false,
	}
	s := setup(t)

	res, err := s.Create(i)
	if err != nil {
		t.Error("could not save data")
	}

	c := true
	u := Update{
		Complete: &c,
	}
	_, err = s.Update(res.ID, u)
	if err != nil {
		t.Fatal("could not resolve data")
	}
}
