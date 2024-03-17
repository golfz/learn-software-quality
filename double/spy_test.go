package double

import "testing"

type spySearcher struct {
	phone           string
	searchWasCalled bool
	whatIsFirstname string
}

func (ss *spySearcher) Search(people []*person, firstname string, lastname string) *person {
	ss.searchWasCalled = true
	ss.whatIsFirstname = firstname

	return &person{
		Firstname: firstname,
		Lastname:  lastname,
		Phone:     ss.phone,
	}
}

func TestFindCallsSearchAndReturnPerson(t *testing.T) {
	// Arrange
	fakePhone := "+31 65 222 333"
	firstname := "John"
	phonebook := &Phonebook{}
	sd := &spySearcher{
		phone: fakePhone,
	}

	// Act
	got, _ := phonebook.Find(sd, firstname, "Doer")

	// Assert
	if !sd.searchWasCalled {
		t.Error("search was not called")
	}

	if sd.whatIsFirstname != firstname {
		t.Errorf("firstname is %s, want %s", sd.whatIsFirstname, firstname)
	}

	if got != fakePhone {
		t.Errorf("got %s, want %s", got, fakePhone)
	}
}
