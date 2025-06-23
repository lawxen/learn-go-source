package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
)

func main() {
	// 读取模型和策略文件
	e, _ := casbin.NewEnforcer("model.conf", "policy.csv")
	// 加载策略文件
	e.LoadPolicy()
	// 将角色和权限添加到策略文件policy.csv
	e.AddGroupingPolicy("alice", "admin")
	e.AddGroupingPolicy("alice", "root")
	e.AddGroupingPolicy("many", "root")
	e.AddPolicy("admin", "/users", "GET")
	e.AddPolicy("admin", "/users/*", "(GET)|(POST)")
	// 保存策略文件
	e.SavePolicy()
	// 获取用户所在的角色
	rfu, _ := e.GetRolesForUser("alice")
	fmt.Println(rfu)
	// 获取角色的所有用户
	ufr, _ := e.GetUsersForRole("admin")
	fmt.Println(ufr)
	// 获取用户权限
	pfu := e.GetPermissionsForUser("admin")
	fmt.Println(pfu)
	// 修改用户和角色
	e.UpdateGroupingPolicy([]string{"many", "root"}, []string{"tom", "admin"})
	// 修改策略
	e.UpdatePolicy([]string{"admin", "/users/*"}, []string{"admin", "/infos/*", "(POST)|(PUT)"})
	// 删除角色
	e.DeleteRole("root")
	// 删除角色的某个权限
	e.DeletePermissionForUser("admin", "/users")
	// 检查访问控制
	allowed, _ := e.Enforce("alice", "/infos/aaa", "GET")
	if allowed {
		println("允许访问")
	} else {
		println("暂无权限")
	}
}
