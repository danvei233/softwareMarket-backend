package handler

import (
	getservice "github.com/danvei233/softwareMarket-backend/app/service/Getservice"
	"github.com/gin-gonic/gin"
	"strconv"
)

type GetHandler struct {
	s *getservice.GetService
	r *gin.RouterGroup
}

func NewGetHandeler(s *getservice.GetService, r *gin.RouterGroup) *GetHandler {
	handler := GetHandler{s: s, r: r}
	r.GET("getmaincategorylist", handler.GetMainCategoryList)
	r.GET("GetSoftwareFromSubcategory", handler.GetSoftwareFromSubcategory)
	r.GET("getsoftwaredetail", handler.GetSoftware)
	r.GET("getsubcategorylist", handler.GetSubcategoryList)
	r.GET("getsoftwareshortcut", handler.GetSoftwareShortCut)
	return &handler

}
func (h *GetHandler) GetSoftwareFromSubcategory(g *gin.Context) {
	var query struct {
		Id    uint64 `form:"id" binding:"required"`
		Page  int    `form:"subpage,default=1"`
		Limit int    `form:"sublimit,default=20"`
	}

	err := g.BindQuery(&query)

	if err != nil {
		g.JSON(500, gin.H{"msg": err.Error()})
		return
	}

	list, err := h.s.GetSoftwareFromSubcategory(g, query.Id, query.Page, query.Limit)
	if err != nil {
		g.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	g.JSON(200, gin.H{"msg": "ok",
		"page":  query.Page,
		"limit": query.Limit,
		"data":  list})
}

func (h *GetHandler) GetMainCategoryList(g *gin.Context) {
	list, err := h.s.GetMainCategory(g)
	if err != nil {
		g.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	g.JSON(200, gin.H{"msg": "ok", "data": list})
}

func (h *GetHandler) GetSoftwareShortCut(g *gin.Context) {
	var query struct {
		Id        uint64 `form:"id" binding:"required"`
		SubPage   int    `form:"subpage,default=1"`
		SubLimit  int    `form:"sublimit,default=20"`
		SoftPage  int    `form:"softpage,default=1"`
		SoftLimit int    `form:"softlimit,default=400"`
	}
	err := g.BindQuery(&query)

	if err != nil {
		g.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	list, err := h.s.GetAllSoftWareShortcut(g, query.Id, query.SubPage, query.SubLimit, query.SoftPage, query.SoftLimit)

	if err != nil {
		g.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	g.JSON(200, gin.H{"msg": "ok",
		"query": query,
		"data":  list})

}
func (h *GetHandler) GetSoftware(g *gin.Context) {
	id, err := strconv.ParseUint(g.Query("id"), 10, 64)
	if err != nil {
		g.JSON(500, gin.H{"msg": "No valid id in query", "Detail": err.Error()})
		return
	}
	software, err := h.s.GetSoftwareDetail(g, id)
	if err != nil {
		g.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	g.JSON(200, gin.H{"msg": "ok", "software": software})

}

func (h *GetHandler) GetSubcategoryList(g *gin.Context) {
	id, err := strconv.ParseUint(g.Query("id"), 10, 64)

	if err != nil {
		g.JSON(500, gin.H{"msg": "No valid id in query", "Detail": err.Error()})
		return
	}
	subcate, err := h.s.GetSubList(g, id)
	if err != nil {
		g.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	g.JSON(200, gin.H{"msg": "ok", "subcategory": subcate})
}
