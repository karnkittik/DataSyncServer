package entities

type ResponseData struct {
	UUID    string  `json:"u"`
	Author  *string `json:"a,omitempty"`
	Message *string `json:"m,omitempty"`
	Likes   *int    `json:"l,omitempty"`
}

type GetRequestBody struct {
	UnixTimestamp int `json:"datetime"`
}

type GetResponseData struct {
	Delete []string       `json:"d"`
	Update []ResponseData `json:"u"`
	Create []ResponseData `json:"c"`
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
