package server

import (
	"auth-template/internal/dto"
	"auth-template/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetAllUsers(ctx *gin.Context) {
	users, err := s.userService.GetAll()

	if err != nil {
		utils.Fail(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.OK(ctx, http.StatusOK, users)
}

func (s *Server) GetUser(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.Fail(ctx, http.StatusBadRequest, "invalid id")
		return
	}

	user, err := s.userService.GetById(int(id))
	if err != nil {
		utils.Fail(ctx, http.StatusNotFound, "user not found")
		return
	}
	utils.OK(ctx, http.StatusOK, user)
}

func (s *Server) CreateUser(ctx *gin.Context) {
	var req dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Fail(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := s.userService.Create(req.Name, req.Email, req.Password)
	if err != nil {
		utils.Fail(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.OK(ctx, http.StatusCreated, user)
}
