package models

import "time"

// composite primary key: BoardID + UserInternalID
type BoardMember struct {
	BoardID  int64     `json:"board_id" db:"board_id" gorm:"column:board_id;primaryKey"`                         // board id
	UserID   int64     `json:"user_internal_id" db:"user_internal_id" gorm:"column:user_internal_id;primaryKey"` // user id
	JoinedAt time.Time `json:"joined_at" db:"joined_at"`                                                         // timestamp when user joined the board
}
