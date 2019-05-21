package a_user

import (
	"first_project/app/service/user"
	"first_project/library/response"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/util/gvalid"
)

// 用户API管理对象
type Controller struct{}

// 用户注册接口
func (c *Controller) Signup(r *ghttp.Request) {
	if err := s_user.Signup(r.GetPostMap()); err != nil {
		response.Json(r, 1, err.Error())
	} else {
		response.Json(r, 0, "ok")
	}
}

// 用户登录接口
func (c *Controller) SignIn(r *ghttp.Request) {
	data := r.GetPostMap()

	rules := map[string]string{
		"password": "required",
		"passport": "required",
	}

	msgs := map[string]interface{}{
		"passport": "账号不能为空",
		"password": "密码不能为空",
	}

	if e := gvalid.CheckMap(data, rules, msgs); e != nil {
		response.Json(r, 1, e.String())
	}

	if err := s_user.SignIn(data["passport"], data["password"], r.Session); err != nil {
		response.Json(r, 1, err.Error())
	} else {
		response.Json(r, 0, "ok")
	}
}
