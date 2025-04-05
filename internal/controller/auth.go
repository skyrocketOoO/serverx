package controller

// @Param   user  body  controller.Login.Req  true  "Login User"
// @Success 200 {object} controller.Login.Resp "token"
// @Failure 500 {string} domain.ErrResp
// @Failure 400 {object} domain.ErrResp
// @Router /login [post]
// func (h *Handler) Login(c *gin.Context) {
// 	type Req struct {
// 		Name     string `json:"Name"     validate:"required"`
// 		Password string `json:"Password" validate:"required"`
// 	}

// 	var req Req
// 	if ok := util.ParseValidate(c, &req); !ok {
// 		return
// 	}

// 	hashedPassword := string(auth.Hash(req.Password, util.GetSalt()))
// 	db := postgres.Get()

// 	var user models.User
// 	if err := db.
// 		Model(&models.User{}).
// 		Where(wh.B(col.Users.Name, ope.Eq), req.Name).
// 		Where(wh.B(col.Users.Password, ope.Eq), hashedPassword).
// 		Take(&user).
// 		Error; err != nil {
// 		util.ErrResp(c, http.StatusInternalServerError, erx.W(domain.ErrLoginFailed))
// 		return
// 	}

// 	token, err := util.GenerateToken(user.ID)
// 	if err != nil {
// 		util.ErrResp(c, http.StatusInternalServerError, erx.W(err))
// 		return
// 	}

// 	token = "Bearer " + token

// 	type Resp struct {
// 		Token string `json:"token"`
// 	}
// 	c.JSON(http.StatusOK, Resp{Token: token})
// }

// // @Param   user  body  controller.Register.Req  true  "Register"
// // @Success 200
// // @Failure 500 {object} domain.ErrResp
// // @Failure 400 {object} domain.ErrResp
// // @Router /register [post]
// func (h *Handler) Register(c *gin.Context) {
// 	type Req struct {
// 		Name     string `json:"Name"     validate:"required,min=6,max=32"`
// 		Password string `json:"Password" validate:"required,min=8,max=32"`
// 	}

// 	var req Req
// 	if ok := util.ParseValidate(c, &req); !ok {
// 		return
// 	}

// 	db := postgres.Get()
// 	var existingUser models.User
// 	if err := db.Where(wh.B(col.Users.Name, ope.Eq), req.Name).
// 		Take(&existingUser).Error; err != nil {
// 		if err != gorm.ErrRecordNotFound {
// 			util.ErrResp(c, http.StatusInternalServerError, erx.W(err))
// 			return
// 		}
// 	} else {
// 		util.ErrResp(c, http.StatusInternalServerError, erx.W(domain.ErrUserNameRepetite))
// 		return
// 	}

// 	if err := db.Create(&models.User{
// 		Name:     req.Name,
// 		Password: string(auth.Hash(req.Password, util.GetSalt())),
// 	}).Error; err != nil {
// 		util.ErrResp(c, http.StatusInternalServerError, erx.W(err))
// 		return
// 	}

// 	c.Status(http.StatusOK)
// }

// func (h *Handler) ForgetPassword(c *gin.Context) {
// 	util.ErrResp(c, http.StatusNotFound, erx.W(domain.ErrNotImplement))
// 	return
// }
