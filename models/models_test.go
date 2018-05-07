package models_test

import (
	"testing"

	"github.com/gobuffalo/packr"
	"github.com/gobuffalo/suite"
)

type ModelSuite struct {
	*suite.Model
}

func Test_ModelSuite(t *testing.T) {
	s, err := suite.NewModelWithFixtures(packr.NewBox("../fixtures"))
	if err != nil {
		t.Fatal(err)
	}
	as := &ModelSuite{s}
	suite.Run(t, as)
}
