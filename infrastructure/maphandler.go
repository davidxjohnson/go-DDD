////
// Contains code for managing and persisting an all-purpose map[string]map[string] data structure
// This type of structure can be used to persist an indexed list of name/value pairs ...
// - i.e. Contact[a38e2a13-a03d-4643-96c7-d50b3d4f2228][Phone] returns 888-555-1212
///

package infrastructure

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
)

// Row represents a dictionary or tupel in a datastore
type Row map[string]string // Field (key) and value pairs

// Table represents a collection of Rows
type Table struct {
	filePath string
	Rows     map[string]Row `json:"rows"`
}

// InitTable function reads data from disk into new Table object
func InitTable(filePath string) (t Table, ok bool) {
	ok = readJSON(&t, filePath)
	if ok {
		t.filePath = filePath
	}
	return
}

// Commit method writes entire datastore to disk
func (t Table) Commit() (ok bool) {
	ok = writeJSON(t, t.filePath) // marshal and write to disk
	return
}

// FindByID method gets index to Rows with matching ID
func (t Table) FindByID(id string) (r Row, ok bool) {
	r, ok = t.Rows[id]
	return
}

// Add method populates a new Row record
func (t Table) Add(r Row) (id string, ok bool) {
	id = makeuuid()
	t.Rows[id] = r // Unique ID
	ok = true
	return
}

// Update method modifies an existing Row record
func (t Table) Update(id string, r Row) (ok bool) {
	_, ok = t.FindByID(id) // Find original Row Info
	if !ok {               // Return if not found
		return
	}
	// Update Row Info
	t.Rows[id] = r // Update Row
	return
}

// Delete method destroys a Row record
func (t Table) Delete(id string) (ok bool) {
	_, ok = t.FindByID(id) // Find Row Info index
	if !ok {               // Return if Not Found
		return
	}
	// Delete Row Info
	delete(t.Rows, id)
	return
}

// Query method finds matching Row records and returns a Table object with all matching rows.
// For query purposes, each key in url.Values[key] is unique.
// Multiple keys in the map are considered "or" as are also multiple values for a single key.
// Keys in url.Values are table column names (fields) in map[key]map[field]string in the Table object.
func (t Table) Query(pairs url.Values) (dt Table, ok bool) {
	dt.Rows = make(map[string]Row)
Rowloop:
	for id, r := range t.Rows { // Scan Rows
		foundMatch := false
	pairsloop:
		for qk, qa := range pairs { // check each key/value against Row Info
			for _, qv := range qa { // Might be more than one value to check
				if rv, idExists := r[qk]; idExists { // Check for matching key name
					if qv == rv { // Check for matching value
						foundMatch = true  // Got a match
						continue pairsloop // Get next pair as this value matches
					}
				}
			}
			continue Rowloop // Get next Row as no values matched
		}
		if foundMatch { // found one
			dt.Rows[id] = r // add to list
		}
	}
	ok = (len(dt.Rows) > 0)
	return
}

////
// supporting functions
////

// makeuuid function Creates a uuid as a unique record id
func makeuuid() (uuid string) {
	b := make([]byte, 16)
	rand.Read(b)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return
}

// readJSON function gets data from a file into a structure
func readJSON(returnData interface{}, path string) (ok bool) {
	rawData, ok := readFileStream(path)
	// convert the raw data into the structure passed by reference to this function
	err := json.Unmarshal(rawData, returnData)
	if err != nil {
		log.Print("readJSON Unmarshal: " + err.Error())
		ok = false
	}
	return
}

// writeJSON function puts data from structure to file
func writeJSON(sourceData interface{}, path string) (ok bool) {
	rawData, err := json.Marshal(sourceData)
	if err != nil {
		log.Print("writeJSON Marshal: '" + path + "' " + err.Error())
		ok = false
		return
	}
	err = ioutil.WriteFile(path, rawData, 0)
	if err != nil {
		log.Print("writeJSON WriteFile: '" + path + "' " + err.Error())
		ok = false
		return
	}
	ok = true
	return
}

// readFileStream function returns a stream of bytes from a file, or bails if unable
func readFileStream(filePath string) (rawData []byte, ok bool) {
	// Open file
	fileHandle, err := os.Open(filePath)
	if err != nil {
		ok = false
		log.Print("readFileStream: " + err.Error())
		return
	}
	defer fileHandle.Close() // closed this later
	// Read file
	rawData, err = ioutil.ReadAll(fileHandle)
	// nearly impossible to throw this error and just as hard to test
	if err != nil {
		ok = false
		log.Print("readFileStream: " + err.Error())
		return
	}
	ok = true
	return
}
