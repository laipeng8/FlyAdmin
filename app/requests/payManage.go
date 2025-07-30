package requests

import "gorm.io/datatypes"

type PayManageRequest struct {
	ID       uint            `json:"id"`
	Name     *string         `json:"name"`
	Value    *datatypes.JSON `json:"value"`
	PlatForm *string         `json:"platform"`
}
