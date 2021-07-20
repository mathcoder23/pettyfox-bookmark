package bookmark

type Bookmark struct {
	Id   string `json:"id"`
	Url  string `json:"url"`
	Desc string `json:"desc"`
	Name string `json:"name"`
}
type ObjResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ObjResponseOk(data interface{}) ObjResponse {
	return ObjResponse{
		Code: 200,
		Data: data,
	}
}

type IdsParams struct {
	Ids []string `json:"ids"`
}
