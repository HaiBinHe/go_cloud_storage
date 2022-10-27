package dao

import (
	"go-cloud/internal/model"
	"go-cloud/pkg/app"
)

func GetFolderByID(folderID uint64) (*model.FileFolder, error) {
	ff := &model.FileFolder{}
	err := model.Db.Where("id = ?", folderID).First(ff).Error
	if err != nil {
		return nil, err
	}
	return ff, nil
}

func ListDirectory(uid, pid, sid uint64, page, pageSize int, order string) ([]*model.FileFolder, int64) {
	var (
		folders []*model.FileFolder
		total   int64
	)
	Dbchain := model.Db.Model(&model.FileFolder{}).Where("user_id = ? AND parent_id = ? AND store_id = ?", uid, pid, sid).Find(&folders)
	Dbchain.Count(&total)

	//分页
	offset := app.GetPageOffset(page, pageSize)
	Dbchain.Limit(pageSize).Offset(offset).Order(order).Find(&folders)
	return folders, total
}