// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Activity struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Details string `json:"details"`
	Type    string `json:"type"`
}

type ActivityInput struct {
	Name     string `json:"name"`
	Details  string `json:"details"`
	Category string `json:"category"`
}
