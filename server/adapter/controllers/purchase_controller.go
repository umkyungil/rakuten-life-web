package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rakuten-life-web/server/entity/service"
	"rakuten-life-web/server/entity/util"
)

// 条件検索: POST /purchase/authorization
func (pc Controllers) GetByAuthController(c *gin.Context) {
	var s service.Service
	p, err := s.GetByAuthService(c)

	if err != nil {
		c.AbortWithStatus(p.ProcessResult)
	} else {
		c.JSON(p.ProcessResult, p)
	}
}

// 条件検索: POST /purchase/search
func (pc Controllers) GetByConditionsController(c *gin.Context) {
	var s service.Service
	p, err := s.GetByConditionsService(c)

	if err != nil {
		c.AbortWithStatus(p.ProcessResult)
	} else {
		c.JSON(p.ProcessResult, p)
	}
}

// 条件検索: GET /purchase/search/:id
func (pc Controllers) GetByIdController(c *gin.Context) {
	var s service.Service
	p, err := s.GetByIdService(c)

	if err != nil {
		c.AbortWithStatus(p.ProcessResult)
	} else {
		c.JSON(p.ProcessResult, p)
	}
}

// 新規登録: POST /purchase/insert
func (pc Controllers) CreateModelController(c *gin.Context) {
	var s service.Service
	p, err := s.CreateModelService(c)

	if err != nil {
		c.AbortWithStatus(p.ProcessResult)
	} else {
		c.JSON(p.ProcessResult, p)
	}
}

// 更新: PUT /purchase/update/:id
func (pc Controllers) UpdateByIdController(c *gin.Context) {
	var s service.Service
	p, err := s.UpdateByIdService(c)

	if err != nil {
		c.AbortWithStatus(p.ProcessResult)
	} else {
		c.JSON(p.ProcessResult, p)
	}
}

// 削除: DELETE /purchase/delete/:id
func (pc Controllers) DeleteByIdController(c *gin.Context) {
	var s service.Service
	p, err := s.DeleteByIdService(c)

	if err != nil {
		c.AbortWithStatus(p.ProcessResult)
	} else {
		c.JSON(p.ProcessResult, p)
	}
}

// CSV: GET /purchase/csv
func (pc Controllers) CsvController(c *gin.Context) {
	var s service.Service
	p, err := s.CsvService(c)
	if err != nil {
		c.AbortWithStatus(p.ProcessResult)
	}
}

// 取引先情報検索（件名システム）: GET /purchse/vendors
func (pc Controllers) GetByAllOfVendorsController(c *gin.Context) {
	var s service.Service
	v, err := s.GetByAllOfVendorService(c)

	if err != nil {
		if err.Error() == util.BAD_REQUEST {
			c.JSON(http.StatusBadRequest, v)
		} else if err.Error() == util.RECORD_NOT_FOUND {
			c.JSON(http.StatusNotFound, v)
		} else {
			c.JSON(http.StatusInternalServerError, v)
		}
	} else {
		c.JSON(http.StatusOK, v)
	}
}

// 取引先情報検索（件名システム）: GET /purchase/vendor
func (pc Controllers) GetByIdOfVendorController(c *gin.Context) {
	var s service.Service
	v, err := s.GetByIdOfVendorService(c)

	if err != nil {
		if err.Error() == util.BAD_REQUEST {
			c.JSON(http.StatusBadRequest, v)
		} else if err.Error() == util.RECORD_NOT_FOUND {
			c.JSON(http.StatusNotFound, v)
		} else {
			c.JSON(http.StatusInternalServerError, v)
		}
	} else {
		c.JSON(http.StatusOK, v)
	}
}
