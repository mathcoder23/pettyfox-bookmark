package bookmark

type Bookmark struct {
	Id   string `json:"id"`
	Url  string `json:"url"`
	Desc string `json:"desc"`
}
type ObjResponse struct {
	code int
	msg  string
	data Bookmark
}
type IdsParams struct {
	Ids []string `json:"ids"`
}
