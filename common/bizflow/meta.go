package bizflow

type Meta struct {
	RetryNum      int
	IsRetryErr    map[error]bool
	IsNonRetryErr map[error]bool
	AbortErr      map[error]bool
}
