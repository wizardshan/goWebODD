package request

import (
	domain2 "goWebODD/test/domain"
	post2 "goWebODD/test/domain/post"
)

type PostOne struct {
	ID      domain2.ID
	Title   post2.Title   `binding:"-"`
	Content post2.Content `binding:"-"`
	User    domain2.User  `binding:"-"`
}
