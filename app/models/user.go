package models

import (
  "labix.org/v2/mgo/bson"
)

type UserLogin struct { 
  UserName string "userName"
  UserPwd  string "userPwd"
}

type User struct { 
  Id bson.ObjectId "_id"
  UserID   int "userID"
  UserName string "userName"
  UserPwd  string "userPwd"
}