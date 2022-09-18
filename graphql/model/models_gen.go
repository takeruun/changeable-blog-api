// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Blog struct {
	ID                 string   `json:"id"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Body               string   `json:"body"`
	NightBody          string   `json:"nightBody"`
	MobileBody         string   `json:"mobileBody"`
	ThumbnailImagePath string   `json:"thumbnailImagePath"`
	Tags               []string `json:"tags"`
	CreatedAt          string   `json:"createdAt"`
	UpdateAt           string   `json:"updateAt"`
}

type BlogList struct {
	ID                 string   `json:"id"`
	Title              string   `json:"title"`
	ThumbnailImagePath string   `json:"thumbnailImagePath"`
	Tags               []string `json:"tags"`
	CreatedAt          string   `json:"createdAt"`
}

type BlogListConnection struct {
	PageInfo *PageInfo   `json:"pageInfo"`
	Nodes    []*BlogList `json:"nodes"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PageCondition struct {
	PageNo int     `json:"pageNo"`
	Limit  int     `json:"limit"`
	Query  *string `json:"query"`
}

type PageInfo struct {
	PageNo int `json:"pageNo"`
	// 検索結果の全件数
	TotalCount int `json:"totalCount"`
}

type RecommendBlogListConnection struct {
	Nodes []*BlogList `json:"nodes"`
}

type SignUp struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	PostalCode string `json:"postalCode"`
}

type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	PostalCode string `json:"postalCode"`
	CreatedAt  string `json:"createdAt"`
}
