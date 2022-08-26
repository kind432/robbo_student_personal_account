package http

import (
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/spf13/viper"
	"strings"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) userIdentity(c *gin.Context) (id string, role models.Role, err error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return "", models.Role(0), auth.ErrTokenNotFound
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return "", models.Role(0), auth.ErrTokenNotFound
	}

	claims, err := h.authDelegate.ParseToken(headerParts[1], []byte(viper.GetString("auth.access_signing_key")))
	if err != nil {
		return "", models.Role(0), auth.ErrTokenNotFound
	}
	return claims.Id, claims.Role, nil
}
