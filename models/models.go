package models

const (
	CMD_NUM     = 10
	DESC_LEN    = 1024
	CMD_MAX_LEN = 128
)

type DataNode struct {
	Cmd     string
	Desc    string
	Handler func()
	Next    *DataNode
}

var (
	Head *DataNode
	cmds [CMD_NUM]DataNode
)

func init() {
	Head = &cmds[0]
	cmds[0] = DataNode{
		Cmd:     "help",
		Desc:    "this is help cmd!",
		Handler: Help,
		Next:    &cmds[1],
	}
	cmds[1] = DataNode{
		Cmd:     "version",
		Desc:    "menu program v1.0",
		Handler: nil,
		Next:    &cmds[2],
	}
	cmds[2] = DataNode{
		Cmd:     "quit",
		Desc:    "Quit from menu",
		Handler: Quit,
		Next:    nil,
	}
}
