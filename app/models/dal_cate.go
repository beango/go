package models

import (
  "errors"
  "labix.org/v2/mgo/bson"
)

/*
 * 单个实体查找
 */
func (d *CateDal) FindByID(id int) Cate { 
  result := []Cate{}
  uc := d.session.DB(DbName).C(CateCollection)
  err := uc.Find(bson.M{"categoryID": id}).All(&result)
  if err != nil {
        panic(err)
  }
  if len(result)>0 {
    return result[0]
  } else {
    return Cate{}
  }
}

/* 
 * 获取列表
 */
func (d *CateDal) List() []Cate { 
  result := []Cate{}
  uc := d.session.DB(DbName).C(CateCollection)
  uc.Find(nil).All(&result)
  
  return result
}

/* 
 * 增加
 */
func (d *CateDal) Add(cate *Cate) error {
  uc := d.session.DB(DbName).C(CateCollection)

  i, _ := uc.Find(bson.M{"categoryName": cate.CategoryName}).Count() 
  if i != 0 { 
    return errors.New("名称已经被使用") 
  }
  
  maxObj := Cate{}
  maxID := 0
  err := uc.Find(nil).Sort("-categoryID").One(&maxObj)
  if err == nil {
    maxID = maxObj.CategoryID
  }

  cate.Id = bson.NewObjectId()
  cate.CategoryID = maxID + 1 
	err = uc.Insert(&cate)
  
  return err
}

/* 
 * 修改
 */
func (d *CateDal) Edit(cate *Cate) error {
  uc := d.session.DB(DbName).C(CateCollection)

  i, _ := uc.Find(bson.M{"cateName": cate.CategoryName,"categoryID": bson.M{"$ne": cate.CategoryID}}).Count() 
  if i != 0 { 
    return errors.New("名称已经被使用") 
  }

  colQuerier := bson.M{"categoryID": cate.CategoryID}
	change := bson.M{"$set": bson.M{"categoryName": cate.CategoryName,"description": cate.Description}}

	err := uc.Update(colQuerier, change)
  return err
}

/* 
 * 删除
 */
func (d *CateDal) Delete(categoryID int) error {
  uc := d.session.DB(DbName).C(CateCollection)

  colQuerier := bson.M{"categoryID": categoryID}

	err := uc.Remove(colQuerier)
  return err
}
