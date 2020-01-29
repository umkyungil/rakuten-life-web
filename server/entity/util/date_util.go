package util

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

/*!
 * @brief 日付整合性チェック
 * @param[in] dateStr 日付
 * @return error エラー情報
 */
func CheckDate(dateStr string) error {
	// パラメータが空白の場合はスキップ
	if dateStr == "" {
		return nil
	}

	// 削除する文字列を定義
	reg := regexp.MustCompile(`[-|/|:| |　]`)

	// 指定文字を削除
	str := reg.ReplaceAllString(dateStr, "")

	// 数値の値に対してフォーマットを定義
	format := string([]rune(FORMAT_DATE)[:len(str)])

	// パース処理 → 日付ではない場合はエラー
	_, error := time.Parse(format, str)
	return error
}

/*!
 * @brief 日付を比較する
 * @param[in] dateStrFrom 日付From
 * @param[in] dateStrTo 日付To
 * @return error エラー情報
 */
func CheckDateCompare(dateStrFrom string, dateStrTo string) error {
	f, _ := strconv.Atoi(dateStrFrom)
	t, _ := strconv.Atoi(dateStrTo)

	if f > t {
		return fmt.Errorf(BAD_REQUEST)
	}

	return nil
}
