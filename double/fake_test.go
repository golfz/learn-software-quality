package double

import "testing"

type fakeSearcher struct{}

func (fs *fakeSearcher) Search(people []*person, firstname string, lastname string) *person {
	if len(people) == 0 {
		return nil
	}

	return people[0]
}

func TestFindCallsSearchAndReturnsEmptyStringForNoPerson(t *testing.T) {
	// Arrange
	phonebook := &Phonebook{}
	fd := &fakeSearcher{}

	// Act
	phone, _ := phonebook.Find(fd, "John", "Doer")

	// Assert
	if phone != "" {
		t.Errorf("got %s, want %s", phone, "")
	}
}

func TestFindCallsSearchAndReturnsString(t *testing.T) {
	// Arrange
	fakePhone := "+31 65 222 333"
	people := []*person{
		{
			Phone: fakePhone,
		},
	}
	phonebook := &Phonebook{
		people: people,
	}
	fd := &fakeSearcher{}

	// Act
	phone, _ := phonebook.Find(fd, "John", "Doer")

	// Assert
	if phone != fakePhone {
		t.Errorf("got %s, want %s", phone, fakePhone)
	}

}
