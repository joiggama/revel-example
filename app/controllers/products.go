package controllers

import "github.com/revel/revel"
//import "fmt"
import "github.com/revel/revel/modules/db/app"

type Products struct {
  *revel.Controller
  //db.Transactional
}

func (c Products) Index() revel.Result {
  db.Init()
  //products, _ := c.Txn.Query
  //fmt.Println("Something")
  //products, _ := c.Txn.Query("select * from products")
  //return c.RenderJson(products)
  return c.RenderJson("{}")
}
