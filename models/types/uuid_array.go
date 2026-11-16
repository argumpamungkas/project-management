package types

import "github.com/google/uuid"

// custom type
type UUIDArray []uuid.UUID

func (u *UUIDArray) Scan(value interface{}) error {
	return nil
}
