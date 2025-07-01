package handler

import (
	"errors"
	"github.com/danvei233/softwareMarket-backend/app/service/DownloadSerivce"
	"github.com/danvei233/softwareMarket-backend/app/utils"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type DownloadHandler struct {
	g *gin.RouterGroup
	s DownloadSerivce.DownloadService
}

func NewDownloadSerivce(g *gin.RouterGroup, s *DownloadSerivce.DownloadService) (*DownloadHandler, error) {
	ds := &DownloadHandler{g: g, s: *s}
	cfg, _ := utils.GetConfig()
	g.GET("download/*path", ds.DownloadFileService)
	g.Static("static/", cfg.Dir.AssertRelativePath)
	return ds, nil
}
func CheckValid(path string) (bool, error) {

	fullPath, err := filepath.Abs(path)
	if err != nil {
		return false, err
	}
	cfg, err := utils.GetConfig()

	if err != nil {
		return false, err
	}
	mainPath, err := filepath.Abs(cfg.Dir.DataRelativePath)

	if err != nil {
		return false, err
	}
	if !strings.HasPrefix(fullPath, mainPath) {
		return false, errors.New("<UNK>")
	}
	return true, nil
}
func (g *DownloadHandler) DownloadFileService(c *gin.Context) {
	path := "./" + c.Param("path")
	_, err := CheckValid(path) //safe
	if err != nil {
		c.JSON(403, gin.H{"msg": err.Error()})
		return
	}
	splitPath := strings.Split(path, "/")
	if len(splitPath) > 1 {
		splitPath = splitPath[:len(splitPath)-2]
	}
	c.File(path)
	pathOfStatus := strings.Join(splitPath, "/") + ".id"
	id, err := readId(pathOfStatus)
	if err == nil {
		//数据库操作
		g.s.AddDownloadNum(c, id)
	}

}
func readId(path string) (uint64, error) {
	fullPath, err := filepath.Abs(path)
	if err != nil {
		return 0, err
	}
	file, _ := os.Open(fullPath)
	defer file.Close()

	buffer := make([]byte, 1024)
	id := []byte{}
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				res, _ := strconv.ParseUint(string(id), 10, 64)
				return res, nil
			}
			return 0, err
		}
		id = append(id, buffer[:n]...)

	}
}
