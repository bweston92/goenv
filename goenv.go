package goenv

import (
	"io/ioutil"
	"os"

	"github.com/spf13/afero"
)

const fileSuffix = "_FILE"

var (
	// OS will be initialised with NewFromOS on init.
	OS *EnvLocator
)

type (
	// EnvResolver takes an expected environment variable
	// name and looks to see if it defined, if it is not
	// we will see if there is a definition with the _FILE
	// suffix and read the value instead.
	EnvResolver func(string) string

	// EnvLocator locates environment configuration.
	EnvLocator struct {
		resolver   EnvResolver
		fs         afero.Fs
		middleware stack
	}
)

func init() {
	OS = NewFromOS(afero.NewOsFs())
}

// New env locator
func New(r EnvResolver, fs afero.Fs) *EnvLocator {
	return &EnvLocator{
		resolver:   r,
		fs:         fs,
		middleware: defaultStack(),
	}
}

// NewFromOS will result in a new EnvLocator using
// the os.Getenv as a value resolver.
func NewFromOS(fs afero.Fs) *EnvLocator {
	return New(os.Getenv, fs)
}

// Get will return value from the resolver for the
// given key or attempt to load from file if "key_FILE"
// exists as a key in the resolver.
// The value then gets passed through a middleware
// stack for stuff like base64 decoding.
func (l *EnvLocator) Get(key string) string {
	v := l.resolver(key)
	if len(v) == 0 {
		if v = l.fileResolve(key); len(v) == 0 {
			return ""
		}
	}

	return l.middleware.Play(v)
}

func (l *EnvLocator) fileResolve(key string) string {
	f := l.resolver(key + "_FILE")
	if len(f) == 0 {
		return ""
	}

	fp, err := l.fs.Open(f)
	if err != nil {
		return ""
	}
	defer fp.Close()

	b, err := ioutil.ReadAll(fp)
	if err != nil {
		return ""
	}

	return string(b)
}

// Get will allow use of goenv.Get which will use goenv.OS.Get
func Get(v string) string {
	return OS.Get(v)
}
