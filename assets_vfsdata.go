// Code generated by vfsgen; DO NOT EDIT.

package superclog

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2020, 9, 10, 0, 50, 41, 181523300, time.UTC),
		},
		"/markdown-external-release-notes.tmpl": &vfsgen۰FileInfo{
			name:    "markdown-external-release-notes.tmpl",
			modTime: time.Date(2020, 9, 8, 2, 15, 36, 223429800, time.UTC),
			content: []byte("\x23\x20\x52\x65\x6c\x65\x61\x73\x65\x20\x4e\x6f\x74\x65\x73\x3a\x0d\x0a\x7b\x7b\x72\x61\x6e\x67\x65\x20\x2e\x43\x6f\x6d\x6d\x69\x74\x73\x7d\x7d\x0d\x0a\x20\x2a\x20\x5b\x20\x5d\x20\x7b\x7b\x20\x2e\x52\x65\x6c\x65\x61\x73\x65\x4e\x6f\x74\x65\x20\x7d\x7d\x0d\x0a\x7b\x7b\x2d\x20\x65\x6e\x64\x7d\x7d\x0d\x0a"),
		},
		"/markdown-internal-qa-release-notes.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "markdown-internal-qa-release-notes.tmpl",
			modTime:          time.Date(2020, 9, 10, 0, 50, 41, 176524700, time.UTC),
			uncompressedSize: 516,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x7c\x90\x3f\x6b\xc3\x30\x10\xc5\x77\x83\xbf\xc3\x81\x3c\x24\xa6\x31\xed\x1a\xe8\x14\x28\x1d\x9a\x0c\xed\x58\x3a\x88\xfa\xd5\x16\xc8\x27\xd7\x3a\x4a\x83\xd0\x77\x2f\xb2\x13\x17\xf7\x4f\x36\x81\xee\xfd\xde\xef\x4e\x29\x7a\x84\x85\xf6\xa0\x83\x13\xf8\x6d\x9e\x85\x40\x83\xe6\x06\x54\xc8\xb1\xc7\x15\x15\xaf\xae\xeb\x1d\x83\x65\xaf\x7b\xda\xde\x52\xb5\xd3\x82\xc6\x0d\x06\x9e\x62\xcc\x33\xa5\x94\xa2\x10\xa6\xf9\xea\x6e\x30\xe0\xda\x1e\x0f\xba\xc3\xf8\x1d\xc2\xe6\x0c\x9c\x49\x13\xb5\x33\xf2\x60\xbc\x24\xe6\xb2\x64\x8a\x51\x61\xc1\x8d\xb4\xe9\xdf\x82\x17\x91\x18\x53\xa3\x79\x23\xbc\xcf\x63\x37\x63\xae\xa4\x67\x7a\xa1\xb2\x4c\xf9\x19\x4a\x31\x96\x65\x4a\xac\x0c\xd7\xf8\x5c\xa0\xae\xd7\xd5\xce\xf1\x07\x58\x8c\x63\x6d\xab\x3d\xbc\xd7\x4d\x72\xa7\xd5\xbf\x89\xa7\xd6\x0d\x72\xaf\x7d\x4b\x31\xae\x47\x59\x58\x8f\x93\xc0\x5f\xe5\x3f\xef\xd0\x99\x79\xef\xef\x9d\xf2\x8c\x68\xf2\x3f\x01\x3a\x23\x17\xe5\xce\x33\xbf\x74\x36\x04\xae\xe7\xf3\x5f\x7e\x7f\x05\x00\x00\xff\xff\xd4\x2f\xea\x9f\x04\x02\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/markdown-external-release-notes.tmpl"].(os.FileInfo),
		fs["/markdown-internal-qa-release-notes.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰FileInfo:
		return &vfsgen۰File{
			vfsgen۰FileInfo: f,
			Reader:          bytes.NewReader(f.content),
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰FileInfo is a static definition of an uncompressed file (because it's not worth gzip compressing).
type vfsgen۰FileInfo struct {
	name    string
	modTime time.Time
	content []byte
}

func (f *vfsgen۰FileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰FileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰FileInfo) NotWorthGzipCompressing() {}

func (f *vfsgen۰FileInfo) Name() string       { return f.name }
func (f *vfsgen۰FileInfo) Size() int64        { return int64(len(f.content)) }
func (f *vfsgen۰FileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰FileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰FileInfo) IsDir() bool        { return false }
func (f *vfsgen۰FileInfo) Sys() interface{}   { return nil }

// vfsgen۰File is an opened file instance.
type vfsgen۰File struct {
	*vfsgen۰FileInfo
	*bytes.Reader
}

func (f *vfsgen۰File) Close() error {
	return nil
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
