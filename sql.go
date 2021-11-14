package uuid

import (
	"database/sql/driver"
	"fmt"
)

// Scan implements sql.Scanner so UUIDs can be read from databases transparently
// Currently, database types that map to string and []byte are supported. Please
// consult database-specific driver documentation for matching types.
func (uuid *UUID) Scan(src interface{}) error {
	switch src.(type) {
	case string:
		// if an empty UUID comes from a table, we return a null UUID
		if src.(string) == "" {
			return nil
		}

		// see Parse for required string format
		u, err := Parse(src.(string))

		if err != nil {
			return fmt.Errorf("Scan: %v", err)
		}

		*uuid = u
	case []byte:
		b := src.([]byte)

		// if an empty UUID comes from a table, we return a null UUID
		if len(b) == 0 {
			return nil
		}

		// assumes a simple slice of bytes if 16 bytes
		// otherwise attempts to parse
		if len(b) != 16 {
			return uuid.Scan(string(b))
		}
		copy((*uuid)[:], b)
	default:
		return fmt.Errorf("Scan: unable to scan type %T into UUID", src)
	}

	return nil
}

// Value implements sql.Valuer so that UUIDs can be written to databases
// transparently. Currently, UUIDs map map to strings. Please consult
// database-specific driver documentation for matching types.
func (uuid UUID) Value() (driver.Value, error) {
	return uuid.String(), nil
}
