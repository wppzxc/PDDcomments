package types

type PageData struct {
	ItemLink string `json:"item_link"`
	CommentNumber string `json:"comment_number"`
	CommentHead string `json:"comment_head"`
	CommentFoot string `json:"comment_foot"`
	CommentFilter string `json:"comment_filter"`
	CommentDiscount string `json:"comment_discount"`
}

type Comment struct {
	Comment        string `json:"comment"`
	Name           string `json:"name"`
	DescScore      int    `json:"desc_score"`
	LogisticsScore int    `json:"logistics_score"`
	ServiceScore   int    `json:"service_score"`
}

type CommentData struct {
	Data []Comment `json:"data"`
}