package models

import (
  "errors"
  "labix.org/v2/mgo/bson"
)

/*
 * 单个实体查找
 */
func (d *AuthDal) FindByID(id int) Auth { 
  result := []Auth{}
  uc := d.session.DB(DbName).C(AuthCollection)
  err := uc.Find(bson.M{"authID": id}).All(&result)
  if err != nil {
        panic(err)
  }
  if len(result)>0 {
    return result[0]
  } else {
    return Auth{}
  }
}

/* 
 * 获取列表
 */
func (d *AuthDal) List() []Auth { 
  result := []Auth{}
  uc := d.session.DB(DbName).C(AuthCollection)
  uc.Find(nil).All(&result)
  
  return result
}

/* 
 * 增加
 */
func (d *AuthDal) Add(auth *Auth) error {
  uc := d.session.DB(DbName).C(AuthCollection)

  i, _ := uc.Find(bson.M{"authName": auth.AuthName}).Count() 
  if i != 0 { 
    return errors.New("权限名已经被使用") 
  }
  
  maxObj := Auth{}
  maxID := 0
  err := uc.Find(nil).Sort("-authID").One(&maxObj)
  if err == nil {
    maxID = maxObj.AuthID
  }

  auth.Id = bson.NewObjectId()
  auth.AuthID = maxID + 1 
	err = uc.Insert(&auth)
  
  return err
}

/* 
 * 修改
 */
func (d *AuthDal) Edit(auth *Auth) error {
  uc := d.session.DB(DbName).C(AuthCollection)

  i, _ := uc.Find(bson.M{"authName": auth.AuthName,"authID": bson.M{"$ne": auth.AuthID}}).Count() 
  if i != 0 { 
    return errors.New("权限名已经被使用") 
  }

  colQuerier := bson.M{"authID": auth.AuthID}
	change := bson.M{"$set": bson.M{"authName": auth.AuthName}}

	err := uc.Update(colQuerier, change)
  return err
}

/* 
 * 删除
 */
func (d *AuthDal) Delete(authID int) error {
  uc := d.session.DB(DbName).C(AuthCollection)

  colQuerier := bson.M{"authID": authID}

	err := uc.Remove(colQuerier)
  return err
}
