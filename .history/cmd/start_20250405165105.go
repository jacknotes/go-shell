package cmd

import (
	"github.com/jacknotes/go-shell.git/conf"
	"github.com/jacknotes/go-shell.git/dao"
	"github.com/spf13/cobra"
)

var (
	confFile string
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "启动 go-shell",
	Long:  "启动 go-shell",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 加载程序配置
		err := conf.LoadConfigFromToml(confFile)
		if err != nil {
			return err
		}

		err := dao.WriteDB(conf.GetFile)
		if err != nil {
			return err
		}
		for i := range conf.GetFile.App.Code {
			err := dao.WriteDB(conf.GetFile.App.Code[i])

		}
		return nil
	},
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f", "code.toml", "code文件路径")
	RootCmd.AddCommand(StartCmd)
}
