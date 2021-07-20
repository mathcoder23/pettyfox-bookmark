package controller

import (
	"github.com/kataras/iris/v12"
	"pettyfox.top/bookmark/modules/bookmark"
	"pettyfox.top/bookmark/modules/bookmark/service"
)

func List(ctx iris.Context) {
	list := service.List()
	ctx.JSON(bookmark.ObjResponseOk(list))
	// TIP: negotiate the response between server's prioritizes
	// and client's requirements, instead of ctx.JSON:
	// ctx.Negotiation().JSON().MsgPack().Protobuf()
	//ctx.Negotiate(books)

}
func Search(ctx iris.Context) {
	list := service.Search(ctx.URLParam("keyword"))

	ctx.JSON(bookmark.ObjResponseOk(list))
	// TIP: negotiate the response between server's prioritizes
	// and client's requirements, instead of ctx.JSON:
	// ctx.Negotiation().JSON().MsgPack().Protobuf()
	//ctx.Negotiate(books)

}
func Save(ctx iris.Context) {
	var body bookmark.Bookmark
	if err := ctx.ReadBody(&body); err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	service.Save(body)
	ctx.JSON(body)
}
func Remove(ctx iris.Context) {
	var params bookmark.IdsParams
	if err := ctx.ReadBody(&params); err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	service.Remove(params)
	ctx.JSON(params)
	// TIP: negotiate the response between server's prioritizes
	// and client's requirements, instead of ctx.JSON:
	// ctx.Negotiation().JSON().MsgPack().Protobuf()
	// ctx.Negotiate(books)
}
