package cm

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	ope "github.com/skyrocketOoO/gorm-plugin/lib/operator"
	wh "github.com/skyrocketOoO/gorm-plugin/lib/where"
	col "github.com/skyrocketOoO/web-server-template/internal/gen/column"
	"github.com/skyrocketOoO/web-server-template/internal/model"
	"github.com/skyrocketOoO/web-server-template/internal/service/db"
	"github.com/spf13/viper"
)

func GetSalt() []byte {
	return []byte(viper.GetString("salt"))
}

func GetJwtSecretKey() []byte {
	return []byte(viper.GetString("jwt-secret-key"))
}

func GenerateToken(username string) (string, error) {
	// Set token claims
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(10000 * time.Hour).Unix(),
		"iss":      "alarm-system",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(GetJwtSecretKey())
}

func GetOperator(c *gin.Context) (user model.User, err error) {
	operatorName, ok := c.Get("username")
	if !ok {
		return model.User{}, errors.New("username not set in jwt")
	}

	db := db.Get()
	if err = db.
		Where(wh.B(col.Users.Name, ope.Eq), operatorName).
		Take(&user).Error; err != nil {
		return
	}

	return user, nil
}
