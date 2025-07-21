package shared

import (
	"github.com/gliderlabs/ssh"
)

type Menu interface {
	Name() string                                                  // 插件名字
	Title() string                                                 // 插件标题，菜单选择时显示
	SubMenuTitle() string                                          // 子菜单标题，进入子菜单后显示
	EmptyMenuMessage() string                                      // 空菜单提示，无子菜单时显示
	Show(session ssh.Session, index int, plugin []Menu) bool       // 是否显示插件
	SubMenus(session ssh.Session, index int, plugin []Menu) []Menu // 获取子菜单列表
	Selected(session ssh.Session, index int, plugin []Menu) error  // 选择菜单后的处理
	BackOptionLabel() string                                       // 返回按钮标签
	AutoReturn() bool                                              // 是否自动返回上级菜单
}
