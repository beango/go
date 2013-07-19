package models

import (
  "labix.org/v2/mgo/bson"
)

type Auth struct { 
  Id bson.ObjectId "_id"
  AuthID int "authID"
  AuthName  string "authName"
}