package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

type BaseClaims struct {
	UUID     uuid.UUID
	ID       uint
	Username string
	NickName string
	RoleId   uint
}
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

func GetClaims(c *gin.Context, key string) (*CustomClaims, error) {
	token := c.Request.Header.Get("x-token")
	j := NewJWT(key)
	claims, err := j.ParseToken(token)
	if err != nil {
		return claims, errors.Errorf("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}


func GetUserID(c *gin.Context, key string) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c, key); err != nil {
			return 0
		} else {
			return cl.BaseClaims.ID
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse.BaseClaims.ID
	}
}


func GetUserUuid(c *gin.Context, key string) uuid.UUID {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c, key); err != nil {
			return uuid.UUID{}
		} else {
			return cl.UUID
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse.UUID
	}
}

func GetUserRoleId(c *gin.Context, key string) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c, key); err != nil {
			return 0
		} else {
			return cl.RoleId
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse.RoleId
	}
}


func GetUserInfo(c *gin.Context, key string) *CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c, key); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse
	}
}
