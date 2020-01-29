package util

const (
	// 日付フォーマット
	FORMAT_DATE       = "20060102"
	FORMAT_DATE_SLASH = "2006/01/02"
	MILLI_FORMAT      = "20060102150405"

	// CSV
	CSV_FILENAME = "purchase"
	EXTENSION    = ".csv"

	// エラーメッセージ
	RECORD_NOT_FOUND = "record not found"
	BAD_REQUEST      = "bad request"
	PERMISSION_ERROR = "permission error"
	DUPLICATE_ERROR  = "client code duplicate check error"

	// データベース
	CONNECTION_NAME_RELEASE = "jrnc-kenmei-release:asia-northeast1:jrnc-kenmei-release-db"
	CONNECTION_NAME_STAGING = "jrnc-kenmei-staging:asia-northeast1:jrnc-kenmei-staging-db"

	// 認可
	AUTH                  = "Authorization"
	CORS_ALL              = "*"
	CORS_LOCAL            = "http://localhost:8080"
	CORS_BPMG_STAGING     = "https://tg-bpmg-staging-dot-jrnc-kenmei-staging.appspot.com"
	CORS_BPMG_API_STAGING = "https://tg-bpmg-api-staging-dot-jrnc-kenmei-staging.appspot.com"
	CORS_BPMG_RELEASE     = "https://bpmg.jrnc.co.jp"
	CORS_BPMG_API_RELEASE = "https://bpmg-api.jrnc.co.jp"
	CORS_KENMEI_RELEASE   = "https://kenmei.jrnc.co.jp"
	JWT_KEY_DEV           = "jrnc-auth-dev-jwt-secret"
	JWT_KEY_STAGING       = "jrnc-auth-staging-jwt-secret"
	JWT_KEY_RELEASE       = "anJuYy1hdXRoLXJlbGVhc2Utand0LXNlY3JldA=="
	//BPMGR_MNGR_EMAIL  = "z-bpmgr-mgnt@jrnc.co.jp"
	BPMGR_EDIT_EMAIL = "z-bpmgr-edit@jrnc.co.jp"
	BPMGR_MNGR_EMAIL = "z-bpmgr-mngr@jrnc.co.jp"
)
