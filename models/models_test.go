package models_test

import (
	"testing"

	"github.com/gobuffalo/packr"
	"github.com/gobuffalo/plush"
	"github.com/gobuffalo/suite"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	plush.Helpers.Add("hash", func(s string) (string, error) {
		ph, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
		return string(ph), err
	})
}

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
