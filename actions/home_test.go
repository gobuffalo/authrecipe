package actions

func (as *ActionSuite) Test_HomeHandler() {
	res := as.HTML("/").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Sign In")
}

func (as *ActionSuite) Test_HomeHandler_LoggedIn() {
	as.Login()

	res := as.HTML("/").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Sign Out")

	as.Session.Clear()
	res = as.HTML("/").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Sign In")
}
