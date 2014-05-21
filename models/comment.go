package models

type CommentModel struct {
	Id        *int
	UserId    *int
	ContentId *int
	Body      *string
	Date      *string
}
