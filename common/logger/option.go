package logger

var (
	defaultOptions = Options{
		logDir:     "/home/www/logs/applogs",
		filename:   "default.log",
		maxSize:    500,
		maxAge:     1,
		maxBackups: 10,
		callerSkip: 1,
	}
)

type Options struct {
	logDir     string
	filename   string
	maxSize    int
	maxBackups int
	maxAge     int
	compress   bool
	callerSkip int
}

type Option interface {
	apply(*Options)
}

type OptionFunc func(*Options)

func (o OptionFunc) apply(opts *Options) {
	o(opts)
}

// WithLogDir ...
func WithLogDir(dir string) Option {
	return OptionFunc(func(options *Options) {
		options.logDir = dir
	})
}

// WithHistoryLogFileName ...
func WithHistoryLogFileName(fileName string) Option {
	return OptionFunc(func(options *Options) {
		options.filename = fileName
	})
}

// WithMaxSize ...
func WithMaxSize(size int) Option {
	return OptionFunc(func(options *Options) {
		options.maxSize = size
	})
}

// WithMaxBackups ...
func WithMaxBackups(backup int) Option {
	return OptionFunc(func(options *Options) {
		options.maxBackups = backup
	})
}

// WithMaxAge ...
func WithMaxAge(maxAge int) Option {
	return OptionFunc(func(options *Options) {
		options.maxAge = maxAge
	})
}

// WithCompress ...
func WithCompress(b bool) Option {
	return OptionFunc(func(options *Options) {
		options.compress = b
	})
}

// WithCallerSkip ...
func WithCallerSkip(skip int) Option {
	return OptionFunc(func(options *Options) {
		options.callerSkip = skip
	})
}
