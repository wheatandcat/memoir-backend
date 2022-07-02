package logging

import (
	"encoding/json"
	"log"
)

// ログレベルのCONSTを定義
// https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#logseverity
var (
	INFO  = "INFO"
	WARN  = "WARNING"
	ERROR = "ERROR"
)

// GCPのLogEntryに則った構造化ログモデル
type LogEntry struct {
	// GCP上でLogLevelを表す
	Severity string `json:"severity"`
	// ログの内容
	Message string `json:"message"`
}

// 構造体をJSON形式の文字列へ変換
// 参考: https://cloud.google.com/run/docs/logging#run_manual_logging-go
func (l LogEntry) String() string {
	if l.Severity == "" {
		l.Severity = INFO
	}
	out, err := json.Marshal(l)
	if err != nil {
		log.Printf("json.Marshal: %v", err)
	}
	return string(out)
}
