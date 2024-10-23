package entity

import "time"

type Message struct {
	ID         uint      `gorm:"primaryKey"`
	SenderID   uint      `gorm:"not null"`
	ReceiverID uint      `gorm:"not null"`
	Content    string    `gorm:"type:text;not null"`
	Timestamp  time.Time `gorm:"autoCreateTime"`
	Sender     User      `gorm:"foreignKey:SenderID"`
	Receiver   User      `gorm:"foreignKey:ReceiverID"`
}

func (Message) TableName() string {
	return "message"
}
