package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/robboUnits"
	"log"
	"net/http"
)

type Handler struct {
	authDelegate       auth.Delegate
	robboUnitsDelegate robboUnits.Delegate
}

func NewRobboUnitsHandler(authDelegate auth.Delegate, robboUnits robboUnits.Delegate) Handler {
	return Handler{
		authDelegate:       authDelegate,
		robboUnitsDelegate: robboUnits,
	}
}

func (h *Handler) InitRobboUnitsRoutes(router *gin.Engine) {
	robboUnits := router.Group("/robboUnits")
	{
		robboUnits.POST("/", h.CreateRobboUnit)
		robboUnits.GET("/:robboUnitId", h.GetRobboUnitById)
		robboUnits.GET("/unitAdmin", h.GetRobboUnitsByUnitAdminId)
		robboUnits.GET("/", h.GetAllRobboUnits)
		robboUnits.PUT("/", h.UpdateRobboUnit)
		robboUnits.DELETE("/:robboUnitId", h.DeleteRobboUnit)
	}
}

func (h *Handler) CreateRobboUnit(c *gin.Context) {
	fmt.Println("Create Robbo Unit")

	robboUnitHttp := models.RobboUnitHTTP{}
	if err := c.BindJSON(&robboUnitHttp); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	robboUnitId, err := h.robboUnitsDelegate.CreateRobboUnit(&robboUnitHttp)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"robboUnitId": robboUnitId,
	})
}

func (h *Handler) GetRobboUnitById(c *gin.Context) {
	fmt.Println("Get RobboUnit By Id")
	robboUnitId := c.Param("robboUnitId")

	robboUnit, err := h.robboUnitsDelegate.GetRobboUnitById(robboUnitId)

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, robboUnit)
}

func (h *Handler) GetAllRobboUnits(c *gin.Context) {
	fmt.Println("Get all robboUnits")

	robboUnits, err := h.robboUnitsDelegate.GetAllRobboUnit()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, robboUnits)
}

func (h *Handler) GetRobboUnitsByUnitAdminId(c *gin.Context) {
	fmt.Println("GetRobboUnitsByUnitAdminId")
	id, role, identityErr := h.userIdentity(c)

	if identityErr != nil || role != models.UnitAdmin {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	robboUnits, err := h.robboUnitsDelegate.GetRobboUnitsByUnitAdminId(id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, robboUnits)
}

func (h *Handler) UpdateRobboUnit(c *gin.Context) {
	fmt.Println("Update RobboUnit")

	robboUnitHttp := models.RobboUnitHTTP{}

	if err := c.BindJSON(&robboUnitHttp); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := h.robboUnitsDelegate.UpdateRobboUnit(&robboUnitHttp)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) DeleteRobboUnit(c *gin.Context) {
	fmt.Println("Delete RobboUnit")

	robboUnitId := c.Param("robboUnitId")
	err := h.robboUnitsDelegate.DeleteRobboUnit(robboUnitId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
