package controllers

import ( 
  "github.com/robfig/revel" 
  "myapp/app/models" 
  "myapp/app/utils" 
  "strconv"
  "encoding/base64"
  "encoding/json"
)

type Role struct { 
  *revel.Controller 
}

func (c Role) checkUser() revel.Result {
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
func (c *Role) List() revel.Result { 
  dal, _ := models.NewRoleDal() 

  defer dal.Close()
  
  list := dal.List()
  return c.Render(list) 
}

/*
 * 修改
 */
func (c *Role) Add(id int) revel.Result { 
  dal, _ := models.NewRoleDal() 
  
  defer dal.Close()

  if id != 0 {
    role := dal.FindByID(id)
    return c.Render(role) 
  }
    return c.Render();
}

/*
 * 提交修改
 */
func (c *Role) Post(role *models.Role) revel.Result { 
  dal, _ := models.NewRoleDal() 
  
  defer dal.Close()
  
  var err error
  var msg string
  if role.RoleID >0 {
    err = dal.Edit(role)
    if (err != nil) {
      msg = "角色修改失败"
    } else {
      msg = "角色修改成功"
    }
  } else {
    err = dal.Add(role)
    if (err != nil) {
      msg = "角色增加失败：" + err.Error()
    } else {
      msg = "角色增加成功"
    }
  }
  
  return c.RenderJson(msg) 
}

/*
 * 删除
 */
func (c *Role) Delete() revel.Result { 
  dal, _ := models.NewRoleDal() 
  
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
