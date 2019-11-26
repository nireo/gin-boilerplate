# Gin boilerplate
##### A fast way to develop gin rest apis with golang

This includes authorization middleware and JWt middleware. It already has user service, but it's up to you to add other things.

### What you need
Make sure you have all these golang packages.
```
go get github.com/jinzhu/gorm
go get github.com/dgrijalva/jwt-go
go get github.com/gin-gonic/gin
go get golang.org/x/crypto/bcrypt
```

After that you need a ```.env``` file which includes the connection string for the database and the port. Like so:
```
DB_CONFIG=user:password@/dbname?charset=utf8&parseTime=True&loc=Local
PORT=8080
```

You can change the database from the default MySQL to whatever, in the ```/database/database.go``` file. All the migrations are handled by gorm upon running ```main.go```. After all this you can type ```go run main.go``` to start the api. 

### Adding things
To add things like models can be done by adding a new file to ```/database/models```. For example let's add a post:
```go
package models

import (
    "github.com/jinzhu/gorm"
    "github.com/nireo/gin-boilerplate/lib/common" // import the json data type alias
)

// Post data model
type Post struct {
    gorm.Model
    Title   string `sql:"type:text;"`
    Content string `sql:"type:text;"`
}

// Serialize the post data
func (p Post) Serialize() common.JSON {
    return common.JSON{
        "id":         p.ID,
        "content":    p.Content,
        "title":      p.Title,
        "created_at": p.CreatedAt,
    }
}
```