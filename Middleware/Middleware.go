package middleware

import (
	db "MECHOPS/Db"
	models "MECHOPS/Models"
	utils "MECHOPS/Utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
)

func Middleware(role ...string) gin.HandlerFunc {
	return func(c *gin.Context) {

		KEY := []byte(os.Getenv("DB_SECRET_KEY"))

		//check Coookie
		CookieToken, err := c.Cookie("Token")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Cookie expired or Missing"})
			c.Abort()
			return
		}

		//unwarp the token with key
		claim := &utils.Claims{}
		Token, err := jwt.ParseWithClaims(CookieToken, claim, func(t *jwt.Token) (interface{}, error) {
			return KEY, nil
		})

		if err == nil && Token.Valid {
			authorized := false
			for _, v := range role {
				if claim.Role == v {
					authorized = true
					break
				}
			}

			if !authorized {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorised"})
				c.Abort()
				return
			}

			c.Set("id", claim.UserId)
			c.Next()
			return
		}

		userID := claim.UserId

		var user models.User
		if err := db.DB.Select("refresh_token").Where("id = ?", userID).First(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to find Refershtoken"})
			c.Abort()
			return
		}
		//unwarp the token with key
		Refershclaim := &utils.Claims{}
		token, err := jwt.ParseWithClaims(user.RefreshToken, Refershclaim, func(t *jwt.Token) (interface{}, error) {
			return KEY, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "RefershToken expired or Missing"})
			c.Abort()
			return
		}

		NewAccessToken, err := utils.AccessToken(Refershclaim.UserId, Refershclaim.Email, Refershclaim.Role)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Unable to Create refersh Token"})
			c.Abort()
			return
		}

		c.SetCookie("Token", NewAccessToken, 3600, "/", "localhost", true, true)
		c.Set("id", Refershclaim.UserId)
		c.Next()
	}
}
