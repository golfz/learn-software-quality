package double

import "errors"

var (
	ErrMissingArgs   = errors.New("arguments not found")
	ErrNoPersonFound = errors.New("no person found")
)

type Queryer interface {
	Search(people []*person, firstname string, lastname string) *person
}

type person struct {
	Firstname string
	Lastname  string
	Phone     string
}

type Phonebook struct {
	people []*person
}

func (p *Phonebook) Find(query Queryer, firstname, lastname string) (string, error) {
	if firstname == "" || lastname == "" {
		return "", ErrMissingArgs
	}

	person := query.Search(p.people, firstname, lastname)
	if person == nil {
		return "", ErrNoPersonFound
	}

	return person.Phone, nil
}
