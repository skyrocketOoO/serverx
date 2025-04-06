package boot

func Run() (err error) {
	if err = InitConfig(); err != nil {
		return
	}

	if err = InitLogger(); err != nil {
		return err
	}

	if err = NewService(); err != nil {
		return err
	}

	InitSwagger()

	return nil
}
