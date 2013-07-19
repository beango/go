package models

import (
  "errors"
  "labix.org/v2/mgo/bson"
)

/*
 * 单个实体查找
 */
func (d *ProductDal) FindByID(id int) Product { 
  result := []Product{}
  uc := d.session.DB(DbName).C(ProductCollection)
  err := uc.Find(bson.M{"productID": id}).All(&result)
  if err != nil {
        panic(err)
  }
  if len(result)>0 {
    return result[0]
  } else {
    return Product{}
  }
}

/* 
 * 获取列表
 */
func (d *ProductDal) List(pageindex int, pagesize int) (*[]Product,int,int) { 
  result := []Product{}
  uc := d.session.DB(DbName).C(ProductCollection)
  uc.Find(nil).Skip((pageindex-1)*pagesize).Limit(pagesize).Select(nil).All(&result)
  totalRecord, _ := uc.Find(nil).Count()
  totalPage := totalRecord / pagesize 
  if totalRecord % pagesize >0 {
    totalPage++
  }
  return &result, totalRecord, totalPage
}

/* 
 * 增加
 */
func (d *ProductDal) Add(prod *Product) error {
  uc := d.session.DB(DbName).C(ProductCollection)

  i, _ := uc.Find(bson.M{"productName": prod.ProductName}).Count() 
  if i != 0 { 
    return errors.New("产品名重复") 
  }
  
  maxObj := Product{}
  maxID := 0
  err := uc.Find(nil).Sort("-productID").One(&maxObj)
  if err == nil {
    maxID = maxObj.ProductID
  }

  prod.Id = bson.NewObjectId()
  prod.ProductID = maxID + 1 
	err = uc.Insert(&prod)
  
  return err
}

/* 
 * 修改
 */
func (d *ProductDal) Edit(prod *Product) error {
  uc := d.session.DB(DbName).C(ProductCollection)

  i, _ := uc.Find(bson.M{"productName": prod.ProductName,"productID": bson.M{"$ne": prod.ProductID}}).Count() 
  if i != 0 { 
    return errors.New("产品名重复") 
  }

  colQuerier := bson.M{"productID": prod.ProductID}
	change := bson.M{"$set": bson.M{
              "productID": prod.ProductID,
              "productName": prod.ProductName,
              "supplierID": prod.SupplierID,
              "categoryID": prod.CategoryID,
              "quantityPerUnit": prod.QuantityPerUnit,
              "unitPrice": prod.UnitPrice,
              "unitsInStock": prod.UnitsInStock,
              "unitsOnOrder": prod.UnitsOnOrder,
              "reorderLevel": prod.ReorderLevel,
              "discontinued": prod.Discontinued,
  }}

	err := uc.Update(colQuerier, change)
  return err
}

/* 
 * 删除
 */
func (d *ProductDal) Delete(productID int) error {
  uc := d.session.DB(DbName).C(ProductCollection)

  colQuerier := bson.M{"productID": productID}

	err := uc.Remove(colQuerier)
  return err
}
