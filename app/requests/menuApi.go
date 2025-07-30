package requests

// CreateMenuApiRequest 创建API请求
type CreateMenuApiRequest struct {
	Code     string `json:"code" binding:"required" validate:"required,max=100" label:"API代码"`
	Url      string `json:"url" binding:"required" validate:"required,max=100" label:"API地址"`
	MenuId   uint   `json:"menu_id" binding:"required" validate:"required" label:"菜单ID"`
	Describe string `json:"describe" validate:"max=255" label:"API描述"`
}

// UpdateMenuApiRequest 更新API请求
type UpdateMenuApiRequest struct {
	ID       uint   `json:"id" binding:"required" validate:"required" label:"API ID"`
	Code     string `json:"code" binding:"required" validate:"required,max=100" label:"API代码"`
	Url      string `json:"url" binding:"required" validate:"required,max=100" label:"API地址"`
	MenuId   uint   `json:"menu_id" binding:"required" validate:"required" label:"菜单ID"`
	Describe string `json:"describe" validate:"max=255" label:"API描述"`
}

// BatchDeleteMenuApiRequest 批量删除API请求
type BatchDeleteMenuApiRequest struct {
	IDs []uint `json:"ids" binding:"required" validate:"required,min=1" label:"API ID列表"`
}
