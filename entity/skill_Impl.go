package entity

import "github.com/sunist-c/genius-invokation-simulator-backend/enum"

type xiaogong1 struct {
	id          uint             // id 角色的ID，由框架确定
	name        string           // name 角色的名称
	description string           // description 角色的描述
	affiliation enum.Affiliation // affiliation 角色的势力归属
}
