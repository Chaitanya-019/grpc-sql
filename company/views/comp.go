package views

import (
	"example.com/grpc-sql/company/companypb"
)

type Comp struct {
	ID      int64
	Name    string
	Creator string
}

func (c *Comp) ToCompany() *companypb.Company {

	res := &companypb.Company{
		Id:     c.ID,
		Name:   c.Name,
		Person: c.Creator,
	}
	return res
}
