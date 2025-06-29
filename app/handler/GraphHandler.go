package handler

import (
	"encoding/json"
	"github.com/danvei233/softwareMarket-backend/app/utils"
	"github.com/dosco/graphjin/core/v3"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GraphqlHandler struct {
	db    *gorm.DB
	graph *core.GraphJin
	g     *gin.RouterGroup
	cfg   *utils.AppConfig
}

func NewGraphqlHandler(db *gorm.DB, group *gin.RouterGroup, cfg *utils.AppConfig) (*GraphqlHandler, error) {
	config := core.Config{
		Production:   !cfg.Debug.DisableProductionMode,
		DefaultLimit: 300}
	sqlDB, err := db.DB()
	g, err := core.NewGraphJin(&config, sqlDB)
	if err != nil {
		return nil, err
	}
	handler := GraphqlHandler{db: db, graph: g, g: group, cfg: cfg}
	group.POST("/", handler.Duty)
	return &handler, nil

}
func (h *GraphqlHandler) Duty(r *gin.Context) {
	var rawQuery struct {
		Query string          `json:"query"`
		Vars  json.RawMessage `json:"variables"`
	}
	if err := r.BindJSON(&rawQuery); err != nil {
		r.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	if res, err := h.graph.GraphQL(r, rawQuery.Query, rawQuery.Vars, nil); err != nil {
		r.JSON(500, gin.H{"msg": err.Error()})
		return

	} else {
		r.JSON(200, gin.H{"msg": "ok", "data": res})
	}
}
