package delegate

import (
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"log"
	"strings"
)

type AuthDelegateImpl struct {
	auth.UseCase
}

type AuthDelegateModule struct {
	fx.Out
	auth.Delegate
}

func SetupAuthDelegate(usecase auth.UseCase) AuthDelegateModule {
	return AuthDelegateModule{
		Delegate: &AuthDelegateImpl{usecase},
	}
}

func (s *AuthDelegateImpl) SignIn(email, password string, role uint) (accessToken, refreshToken string, err error) {
	return s.UseCase.SignIn(email, password, role)
}

func (s *AuthDelegateImpl) SignUp(userHttp *models.UserHTTP) (accessToken, refreshToken string, err error) {
	userCore := userHttp.ToCore()
	return s.UseCase.SignUp(&userCore)
}

//func (s *AuthDelegateImpl) ParseToken(token string, key []byte) (claims *models.UserClaims, err error) {
//	return s.UseCase.ParseToken(token, key)
//}

func (s *AuthDelegateImpl) RefreshToken(refreshToken string) (newAccessToken string, err error) {
	return s.UseCase.RefreshToken(refreshToken)
}

const (
	authorizationHeader = "Authorization"
)

func (s *AuthDelegateImpl) UserIdentity(c *gin.Context) (id string, role models.Role, err error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return "", models.Anonymous, auth.ErrTokenNotFound
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return "", models.Anonymous, auth.ErrTokenNotFound
		return
	}

	claims, err := s.UseCase.ParseToken(headerParts[1], []byte(viper.GetString("auth.access_signing_key")))
	if err != nil {
		log.Println(err)
		return "", models.Anonymous, auth.ErrInvalidAccessToken
	}
	return claims.Id, claims.Role, nil
}

func (s *AuthDelegateImpl) UserAccess(currentRole models.Role, roles []models.Role) (err error) {
	for _, role := range roles {
		if currentRole == role {
			return nil
		}
	}
	err = auth.ErrNotAccess
	return
}

//func (s *AuthDelegateImpl) UserIdentityGraphQL(c *context.Context) (id string, role models.Role, err error) {
//	header := c.(authorizationHeader)
//	if header == "" {
//		return "", models.Anonymous, auth.ErrTokenNotFound
//	}
//
//	headerParts := strings.Split(header, " ")
//	if len(headerParts) != 2 {
//		return "", models.Anonymous, auth.ErrTokenNotFound
//		return
//	}
//
//	claims, err := s.UseCase.ParseToken(headerParts[1], []byte(viper.GetString("auth.access_signing_key")))
//	if err != nil {
//		return "", models.Anonymous, auth.ErrInvalidAccessToken
//	}
//	return claims.Id, claims.Role, nil
//}
