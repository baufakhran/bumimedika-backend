package middleware

import (
	"net/http"
	"strings"
	"time"

	redisClient "bumimedika-backend/pkg/redis"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	redis "github.com/redis/go-redis/v9"
)

// AuthMiddleware validates JWT token and caches valid tokens in Redis for 3 minutes
func AuthMiddleware(secret string, rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid header format"})
			return
		}

		tokenString := parts[1]

		// ✅ Check Redis cache first
		if redisClient.IsTokenCached(rdb, tokenString) {
			c.Next()
			return
		}

		// ✅ Parse and validate JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// ✅ Cache token for 3 minutes
		_ = redisClient.CacheToken(rdb, tokenString, 3*time.Minute)

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user_id", claims["user_id"])
		}

		c.Next()
	}
}
