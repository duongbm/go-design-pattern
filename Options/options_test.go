package OptionsPattern

import (
	"testing"
)

// Refer API here: https://dummyjson.com/docs/auth

var r = Requester{
	BaseUrl: "https://dummyjson.com",
}

func TestPost(t *testing.T) {

	var response struct {
		Id       int    `json:"id"`
		Username string `json:"username"`
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

func TestGet(t *testing.T) {
	var token struct {
		AccessToken string `json:"token"`
	}
	r.Post(
		"/auth/login",
		WithHeader("Content-type", "application/json"),
		WithData("username", "emilys"),
		WithData("password", "emilyspass"),
		WithResponse(&token),
	)

	var response struct {
		Id       int    `json:"id"`
		Username string `json:"username"`
		Gender   string `json:"gender"`
	}
	r.Get("/auth/me",
		WithAuthorization(token.AccessToken),
		WithResponse(&response),
	)

	if response.Username != "emilys" {
		t.Errorf("wrong username")
	}
	if response.Gender != "female" {
		t.Errorf("wrong gender")
	}
	if response.Id != 1 {
		t.Errorf("wrong id")
	}
}
