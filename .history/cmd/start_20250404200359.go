
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
	}
}


func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f", "codelist.txt", "code list文件路径")
	RootCmd.AddCommand(StartCmd)
}