package util

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	ope "github.com/skyrocketOoO/gorm-plugin/lib/operator"
	wh "github.com/skyrocketOoO/gorm-plugin/lib/where"
	col "github.com/skyrocketOoO/serverx/internal/gen/column"
	models "github.com/skyrocketOoO/serverx/internal/model"
	"github.com/skyrocketOoO/serverx/internal/service/postgres"
	"github.com/spf13/viper"
)

func GetSalt() []byte {
	return []byte(viper.GetString("salt"))
}

func GetJwtSecretKey() []byte {
	return []byte(viper.GetString("jwt-secret-key"))
}

func GenerateToken(userID uint) (string, error) {
	// Set token claims
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(10000 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(GetJwtSecretKey())
}

func GetOperator(c *gin.Context) (user models.User, err error) {
	operatorName, ok := c.Get("username")
	if !ok {
		return models.User{}, errors.New("username not set in jwt")
	}

	db := postgres.Get()
	if err = db.
		Where(wh.B(col.Users.Name, ope.Eq), operatorName).
		Take(&user).Error; err != nil {
		return
	}

	return user, nil
}
