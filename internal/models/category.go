package models

type Category struct {
	ID       int        `db:"id" json:"id"`
	ParentID *int       `db:"parent_id" json:"parent_id,omitempty"`
	Name     string     `db:"name" json:"name"`
	Slug     string     `db:"slug" json:"slug"`
	Children []*Category `json:"children,omitempty"`
}
