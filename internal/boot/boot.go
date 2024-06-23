package boot

func InitAll() (err error) {
	InitLogger()
	if err = InitConfig(); err != nil {
		return
	}
	InitSwagger()

	return nil
}
