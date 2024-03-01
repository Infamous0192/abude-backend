package file

import (
	"abude-backend/pkg/exception"
	"abude-backend/pkg/pagination"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"path/filepath"
	"sync"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FileServiceConfig struct {
	Url        string // Server URL
	UrlPath    string // URL Path
	UploadPath string // Uploaded Path
}

type FileService struct {
	db     *gorm.DB
	config FileServiceConfig
}

func NewService(db *gorm.DB, config FileServiceConfig) *FileService {
	return &FileService{
		db:     db,
		config: config,
	}
}

func (s *FileService) FindOne(id uint) (*File, error) {
	var file File
	if err := s.db.Find(&file, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &file, nil
}

func (s *FileService) FindAll(query pagination.Pagination) *pagination.Result[File] {
	result := pagination.New[File](query)

	db := s.db.Model(&File{})

	db.Order("created_at DESC")

	return result.Paginate(db)
}

func (s *FileService) Save(fh *multipart.FileHeader) (*File, error) {
	extension := filepath.Ext(fh.Filename)
	filename := uuid.New().String() + extension

	err := SaveFile(fh, filepath.Join(s.config.UploadPath, filename))
	if err != nil {
		return nil, exception.BadRequest("Failed to upload files")
	}

	path, _ := url.JoinPath(s.config.Url, s.config.UrlPath, filename)

	file := File{
		Filename:     filename,
		OriginalName: fh.Filename,
		Size:         int(fh.Size),
		Path:         path,
		Extension:    extension,
	}

	if err := s.db.Create(&file).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &file, nil
}

func (s *FileService) Update(id uint, fh *multipart.FileHeader) (*File, error) {
	file, err := s.FindOne(id)
	if err != nil {
		return nil, err
	}

	extension := filepath.Ext(fh.Filename)
	filename := uuid.New().String() + extension

	err = SaveFile(fh, filepath.Join(s.config.UploadPath, filename))
	if err != nil {
		return nil, exception.BadRequest("Failed to upload files")
	}

	os.Remove(filepath.Join(s.config.UploadPath, file.Filename))

	file.Filename = filename
	file.Path, _ = url.JoinPath(s.config.Url, s.config.UrlPath, filename)

	if err := s.db.Save(&file).Error; err != nil {
		return nil, exception.DB(err)
	}

	return file, nil
}

func (s *FileService) Delete(id uint) (*File, error) {
	file, err := s.FindOne(id)
	if err != nil {
		return nil, err
	}

	if err := s.db.Delete(&file).Error; err != nil {
		return file, exception.DB(err)
	}

	os.Remove(filepath.Join(s.config.UploadPath, file.Filename))

	return file, nil
}

// SaveFile saves multipart file fh under the given filename path.
func SaveFile(fh *multipart.FileHeader, path string) (err error) {
	var (
		f  multipart.File
		ff *os.File
	)
	f, err = fh.Open()
	if err != nil {
		return
	}

	var ok bool
	if ff, ok = f.(*os.File); ok {
		// Windows can't rename files that are opened.
		if err = f.Close(); err != nil {
			return
		}

		// If renaming fails we try the normal copying method.
		// Renaming could fail if the files are on different devices.
		if os.Rename(ff.Name(), path) == nil {
			return nil
		}

		// Reopen f for the code below.
		if f, err = fh.Open(); err != nil {
			return
		}
	}

	defer func() {
		e := f.Close()
		if err == nil {
			err = e
		}
	}()

	if ff, err = os.Create(path); err != nil {
		return
	}
	defer func() {
		e := ff.Close()
		if err == nil {
			err = e
		}
	}()
	_, err = copyZeroAlloc(ff, f)
	return
}

func (s *FileService) Using(tx *gorm.DB) *FileService {
	db := s.db

	defer func() {
		s.db = db
	}()

	s.db = tx

	return s
}

func copyZeroAlloc(w io.Writer, r io.Reader) (int64, error) {
	vbuf := copyBufPool.Get()
	buf := vbuf.([]byte)
	n, err := io.CopyBuffer(w, r, buf)
	copyBufPool.Put(vbuf)
	return n, err
}

var copyBufPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 4096)
	},
}
