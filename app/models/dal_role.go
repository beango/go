package models

import (
  "errors"
  "labix.org/v2/mgo/bson"
)

/*
 * 单个实体查找
 */
func (d *RoleDal) FindByID(id int) Role { 
  result := []Role{}
  uc := d.session.DB(DbName).C(RoleCollection)
  err := uc.Find(bson.M{"roleID": id}).All(&result)
  if err != nil {
        panic(err)
  }
  if len(result)>0 {
    return result[0]
  } else {
    return Role{}
  }
}

/* 
 * 获取列表
 */
func (d *RoleDal) List() []Role { 
  result := []Role{}
  uc := d.session.DB(DbName).C(RoleCollection)
  uc.Find(nil).All(&result)
  
  return result
}

/* 
 * 增加
 */
func (d *RoleDal) Add(role *Role) error {
  uc := d.session.DB(DbName).C(RoleCollection)

  i, _ := uc.Find(bson.M{"roleName": role.RoleName}).Count() 
  if i != 0 { 
    return errors.New("角色名已经被使用") 
  }
  
  maxObj := Role{}
  maxID := 0
  err := uc.Find(nil).Sort("-roleID").One(&maxObj)
  if err == nil {
    maxID = maxObj.RoleID
  }

  role.Id = bson.NewObjectId()
  role.RoleID = maxID + 1 
	err = uc.Insert(&role)
  
  return err
}

/* 
 * 修改
 */
func (d *RoleDal) Edit(role *Role) error {
  uc := d.session.DB(DbName).C(RoleCollection)

  i, _ := uc.Find(bson.M{"roleName": role.RoleName,"roleID": bson.M{"$ne": role.RoleID}}).Count() 
  if i != 0 { 
    return errors.New("角色名已经被使用") 
  }

  colQuerier := bson.M{"roleID": role.RoleID}
	change := bson.M{"$set": bson.M{"roleName": role.RoleName}}

	err := uc.Update(colQuerier, change)
  return err
}

/* 
 * 删除
 */
func (d *RoleDal) Delete(roleID int) error {
  uc := d.session.DB(DbName).C(RoleCollection)

  colQuerier := bson.M{"roleID": roleID}

	err := uc.Remove(colQuerier)
  return err
}
