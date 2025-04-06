package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
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

// func GetOperator(c *gin.Context) (user models.User, err error) {
// 	operatorName, ok := c.Get("username")
// 	if !ok {
// 		return models.User{}, errors.New("username not set in jwt")
// 	}

// 	db := postgres.Get()
// 	if err = db.
// 		Where(wh.B(col.Users.Name, ope.Eq), operatorName).
// 		Take(&user).Error; err != nil {
// 		return
// 	}

// 	return user, nil
// }
