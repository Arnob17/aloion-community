package dto

import "time"

type CreateExamRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	CourseID    *uint     `json:"course_id"`
	BatchID     *uint     `json:"batch_id"`
	MaxScore    float64   `json:"max_score" binding:"required"`
	ExamDate    time.Time `json:"exam_date" binding:"required"`
}

type ExamResponse struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	CourseID    *uint   `json:"course_id"`
	BatchID     *uint   `json:"batch_id"`
	MaxScore    float64 `json:"max_score"`
	ExamDate    string  `json:"exam_date"`
	CreatedAt   string  `json:"created_at"`
}

type Question struct {
	ID            uint     `json:"id"`
	ExamID        uint     `json:"exam_id"`
	QuestionTitle string   `json:"question_title"`
	Options       []string `json:"options"`
	CorrectAnswer string   `json:"correct_answer"`
	Subject       string   `json:"subject"`
	Topic         string   `json:"topic?"`
	FullMarks     float64  `json:"full_marks"`
}

type CreateExamResultRequest struct {
	ExamID   uint      `json:"exam_id" binding:"required"`
	UserID   uint      `json:"user_id" binding:"required"`
	Score    float64   `json:"score" binding:"required"`
	MaxScore float64   `json:"max_score" binding:"required"`
	Feedback string    `json:"feedback"`
	ExamDate time.Time `json:"exam_date" binding:"required"`
}

type ExamResultResponse struct {
	ID         uint          `json:"id"`
	ExamID     uint          `json:"exam_id"`
	UserID     uint          `json:"user_id"`
	Score      float64       `json:"score"`
	MaxScore   float64       `json:"max_score"`
	Percentage float64       `json:"percentage"`
	Grade      string        `json:"grade"`
	Feedback   string        `json:"feedback"`
	ExamDate   string        `json:"exam_date"`
	ResultDate string        `json:"result_date"`
	Exam       *ExamResponse `json:"exam,omitempty"`
	User       *UserResponse `json:"user,omitempty"`
}
