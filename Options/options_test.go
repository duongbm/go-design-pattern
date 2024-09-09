package OptionsPattern

import "testing"

func TestPost(t *testing.T) {
	r := Requester{
		BaseUrl: "https://dummyjson.com",
	}

	var response struct {
		Id       int    `json:"id"`
		Username string `json:"username"`
		Token    string `json:"token"`
	}
	r.Post(
		"/auth/login",
		WithHeader("Content-type", "application/json"),
		WithTimeout(3),
		WithData("username", "emilys"),
		WithData("password", "emilyspass"),
		WithResponse(&response),
	)

	if response.Username != "emilys" {
		t.Errorf("wrong username")
	}
	if response.Id != 1 {
		t.Errorf("wrong id")
	}
}
