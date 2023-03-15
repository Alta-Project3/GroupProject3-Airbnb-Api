package handler

import (
	"groupproject3-airbnb-api/features/rooms"
	"groupproject3-airbnb-api/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RoomHandler struct {
	Service rooms.RoomServiceInterface
}

func New(s rooms.RoomServiceInterface) *RoomHandler {
	return &RoomHandler{
		Service: s,
	}
}

func (h *RoomHandler) GetAll(c echo.Context) error {
	classEntity, err := h.Service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFail("error read data"))
	}
	listClassResponse := ListRoomEntityToRoomResponse(classEntity)
	return c.JSON(http.StatusOK, helper.ResponseSuccess("-", listClassResponse))
}

func (h *RoomHandler) GetById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	roomEntity, err := h.Service.GetById(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.ResponseFail("data not found"))
	}
	teamResponse := RoomEntityToRoomResponse(roomEntity)
	return c.JSON(http.StatusOK, helper.ResponseSuccess("-", teamResponse))
}

func (h *RoomHandler) GetByUserId(c echo.Context) error {
	userIdToken := helper.ClaimToken(c.Get("user"))
	userId, _ := strconv.Atoi(c.Param("id"))

	classEntity, err := h.Service.GetByUserId(uint(userId), uint(userIdToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFail(err.Error()))
	}
	listClassResponse := ListRoomEntityToRoomResponse(classEntity)
	return c.JSON(http.StatusOK, helper.ResponseSuccess("-", listClassResponse))
}

func (h *RoomHandler) Create(c echo.Context) error {
	var formInput RoomRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFail("error bind data"))
	}

	userId := helper.ClaimToken(c.Get("user"))

	team, err := h.Service.Create(RoomRequestToRoomEntity(&formInput), uint(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusCreated, helper.ResponseSuccess("Create Data Success", RoomEntityToRoomResponse(team)))
}

func (h *RoomHandler) Update(c echo.Context) error {
	var formInput RoomRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFail("error bind data"))
	}

	userId := helper.ClaimToken(c.Get("user"))

	id, _ := strconv.Atoi(c.Param("id"))
	team, err := h.Service.Update(RoomRequestToRoomEntity(&formInput), uint(id), uint(userId))
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccess("Update Data Success", RoomEntityToRoomResponse(team)))
}

func (h *RoomHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	userId := helper.ClaimToken(c.Get("user"))

	if err := h.Service.Delete(uint(id), uint(userId)); err != nil {
		return c.JSON(http.StatusNotFound, helper.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccess("Delete Data Success", nil))
}
