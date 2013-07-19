package models

import (
  "labix.org/v2/mgo/bson"
)

type Cate struct { 
  Id bson.ObjectId "_id"
  CategoryID int "categoryID"
  CategoryName  string "categoryName"
  Description  string "description"
}