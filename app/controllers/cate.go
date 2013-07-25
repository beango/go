package controllers

import ( 
  "github.com/robfig/revel" 
  "myapp/app/models" 
  "myapp/app/utils" 
  "strconv"
  "encoding/base64"
  "encoding/json"
)

type Cate struct { 
  *revel.Controller 
}

func (c Cate) checkUser() revel.Result {
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
func (c *Cate) List() revel.Result { 
  dal, _ := models.NewCateDal() 

  defer dal.Close()
  
  list := dal.List()
  return c.Render(list) 
}

/*
 * 修改 
 */
func (c *Cate) Add(id int) revel.Result { 
  dal, _ := models.NewCateDal() 
  
  defer dal.Close()

  if id != 0 {
    cate := dal.FindByID(id)
    return c.Render(cate) 
  }
    return c.Render();
}

/*
 * 提交修改
 */
func (c *Cate) Post(cate *models.Cate) revel.Result { 
  dal, _ := models.NewCateDal() 
  
  defer dal.Close()
  
  var err error
  var msg string
  if cate.CategoryID >0 {
    err = dal.Edit(cate)
    if (err != nil) {
      msg = "修改失败"
    } else {
      msg = "修改成功"
    }
  } else {
    err = dal.Add(cate)
    if (err != nil) {
      msg = "增加失败：" + err.Error()
    } else {
      msg = "增加成功"
    }
  }
  
  return c.RenderJson(msg) 
}

/*
 * 删除
 */
func (c *Cate) Delete() revel.Result { 
  dal, _ := models.NewCateDal() 
  
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
