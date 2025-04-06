package boot

func Run() (err error) {
	if err = InitConfig(); err != nil {
		return
	}

	InitLogger()

	if err = NewService(); err != nil {
		return err
	}

	InitSwagger()

	return nil
}
