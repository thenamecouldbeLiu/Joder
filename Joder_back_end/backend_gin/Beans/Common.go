package beans

type RequestWrapper struct {
	Model any `json:"model"`
}

type ResponseWrapper struct {
	ResultCode    string `json:"resultCode"`
	ResultMessage string `json:"resultMessage"`
	ResultBody    any    `json:"resultBody"`
}

type UserInfo struct {
	Name      string   `json:"name"`
	UserId    string   `json:"userId"`
	PicUrl    []string `json:"picUrl"`
	Email     string   `json:"email"`
	SelfIntro string   `json:"selfIntro"`
	UserType  int      `json:"userType"`
	Tags      []Tag    `json:"tags"`
}
type InsertUserDto struct {
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	SelfIntro string  `json:"selfIntro"`
	Birthday  string  `json:"birthday"`
	Usertype  int     `json:"userType"`
	Gender    int     `json:"gender"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
type Tag struct {
	Content string
}
type PaginateDto struct {
	PaginateNum int `json:"paginateNum"`
}

type UnmatchedUserDto struct {
	PaginateDto
}
