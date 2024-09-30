package response

type Posts []*Post

type Post struct {
	ID      *int64  `json:",omitempty"`
	Title   *string `json:",omitempty"`
	Content *string `json:",omitempty"`
	View    *int64  `json:",omitempty"`

	User *User `json:",omitempty"`
}
