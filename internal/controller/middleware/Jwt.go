package middleware

// func Jwt() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Get the token from the Authorization header
// 		authHeader := c.GetHeader("Authorization")
// 		if authHeader == "" || len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
// 			c.JSON(http.StatusUnauthorized, er.APIError{Error: "missing or invalid Authorization header"})
// 			c.Abort()
// 			return
// 		}

// 		tokenString := authHeader[7:]

// 		// Parse and validate the token
// 		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 			// Validate the signing method
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, jwt.ErrSignatureInvalid
// 			}
// 			return util.GetJwtSecretKey(), nil
// 		})
// 		if err != nil || !token.Valid {
// 			c.JSON(http.StatusUnauthorized, er.APIError{Error: "invalid or expired token"})
// 			c.Abort()
// 			return
// 		}

// 		// Token is valid, store the claims in the context
// 		if claims, ok := token.Claims.(jwt.MapClaims); ok {
// 			c.Set("userID", claims["userID"])
// 		}

// 		// Continue to the next handler
// 		c.Next()
// 	}
// }
