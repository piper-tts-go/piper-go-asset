package asset

import (
	"io/fs"
)

// Asset describes an embedded asset used in github.com/amitybell/piper
type Asset struct {
	// Name is the name of the asset
	Name string
	// FS is the embedded FS containing the distributed files
	FS fs.FS
}
