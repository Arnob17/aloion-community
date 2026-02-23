package dto

type CreateCourseRequest struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description"`
	Subject     string  `json:"subject" binding:"required"`
	Class       int     `json:"class" binding:"required"`
	Level       string  `json:"level" binding:"required,oneof=beginner intermediate advanced olympiad"`
	Duration    int     `json:"duration"`
	Price       float64 `json:"price"`
	IsFree      bool    `json:"is_free"`
	Language    string  `json:"language"`
}

type UpdateCourseRequest struct {
	Title       *string  `json:"title"`
	Description *string  `json:"description"`
	Status      *string  `json:"status"`
	Price       *float64 `json:"price"`
	Discount    *float64 `json:"discount"`
}

type CourseResponse struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Subject     string  `json:"subject"`
	Class       int     `json:"class"`
	Level       string  `json:"level"`
	Status      string  `json:"status"`
	Duration    int     `json:"duration"`
	Price       float64 `json:"price"`
	IsFree      bool    `json:"is_free"`
	TeacherID   uint    `json:"teacher_id"`
	CreatedAt   string  `json:"created_at"`
}

type EnrollmentRequest struct {
	CourseID uint `json:"course_id" binding:"required"`
}
