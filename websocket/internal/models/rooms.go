package models

// Struct cho bảng "rooms"
type Rooms struct {
	ID          int64 `json:"id"`
	SenderID    int64 `json:"sender_id"`
	RecipientID int64 `json:"recipient_id"`
	CreatedAt   int   `json:"created_at"`
}
