package double

import "testing"

type DummySearcher struct{}

func (ds DummySearcher) Search(people []*person, firstname string, lastname string) *person {
	return &person{}
}

func TestFindShouldReturnErrorWhenFirstnameOrLastnameEmpty(t *testing.T) {
	// Arrange
	phonebook := &Phonebook{}
	dd := DummySearcher{}
	want := ErrMissingArgs

	// Act
	_, got := phonebook.Find(dd, "", "")

	// Assert
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
