package middlewares


import (
    "net/http"
    "gopkg.in/gin-gonic/gin.v1"
    "github.com/dgrijalva/jwt-go"
    "github.com/dgrijalva/jwt-go/request"
)

func Auth(secret string) gin.HandlerFunc {
    return func(c *gin.Context) {
        _, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
            b := ([]byte(secret))
            return b, nil
        })

        if err != nil {
            c.AbortWithError(http.StatusUnauthorized, err)
        }
    }
}
