package entities

type ResponseData struct {
	UUID    string  `json:"uuid"`
	Author  *string `json:"author"`
	Message *string `json:"message"`
	Likes   *int    `json:"likes"`
}

type GetRequestBody struct {
	UnixTimestamp int `json:"datetime"`
}

type GetResponseData struct {
	Delete []string       `json:"delete"`
	Update []ResponseData `json:"update"`
	Create []ResponseData `json:"create"`
}

type PostRequestBody struct {
	UUID    string `json:"uuid"`
	Author  string `json:"author"`
	Message string `json:"message"`
	Likes   int    `json:"likes"`
}

type PutRequestBody struct {
	Author  string `json:"author"`
	Message string `json:"message"`
	Likes   int    `json:"likes"`
}
