package models

import (
  "labix.org/v2/mgo/bson"
)

type Product struct { 
  Id bson.ObjectId "_id"
	ProductID int "productID"
	ProductName string "productName"
	SupplierID int "supplierID"
	CategoryID int "categoryID"
	QuantityPerUnit string "quantityPerUnit"
	UnitPrice float32 "unitPrice"
	UnitsInStock int "unitsInStock"
	UnitsOnOrder int "unitsOnOrder"
	ReorderLevel int "reorderLevel"
	Discontinued string "discontinued"
  CateName string 
}