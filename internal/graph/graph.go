package graph

import (
	myarcs "github.com/joshuanario/arcs"
)

func Tree(arcs []myarcs.Arc, arcrole string) *myarcs.RArc {
	return myarcs.NewRArc(arcs, arcrole)
}
