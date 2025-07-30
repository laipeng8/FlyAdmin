package system

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"server/app/models"
	"server/app/repositorys"
	"server/app/requests"
	"server/global"
	"server/global/response"
	"strconv"
)

type UserController struct{}

// Login 登陆
func (u *UserController) Login(c *gin.Context) {
	var (
		LoginForm      requests.Login
		roles          []string
		permission     []string
		userRepository repositorys.AdminRepository
		menuRepository repositorys.SystemMenuRepository
	)
	err := c.ShouldBind(&LoginForm)
	if err != nil {
		response.Failed(c, global.GetError(err, LoginForm))
		return
	}
	isLogin, user := userRepository.Login(LoginForm.PassWord, LoginForm.Name, c)

	if isLogin {

		for _, role := range user.Roles {
			roles = append(roles, role.Alias)
		}
		_ = menuRepository.GetPermissionByUser(user, &permission)

		token, _ := models.GenToken(models.JwtUser{}.NewJwtUser(
			user.ID,
			user.Name,
			roles,
			permission,
		), global.Config.MyJwt.Secret)
		response.Success(c, "登陆成功", gin.H{
			"token":    token,
			"userInfo": user,
		})
	} else {
		response.Failed(c, "用户名或密码错误")
	}

}

// Add 注册用户
func (u *UserController) Add(c *gin.Context) {
	var (
		userAdd        requests.UserAdd
		userRepository repositorys.AdminRepository
	)
	err := c.ShouldBind(&userAdd)
	if err != nil {
		response.Failed(c, err.Error())
		return
	}
	result, model := userRepository.Add(userAdd.PassWord, userAdd.Name, userAdd)

	if result.Error == nil {
		response.Success(c, "ok", model)
	} else {
		response.Failed(c, result.Error.Error())
	}
}

func (u *UserController) List(c *gin.Context) {
	var (
		params         requests.AdminList
		userRepository repositorys.AdminRepository
	)
	_ = c.ShouldBind(&params)

	userRepository.Where = params.Where
	response.Success(c, "ok", userRepository.List(params.Page, params.PageSize, "created_at"))
}

func (u *UserController) All(c *gin.Context) {
	// 定义返回数据结构
	type UserResponse struct {
		ID       uint   `json:"id"`
		RealName string `json:"real_name"`
	}

	var users []UserResponse // 定义接收查询结果的变量

	// 执行查询
	if err := global.Db.Model(&models.AdminUser{}).
		Select("id", "real_name").
		Order("id desc").
		Find(&users).Error; err != nil {
		response.InsetErrorMsg(err, c)
		return
	}

	response.OkWithData(users, c)
}

func (u *UserController) Up(c *gin.Context) {

	var (
		data           requests.AdminUpdate
		userRepository repositorys.AdminRepository
	)
	err := c.ShouldBind(&data)
	if err != nil {
		response.Failed(c, global.GetError(err, data))
		return
	}

	reErr := userRepository.Update(data)
	if reErr == nil {
		response.Success(c, "ok", "")
	} else {
		response.Failed(c, reErr.Error())
	}
}

func (u *UserController) Dels(c *gin.Context) {
	var delIds global.Del
	_ = c.ShouldBind(&delIds)
	result := global.Db.Delete(&models.AdminUser{}, delIds.Ids)
	if result.Error == nil {
		response.Success(c, "ok", "")
	} else {
		response.Failed(c, result.Error.Error())
	}
}

func (u *UserController) Del(c *gin.Context) {
	// 1. 参数校验
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		response.FrontDataError(err, c)
		return
	}

	// 2. 检查用户是否存在
	var user models.AdminUser
	if err := global.Db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Failed(c, "用户不存在")
		} else {
			response.Failed(c, "查询用户失败")
		}
		return
	}

	// 3. 执行删除（物理删除）
	result := global.Db.Delete(&user) // 如果用软删除，去掉 Unscoped()
	if result.Error != nil {
		response.Failed(c, result.Error.Error())
		return
	}

	// 4. 成功返回（204 无内容）
	response.DeleteSuccessMsg(c)
}
