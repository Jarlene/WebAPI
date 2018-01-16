package login

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
	"../base"
	"github.com/wspl/creeper"
)

type LoginInfo struct {
	UserName string
	From     string
	UserId   string
	ExtendId string
	Avatar   string
	Gender   int
	Mobile   string
	Token    string
	base.Info
}

func Login(c *gin.Context) {
	info := LoginInfo{}
	errorSts := base.LoginParamError()
	token, err := c.Cookie("login_token") // get login token
	redis, _ := base.NewRedisClient()
	if err != nil || !redis.Exist(token) {
		login_type, err := c.Cookie("login_type")
		if err != nil {
			c.JSON(http.StatusOK, errorSts)
			return
		}

		logintype, err := strconv.Atoi(login_type)
		if err != nil {
			c.JSON(http.StatusOK, errorSts)
			return
		}
		if logintype == 0 {
			name, _ := c.Cookie("name")
			pass, _ := c.Cookie("pass")
			sql := "select user_id from user where user_name='" + name + "' and pass='" + creeper.MD5(pass) + "'"
			conf, _ := base.NewMysqlConf("root:password@/database")
			res, _ := conf.Sql(sql)

			info.UserName = name
			if res != nil {
				ss := res.([]map[string][]byte)
				info.UserId = string(ss[0]["user_id"])
			}
			info.Token = base.GenToken(info.UserId, info.UserName)

		} else if logintype == 1 {

		} else if logintype == 2 {

		} else {

		}
		s := base.Success()
		info.Code = s.Code
		info.Msg = s.Msg

		redis.SetTimeout(info.Token, info, 1200)
		c.JSON(http.StatusOK, info)
	} else {
		res, err := redis.Get(token)
		errorSts := base.LoginTokenError()
		if err != nil {
			c.JSON(http.StatusOK, errorSts)
			return
		}
		c.JSON(http.StatusOK, res)
	}

}

func Logout(c *gin.Context) {
	token, err := c.Cookie("login_token") // get login token
	info := base.Success()
	if err != nil {
		info.Code = 22001
		info.Msg = "please login first"
		c.JSON(http.StatusOK, info)
		return
	}
	redis, err := base.NewRedisClient()
	if redis != nil {
		defer redis.Close()
	}
	if err == nil {
		redis.Del(token)
	}
	c.JSON(http.StatusOK, info)
}


func Register(c *gin.Context) {
	email, err := c.Cookie("email")
	ests := base.RegisterParamError()
	if err != nil {
		ests.Msg = "param email not find"
		c.JSON(http.StatusOK, ests)
		return
	}

	username, err := c.Cookie("username")
	if err != nil {
		ests.Msg = "param username not find"
		c.JSON(http.StatusOK, ests)
		return
	}

	pass, err:= c.Cookie("pass")
	if err != nil {
		ests.Msg = "param password not find"
		c.JSON(http.StatusOK, ests)
		return
	}

	conf, _ := base.NewMysqlConf("root:password@/database")
	sql := "select user_id from user where user_email ='" + email + "'"
	res, err :=conf.Sql(sql)
	if err == nil && res == nil {
		sql = "insert into user(user_name, user_email, pass) values ('" + username + "' , '" + email + "' , '" + creeper.MD5(pass) + "')"
		res, _ := conf.Insert(sql)
		userid := strconv.FormatInt(res.(int64), 10)
		token := base.GenToken(userid, username)
		info := LoginInfo{}
		s := base.Success()
		info.Code = s.Code
		info.Msg = s.Msg
		info.UserName = username
		info.Token = token
		info.UserId = userid
		c.JSON(http.StatusOK, info)
	} else {
		ests.Msg = "email has register, please login"
		c.JSON(http.StatusOK, ests)
	}

}