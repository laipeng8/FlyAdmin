package repositorys

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"server/app/event"
	"server/app/models"
	"server/app/requests"
	"server/global"
)

type AdminRepository struct {
	AdminUserModel models.AdminUser
	Where          map[string]interface{}
}

// Add 添加一个用户
func (u *AdminRepository) Add(password string, name string, data requests.UserAdd) (*gorm.DB, models.AdminUser) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	u.AdminUserModel.Password = string(pwd)
	u.AdminUserModel.Name = name
	u.AdminUserModel.RealName = data.RealName
	u.AdminUserModel.Email = data.Email
	u.AdminUserModel.Phone = data.Phone
	u.AdminUserModel.Gender = data.Gender
	u.AdminUserModel.Avatar = "https://www.lpmyblog.cn/applet/static/init1.png"
	for _, v := range data.Roles {
		var role models.Role
		role.ID = v
		u.AdminUserModel.Roles = append(u.AdminUserModel.Roles, role)
	}
	return global.Db.Create(&u.AdminUserModel), u.AdminUserModel
}

// Update 更新用户
func (u *AdminRepository) Update(data requests.AdminUpdate) error {

	return global.Db.Transaction(func(sessionDb *gorm.DB) error {
		var model models.AdminUser
		model.ID = data.Id
		if len(data.PassWord) > 0 {
			pwd, err := bcrypt.GenerateFromPassword([]byte(data.PassWord), bcrypt.MinCost)
			if err != nil {
				fmt.Println(err)
			}
			u.AdminUserModel.Password = string(pwd)
		}
		u.AdminUserModel.Name = data.Name
		u.AdminUserModel.RealName = data.RealName
		u.AdminUserModel.Email = data.Email
		u.AdminUserModel.Phone = data.Phone
		u.AdminUserModel.Gender = data.Gender
		u.AdminUserModel.Avatar = data.Avatar

		db := sessionDb.Where("id = ?", data.Id).Updates(&u.AdminUserModel)
		if db.Error == nil {
			var replace []models.Role
			for _, v := range data.Roles {
				var role models.Role
				role.ID = v
				replace = append(replace, role)
			}
			return sessionDb.Model(&model).Omit("Roles.*").Association("Roles").Replace(replace)

		} else {
			return db.Error
		}
	})

}

// Login 登陆用户
func (u *AdminRepository) Login(password string, name string, c *gin.Context) (bool, models.AdminUser) {
	re := global.Db.Where("name = ?", name).Preload("Roles").First(&u.AdminUserModel)

	_ = global.GetEventDispatcher(c).Dispatch(event.NewLoginEvent("login", u.AdminUserModel))

	if re.Error == nil && bcrypt.CompareHashAndPassword([]byte(u.AdminUserModel.Password), []byte(password)) == nil {
		return true, u.AdminUserModel
	} else {
		return false, u.AdminUserModel
	}
}

func (u *AdminRepository) List(page int, pageSize int, sortField string) map[string]interface{} {
	var (
		total  int64
		data   []models.AdminUser
		offSet int
	)
	db := global.Db.Model(&u.AdminUserModel)

	if len(u.Where) > 0 {
		for key, value := range u.Where {
			if value == nil || value == "" {
				continue // 跳过空值
			}

			switch key {
			case "nickname":
				db = db.Where("name LIKE ?", "%"+value.(string)+"%") // 模糊查询
			case "name", "real_name":
				db = db.Where(key+" LIKE ?", "%"+value.(string)+"%") // 支持按 name 或 real_name 模糊查询
			default:
				db = db.Where(key+" = ?", value) // 其他字段使用等值查询
			}
		}
	}

	db.Count(&total)
	offSet = (page - 1) * pageSize
	db.Preload("Roles").
		Limit(pageSize).
		Order(sortField + " desc, id desc").
		Offset(offSet).
		Find(&data)

	return global.Pages(page, pageSize, int(total), data)
}
