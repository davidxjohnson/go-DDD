package utility

import "testing"

func TestMakeuuid(t *testing.T) {
        key := Makeuuid()
        if key == "" {
                t.Errorf("TestMakeuuid: generated key was blank.")
        }
}
