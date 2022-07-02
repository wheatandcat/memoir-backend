package logging

// INFOレベルのログ出力
func InfoLogEntry(message string) string {
	entry := &LogEntry{
		Severity: INFO,
		Message:  message,
	}

	return entry.String()
}

// WARNレベルのログ出力
func WarnLogEntry(message string) string {
	entry := &LogEntry{
		Severity: WARN,
		Message:  message,
	}

	return entry.String()
}

// ERRORレベルのログ出力
func ErrorLogEntry(message string) string {
	entry := &LogEntry{
		Severity: ERROR,
		Message:  message,
	}

	return entry.String()
}
