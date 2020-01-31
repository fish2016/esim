package new

func ServiceInit() {
	fc1 := &FileContent{
		FileName: "user.go",
		Dir:      "internal/app/service",
		Content: `package app

import (
	"context"
	"{{PROPATH}}{{service_name}}/internal/infra"
	"{{PROPATH}}{{service_name}}/internal/domain/entity"
)

type UserService struct {
	*infra.Infra
}

func NewUserSvc(infra *infra.Infra)*UserService{
	svc := &UserService{infra}

	return svc
}

func (svc *UserService) GetUserInfo(ctx context.Context, username string) (user entity.User) {

	user = svc.UserRepo.FindByUserName(ctx, username)

	return
}
`,
	}



	fc2 := &FileContent{
		FileName: "README.md",
		Dir:      "internal/app/dto",
		Content:  `for DTO`,
	}


	fc3 := &FileContent{
		FileName: "user.go",
		Dir:      "internal/app/dto",
		Content: `package dto

import "gitlab.etcchebao.cn/go_service/{{service_name}}/internal/domain/entity"

type User struct {

	//用户名称
	UserName string {{!}}json:"user_name"{{!}}

	//密码
	PassWord string {{!}}json:"pass_word"{{!}}
}

func NewUser(user entity.User) User {
	dto := User{}

	dto.UserName = user.UserNmae
	dto.PassWord = user.PassWord
	return dto
}`,
	}

	Files = append(Files, fc1, fc2, fc3)
}
