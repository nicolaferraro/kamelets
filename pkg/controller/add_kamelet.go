package controller

import (
	"github.com/nicolaferraro/kamelets/pkg/controller/kamelet"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, kamelet.Add)
}
