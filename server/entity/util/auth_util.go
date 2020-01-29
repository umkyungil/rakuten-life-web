package util

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
)

// GroupModel はGroup情報を格納するjwt用の構造体
type GroupModel struct {
	GroupID    string `json:"group_id"`
	GroupEmail string `json:"group_email"`
}

// ResultClaims はJWTパース後のJSONをbindする構造体
type ResultClaims struct {
	Email  string       `json:"email"`
	Groups []GroupModel `json:"groups"`
	Exp    int          `json:"exp"`
	jwt.StandardClaims
}

/*!
 * @brief JWTから該当E-Mailの存在チェックを行う
 * @param[in] コンテキスト
 * @return error エラー情報
 */
func ChkAuth(c *gin.Context) error {
	// ヘッダ取得
	rawToken := c.Request.Header.Get(AUTH)

	// token存在チェック
	if rawToken == "" {
		return fmt.Errorf(PERMISSION_ERROR)
	}
	// BeareとJWTを分月して配列格納
	token, err := makeIDToken((strings.Split(rawToken, " ")))
	if err != nil {
		return fmt.Errorf(PERMISSION_ERROR)
	}
	// JWTパーシング
	parsedJWT, err := parseJWT(token)
	if err != nil {
		return fmt.Errorf(PERMISSION_ERROR)
	}

	models := parsedJWT.Groups
	bol := false
	for _, v := range models {
		if v.GroupEmail == BPMGR_EDIT_EMAIL {
			bol = true
			break
		} else if v.GroupEmail == BPMGR_MNGR_EMAIL {
			bol = true
			break
		}
	}

	// 戻り値設定
	if bol == true {
		return nil
	} else {
		return fmt.Errorf(PERMISSION_ERROR)
	}
}

/*!
 * @brief JWTをJSON型にパース
 * @param[in] JWT
 * @return error エラー情報
 */
func parseJWT(requestJWT string) (*ResultClaims, error) {
	var rc ResultClaims

	token, err := jwt.ParseWithClaims(requestJWT, &rc, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_KEY_RELEASE), nil
	})

	if err != nil {
		return nil, fmt.Errorf(PERMISSION_ERROR)
	}

	result, ok := token.Claims.(*ResultClaims)
	if !ok && !token.Valid {
		return result, err
	}
	return result, nil
}

/*!
 * @brief JWT取得
 * @param[in] Bearer+JWT（文字列)
 * @return string JWT
 * @return error エラー情報
 */
func makeIDToken(splitted []string) (string, error) {
	if len(splitted) == 2 {
		return splitted[1], nil
	} else if len(splitted) == 1 {
		return splitted[0], nil
	} else {
		return "", fmt.Errorf(PERMISSION_ERROR)
	}
}
