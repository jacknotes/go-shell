package conf

func LoadFile() {

}

func LoadFile(filePath string) error {
	config = NewDefaultConfig()

	// 读取Toml格式的配置
	_, err := toml.DecodeFile(filePath, config)
	if err != nil {
		return fmt.Errorf("load config from file error, path:%s, %s", filePath, err)
	}

	return nil
}
