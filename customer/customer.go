package customer

import "go-DDD/utility"

type Phones map[string]string

type Customer struct {
	id           string
	nameFirst    string
	nameLast     string
	phones       Phones
	streetNumber string
	streetName   string
	city         string
	state        string
	zipCode      string
	countryCode  string
}

func (c Customer) New(
	_nameFirst string,
	_nameLast string,
	_phones Phones,
	_streetNumber string,
	_streetName string,
	_city string,
	_state string,
	_zipCode string,
	_countryCode string) (newCust Customer, err error) {
	c.id = utility.Makeuuid()
	c.nameFirst = _nameFirst
	c.nameLast = _nameLast
	c.phones = _phones
	c.streetNumber = _streetNumber
	c.streetName = _streetName
	c.city = _city
	c.state = _state
	c.zipCode = _zipCode
	c.countryCode = _countryCode
	newCust = c
	return
}
