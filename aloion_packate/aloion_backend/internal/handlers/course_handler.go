package handlers

import (
	"net/http"
	"strconv"

	"github.com/arnob17/aloion_backend/internal/dto"
	"github.com/arnob17/aloion_backend/internal/models"
	"github.com/arnob17/aloion_backend/internal/services"
	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	courseService *services.CourseService
}

func NewCourseHandler(courseService *services.CourseService) *CourseHandler {
	return &CourseHandler{
		courseService: courseService,
	}
}

func (h *CourseHandler) Create(c *gin.Context) {
	var req dto.CreateCourseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	teacherID, _ := c.Get("user_id")
	course, err := h.courseService.Create(req, teacherID.(uint))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, course)
}

func (h *CourseHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	course, err := h.courseService.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	c.JSON(http.StatusOK, course)
}

func (h *CourseHandler) GetAll(c *gin.Context) {
	limit := 20
	offset := 0

	if l := c.Query("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil {
			limit = parsed
		}
	}
	if o := c.Query("offset"); o != "" {
		if parsed, err := strconv.Atoi(o); err == nil {
			offset = parsed
		}
	}

	filters := make(map[string]interface{})
	if subject := c.Query("subject"); subject != "" {
		filters["subject"] = subject
	}
	if class := c.Query("class"); class != "" {
		if parsed, err := strconv.Atoi(class); err == nil {
			filters["class"] = parsed
		}
	}
	if level := c.Query("level"); level != "" {
		filters["level"] = level
	}
	if status := c.Query("status"); status != "" {
		filters["status"] = status
	}
	if isFree := c.Query("is_free"); isFree != "" {
		if parsed, err := strconv.ParseBool(isFree); err == nil {
			filters["is_free"] = parsed
		}
	}

	courses, total, err := h.courseService.GetAll(limit, offset, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"courses": courses,
		"total":   total,
		"limit":   limit,
		"offset":  offset,
	})
}

func (h *CourseHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	var req dto.UpdateCourseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")
	course, err := h.courseService.Update(uint(id), userID.(uint), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, course)
}

func (h *CourseHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	userID, _ := c.Get("user_id")
	if err := h.courseService.Delete(uint(id), userID.(uint)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}

func (h *CourseHandler) Enroll(c *gin.Context) {
	var req dto.EnrollmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")
	enrollment, err := h.courseService.Enroll(userID.(uint), req.CourseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, enrollment)
}

func (h *CourseHandler) GetMyEnrollments(c *gin.Context) {
	userID, _ := c.Get("user_id")
	enrollments, err := h.courseService.GetUserEnrollments(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"enrollments": enrollments})
}

func ToCourseResponse(course *models.Course) dto.CourseResponse {
	return dto.CourseResponse{
		ID:          course.ID,
		Title:       course.Title,
		Description: course.Description,
		Subject:     course.Subject,
		Class:       course.Class,
		Level:       string(course.Level),
		Status:      string(course.Status),
		Duration:    course.Duration,
		Price:       course.Price,
		IsFree:      course.IsFree,
		TeacherID:   course.TeacherID,
		CreatedAt:   course.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}
