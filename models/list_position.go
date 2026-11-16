package models

type ListPosition struct {
	InternalID int64       `json:"internal_id" db:"internal_id" gorm:"primaryKey;autoIncrement"`
	PublicID   int64       `json:"public_id" db:"public_id" gorm:"public_id"`
	BoardId    int64       `json:"-" db:"board_internal_id" gorm:"column:board_internal_id"` // board internal id
	ListOrder  []UUIDArray // array uuid
}
