package entity

type User struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Friends []string `json:"friends"`
	Id      int      `json:"id"`
}
