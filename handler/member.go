package handler

import (
	"errors"
	"net/http"
	"strconv"
	"technical-test/domain/requests/member"
	"technical-test/pkg/token"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *Handler) initMemberGroup(api *gin.Engine) {
	member := api.Group("/api/member")

	member.Use(authMiddleware(h.tokenMaker))
	member.POST("/create", h.handlerCreateMember)
	member.GET("/showAll", h.handlerGetMemberAll)
	member.GET("/showById/:id", h.handlerGetMemberById)
	// member.GET("/showByUserId/:userId", h.handlerGetMemberByUserId)
}

// handleCreateMember
// @Summary Create new member
// @Description Create new member
// @Tags member
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param data body member.CreateMemberRequest true "data"
// @Success 201 {object} responses.MemberResponse
// @Failure 400 {object} responses.ErrorMessage
// @Failure 500 {object} responses.ErrorMessage
// @Router /member/create [post]
func (h *Handler) handlerCreateMember(c *gin.Context) {
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	userId, err := strconv.Atoi(authPayload.Subject)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	privilege := authPayload.Privilege
	if privilege == "member" {
		c.JSON(http.StatusUnauthorized, errorResponse(errors.New("user with privilege member cannot create new member")))
		return
	}

	var body member.CreateMemberRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := body.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	genderName, err := h.services.Gender.GetGenderByName(body.Gender)
	if err != nil {
		// create new gender if not exist
		if err == gorm.ErrRecordNotFound {
			genderName, err = h.services.Gender.CreateGender(body.Gender)
			if err != nil {
				c.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	levelName, err := h.services.Level.GetLevelByName(body.Level)
	if err != nil {
		// create new level if not exist
		if err == gorm.ErrRecordNotFound {
			levelName, err = h.services.Level.CreateLevel(body.Level)
			if err != nil {
				c.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	res, err := h.services.Member.CreateMember(userId, genderName.ID, levelName.ID, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// handleGetMemberAll
// @Summary Get all member
// @Description Get all member
// @Tags member
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} responses.MemberResponse
// @Failure 400 {object} responses.ErrorMessage
// @Failure 500 {object} responses.ErrorMessage
// @Router /member/showAll [get]
func (h *Handler) handlerGetMemberAll(c *gin.Context) {
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)

	privilege := authPayload.Privilege
	if privilege == "member" {
		c.JSON(http.StatusUnauthorized, errorResponse(errors.New("user with privilege member cannot get member data")))
		return
	}

	members, err := h.services.Member.GetMemberAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, members)
}

// handleGetMemberById
// @Summary Get member by id
// @Description Get member by id
// @Tags member
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "id"
// @Success 200 {object} responses.MemberResponse
// @Failure 400 {object} responses.ErrorMessage
// @Failure 500 {object} responses.ErrorMessage
// @Router /member/showById/{id} [get]
func (h *Handler) handlerGetMemberById(c *gin.Context) {
	idStr := c.Param("id")
	memberId, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	userId, err := strconv.Atoi(authPayload.Subject)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	privilege := authPayload.Privilege

	if privilege == "member" {
		res, err := h.services.Member.GetMemberByUserId(userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		for _, v := range *res {
			if v.ID == memberId {
				c.JSON(http.StatusOK, v)
				return
			}
		}
		c.JSON(http.StatusNotFound, errorResponse(errors.New("member not found")))
	} else {
		member, err := h.services.Member.GetMemberById(memberId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		c.JSON(http.StatusOK, member)
	}

}

// func (h *Handler) handlerGetMemberById(c *gin.Context) {

// }

// func (h *Handler) handlerGetMemberByUserId(c *gin.Context) {

// }
