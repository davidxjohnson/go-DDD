package customer

import     "go-DDD/utility"

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

func NewCustomer (
	_nameFirst 	string, 
	_nameLast 	string, 
	_phones 	Phones, 
	_streetNumber 	string, 
	_streetName 	string, 
	_city 		string, 
	_state 		string, 
	_zipCode 	string, 
	_countryCode	string ) ( aCustomer Customer, err error ) {
   aCustomer = Customer{
	id: utility.Makeuuid(), 
	nameFirst: _nameFirst, 
	nameLast: _nameLast, 
	phones: _phones, 
	streetNumber: _streetNumber, 
	streetName: _streetName, 
	city: _city, 
	state: _state, 
	zipCode: _zipCode, 
	countryCode: _countryCode }
    return 
}
