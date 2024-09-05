package api

import (
	"demo/types"
	"net/http"
)

// @Summary Generic Response
// @Produce  json
// @Success 200 {object} types.GenericResponse[api.GetGeneric.User]
// @Success 201 {object} types.GenericResponse[api.GetGeneric.Post]
// @Router / [get]
func GetGeneric(w http.ResponseWriter, r *http.Request) {
	type User struct {
		Username int    `json:"username"`
		Email    string `json:"email"`
	}
	type Post struct {
		Slug  int    `json:"slug"`
		Title string `json:"title"`
	}

	_ = types.GenericResponse[any]{}
}

// @Summary Multiple Generic Response
// @Produce  json
// @Success 200 {object} types.GenericMultiResponse[api.GetGenericMulti.Foo, api.GetGenericMulti.Bar]
// @Success 201 {object} types.GenericMultiResponse[api.GetGenericMulti.Bar, api.GetGenericMulti.Foo]
// @Router /multi [get]
func GetGenericMulti(w http.ResponseWriter, r *http.Request) {
	type Foo struct {
		SomeFieldA string `json:"some_field_a"`
	}
	type Bar struct {
		SomeFieldB string `json:"some_field_b"`
	}

	_ = types.GenericMultiResponse[any, any]{}
}
