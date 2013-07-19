package controllers

import (
  "strconv"
  "github.com/robfig/revel"
)

func init() {
	//revel.OnAppStart(Init)
	//revel.InterceptMethod(App.checkUser, revel.BEFORE)
	revel.InterceptMethod(UserAccount.checkUser, revel.BEFORE)
	revel.InterceptMethod(Role.checkUser, revel.BEFORE)
	revel.InterceptMethod(Auth.checkUser, revel.BEFORE)
  revel.InterceptMethod(Product.checkUser, revel.BEFORE)

	//revel.InterceptMethod(Application.AddUser, revel.BEFORE)
	//revel.InterceptMethod(Hotels.checkUser, revel.BEFORE)
	//revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
	//revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)

  revel.TemplateFuncs["isNotNil"] = func(a interface{}) bool { return a != nil }
  revel.TemplateFuncs["lte"] = func(a, b int) bool { return a <= b }
  revel.TemplateFuncs["gte"] = func(a, b int) bool { return a >= b }
  revel.TemplateFuncs["de"] = func(a, b int) int { return a - b }
  revel.TemplateFuncs["equ"] = func(a, b interface{}) bool { return a == b }
  revel.TemplateFuncs["convert2int"] = func(a interface{}) int { 
    res, _ := strconv.Atoi("2")
    return res 
  }
}