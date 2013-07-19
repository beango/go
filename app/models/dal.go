package models

import ( 
  "github.com/robfig/revel" 
  "labix.org/v2/mgo" 
)

const ( 
  DbName                         = "Northwind" 
  UserCollection                 = "UserAccount" 
  RoleCollection                 = "Role" 
  AuthCollection                 = "Auth" 
  ProductCollection              = "Product"
  CateCollection                 = "Categories"
)

type UserDal struct { 
  session *mgo.Session
}

type RoleDal struct { 
  session *mgo.Session
}

type AuthDal struct { 
  session *mgo.Session
}

type ProductDal struct { 
  session *mgo.Session
}

type CateDal struct { 
  session *mgo.Session
}

func NewUserDal() (*UserDal, error) { 
  revel.Config.SetSection("db") 
  ip, found := revel.Config.String("ip") 
  if !found { 
    revel.ERROR.Fatal("Cannot load database ip from app.conf") 
  }

  session, err := mgo.Dial(ip) 
  if err != nil { 
    return nil, err 
  }

  return &UserDal{session}, nil 
}

func (d *UserDal) Close() { 
  d.session.Close() 
} 

func NewRoleDal() (*RoleDal, error) { 
  revel.Config.SetSection("db") 
  ip, found := revel.Config.String("ip") 
  if !found { 
    revel.ERROR.Fatal("Cannot load database ip from app.conf") 
  }

  session, err := mgo.Dial(ip) 
  if err != nil { 
    return nil, err 
  }

  return &RoleDal{session}, nil 
}

func (d *RoleDal) Close() { 
  d.session.Close() 
} 

func NewAuthDal() (*AuthDal, error) { 
  revel.Config.SetSection("db") 
  ip, found := revel.Config.String("ip") 
  if !found { 
    revel.ERROR.Fatal("Cannot load database ip from app.conf") 
  }

  session, err := mgo.Dial(ip) 
  if err != nil { 
    return nil, err 
  }

  return &AuthDal{session}, nil 
}

func (d *AuthDal) Close() { 
  d.session.Close() 
} 

func NewProductDal() (*ProductDal, error) { 
  revel.Config.SetSection("db") 
  ip, found := revel.Config.String("ip") 
  if !found { 
    revel.ERROR.Fatal("Cannot load database ip from app.conf") 
  }

  session, err := mgo.Dial(ip) 
  if err != nil { 
    return nil, err 
  }

  return &ProductDal{session}, nil 
}

func (d *ProductDal) Close() { 
  d.session.Close() 
} 


func NewCateDal() (*CateDal, error) { 
  revel.Config.SetSection("db") 
  ip, found := revel.Config.String("ip") 
  if !found { 
    revel.ERROR.Fatal("Cannot load database ip from app.conf") 
  }

  session, err := mgo.Dial(ip) 
  if err != nil { 
    return nil, err 
  }

  return &CateDal{session}, nil 
}

func (d *CateDal) Close() { 
  d.session.Close() 
} 