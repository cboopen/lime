package controllers

import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/dto"
	"lime/pkg/api/service"
)

var BooksService = service.BooksService{}

type BooksController struct {
	BaseController
}

func (Book *BooksController) List(c *gin.Context) {
	var Dto dto.BooksListDto
	if Book.BindAndValidate(c, &Dto) {
		data, total := BooksService.List(Dto)
		resp(c, map[string]interface{}{
			"list":      data,
			"total":     total,
			"page":      Dto.Page,
			"page_size": Dto.Limit,
		})
	}
}
func (Book *BooksController) Create(c *gin.Context) {
	var Dto dto.BooksCreateDto
	if Book.BindAndValidate(c, &Dto) {
		created, _ := BooksService.Create(Dto)
		if created.Id <= 0 {
			fail(c, ErrAddFail)
			return
		}
		resp(c, map[string]interface{}{
			"id": created.Id,
		})
	}
}

func (Book *BooksController) Edit(c *gin.Context) {
	var Dto dto.BooksEditDto
	if Book.BindAndValidate(c, &Dto) {
		affected := BooksService.Update(Dto)
		if affected > 0 {
			ok(c, "更新成功")
			return
		}
		fail(c, ErrEditFail)
		return
	}
}

func (Book *BooksController) Get(c *gin.Context) {
	var Dto dto.GeneralGetDto
	if Book.BindAndValidate(c, &Dto) {
		data := BooksService.InfoOfId(Dto)
		if data.Id < 1 {
			fail(c, ErrIdData)
			return
		}
		resp(c, map[string]interface{}{
			"result": data,
		})
	}
}

func (Book *BooksController) Delete(c *gin.Context) {
	var Dto dto.GeneralDelDto
	if Book.BindAndValidate(c, &Dto) {
		affected := BooksService.Delete(Dto)
		if affected <= 0 {
			fail(c, ErrDelFail)
			return
		}
		ok(c, "删除成功")
	}
}