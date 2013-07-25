package models

import (
  "errors"
  //"code.google.com/p/go.crypto/bcrypt"
  "labix.org/v2/mgo/bson"
)

/*
 * 单个实体查找
 */
func (d *UserDal) FindByID(id int) User { 
  result := []User{}
  uc := d.session.DB(DbName).C(UserCollection)
  err := uc.Find(bson.M{"userID": id}).All(&result)
  if err != nil {
        panic(err)
  }
  if len(result)>0 {
    return result[0]
  } else {
    return User{}
  }
}

/*
 * 检查用户登录
 */
func (d *UserDal) CheckLogin(user *UserLogin) (User, error) { 
  result := User{}

  uc := d.session.DB(DbName).C(UserCollection)
  err := uc.Find(bson.M{"userName": user.UserName, "userPwd": user.UserPwd}).One(&result)
  if err == nil {
    return result, nil
  } else {
    return result, err
  }
}

/* 
 * 获取列表
 */
func (d *UserDal) List() []User { 
  result := []User{}
  uc := d.session.DB(DbName).C(UserCollection)
  uc.Find(nil).All(&result)
  
  return result
}

/* 
 * 增加
 */
func (d *UserDal) Add(user *User) error {
  uc := d.session.DB(DbName).C(UserCollection)

  i, _ := uc.Find(bson.M{"userName": user.UserName}).Count() 
  if i != 0 { 
    return errors.New("用户名已经被使用") 
  }
  
  maxUser := User{}
  maxUserID := 0
  err := uc.Find(nil).Sort("-userID").One(&maxUser)
  if err == nil {
    maxUserID = maxUser.UserID
  }

  user.Id = bson.NewObjectId()
  user.UserID = maxUserID + 1 
	err = uc.Insert(&user)//&User{UserID:maxUserID+1, UserName: user.UserName, UserPwd: user.UserPwd}
  
  return err
}

/* 
 * 修改
 */
func (d *UserDal) Edit(user *User) error {
  uc := d.session.DB(DbName).C(UserCollection)

  i, _ := uc.Find(bson.M{"userName": user.UserName,"userID": bson.M{"$ne": user.UserID}}).Count() 
  if i != 0 { 
    return errors.New("用户名已经被使用") 
  }

  colQuerier := bson.M{"userID": user.UserID}
	change := bson.M{"$set": bson.M{"userName": user.UserName, "userPwd": user.UserPwd}}

	err := uc.Update(colQuerier, change)
  return err
}

/* 
 * 删除
 */
func (d *UserDal) Delete(userID int) error {
  uc := d.session.DB(DbName).C(UserCollection)

  colQuerier := bson.M{"userID": userID}

	err := uc.Remove(colQuerier)
  return err
}

/* 
 * 用户注册
 */
func (d *UserDal) RegisterUser(mu *User) error { 
  println("111")
  uc := d.session.DB(DbName).C(UserCollection)

  //先检查email和nickname是否已经被使用 
  i, _ := uc.Find(bson.M{"nickname": mu.UserName}).Count() 
  if i != 0 { 
    return errors.New("用户昵称已经被使用") 
  }

  /*
  var u User 
  u.Email = mu.Email 
  u.Nickname = mu.Nickname 
  u.Password, _ = bcrypt.GenerateFromPassword([]byte(mu.Password), bcrypt.DefaultCost)
  */

  err := uc.Insert(mu)

  return err 
}