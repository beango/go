package controllers

import ( 
  "github.com/robfig/revel" 
  "myapp/app/models" 
  "myapp/app/utils" 
  "strconv"
  "encoding/base64"
  "encoding/json"
  "net/http"
)

type UserAccount struct { 
  *revel.Controller 
}

func (c UserAccount) checkUser() revel.Result {
  userInfo := c.Session["UserInfo"]
  if userInfo == "" {
      c.Flash.Error("Please login in first")
      return c.Redirect(App.Login)
  }

  deskey, _ := revel.Config.String("deskey") 
  key := []byte(deskey)

  origData, err := base64.StdEncoding.DecodeString(userInfo)
  userstats, _ := utils.DesDecrypt(origData, key)
  if err != nil {
    panic(err)
  }
  us := models.UserStats{}
  err = json.Unmarshal(userstats, &us)
  if err != nil {
    panic(err)
  }
  c.Flash.Data["UserName"] = us.UserName
  return nil
}

/*
 * 用户列表 
 */
func (c *UserAccount) List(w http.ResponseWriter, req *http.Request) revel.Result { 
  dal, _ := models.NewUserDal() 

  defer dal.Close()
  
  list := dal.List()
  return c.Render(list) 
}

/*
 * 修改用户资料 
 */
func (c *UserAccount) Add(id int) revel.Result { 
  dal, _ := models.NewUserDal() 
  
  defer dal.Close()

  if id != 0 {
    entity := dal.FindByID(id)
    return c.Render(entity) 
  }
    return c.Render();
}

/*
 * 提交用户资料修改
 */
func (c *UserAccount) UserPost(user *models.User) revel.Result { 
  dal, _ := models.NewUserDal() 
  
  defer dal.Close()
  
  var err error
  var msg string
  if user.UserID >0 {
    err = dal.Edit(user)
    if (err != nil) {
      msg = "用户资料修改失败"
    } else {
      msg = "用户资料修改成功"
    }
  } else {
    err = dal.Add(user)
    if (err != nil) {
      msg = "用户资料增加失败：" + err.Error()
    } else {
      msg = "用户资料增加成功"
    }
  }
  
  return c.RenderJson(msg) 
}

/*
 * 删除用户资料 
 */
func (c *UserAccount) Delete() revel.Result { 
  dal, _ := models.NewUserDal() 
  
  defer dal.Close()
  
  if c.Params.Get("id") != "" {
    println(c.Params.Get("id"))

    userid, err := strconv.Atoi(c.Params.Get("id"))
    if err != nil {
      return c.RenderJson(c.Params.Get("id")+"转换为整数失败！")
    }
    err = dal.Delete(userid)
    if err == nil {
      return c.RenderJson("用户删除成功");
    } else {
      return c.RenderJson("用户删除失败");
    }
  } else {
    return c.RenderJson("提交失败");
  }
}

/*
 * 用户注册 
 */
func (c *UserAccount) Register() revel.Result { 
  return c.Render() 
}

/*
 * 提交用户注册资料 
 */
func (c *UserAccount) PostRegister(user *models.User) revel.Result { 
  return c.Render() 
}