package ulid

import (
	"crypto/rand"

	"github.com/oklog/ulid"
)

// New returns a new ULID
func New() string {
	return ulid.MustNew(ulid.Now(), rand.Reader).String()
}
