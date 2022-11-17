package http

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/courses"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"io/ioutil"
	"log"
	"net/http"
)

type Handler struct {
	authDelegate    auth.Delegate
	coursesDelegate courses.Delegate
}

func NewCoursesHandler(
	authDelegate auth.Delegate,
	coursesDelegate courses.Delegate,
) Handler {
	return Handler{
		authDelegate:    authDelegate,
		coursesDelegate: coursesDelegate,
	}
}

type testCourseResponse struct {
	CourseId string `json:"courseId"`
}

func (h *Handler) InitCourseRoutes(router *gin.Engine) {
	course := router.Group("/course")
	{
		course.POST("/createCourse/:courseId", h.CreateCourse)
		course.GET("/getCourseContent/:courseId", h.GetCourseContent)
		course.GET("/getCoursesByUser", h.GetCoursesByUser)
		course.GET("/getAllPublicCourses/:pageNumber", h.GetAllPublicCourses)
		course.GET("/getEnrollments/:username", h.GetEnrollments)
		course.PUT("/updateCourse", h.UpdateCourse)
		course.DELETE("/deleteCourse/:courseId", h.DeleteCourse)
	}
}

func (h *Handler) UpdateCourse(c *gin.Context) {
	log.Println("Update Course")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	courseHTTP := models.CourseHTTP{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		err = courses.ErrBadRequestBody
		ErrorHandling(err, c)
		return
	}

	err = json.Unmarshal(body, &courseHTTP)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	err = h.coursesDelegate.UpdateCourse(&courseHTTP)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) CreateCourse(c *gin.Context) {
	log.Println("Create Course")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if role < models.UnitAdmin || userIdentityErr != nil {
		//if added to avoid panic
		if role < models.UnitAdmin && userIdentityErr == nil {
			err := errors.New("No access")
			ErrorHandling(err, c)
			return
		}
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	courseId := c.Param("courseId")
	courseHTTP := models.CourseHTTP{}
	courseId, err := h.coursesDelegate.CreateCourse(&courseHTTP, courseId)

	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}

	c.JSON(http.StatusOK, testCourseResponse{
		courseId,
	})
}

func (h *Handler) GetCourseContent(c *gin.Context) {
	log.Println("Get Course Content")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if role == models.Parent || userIdentityErr != nil {
		//if added to avoid panic
		if role == models.Parent && userIdentityErr == nil {
			err := errors.New("No access")
			ErrorHandling(err, c)
			return
		}
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	courseId := c.Param("courseId")
	courseHTTP, err := h.coursesDelegate.GetCourseContent(courseId)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}
	c.JSON(http.StatusOK, courseHTTP)
}

func (h *Handler) GetCoursesByUser(c *gin.Context) {
	log.Println("Get Courses By User")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	coursesHTTP, err := h.coursesDelegate.GetCoursesByUser()
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}
	c.JSON(http.StatusOK, coursesHTTP)
}

func (h *Handler) GetAllPublicCourses(c *gin.Context) {
	log.Println("Get All Public Courses")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	pageNumber := c.Param("pageNumber")
	coursesListHTTP, err := h.coursesDelegate.GetAllPublicCourses(pageNumber)

	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}
	c.JSON(http.StatusOK, coursesListHTTP)
}

func (h *Handler) GetEnrollments(c *gin.Context) {
	log.Println("Get Enrollments")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if role < models.UnitAdmin || userIdentityErr != nil {
		//if added to avoid panic
		if role < models.UnitAdmin && userIdentityErr == nil {
			err := errors.New("No access")
			ErrorHandling(err, c)
			return
		}
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	username := c.Param("username")

	enrollmentsHTTP, err := h.coursesDelegate.GetEnrollments(username)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}
	c.JSON(http.StatusOK, enrollmentsHTTP)
}

func (h *Handler) DeleteCourse(c *gin.Context) {
	log.Println("Delete Course")
	_, _, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	courseId := c.Param("courseId")
	err := h.coursesDelegate.DeleteCourse(courseId)

	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}
}

func ErrorHandling(err error, c *gin.Context) {
	switch err {
	case courses.ErrBadRequest:
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	case courses.ErrInternalServer:
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	case courses.ErrBadRequestBody:
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	case auth.ErrInvalidAccessToken:
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
	case auth.ErrTokenNotFound:
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
}
