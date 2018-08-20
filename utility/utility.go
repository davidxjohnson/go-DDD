package utility

import(
    "crypto/rand"
    "fmt"
)

// utilities

// makeuuid function Creates a uuid as a unique record id
func Makeuuid() (uuid string) {
        b := make([]byte, 16)
        rand.Read(b)
        uuid = fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
        return
}
