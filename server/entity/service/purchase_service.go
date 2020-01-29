package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rakuten-life-web/server/entity/models"
	"rakuten-life-web/server/entity/repository"
	"rakuten-life-web/server/entity/util"
	"strconv"
)

// 複数行レスポンス
type PurchaseResults struct {
	Purchase      []models.Purchase `json:"purchase"`
	ProcessResult int               `json:"process_result"`
}

// 単1行レスポンス
type PurchaseResult struct {
	Purchase      models.Purchase `json:"purchase"`
	ProcessResult int             `json:"process_result"`
}

// 検索条件
type PurchaseRequest struct {
	ClientCode              string `json:"client_code"`
	ClientName              string `json:"client_name"`
	ClientSpecializedField  string `json:"client_specialized_field"`
	EvaluateDateFrom        string `json:"evaluate_date_from"`
	EvaluateDateTo          string `json:"evaluate_date_to"`
	ContinuedEvaluateResult string `json:"continued_evaluate_result"`
}

// 件名システム：複数行レスポンス
type VendorResults struct {
	Vendor        []models.Vendor `json:"vendor"`
	ProcessResult int             `json:"process_result"`
}

// 件名システム：単1行レスポンス
type VendorResult struct {
	Vendor        models.Vendor `json:"vendor"`
	ProcessResult int           `json:"process_result"`
}

/*!
 * @brief 認可サービス
 * @param[in] コンテキスト
 * @return PurchaseResult　認可判定結果の構造体で返す
 * @return error エラー情報
 */
func (s Service) GetByAuthService(c *gin.Context) (PurchaseResult, error) {
	var e models.Purchase

	// JWTチェック
	authResult := util.ChkAuth(c)
	if authResult != nil {
		var r = PurchaseResult{e, http.StatusUnauthorized}
		return r, fmt.Errorf(util.PERMISSION_ERROR)
	} else {
		var r = PurchaseResult{e, http.StatusOK}
		return r, nil
	}
}

/*!
 * @brief 条件検索サービス
 * @param[in] コンテキスト
 * @return PurchaseResults　検索結果を構造体で返す（配列）
 * @return error エラー情報
 */
func (s Service) GetByConditionsService(c *gin.Context) (PurchaseResults, error) {
	var rep repository.Repository
	var req PurchaseRequest

	// JWTチェック
	authResult := util.ChkAuth(c)
	if authResult != nil {
		var r = PurchaseResults{nil, http.StatusUnauthorized}
		return r, fmt.Errorf(util.PERMISSION_ERROR)
	}

	c.BindJSON(&req)

	// 日付整合性チェック
	if err := util.CheckDate(req.EvaluateDateFrom); err != nil {
		var r = PurchaseResults{nil, http.StatusBadRequest}
		return r, err
	}
	if err := util.CheckDate(req.EvaluateDateTo); err != nil {
		var r = PurchaseResults{nil, http.StatusBadRequest}
		return r, err
	}
	if req.EvaluateDateFrom != "" && req.EvaluateDateTo != "" {
		if err := util.CheckDateCompare(req.EvaluateDateFrom, req.EvaluateDateTo); err != nil {
			var r = PurchaseResults{nil, http.StatusBadRequest}
			return r, err
		}
	}

	// 検索処理
	/*res, err := rep.GetByConditionsRepository(req.ClientCode, req.ClientName, req.ClientSpecializedField, req.EvaluateDateFrom, req.EvaluateDateTo, req.ContinuedEvaluateResult)

	// 戻り値処理
	if err != nil {
		if err.Error() == util.RECORD_NOT_FOUND {
			var r = PurchaseResults{nil, http.StatusNotFound}
			return r, err
		} else {
			var r = PurchaseResults{nil, http.StatusInternalServerError}
			return r, err
		}
	} else {
		var r = PurchaseResults{res, http.StatusOK}
		return r, nil
	}*/
	return r, nil
}

/*!
 * @brief ID検索サービス
 * @param[in] コンテキスト
 * @return PurchaseResult　検索結果を構造体で返す（単一検索）
 * @return error エラー情報
 */
func (s Service) GetByIdService(c *gin.Context) (PurchaseResult, error) {
	var rep repository.Repository
	var e models.Purchase

	// JWTチェック
	authResult := util.ChkAuth(c)
	if authResult != nil {
		var r = PurchaseResult{e, http.StatusUnauthorized}
		return r, fmt.Errorf(util.PERMISSION_ERROR)
	}

	// パラメータ取得
	res, err := rep.GetByIdRepository(c.Param("id"))

	// 戻り値設定
	if err != nil {
		if err.Error() == util.RECORD_NOT_FOUND {
			var r = PurchaseResult{e, http.StatusNotFound}
			return r, err
		} else {
			var r = PurchaseResult{e, http.StatusInternalServerError}
			return r, err
		}
	} else {
		var r = PurchaseResult{res, http.StatusOK}
		return r, nil
	}
}

/*!
 * @brief 新規登録サービス
 * @param[in] コンテキスト
 * @return PurchaseResult　登録したデータを構造帯で返す
 * @return error エラー情報
 */
func (s Service) CreateModelService(c *gin.Context) (PurchaseResult, error) {
	var rep repository.Repository
	var p models.Purchase
	var e models.Purchase

	// JWTチェック
	authResult := util.ChkAuth(c)
	if authResult != nil {
		var r = PurchaseResult{e, http.StatusUnauthorized}
		return r, fmt.Errorf(util.PERMISSION_ERROR)
	}

	// リクエストバインディングチェック
	if err := c.BindJSON(&p); err != nil {
		var r = PurchaseResult{e, http.StatusBadRequest}
		return r, err
	}
	// 日付整合性チェック
	if err := util.CheckDate(p.RequestDate); err != nil {
		var r = PurchaseResult{e, http.StatusBadRequest}
		return r, err
	}
	if err := util.CheckDate(p.EvaluateDate); err != nil {
		var r = PurchaseResult{e, http.StatusBadRequest}
		return r, err
	}
	if p.RequestDate != "" && p.EvaluateDate != "" {
		if err := util.CheckDateCompare(p.RequestDate, p.EvaluateDate); err != nil {
			var r = PurchaseResult{e, http.StatusBadRequest}
			return r, err
		}
	}

	// 新規登録処理
	res, err := rep.CreateModelRepository(&p)

	// 戻り値処理
	if err != nil {
		// 重複チェックエラー
		if err.Error() == util.DUPLICATE_ERROR {
			var r = PurchaseResult{e, http.StatusConflict}
			return r, err
		} else {
			var r = PurchaseResult{e, http.StatusInternalServerError}
			return r, err
		}
	} else {
		var r = PurchaseResult{*res, http.StatusCreated}
		return r, err
	}
}

/*!
 * @brief 更新登録サービス
 * @param[in] コンテキスト
 * @return PurchaseResult　更新したデータを構造体で返す
 * @return error エラー情報
 */
func (s Service) UpdateByIdService(c *gin.Context) (PurchaseResult, error) {
	var rep repository.Repository
	var p models.Purchase
	var e models.Purchase

	// JWTチェック
	authResult := util.ChkAuth(c)
	if authResult != nil {
		var r = PurchaseResult{e, http.StatusUnauthorized}
		return r, fmt.Errorf(util.PERMISSION_ERROR)
	}

	// 日付整合性チェック
	if err := util.CheckDate(p.RequestDate); err != nil {
		var r = PurchaseResult{e, http.StatusBadRequest}
		return r, err
	}
	if err := util.CheckDate(p.EvaluateDate); err != nil {
		var r = PurchaseResult{e, http.StatusBadRequest}
		return r, err
	}
	if p.RequestDate != "" && p.EvaluateDate != "" {
		if err := util.CheckDateCompare(p.RequestDate, p.EvaluateDate); err != nil {
			var r = PurchaseResult{e, http.StatusBadRequest}
			return r, err
		}
	}

	// 更新登録処理
	p, err := rep.UpdateByIdRepository(c.Param("id"), c)

	// 戻り値処理
	if err != nil {
		// 重複チェックエラー
		if err.Error() == util.DUPLICATE_ERROR {
			var r = PurchaseResult{e, http.StatusConflict}
			return r, err
		} else {
			var r = PurchaseResult{e, http.StatusInternalServerError}
			return r, err
		}
	} else {
		var r = PurchaseResult{p, http.StatusOK}
		return r, err
	}
}

/*!
 * @brief 削除サービス
 * @param[in] コンテキスト
 * @return PurchaseResult　削除したデータを構造体で返す
 * @return error エラー情報
 */
func (s Service) DeleteByIdService(c *gin.Context) (PurchaseResult, error) {
	var rep repository.Repository
	var e models.Purchase

	// JWTチェック
	authResult := util.ChkAuth(c)
	if authResult != nil {
		var r = PurchaseResult{e, http.StatusUnauthorized}
		return r, fmt.Errorf(util.PERMISSION_ERROR)
	}

	// 削除処理
	res, err := rep.DeleteByIdRepository(c.Param("id"), c)

	// 戻り値処理
	if err != nil {
		var r = PurchaseResult{e, http.StatusInternalServerError}
		return r, err
	} else {
		var r = PurchaseResult{res, http.StatusOK}
		return r, err
	}
}

/*!
 * @brief CSVサービス
 * @param[in] コンテキスト
 * @return PurchaseResults　CSV出力結果を返す
 * @return error エラー情報
 */
func (s Service) CsvService(c *gin.Context) (PurchaseResults, error) {
	var rep repository.Repository
	var req PurchaseRequest

	// JWTチェック
	authResult := util.ChkAuth(c)
	if authResult != nil {
		var r = PurchaseResults{nil, http.StatusUnauthorized}
		return r, fmt.Errorf(util.PERMISSION_ERROR)
	}

	c.BindJSON(&req)

	// 日付整合性チェック
	if err := util.CheckDate(req.EvaluateDateFrom); err != nil {
		var r = PurchaseResults{nil, http.StatusBadRequest}
		return r, err
	}
	if err := util.CheckDate(req.EvaluateDateTo); err != nil {
		var r = PurchaseResults{nil, http.StatusBadRequest}
		return r, err
	}
	if req.EvaluateDateFrom != "" && req.EvaluateDateTo != "" {
		if err := util.CheckDateCompare(req.EvaluateDateFrom, req.EvaluateDateTo); err != nil {
			var r = PurchaseResults{nil, http.StatusBadRequest}
			return r, err
		}
	}

	// 検索処理
	res, err := rep.GetByConditionsRepository(req.ClientCode, req.ClientName, req.ClientSpecializedField, req.EvaluateDateFrom, req.EvaluateDateTo, req.ContinuedEvaluateResult)

	// 検索エラー処理
	if err != nil {
		if err.Error() == util.RECORD_NOT_FOUND {
			var r = PurchaseResults{nil, http.StatusNotFound}
			return r, err
		} else {
			var r = PurchaseResults{nil, http.StatusInternalServerError}
			return r, err
		}
	}

	// CSV出力処理
	if err := rep.PurchaseCSV(c, res); err != nil {
		var r = PurchaseResults{nil, http.StatusInternalServerError}
		return r, err
	} else {
		var r = PurchaseResults{res, http.StatusOK}
		return r, err
	}
}

/*!
 * @brief 件名システム：取引先コード検索サービス
 * @param[in] コンテキスト
 * @return VendorResult　検索結果を構造体で返す（単一）
 * @return error エラー情報
 */
func (s Service) GetByIdOfVendorService(c *gin.Context) (VendorResult, error) {
	var rep repository.Repository
	var e models.Vendor

	// パラメータ取得
	clientCode := c.DefaultQuery("vendor_id", "")

	//パラメータチェック
	id, err := strconv.Atoi(clientCode)
	if err != nil {
		var r = VendorResult{e, 100}
		return r, fmt.Errorf(util.BAD_REQUEST)
	}

	res, err := rep.GetByIdOfVendorRepository(id)

	// 戻り値設定
	if err != nil {
		if err.Error() == util.RECORD_NOT_FOUND {
			var r = VendorResult{e, 400}
			return r, err
		} else {
			var r = VendorResult{e, 900}
			return r, err
		}
	} else {
		var r = VendorResult{res, 0}
		return r, nil
	}
}

/*!
 * @brief 件名システム：全件検索サービス
 * @param[in] コンテキスト
 * @return VendorResults　検索結果を構造体で返す（複数）
 * @return error エラー情報
 */
func (s Service) GetByAllOfVendorService(c *gin.Context) (VendorResults, error) {
	var rep repository.Repository
	var e []models.Vendor

	// パラメータ取得
	switch_flg := c.DefaultQuery("switch_flg", "")

	//パラメータチェック
	flg, err := strconv.Atoi(switch_flg)
	if err != nil {
		e = make([]models.Vendor, 1)
		var r = VendorResults{e, 100}
		return r, fmt.Errorf(util.BAD_REQUEST)
	}
	if flg < 1 || flg > 2 {
		e = make([]models.Vendor, 1)
		var r = VendorResults{e, 100}
		return r, fmt.Errorf(util.BAD_REQUEST)
	}

	res, err := rep.GetAllOfVendorRepository(flg)

	// 戻り値設定
	if err != nil {
		if err.Error() == util.RECORD_NOT_FOUND {
			e = make([]models.Vendor, 1)
			var r = VendorResults{e, 400}
			return r, err
		} else {
			e = make([]models.Vendor, 1)
			var r = VendorResults{e, 900}
			return r, err
		}
	} else {
		var r = VendorResults{res, 0}
		return r, nil
	}
}
