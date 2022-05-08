package custom_error

const (
	CodeDefault = "-1"
	// バリデーションエラー
	CodeValidation = "000001"
	// 無効な認証
	CodeInvalidAuthorization = "000002"
	// Not Found
	CodeNotFound = "000003"
	// Already Exists
	CodeAlreadyExists = "000004"
	// 自身の招待コード
	CodeMyInviteCode = "000005"
	// ユーザー削除時に共有メンバーが存在する
	HasRelationshipByDeleteUser = "000006"
)
