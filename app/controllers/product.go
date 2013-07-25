package controllers

import ( 
  "github.com/robfig/revel" 
  "myapp/app/models" 
  "myapp/app/utils" 
  "strconv"
  "encoding/base64"
  "encoding/json"
)

type Product struct { 
  *revel.Controller 
}

func (c Product) checkUser() revel.Result {
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
func (c *Product) List(pageindex int) revel.Result { 
  dal, _ := models.NewProductDal() 
  catedal, _ := models.NewCateDal()

  defer dal.Close()
  
  if pageindex<=0 {
    pageindex = 1
  }

  list, _, totalPage := dal.List(pageindex,10)//totalRecord
  var pagelist =make([]int, totalPage)
  for k, _ := range pagelist{
    pagelist[k] = k+1
  }

  /*for _, item := range *list {
    c := catedal.FindByID(item.CategoryID)
    if c.CategoryID > 0 {
      item.CateName = c.CategoryName
      println("found................")
    }
  }*/
  
  catelist := catedal.List()
  return c.Render(list, catelist, pageindex , totalPage, pagelist) 
}

/*
 * 修改 
 */
func (c *Product) Add(id int) revel.Result { 
  dal, _ := models.NewProductDal() 
  catedal, _ := models.NewCateDal()

  defer dal.Close()

  catelist := catedal.List()

  if id != 0 {
    prod := dal.FindByID(id)
    return c.Render(prod, catelist) 
  }
  return c.Render(catelist);
}

/*
 * 提交修改
 */
func (c *Product) Post(prod *models.Product) revel.Result { 
  dal, _ := models.NewProductDal() 
  
  defer dal.Close()
  
  var err error
  var msg string
  if prod.ProductID >0 {
    err = dal.Edit(prod)
    if (err != nil) {
      msg = "修改失败"
    } else {
      msg = "修改成功"
    }
  } else {
    err = dal.Add(prod)
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
func (c *Product) Delete() revel.Result { 
  dal, _ := models.NewProductDal() 
  
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
