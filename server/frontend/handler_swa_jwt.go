package frontend

import (
	"context"
	"demo1/service/service_swa"
	"demo1/utils"
	"demo1/utils/response"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"strconv"
	"time"
)


func (s *server) InitRouterSwaBlacklist(Router *gin.RouterGroup) {
	router := Router.Group("jwt").Use(s.OperationRecord())

	router.POST("jsonInBlacklist", s.JsonInBlacklist)
}


func (s *server) JsonInBlacklist(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	jwtList := service_swa.JwtBlacklist{Jwt: token}
	err := s.serviceSwa.Get().CreateBlacklist(c, jwtList)
	if err != nil {
		s.Logger(c).Error("jwt作废失败!", zap.Error(err))
		response.FailWithMessage("jwt作废失败", c)
		return
	}
	s.frontCache.SetDefault(jwtList.Jwt, struct{}{})
	response.OkWithMessage("jwt作废成功", c)
}

func (s *server) JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		if s.IsBlacklist(token) {
			response.FailWithDetailed(gin.H{"reload": true}, "您的帐户异地登陆或令牌失效", c)
			c.Abort()
			return
		}
		j := utils.NewJWT(s.swaConfig.Jwt.SigningKey)
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, utils.TokenExpired) {
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}


		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			dr, _ := utils.ParseDuration(s.swaConfig.Jwt.ExpiresTime)
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(dr))
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
			if s.swaConfig.System.UseMultipoint {
				RedisJwtToken, err := s.GetRedisJWT(newClaims.Username)
				if err != nil {
					s.Logger(c).Error("get redis jwt failed", zap.Error(err))
				} else { // --当之前的取成功时才进行拉黑操作
					err = s.serviceSwa.Get().CreateBlacklist(c, service_swa.JwtBlacklist{Jwt: RedisJwtToken})
					if err != nil {
						return
					}
					s.frontCache.SetDefault(RedisJwtToken, struct{}{})
				}
				_ = s.SetRedisJWT(newToken, newClaims.Username)
			}
		}
		c.Set("claims", claims)
		c.Next()
	}
}

func (s *server) IsBlacklist(jwt string) bool {
	_, ok := s.frontCache.Get(jwt)
	return ok
}


func (s *server) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = s.frontRedis.Get(context.Background(), userName).Result()
	return redisJWT, err
}


func (s *server) SetRedisJWT(jwt string, userName string) (err error) {
	dr, err := utils.ParseDuration(s.swaConfig.Jwt.ExpiresTime)
	if err != nil {
		return err
	}
	timer := dr
	err = s.frontRedis.Set(context.Background(), userName, jwt, timer).Err()
	return err
}
