package controllers

import ( 
  "github.com/robfig/revel" 
  "myapp/app/models"
  "myapp/app/utils"
  "net/http"
  "encoding/json"
  //"time"
  "encoding/base64"
  //"fmt"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
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

  u, err := dal.CheckLogin(user);
  if err == nil {
    us := models.UserStats{UserID: u.UserID, UserName: u.UserName}
    ustr, err := json.Marshal(us)
    if err != nil {
      panic("登录失败")
    }
    deskey, _ := revel.Config.String("deskey") 
    key := []byte(deskey)
    result, err := utils.DesEncrypt([]byte(ustr), key)
    if err != nil {
      panic(err)
    }

    /*expire := time.Now().AddDate(0, 0, 1)
    cookie := http.Cookie{Name: "userstats", Value: base64.StdEncoding.EncodeToString(result), Expires: expire}
    expire := time.Now().AddDate(0, 0, 1)
    cookie := http.Cookie{Name: "userstats", Value: base64.StdEncoding.EncodeToString(result), Expires: expire, Path: "/"}
    c.SetCookie(&cookie)
    expiration := time.Now().AddDate(0, 0, 1)
    cookie := http.Cookie{Name: "userstats", Value: base64.StdEncoding.EncodeToString(result), Expires: expiration}
    http.SetCookie(w, &cookie)
    http.SetCookie(w, &cookie)*/
    c.Session["UserInfo"] = base64.StdEncoding.EncodeToString(result)
    return c.RenderJson("登录成功")
  } else {
    return c.RenderJson("登录失败")
  }
}

func (c App) LogOff(w http.ResponseWriter, req *http.Request) revel.Result {
  for k := range c.Session {
		delete(c.Session, k)
	}
  /*for c := range c.Request.Cookies {
  
  cookie, err := c.Request.Cookie("userstats")
  
  if(err!=nil){
    println(err.Error())
  } else {
    println(cookie.Name)
    expiration := time.Now().AddDate(0, 0, -1)
    cookie.Expires = expiration
    c.SetCookie(cookie)
  }
  */
  return c.Redirect(App.Index)
} 

func (c App) checkUser() revel.Result {
  userInfo, err := c.Request.Cookie("userstats")
  if err==nil && userInfo.Value != "" {
      deskey, _ := revel.Config.String("deskey") 
      key := []byte(deskey)

      origData, err := base64.StdEncoding.DecodeString(userInfo.Value)
      userstats, _ := utils.DesDecrypt(origData, key)
      if err != nil {
        println("111:"+err.Error())
        //panic(err)
      }
      us := models.UserStats{}
      err = json.Unmarshal(userstats, &us)
      if err != nil {
        panic(err)
      }
      c.Flash.Data["UserName"] = us.UserName
  }
  return nil
}