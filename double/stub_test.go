package double

import "testing"

type stubSearcher struct {
	phone string
}

func (ss stubSearcher) Search(people []*person, firstname string, lastname string) *person {
	return &person{
		Firstname: firstname,
		Lastname:  lastname,
		Phone:     ss.phone,
	}
}

func TestFindReturnPerson(t *testing.T) {
	// Arrange
	fakePhone := "+31 65 222 333"
	phonebook := &Phonebook{}
	sd := stubSearcher{
		phone: fakePhone,
	}

	// Act
	got, _ := phonebook.Find(sd, "John", "Doer")

	// Assert
	if got != fakePhone {
		t.Errorf("got %s, want %s", got, fakePhone)
	}
}
