package upload

import (
	"go-cloud/internal/model"
	"strconv"
)

type UserUpload struct {
	FileStoreID uint64 `json:"file_store_id" binding:"required"`
	ParentID    uint64 `json:"parent_id,omitempty" `
	FileHash    string `json:"file_hash" binding:"required"`
	FileName    string `binding:"-"`
	FileSize    int64  `binding:"-"`
	FileType    int    `binding:"-"`
	FileExt     string `binding:"-"`
	ChunkInfo
}

type ChunkInfo struct {
	UploadID    string `json:"upload_id" binding:"required"`
	ChunkNumber int    `json:"chunk_number" binding:"required"`
	TotalChunk  int    `json:"total_chunk" binding:"required"`
	ChunkSize   int    `json:"chunk_size" binding:"required"`
}

func SaveToDB(u UserUpload) error {
	uf := model.UserFile{
		FileName:     u.FileName,
		FileHash:     u.FileHash,
		FileStoreID:  u.FileStoreID,
		FileFolderID: u.ParentID,
		FileSize:     u.FileSize,
		FileExt:      u.FileExt,
		FileType:     strconv.Itoa(u.FileType),
	}
	err := uf.Create()
	return err
}
