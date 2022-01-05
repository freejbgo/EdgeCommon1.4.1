// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package userconfigs

type UserRegisterConfig struct {
	IsOn                bool     `yaml:"isOn" json:"isOn"`                               // 是否启用用户注册
	ClusterId           int64    `yaml:"clusterId" json:"clusterId"`                     // 用户创建服务集群
	ServerGroupIds      []int64  `yaml:"serverGroupIds" json:"serverGroupIds"`           // 用户注册的服务所在分组
	ComplexPassword     bool     `yaml:"complexPassword" json:"complexPassword"`         // 必须使用复杂密码
	Features            []string `yaml:"features" json:"features"`                       // 默认启用的功能
	RequireVerification bool     `yaml:"requireVerification" json:"requireVerification"` // 是否需要审核
}

func DefaultUserRegisterConfig() *UserRegisterConfig {
	return &UserRegisterConfig{
		IsOn:            false,
		ComplexPassword: true,
		Features: []string{
			UserFeatureCodeServerAccessLog,
			UserFeatureCodeServerViewAccessLog,
			UserFeatureCodeServerWAF,
			UserFeatureCodeFinance,
			UserFeatureCodePlan,
		},
		RequireVerification: false,
	}
}
