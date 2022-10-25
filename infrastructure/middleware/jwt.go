package middleware

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/chichiton/sweaterSocialNetwork/domain"
	"github.com/chichiton/sweaterSocialNetwork/infrastructure/repositories"
	"github.com/gin-gonic/gin"
	"time"
)

type AuthMiddleware struct {
	userRepository *repositories.UserRepositoryImp
}

func NewAuthMiddleware(userRepository *repositories.UserRepositoryImp) *AuthMiddleware {
	return &AuthMiddleware{userRepository: userRepository}
}

func (m AuthMiddleware) GetInstance(secretKey string, identityKey string) (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "sweater social network",
		Key:         []byte(secretKey),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*domain.UserId); ok {
				return jwt.MapClaims{
					identityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			userId := claims[identityKey].(float64)
			v := domain.UserId(userId)

			return &v
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var userAuth domain.Auth
			if err := c.ShouldBind(&userAuth); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			login := userAuth.Login
			password := userAuth.Password

			auth, err := m.userRepository.GetUserAuthByLogin(userAuth.Login)
			if err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			if login == auth.Login && password.CheckPasswordHash(auth.PasswordHash) {

				/*userProfile, err := repositories.GetUserProfile(auth.UserId)
				if err != nil {
					return nil, jwt.ErrMissingLoginValues
				}*/

				return &auth.UserId, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			/*if v, ok := data.(*User); ok && v.UserName == "admin" {
				return true
			}

			return false*/

			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
}
