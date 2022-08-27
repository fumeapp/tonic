
# Tonic Web Framework

<p align="center">
  <img src="https://raw.githubusercontent.com/fumeapp/tonic/main/tonic.jpg" width="300" />
</p>

Tonic is a web application framework that supplies tools and libraries on top of Gin to give an expressive, elegant syntax.

[![Go Reference](https://pkg.go.dev/badge/github.com/fumeapp/tonic.svg)](https://pkg.go.dev/github.com/fumeapp/tonic)
[![Go Report Card](https://goreportcard.com/badge/github.com/fumeapp/tonic)](https://goreportcard.com/report/github.com/fumeapp/tonic)
[![format](https://github.com/fumeapp/tonic/actions/workflows/format.yml/badge.svg)](https://github.com/fumeapp/tonic/actions/workflows/format.yml)
[![lint](https://github.com/fumeapp/tonic/actions/workflows/lint.yml/badge.svg)](https://github.com/fumeapp/tonic/actions/workflows/lint.yml)
[![GitHub issues](https://img.shields.io/github/issues/fumeapp/tonic)](https://github.com/fumeapp/tonic/issues)
[![GitHub license](https://img.shields.io/github/license/fumeapp/tonic)](https://github.com/fumeapp/tonic/blob/main/license)

## Getting Started

### Environment variables

Take a look at [.env.example](https://github.com/fumeapp/tonic/blob/main/.env.example) for an example of how to set up environment variables. Copy it to `.env` and modify them to your needs
* Any ENV variable you do not set will default to what is in `.env.example`


### Endpoint Configuration
Define an endpoint that you can call locally and remotely.
* In this example we bind `/` to a standard JSON response
* You can specify params in the bind with `:` - ex: `/search/:query` can be access via `c.Param("query")`
* With `tonic.Init()` your database and other connections are ready to use

```go
import (
   fume "github.com/fumeapp/gin"
   "github.com/fumeapp/tonic"
   "github.com/fumeapp/tonic/render"
   "github.com/gin-gonic/gin"
)

func main() {
   routes := tonic.Init()
   routes.GET("/", func(c *gin.Context) { render.Render(c, {"message": "Hello World"}) })
   fume.Start(routes, fume.Options{})
}
```

### Database Connectivity
Connect to both a MySQL and OpenSearch database - other engines coming soon.
* access your databases via the `database` package
* Make sure to set `DB_CONNECT=true` if you want to use the MySQL database
* `database.Db` is your mysql connection
* Make sure to set `OS_CONNECT=true` if you want to use the Opensearch database
* `database.Os` is your opensearch connection

```go
import (
  . "github.com/fumeapp/tonic/database"
)

func main() {
  tonic.Init()
  Db.Create(&models.User{Name: "John Doe"})
  fmt.Println(Os.Info())
}
```

## AWS Access and Configuration
You can find the following `.env` variables in [.env.example](https://github.com/fumeapp/tonic/blob/main/.env.example)

```
AWS_CONNECT=true
AWS_ACCESS_KEY_ID=
AWS_SECRET_ACCESS_KEY=
AWS_REGION=us-east-1
AWS_BUCKET=
```
* You can access your AWS Config via `aws.Config`, your AWS S3 client via `aws.S3`, and your [Upload Manager](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/feature/s3/manager) via `aws.Uploader`
* An upload helper can be found as `aws.Upload(url)` - this will upload a file to your bucket, mark it public, and return the URL

## Router Features
### Route Model Binding

Easily bind a gorm Model to a controller and have Index, Show, Update, and Delete methods

```go
route.Init(engine)
route.ApiResource(engine, "user", &models.User{}, controllers.UserResources())
```
* 4 Routes will be created:
* `GET /user` binds to `Index`
* `GET /user/:id` binds to `Show`
* `PUT /user/:id` binds to `Update`
* `DELETE /user/:id` binds to `Delete`

* Show, Update, and Delete will have the model passed in as a parameter ready to be re-casted, otherwise return a 404

```go
func index(c *gin.Context) {
  var users = []models.User{}
  database.Db.Find(&users)
  render.Render(c, users)
}

func show(c *gin.Context, value any) {
  user := value.(*models.User)
  render.Render(c, user)
}

func update(c *gin.Context, value any) {
  user := value.(*models.User)
  render.Render(c, user)
}

func UserResources() route.ApiResourceStruct {
  return route.ApiResourceStruct{Index: index, Show: show, Update: update}
}
```