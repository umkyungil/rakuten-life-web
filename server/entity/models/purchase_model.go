package models

import (
	"time"
)

// 購買取引台帳マスター
type Purchase struct {
	Id                 int64  `json:"id" gorm:"primary_key; AUTO_INCREMENT"` // 購買取引先台帳マスターID
	CustomerNo         int16  `json:"customer_no"`                           // 顧客番号（現行システムカーラムで不使用）
	IntroducerType     int8   `json:"introducer_type"`                       // 紹介者区分（1:無し、2:顧客、3:その他）
	Introducer         string `json:"introducer" gorm:"size:100"`            // 紹介者
	IntroducerLocated  string `json:"introducer_located" gorm:"size:100"`    // 紹介者所属先
	RequestSection     string `json:"request_section" gorm:"size:100"`       // 評価依頼部門
	RequestDate        string `json:"request_date" gorm:"size:8"`            // 評価依頼日
	EvaluateDate       string `json:"evaluate_date" gorm:"size:8"`           // 評価実施日
	Evaluate1          int8   `json:"evaluate1"`                             // 評価１区分（1:一般、2:個人）
	Evaluate1no1       int8   `json:"evaluate1_1" gorm:"column:evaluate1_1"` // 評価１ポイント（一般）
	Evaluate1no2       int8   `json:"evaluate1_2" gorm:"column:evaluate1_2"` // 評価１ポイント（個人商店など）
	Evaluate2          int8   `json:"evaluate2"`                             // 評価２ポイント
	Evaluate3          int8   `json:"evaluate3"`                             // 評価３ポイント
	EvaluatePoint      int8   `json:"evaluate_point"`                        // 評価ポイント計
	EvaluateResult     int8   `json:"evaluate_result"`                       // 判定結果（1:取引不可、2:取引可）
	ContractType       int8   `json:"contract_type"`                         // 契約形態（1:基本契約、2:随時契約）
	ManagementDocument int8   `json:"management_document"`                   // 経営評価資料（1:無し、2:有り）
	Comment            string `json:"comment" gorm:"size:500"`               // 特記事故

	// 継続評価情報
	ContinuedEvaluate1         int8   `json:"continued_evaluate1"`                   // 取組姿勢
	ContinuedEvaluate2         int8   `json:"continued_evaluate2"`                   // 専門技術
	ContinuedEvaluate3         int8   `json:"continued_evaluate3"`                   // 照査体制
	ContinuedEvaluate4         int8   `json:"continued_evaluate4"`                   // 成果レベル
	ContinuedEvaluatePoint     int8   `json:"continued_evaluate_point"`              // 継続評価ポイント計
	ContinuedEvaluateResult    int8   `json:"continued_evaluate_result"`             // 注文停止・抹殺（1:注文停止、2:抹消）
	BasicContractDate          string `json:"basic_contract_date" gorm:"size:8"`     // 基本契約日
	BusinessEvaluationDocument int8   `json:"business_evaluation_document"`          // 業務評価無し（1:無し、2:あり）
	TradingField               string `json:"trading_field" gorm:"size:100"`         // 取引分野
	TradingFieldDetails        string `json:"trading_field_details" gorm:"size:300"` // 取引分野細目

	//取引先情報
	ClientCode                int16  `json:"client_code"`                                  // 取引先コード
	ClientSpecializedField    string `json:"client_specialized_field" gorm:"size:100"`     // 取引先：専門分野
	ClientOfficer             string `json:"client_officer" gorm:"size:100"`               // 取引先：担当者
	ClientJapaneseSyllabary   string `json:"client_japanese_syllabary" gorm:"size:3"`      // 五十音
	ClientName                string `json:"client_name" gorm:"size:100"`                  // 取引先名称
	ClientNameK               string `json:"client_name_k" gorm:"size:100"`                // 取引先名称カナ
	ClientRepresentativeName  string `json:"client_representative_name" gorm:"size:100"`   // 代表者名
	ClientRepresentativeNameK string `json:"client_representative_name_k" gorm:"size:100"` // 代表者名カナ
	ClientZipCode1            string `json:"client_zip_code1" gorm:"size:10"`              // 郵便番号1
	ClientZipCode2            string `json:"client_zip_code2" gorm:"size:10"`              // 郵便番号2
	ClientAddress1            string `json:"client_address1" gorm:"size:200"`              // 住所1
	ClientAddress2            string `json:"client_address2" gorm:"size:200"`              // 住所2
	ClientAddressK            string `json:"client_address_k" gorm:"size:200"`             // 住所カナ

	ClientTelNo     string `json:"client_tel_no" gorm:"size:20"` // 電話番号
	ClientFaxNo     string `json:"client_fax_no" gorm:"size:20"` // ファックス番号
	ClientCapital   string `json:"client_capital" gorm:"size:8"` // 資本金（コンマ無し）
	ClientISO       int8   `json:"client_iso"`                   // ISO有無（1:無し、2:あり）
	ClientOrderFlag int8   `json:"client_order_flag"`            // 注文FLG（1:注文不可、2:注文可）

	//その他情報
	DeletedDate *time.Time `json:"deleted_date"`
	CreatedDate *time.Time `json:"created_date"`
	UpdatedDate *time.Time `json:"updated_date"`
}
