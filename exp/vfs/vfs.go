package vfs

import (
	"github.com/gabriel-vasile/mimetype"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gopub/errors"
	"github.com/gopub/log"
	"github.com/gopub/types"
)

var logger *log.Logger

func init() {
	logger = log.Default().Derive("Wine.vfs")
	logger.SetFlags(log.LstdFlags - log.Lfunction - log.Lshortfile)
}

func SetLogger(l *log.Logger) {
	logger = l
}

const (
	keySize         = 32
	DefaultPageSize = int64(types.MB)
	MinPageSize     = int64(32 * types.KB)
)

const (
	ErrAuth errors.String = "invalid password"
)

const (
	keyFSRootDir    = "filesystem.root"
	keyFSCredential = "filesystem.credential"
	keyFSConfig     = "filesystem.config"
	keyFSPageSize   = "filesystem.page_size"
)

type Storage interface {
	// Get returns os.ErrNotExist if key doesn't exist
	Get(key string) ([]byte, error)
	Put(key string, val []byte) error
	Delete(key string) error
	Close() error
}

func cleanName(name string) string {
	name = filepath.Clean(name)
	name = strings.TrimPrefix(name, "/")
	name = strings.TrimSuffix(name, "/")
	return name
}

func splitPath(path string) []string {
	path = filepath.Clean(path)
	path = strings.TrimPrefix(path, "/")
	path = strings.TrimSuffix(path, "/")
	paths := strings.Split(path, "/")
	return paths
}

func validateFileName(name string) bool {
	if name == "" {
		return false
	}

	if strings.Contains(name, "/") {
		return false
	}

	return true
}

type Flag int

const (
	ReadOnly  = Flag(os.O_RDONLY)
	WriteOnly = Flag(os.O_WRONLY)
	Create    = Flag(os.O_CREATE)
)

const contentTypeOctetStream = "application/octet-stream"

func DetectContentType(b []byte) string {
	t := http.DetectContentType(b)
	if t == "" {
		return mimetype.Detect(b).String()
	}

	if t != contentTypeOctetStream {
		return t
	}

	if mt := mimetype.Detect(b).String(); mt != contentTypeOctetStream {
		return mt
	}

	return t
}
