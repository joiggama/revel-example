package models

import "github.com/revel/revel"

type Product struct {
  Id   int32
  Code string
  Name string
}

func (p Product) Validate(v revel.Validation) {
  v.Required(Id)
  v.Required(Code)
  v.MaxSize(Name, 10)
}
