package models

// 件名システム用
type Vendor struct {
	//取引先情報
	VendorId                string `json:"vendor_id"`
	VendorName              string `json:"vendor_name"`
	ContinuedEvaluateResult string `json:"continued_evaluate_result"`         // 注文停止・抹殺（1:注文停止、2:抹消）
	ContractType            string `json:"contract_type"`                     // 契約形態（1:基本契約、2:随時契約）
	BasicContractDate       string `json:"basic_contract_date" gorm:"size:8"` // 基本契約日
	DeletedDate             string `json:"deleted_date"`
}
