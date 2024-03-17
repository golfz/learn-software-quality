package double

import "testing"

type mockSearcher struct {
	phone        string
	methodToCall map[string]bool
}

func (ms *mockSearcher) Search(people []*person, firstname string, lastname string) *person {
	ms.methodToCall["Search"] = true
	return &person{
		Firstname: firstname,
		Lastname:  lastname,
		Phone:     ms.phone,
	}
}

func (ms *mockSearcher) Create(people []*person, firstname string, lastname string) *person {
	ms.methodToCall["Create"] = true
	return &person{
		Firstname: firstname,
		Lastname:  lastname,
		Phone:     ms.phone,
	}
}

func (ms *mockSearcher) ExpectToCall(methodName string) {
	if ms.methodToCall == nil {
		ms.methodToCall = make(map[string]bool)
	}
	ms.methodToCall[methodName] = false
}

func (ms *mockSearcher) Verify(t *testing.T) {
	for methodName, called := range ms.methodToCall {
		if !called {
			t.Errorf("Expected to call '%s', but it wasn't", methodName)
		}
	}
}

func TestFindCallsSearchAndReturnPersonUsingMock(t *testing.T) {
	// Arrange
	fakePhone := "+31 65 222 333"
	phonebook := &Phonebook{}
	mock := &mockSearcher{
		phone: fakePhone,
	}
	mock.ExpectToCall("Search")

	// Act
	got, _ := phonebook.Find(mock, "John", "Doer")

	// Assert
	if got != fakePhone {
		t.Errorf("got %s, want %s", got, fakePhone)
	}

	mock.Verify(t)
}
