package controllers

import ( 
  "github.com/robfig/revel" 
  "myapp/app/models"
  "net/http"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
  c.checkUser()
	return c.Render()
}

func (c App) checkUser() revel.Result {
  if userInfo := c.Session["UserInfo"]; userInfo == "" {
      c.Flash.Error("Please login in first")
      return c.Redirect(App.Login)
  }
  return nil
}

/*
 * 用户登录 
 */
func (c App) Login() revel.Result { 
  return c.Render() 
}

/*
 * 提交用户登录 
 */
func (c App) UserLog(user *models.UserLogin, w http.ResponseWriter, req *http.Request) revel.Result { 
  dal, _ := models.NewUserDal() 
  defer dal.Close()

  err := dal.CheckLogin(user);
  if err == nil {
    c.Session["UserInfo"] = user.UserName
    return c.RenderJson("登录成功")
  } else {
    return c.RenderJson("登录失败")
  }
}

func (c App) LogOff() revel.Result {
  for k := range c.Session {
		delete(c.Session, k)
	}
  return c.Redirect(App.Index)
} 