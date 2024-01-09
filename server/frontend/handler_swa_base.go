package frontend

import (
	"demo1/service/service_swa"
	"demo1/utils"
	"demo1/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"time"
)

var store = base64Captcha.DefaultMemStore


func (s *server) InitRouterSwaBase(Router *gin.RouterGroup) {
	router := Router.Group("base")

	router.POST("captcha", s.Captcha)
	router.POST("login", s.Login)
	router.POST("getToken", s.GetToken)
}


type captchaResponse struct {
	CaptchaId     string `json:"captchaId"`
	PicPath       string `json:"picPath"`
	CaptchaLength int    `json:"captchaLength"`
	OpenCaptcha   bool   `json:"openCaptcha"`
}

type ModLogin struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}

type loginResponse struct {
	User      service_swa.SysUser `json:"user"`
	Token     string              `json:"token"`
	ExpiresAt int64               `json:"expiresAt"`
}


func (s *server) Captcha(c *gin.Context) {
	openCaptcha := s.swaConfig.Captcha.OpenCaptcha
	openCaptchaTimeOut := s.swaConfig.Captcha.OpenCaptchaTimeOut
	key := c.ClientIP()
	v, ok := s.frontCache.Get(key)
	if !ok {
		s.frontCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	var oc bool
	if openCaptcha == 0 || openCaptcha < utils.InterfaceToInt(v) {
		oc = true
	}
	driver := base64Captcha.NewDriverDigit(s.swaConfig.Captcha.ImgHeight, s.swaConfig.Captcha.ImgWidth, s.swaConfig.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := cp.Generate()
	if err != nil {
		s.Logger(c).Error("验证码获取失败!", zap.Error(err))
		response.FailWithMessage("验证码获取失败", c)
		return
	}
	response.OkWithDetailed(captchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: s.swaConfig.Captcha.KeyLong,
		OpenCaptcha:   oc,
	}, "验证码获取成功", c)
}

func (s *server) Login(c *gin.Context) {
	var l ModLogin
	err := c.ShouldBindJSON(&l)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(l, utils.LoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	key := c.ClientIP()
	openCaptcha := s.swaConfig.Captcha.OpenCaptcha
	openCaptchaTimeOut := s.swaConfig.Captcha.OpenCaptchaTimeOut
	v, ok := s.frontCache.Get(key)
	if !ok {
		s.frontCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	var oc = openCaptcha == 0 || openCaptcha < utils.InterfaceToInt(v)

	if !oc || store.Verify(l.CaptchaId, l.Captcha, true) {
		u := service_swa.SysUser{Username: l.Username, Password: l.Password}
		user, err := s.serviceSwa.Get().Login(c, u)
		if err != nil {
			s.Logger(c).Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			err := s.frontCache.Increment(key, 1)
			if err != nil {
				s.Logger(c).Info("登陆失败次数加1失败")
			}
			response.FailWithMessage("用户名不存在或者密码错误", c)
			return
		}
		if user.Enable != 1 {
			s.Logger(c).Error("登陆失败! 用户被禁止登录!")
			err := s.frontCache.Increment(key, 1)
			if err != nil {
				s.Logger(c).Info("登陆失败次数加1失败")
			}
			response.FailWithMessage("用户被禁止登录", c)
			return
		}
		s.TokenNext(c, user)
		return
	}
}

func (s *server) GetToken(c *gin.Context) {
	var postUser ModLogin
	err := c.ShouldBindJSON(&postUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	u := service_swa.SysUser{Username: postUser.Username, Password: postUser.Password}
	user, err1 := s.serviceSwa.Get().Login(c, u)
	if err1 != nil {
		s.Logger(c).Error("API获取用户登陆失败! 用户名不存在或者密码错误!", zap.Error(err1))
		response.FailWithMessage("用户名不存在或者密码错误", c)
		return
	}
	if user.Enable != 1 {
		s.Logger(c).Error("API获取用户登陆失败! 用户被禁止登录!")
		response.FailWithMessage("用户被禁止登录", c)
		return
	}
	s.TokenNext(c, user)

	return
}

func (s *server) TokenNext(c *gin.Context, user service_swa.SysUser) {
	j := utils.NewJWT(s.swaConfig.Jwt.SigningKey)
	claims := j.CreateClaims(utils.BaseClaims{
		UUID:     user.UUID,
		ID:       user.ID,
		NickName: user.NickName,
		Username: user.Username,
		RoleId:   user.RoleId,
	}, s.swaConfig.Jwt.BufferTime, s.swaConfig.Jwt.ExpiresTime, s.swaConfig.Jwt.Issuer)

	token, err := j.CreateToken(claims)
	if err != nil {
		s.Logger(c).Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}

	if !s.swaConfig.System.UseMultipoint {
		response.OkWithDetailed(loginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功0", c)
		return
	}

	if jwtStr, err := s.GetRedisJWT(user.Username); err == redis.Nil {
		if err := s.SetRedisJWT(token, user.Username); err != nil {
			s.Logger(c).Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(loginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功1", c)
	} else if err != nil {
		s.Logger(c).Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT service_swa.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := s.serviceSwa.Get().CreateBlacklist(c, blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := s.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		s.frontCache.SetDefault(blackJWT.Jwt, struct{}{})
		response.OkWithDetailed(loginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功2", c)
	}
}

