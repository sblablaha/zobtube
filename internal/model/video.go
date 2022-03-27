package model

// Video main type
type Video struct {
	Base
	Title      string `json:"title"`
	Production string `json:"production"`
}
