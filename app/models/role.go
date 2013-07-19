package models

import (
  "labix.org/v2/mgo/bson"
)

type Role struct { 
  Id bson.ObjectId "_id"
  RoleID int "roleID"
  RoleName  string "roleName"
}