package controllers

import ( 
  "github.com/robfig/revel" 
  "myapp/app/models" 
  "myapp/app/utils" 
  "strconv"
  "encoding/base64"
  "encoding/json"
)

type Auth struct { 
  *revel.Controller 
}

func (c Auth) checkUser() revel.Result {
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
 * 列表 
 */
func (c *Auth) List() revel.Result { 
  dal, _ := models.NewAuthDal() 

  defer dal.Close()
  
  list := dal.List()
  return c.Render(list) 
}

/*
 * 修改 
 */
func (c *Auth) Add(id int) revel.Result { 
  dal, _ := models.NewAuthDal() 
  
  defer dal.Close()
  
  if id != 0 {
    auth := dal.FindByID(id)
    return c.Render(auth) 
  }
    return c.Render();
}

/*
 * 提交修改
 */
func (c *Auth) Post(auth *models.Auth) revel.Result { 
  dal, _ := models.NewAuthDal() 
  
  defer dal.Close()
  
  var err error
  var msg string
  if auth.AuthID >0 {
    err = dal.Edit(auth)
    if (err != nil) {
      msg = "权限修改失败"
    } else {
      msg = "权限修改成功"
    }
  } else {
    err = dal.Add(auth)
    if (err != nil) {
      msg = "权限增加失败：" + err.Error()
    } else {
      msg = "权限增加成功"
    }
  }
  
  return c.RenderJson(msg) 
}

/*
 * 删除
 */
func (c *Auth) Delete() revel.Result { 
  dal, _ := models.NewAuthDal() 
  
  defer dal.Close()
  
  if c.Params.Get("id") != "" {
    objid, err := strconv.Atoi(c.Params.Get("id"))
    if err != nil {
      return c.RenderJson(c.Params.Get("id")+"转换为整数失败！")
    }
    err = dal.Delete(objid)
    if err == nil {
      return c.RenderJson("删除成功");
    } else {
      return c.RenderJson("删除失败");
    }
  } else {
    return c.RenderJson("提交失败");
  }
}
