// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package db

import ()

type Person struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}
