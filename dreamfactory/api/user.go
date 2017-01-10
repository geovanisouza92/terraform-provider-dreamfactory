
func (c *Client) UserCreate(u types.UsersRequest) (ur types.UsersResponse, err error) {
	err = c.send("POST", "/api/v2/system/user", 200, u, &ur)
	return
}

func (c *Client) UserRead(id string) (u types.User, err error) {
	err = c.send("GET", "/api/v2/system/user/"+id, 200, nil, &u)
	return
}

func (c *Client) UserUpdate(id string, u types.User) error {
	return c.send("PATCH", "/api/v2/system/user/"+id, 200, u, nil)
}

func (c *Client) UserDelete(id string) error {
	return c.send("DELETE", "/api/v2/system/user/"+id, 200, nil, nil)
}

func (c *Client) UserExists(id string) error {
	return c.send("GET", "/api/v2/system/user/"+id, 200, nil, nil)
}