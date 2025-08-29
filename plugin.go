package jump_access_plugin

import (
	"context"
)

type Plugin interface {
	Name(ctx context.Context) string                                     // 插件名字
	Title(ctx context.Context) string                                    // 插件标题，选择时候显示
	SubMenuTitle(ctx context.Context) string                             // 子菜单标题，进入子菜单后显示
	EmptyMenuInfo(ctx context.Context) string                            // 空菜单提示，无子菜单时显示
	IsShow(ctx context.Context, index int, plugin []Plugin) bool         // 是否显示插件
	GetSubMenu(ctx context.Context, index int, plugin []Plugin) []Plugin // 获取子菜单列表
	OnSelected(ctx context.Context, index int, plugin []Plugin) error    // 选择菜单后的处理
	BackOptionLabel(ctx context.Context) string                          // 返回按钮标签
	AutoReturn(ctx context.Context) bool                                 // 是否自动返回上级菜单, true的话，将会在处理完成后自动返回上级菜单
}
