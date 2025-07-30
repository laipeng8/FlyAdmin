package requests

type FileGroupAdd struct {
	Name     string `json:"name" binding:"required" msg:"文件组名不能为空"`
	ParentID uint   `json:"parent_id"`
}

type FileGroupUpdate struct {
	ID uint `json:"id"`
	FileGroupAdd
}

type FileGroupDelete struct {
	IDS []uint `json:"ids"`
}

type FileAdd struct {
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
	FileUrl  string `json:"file_url"`
	Type     uint   `json:"type"` //1图片，2视频，3html
	Uploader uint   `json:"uploader"`
	GroupID  uint   `json:"group_id"`
}

type FileUpdate struct {
	ID uint `json:"id"`
	FileAdd
}
