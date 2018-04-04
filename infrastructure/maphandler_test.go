package infrastructure

import (
	"fmt"
	"net/url"
	"testing"
)

var myTable Table // global var representing our test table
var addKey string
var addFieldName = "email"
var addRow = Row{"city": "Hammonton", "email": "lthomas26@go.com", "first": "Janie", "last": "Mitchell", "number": "13", "phone": "718-864-5281", "state": "NJ", "street": "E. Armstrong Rd.", "zip": "08037"}
var addFieldValue string

func TestInit(t *testing.T) {
	filePath := "./" + makeuuid() + ".json"
	_, ok := InitTable(filePath)
	if ok {
		t.Errorf("TestInit: received false 'ok' reading non-existing file path '%s'", filePath)
	}
	filePath = "./badtestdata.json"
	_, ok = InitTable(filePath)
	if ok {
		t.Errorf("TestInit: received false 'ok' reading malformed JSON file '%s'", filePath)
	}
	// read the test data for the rest of the test cases
	filePath = "./testdata.json"
	myTable, ok = InitTable(filePath)
	if !ok {
		t.Errorf("TestInit: Table create failed using filePath '%s'", filePath)
	}
	fmt.Printf("TestInit: file path is '%s'", myTable.filePath)
}

func TestAdd(t *testing.T) {
	var myRow = addRow
	key, ok := myTable.Add(myRow)
	if !ok {
		t.Errorf("TestAdd: failed to add new row %+v", myRow)
	}
	addKey = key // store in global var for later use
	addFieldValue = myRow[addFieldName]
}

func TestFindById(t *testing.T) {
	var id = addKey
	var fieldName = addFieldName
	var fieldValue = addFieldValue
	myRow, ok := myTable.FindByID(id)
	if !ok {
		t.Errorf("TestFindById: myRow[%s]: not found.", id)
	}
	if myRow[fieldName] != fieldValue {
		t.Errorf("TestFindById: myRow[%s][%s]: not does not match \"%s\".", id, fieldName, fieldValue)
	}
}

func TestUpdate(t *testing.T) {
	var id = addKey
	bogusID := makeuuid()
	var fieldName = addFieldName
	var fieldValue = "jboyd12@telegraph.co.uk"
	myRow := addRow
	myRow[fieldName] = fieldValue
	ok := myTable.Update(bogusID, myRow)
	if ok {
		t.Errorf("TestUpdate: myRow[%s]: was found, but shouldn't have (key is bogus).", bogusID)
	}
	ok = myTable.Update(id, myRow)
	if !ok {
		t.Errorf("TestUpdate: myRow[%s]: not found.", id)
	}
	if myRow[fieldName] != fieldValue {
		t.Errorf("TestUpdate: myRow[%s][%s]: not does not match \"%s\".", id, fieldName, fieldValue)
	}
}

func TestDelete(t *testing.T) {
	var id = addKey
	bogusID := makeuuid()
	ok := myTable.Delete(bogusID)
	if ok {
		t.Errorf("TestDelete: myRow[%s]: returned false 'ok for deletion of bogus key.", bogusID)
	}
	ok = myTable.Delete(id)
	if !ok {
		t.Errorf("TestDelete: myRow[%s]: was not found for deletion.", id)
	}
	_, ok = myTable.FindByID(id)
	if ok {
		t.Errorf("TestDelete: myRow[%s]: was found (should be deleted).", id)
	}
}

func TestQuery(t *testing.T) {
	myQuery := url.Values{}
	myQuery.Add("state", "VA")
	myQuery.Add("state", "NJ")
	fmt.Printf("%v+", myQuery)
	_, ok := myTable.Query(myQuery)
	if !ok {
		t.Errorf("TestQuery: No results found.")
	}
}

func TestCommit(t *testing.T) {
	ok := myTable.Commit()
	if !ok {
		t.Errorf("TestCommit: failed to write '%s'.", myTable.filePath)
	}
	myTable.filePath = "/dev/sdc"
	ok = myTable.Commit()
	if ok {
		t.Errorf("TestCommit: false 'ok' on write to bogus file '%s'.", myTable.filePath)
	}
}

func TestWriteJSON(t *testing.T) {
	ok := writeJSON(func() {}, "bogus")
	if ok {
		t.Errorf("TestWriteJSON: false 'ok' on write using invalid object func()")
	}
}
