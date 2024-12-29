package models

import "time"

// AnswerRecord represents a record of a user's answer to a question.
type AnswerRecord struct {
    ID        int       `json:"id"`        // Unique identifier for the answer record
    QuestionID int      `json:"question_id"` // ID of the question being answered
    UserID    int       `json:"user_id"`   // ID of the user who provided the answer
    Answer    string    `json:"answer"`    // The answer provided by the user
    CreatedAt time.Time `json:"created_at"` // Timestamp when the answer was recorded
    UpdatedAt time.Time `json:"updated_at"` // Timestamp when the answer was last updated
}
