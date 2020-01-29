package repository

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	_ "github.com/djimenez/iconv-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"net/http"
	//db "rakuten-life-web/server/driver/mysql"
	"rakuten-life-web/server/entity/models"
	"rakuten-life-web/server/entity/util"
	"strconv"
	"time"
)

/*!
 * @brief ID検索をする
 * @param[in] id 連番
 * @return models.Purchase 検索結果情報
 * @return error エラー情報
 */
/*func (r Repository) GetByIdRepository(id string) (models.Purchase, error) {
	db := db.GetDB()
	var p models.Purchase

	// 検索処理
	if err := db.Where("id = ?", id).First(&p).Error; err != nil {
		return p, err
	} else {
		return p, err
	}
}*/

/*!
 * @brief 条件検索を行う。（削除されたデータは対象外）
 * @param[in] clientCode 取引先コード
 * @param[in] clientName 取引先名
 * @param[in] clientSpecializedField 専門分野
 * @param[in] evaluateDateFrom 評価実施日
 * @param[in] evaluateDateTo 評価実施日
 * @param[in] continuedEvaluateResult 継続評価結果
 * @return []models.Purchase 検索結果情報
 * @return error エラー情報
 */
/*func (r Repository) GetByConditionsRepository(clientCode string, clientName string, clientSpecializedField string, evaluateDateFrom string,
	evaluateDateTo string, continuedEvaluateResult string) ([]models.Purchase, error) {
	db := db.GetDB()
	var res []models.Purchase

	// SQL作成
	tx := db.Where("deleted_date IS NULL")
	// 取引先コード
	if clientCode != "" {
		c, _ := strconv.Atoi(clientCode)
		tx = tx.Where("client_code = ?", c)
	}
	// 取引先名称
	if clientName != "" {
		tx = tx.Where("client_name LIKE ?", "%"+clientName+"%")
	}
	// 専門分野
	if clientSpecializedField != "" {
		tx = tx.Where("client_specialized_field LIKE ?", "%"+clientSpecializedField+"%")
	}
	// 評価実施日
	if evaluateDateFrom != "" && evaluateDateTo != "" {
		tx = tx.Where("evaluate_date BETWEEN ? AND ?", evaluateDateFrom, evaluateDateTo)
	}
	// 評価実施日
	if evaluateDateFrom != "" && evaluateDateTo == "" {
		// 現在日付取得
		today := time.Now().Format(util.FORMAT_DATE)
		tx = tx.Where("evaluate_date BETWEEN ? AND ?", evaluateDateFrom, today)
	}
	// 評価実施日（１年前〜）
	if evaluateDateFrom == "" && evaluateDateTo != "" {
		t, _ := time.Parse(util.FORMAT_DATE, evaluateDateTo)
		base_date := t.AddDate(-1, 0, 0).Format(util.FORMAT_DATE)
		tx = tx.Where("evaluate_date BETWEEN ? AND ?", base_date, evaluateDateTo)
	}
	// 注文停止・抹殺（１：注文停止、２：抹消）
	if continuedEvaluateResult != "" {
		c, _ := strconv.Atoi(continuedEvaluateResult)
		tx = tx.Where("continued_evaluate_result = ?", c)
	}
	// 昇順
	tx = tx.Order("client_japanese_syllabary")

	// 検索処理
	var count = 0
	tx.Find(&res).Count(&count)
	if count == 0 {
		return res, errors.New(util.RECORD_NOT_FOUND)
	}

	if err := tx.Find(&res).Error; err != nil {
		return res, err
	} else {
		return res, err
	}
}*/

/*!
 * @brief 新規登録を行う。
 * @param[in] models.Purchase requestがバインディングされた構造体
 * @return models.Purchase 新規登録結果情報
 * @return error エラー情報
 */
/*func (r Repository) CreateModelRepository(req *models.Purchase) (d *models.Purchase, e error) {
	db := db.GetDB()
	var p models.Purchase
	var count = 0

	tx := db.Begin()
	if tx.Error != nil {
		return req, tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic: %+v\n", r)
			tx.Rollback()
			d = req
			e = errors.New("Unexpected error")
		}
	}()

	// 取引先コード重複チェック
	if req.ClientCode != 0 {
		tx.Where("client_code = ?", req.ClientCode).Find(&p).Count(&count)
	}
	if count > 0 {
		return req, errors.New("client code duplicate check error")
	}

	// 作成日付設定
	time := time.Now()
	req.CreatedDate = &time

	// 新規登録処理
	res := tx.Create(&req)
	if res.Error != nil {
		tx.Rollback()
		return req, res.Error
	}

	tx.Commit()
	return req, nil
}*/

/*!
 * @brief 更新登録を行う。
 * @param[in] id 連番
 * @param[in] コンテキスト
 * @return models.Purchase 新規登録結果情報
 * @return error エラー情報
 */
/*func (r Repository) UpdateByIdRepository(id string, c *gin.Context) (d models.Purchase, e error) {
	db := db.GetDB()
	var p models.Purchase

	tx := db.Begin()
	if tx.Error != nil {
		return p, tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			d = p
			e = errors.New("Unexpected error")
		}
	}()

	// 検索処理
	if err := db.Where("id = ?", id).First(&p).Error; err != nil {
		return p, err
	}
	// 既存の取引先コード取得
	oldClientCode := p.ClientCode
	// フォーム値を設定
	if err := c.BindJSON(&p); err != nil {
		return p, err
	}
	// 新しい取引先コード取得
	newClientCode := p.ClientCode

	// 取引先コード比較
	if oldClientCode != newClientCode {
		var count = 0
		var c models.Purchase

		// 取引先コード重複チェック
		if newClientCode != 0 {
			db.Where("client_code = ?", newClientCode).Find(&c).Count(&count)
		}
		if count > 0 {
			return p, errors.New("client code duplicate check error")
		}
	}

	// 更新日付設定
	time := time.Now()
	p.UpdatedDate = &time

	// 更新処理
	result := db.Save(&p)
	if result.Error != nil {
		tx.Rollback()
		return p, result.Error
	}

	tx.Commit()
	return p, nil
}*/

/*!
 * @brief 削除を行う。
 * @param[in] id 連番
 * @param[in] コンテキスト
 * @return error エラー情報
 */
/*func (r Repository) DeleteByIdRepository(id string, c *gin.Context) (d models.Purchase, e error) {
	db := db.GetDB()
	var p models.Purchase

	tx := db.Begin()
	if tx.Error != nil {
		return p, tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			e = errors.New("Unexpected error")
		}
	}()

	// 検索処理
	if err := tx.Where("id = ?", id).First(&p).Error; err != nil {
		return p, err
	}

	// 削除日付設定
	time := time.Now()
	p.DeletedDate = &time

	// 削除処理
	result := tx.Save(&p)
	if result.Error != nil {
		tx.Rollback()
		return p, result.Error
	}

	tx.Commit()
	return p, nil
}*/

/*!
 * @brief CSV出力を行う。
 * @param[in] コンテキスト
 * @param[in] []models.Purchase 検索結果
 * @return error エラー情報
 */
func (r Repository) PurchaseCSV(c *gin.Context, data []models.Purchase) error {
	headers := []string{
		"発効日", "No", "五十音", "コード", "取引先名", "代表者名", "専門分野", "資本金", "取引先", "基本契約",
		"ISO9001", "経営資料評価の要否", "財務内容", "過去実績", "将来性", "計", "専門技術", "照査技術", "成果レベル", "取組姿勢",
		"計", "登録抹消・注文停止", "評価実施年月日", "基本契約年月日", "業務評価無", "削除日", "作成日", "更新日"}

	b := &bytes.Buffer{}
	w := csv.NewWriter(transform.NewWriter(b, japanese.ShiftJIS.NewEncoder()))

	// ヘッダ作成
	if err := w.Write(headers); err != nil {
		return err
	}
	// データ作成
	for _, purchase := range data {
		record := make([]string, len(headers))
		// 発効日
		today := time.Now().Format(util.FORMAT_DATE_SLASH)
		record[0] = today
		// ID
		record[1] = strconv.Itoa(int(purchase.Id))
		// 五十音
		record[2] = purchase.ClientJapaneseSyllabary
		// 取引先コード
		if purchase.ClientCode != 0 {
			record[3] = strconv.Itoa(int(purchase.ClientCode))
		} else {
			record[3] = ""
		}
		// 取引先名
		record[4] = purchase.ClientName
		// 代表者名
		record[5] = purchase.ClientRepresentativeName
		// 専門分野
		record[6] = purchase.ClientSpecializedField
		// 資本金
		record[7] = purchase.ClientCapital
		// 取引先
		if purchase.ContinuedEvaluatePoint > 0 {
			record[8] = "継続"
		} else {
			record[8] = "新規"
		}
		// 契約形態
		if purchase.ContractType != 0 {
			if purchase.ContractType == 1 {
				record[9] = "有"
			} else {
				record[9] = "無"
			}
		} else {
			record[9] = ""
		}
		// ISO9001
		if purchase.ClientISO != 0 {
			if purchase.ClientISO == 1 {
				record[10] = "無"
			} else if purchase.ClientISO == 2 {
				record[10] = "有"
			}
		} else {
			record[10] = ""
		}
		// 経営評価資料（１：無し、２：有り）
		if purchase.ManagementDocument != 0 {
			if purchase.ManagementDocument == 1 {
				record[11] = "無"
			} else if purchase.ManagementDocument == 2 {
				record[11] = "有"
			}
		} else {
			record[11] = ""
		}
		// 財務内容
		if purchase.Evaluate1no1 != 0 {
			record[12] = strconv.Itoa(int(purchase.Evaluate1no1))
		} else if purchase.Evaluate1no2 != 0 {
			record[12] = strconv.Itoa(int(purchase.Evaluate1no2))
		} else {
			record[12] = ""
		}
		// 過去実績
		if purchase.Evaluate2 != 0 {
			record[13] = strconv.Itoa(int(purchase.Evaluate2))
		} else {
			record[13] = ""
		}
		// 将来性
		if purchase.Evaluate3 != 0 {
			record[14] = strconv.Itoa(int(purchase.Evaluate3))
		} else {
			record[14] = ""
		}
		// 計
		if purchase.EvaluatePoint != 0 {
			record[15] = strconv.Itoa(int(purchase.EvaluatePoint))
		} else {
			record[15] = "0"
		}
		// 専門技術
		if purchase.ContinuedEvaluate2 != 0 {
			record[16] = strconv.Itoa(int(purchase.ContinuedEvaluate2))
		} else {
			record[16] = ""
		}
		// 照査体制
		if purchase.ContinuedEvaluate3 != 0 {
			record[17] = strconv.Itoa(int(purchase.ContinuedEvaluate3))
		} else {
			record[17] = ""
		}
		// 成果レベル
		if purchase.ContinuedEvaluate4 != 0 {
			record[18] = strconv.Itoa(int(purchase.ContinuedEvaluate4))
		} else {
			record[18] = ""
		}
		// 取組姿勢
		if purchase.ContinuedEvaluate1 != 0 {
			record[19] = strconv.Itoa(int(purchase.ContinuedEvaluate1))
		} else {
			record[19] = ""
		}
		// 計
		if purchase.ContinuedEvaluatePoint != 0 {
			record[20] = strconv.Itoa(int(purchase.ContinuedEvaluatePoint))
		} else {
			record[20] = "0"
		}
		// 注文停止・抹殺（１：注文停止、２：抹消）
		if purchase.ContinuedEvaluateResult != 0 {
			if purchase.ContinuedEvaluateResult == 1 {
				record[21] = "注文停止"
			} else if purchase.ContinuedEvaluateResult == 2 {
				record[21] = "抹消"
			}
		} else {
			record[21] = ""
		}
		// 評価実施日
		e, _ := time.Parse(util.FORMAT_DATE, purchase.EvaluateDate)
		record[22] = e.Format(util.FORMAT_DATE_SLASH)
		// 基本契約日
		if purchase.BasicContractDate == "" {
			record[23] = ""
		} else {
			b, _ := time.Parse(util.FORMAT_DATE, purchase.BasicContractDate)
			record[23] = b.Format(util.FORMAT_DATE_SLASH)
		}
		// 業務評価無し
		if purchase.BusinessEvaluationDocument != 0 {
			if purchase.BusinessEvaluationDocument == 1 {
				record[24] = "無"
			} else if purchase.BusinessEvaluationDocument == 2 {
				record[24] = "有"
			}
		} else {
			record[24] = ""
		}
		// 削除日
		if purchase.DeletedDate == nil {
			record[25] = ""
		} else {
			record[25] = purchase.DeletedDate.Format(util.FORMAT_DATE_SLASH)
		}
		// 作成日
		if purchase.CreatedDate == nil {
			record[26] = ""
		} else {
			record[26] = purchase.CreatedDate.Format(util.FORMAT_DATE_SLASH)
		}
		// 更新日
		if purchase.UpdatedDate == nil {
			record[27] = ""
		} else {
			record[27] = purchase.UpdatedDate.Format(util.FORMAT_DATE_SLASH)
		}

		// データ書き込む
		if err := w.Write(record); err != nil {
			return err
		}
	}
	w.Flush()

	//ファイル名作成（ミリセコンド）
	filename := util.CSV_FILENAME + time.Now().Format(util.MILLI_FORMAT) + util.EXTENSION

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Data(http.StatusOK, "text/csv", b.Bytes())

	return nil
}

/*!
 * @brief 件名システムAPI：全件検索を行う
 * @param[in] flg：取引先の有効（1）／無効（0）フラグ
 * @return []models.Vendor 検索結果情報
 * @return error エラー情報
 */
/*func (r Repository) GetAllOfVendorRepository(flg int) ([]models.Vendor, error) {
	db := db.GetDB()
	var v []models.Vendor
	var p []models.Purchase

	tx := db.Select("client_code, client_name, continued_evaluate_result, contract_type, basic_contract_date, deleted_date")
	if flg == 1 {
		tx = tx.Where("client_code != ?", 0)
		tx = tx.Where("evaluate_result = ?", 2)           // 新規評価結果判定：取引可
		tx = tx.Where("continued_evaluate_result = ?", 0) // 継続評価判定：0（初期値）
		tx = tx.Where("deleted_date IS NULL")
	}
	tx = tx.Order("client_japanese_syllabary, client_name_k")

	var count int = 0
	tx.Find(&p).Count(&count)

	if count == 0 {
		return v, fmt.Errorf(util.RECORD_NOT_FOUND)
	}

	// 検索処理
	if err := tx.Find(&p).Error; err != nil {
		return v, err
	} else {
		v = make([]models.Vendor, len(p), cap(p))

		for i := 0; i < len(p); i++ {
			// 取引先コード
			if p[i].ClientCode != 0 {
				v[i].VendorId = strconv.Itoa(int(p[i].ClientCode))
			} else {
				v[i].VendorId = ""
			}
			// 取引先名
			v[i].VendorName = p[i].ClientName
			// 注文停止・抹殺（1:注文停止、2:抹消）
			if p[i].ContinuedEvaluateResult != 0 {
				if p[i].ContinuedEvaluateResult == 1 {
					v[i].ContinuedEvaluateResult = "注文停止"
				} else if p[i].ContinuedEvaluateResult == 2 {
					v[i].ContinuedEvaluateResult = "抹消"
				}
			} else {
				v[i].ContinuedEvaluateResult = ""
			}
			// 契約形態（1:基本契約、2:随時契約）
			if p[i].ContractType != 0 {
				if p[i].ContractType == 1 {
					v[i].ContractType = "基本契約"
				} else if p[i].ContractType == 2 {
					v[i].ContractType = "随時契約"
				}
			} else {
				v[i].ContractType = ""
			}
			// 基本契約日
			v[i].BasicContractDate = p[i].BasicContractDate
			// 削除日
			if p[i].DeletedDate != nil {
				v[i].DeletedDate = p[i].DeletedDate.Format(util.FORMAT_DATE_SLASH)
			} else {
				v[i].DeletedDate = ""
			}
		}
		return v, nil
	}
}*/

/*!
 * @brief 件名システムAPI：取引先コード検索
 * @param[in] id 連番
 * @return models.Vendor 検索結果情報
 * @return error エラー情報
 */
/*func (r Repository) GetByIdOfVendorRepository(id int) (models.Vendor, error) {
	db := db.GetDB()
	var v models.Vendor
	var p models.Purchase

	tx := db.Select("client_code, client_name, continued_evaluate_result, contract_type, basic_contract_date, deleted_date")
	tx = tx.Where("client_code = ?", id)

	// 検索処理
	if err := tx.Find(&p).Error; err != nil {
		return v, err
	} else {
		// 取引先コード
		if p.ClientCode != 0 {
			v.VendorId = strconv.Itoa(int(p.ClientCode))
		} else {
			v.VendorId = ""
		}
		// 取引先名
		v.VendorName = p.ClientName
		// 注文停止・抹殺（1:注文停止、2:抹消）
		if p.ContinuedEvaluateResult != 0 {
			if p.ContinuedEvaluateResult == 1 {
				v.ContinuedEvaluateResult = "注文停止"
			} else if p.ContinuedEvaluateResult == 2 {
				v.ContinuedEvaluateResult = "抹消"
			}
		} else {
			v.ContinuedEvaluateResult = ""
		}
		// 契約形態（1:基本契約、2:随時契約）
		if p.ContractType != 0 {
			if p.ContractType == 1 {
				v.ContractType = "基本契約"
			} else if p.ContractType == 2 {
				v.ContractType = "随時契約"
			}
		} else {
			v.ContractType = ""
		}
		// 基本契約日
		v.BasicContractDate = p.BasicContractDate
		// 削除日
		if p.DeletedDate != nil {
			v.DeletedDate = p.DeletedDate.Format(util.FORMAT_DATE_SLASH)
		} else {
			v.DeletedDate = ""
		}
		return v, nil
	}
}
*/
