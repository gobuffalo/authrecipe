package models_test

import (
	"github.com/gobuffalo/authrecipe/models"
	"github.com/gobuffalo/uuid"
)

func (ms *ModelSuite) Test_User_Create() {
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	u := &models.UserForm{
		User: models.User{
			Email: "mark@example.com",
		},
		Password:             "password",
		PasswordConfirmation: "password",
	}
	ms.Zero(u.PasswordHash)

	verrs, err := u.ValidateAndCreate(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	ms.NotZero(u.PasswordHash)

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)
}

func (ms *ModelSuite) Test_User_Create_ValidationErrors() {
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	u := &models.UserForm{
		User: models.User{
			Email: "mark@example.com",
		},
	}
	ms.Zero(u.PasswordHash)

	verrs, err := u.ValidateAndCreate(ms.DB)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)
}

func (ms *ModelSuite) Test_User_Create_UserExists() {
	ms.LoadFixture("user")
	ms.DBDelta(0, "users", func() {
		u := models.User{}
		ms.NoError(ms.DB.First(&u))
		u.ID = uuid.UUID{}

		form := &models.UserForm{
			User:                 u,
			Password:             "password",
			PasswordConfirmation: "password",
		}
		verrs, err := form.ValidateAndCreate(ms.DB)
		ms.NoError(err)
		ms.True(verrs.HasAny())
	})

}
