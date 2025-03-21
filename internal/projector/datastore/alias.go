package datastore

import (
	"context"

	"github.com/peb-adr/openslides-go/datastore/dsfetch"
	"github.com/peb-adr/openslides-go/datastore/dskey"
)

// DoesNotExistError is a type alias from datastore.DoesNotExistError
type DoesNotExistError = dsfetch.DoesNotExistError

// Key is a type alias from dskey.Key
type Key = dskey.Key

// KeyFromString from package dskey.
var KeyFromString = dskey.FromString

// Getter is the same as datastore.Getter
type Getter interface {
	Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error)
}
