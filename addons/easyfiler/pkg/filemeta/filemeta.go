package filemeta

import (
	"dtstack.com/dtstack/easymatrix/addons/easyfiler/pkg/model"
)

type FileMeta struct {
	FileSha1 string
	FileMD5  string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

func UpdateFileMeta(f FileMeta) bool {
	return model.UploadAndFinished(f.FileSha1, f.FileName, f.Location, f.FileSize)
}

func GetFileMeta(filesha1 string) (filemeta FileMeta, err error) {
	tablefile, err := model.GetFileMeta(filesha1)
	if err != nil {
		return FileMeta{}, err
	}
	filemeta = FileMeta{
		FileSha1: tablefile.FileHash,
		FileName: tablefile.FileName.String,
		FileSize: tablefile.FileSize.Int64,
		Location: tablefile.FileAddr.String,
	}
	return
}

func DelteFileMeta(fileSha1 string) {
	model.DeleteFileMeta(fileSha1)
}
