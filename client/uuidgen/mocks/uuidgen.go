package mock_uuidgen

// UUID has generating method.
type UUID struct {
}

// Get UUIDを取得する
func (*UUID) Get() string {
	return "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
}
