



func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f", "codelist.txt", "code list文件路径")
	RootCmd.AddCommand(StartCmd)
}