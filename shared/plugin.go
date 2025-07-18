package plugin

import "github.com/gliderlabs/ssh"

type ShowFunc func(s ssh.Session, idx int, items []*MenuItem) bool              // 是否显示
type GetSubMenuFunc func(s ssh.Session, idx int, items []*MenuItem) []*MenuItem // 获取子菜单
type SelectedFunc func(s ssh.Session, idx int, items []*MenuItem) error         // 选择子菜单

type MenuItem struct {
	Name              string         // 插件名
	ShowFunc          ShowFunc       // 控制是否显示插件
	Title             string         // 标题
	MenuTitle         string         // 菜单标题
	NoSubMenuTitle    string         // 子菜单为空时候的标题
	GetSubMenuFunc    GetSubMenuFunc // 获取子菜单函数
	SelectedFunc      SelectedFunc   // 选择子菜单函数
	BackAfterSelected bool           // 选择后是否返回
	BackOptionLabel   string         // 返回的标签
}

func DefaultShow(s ssh.Session, idx int, items []*MenuItem) bool { return true }
