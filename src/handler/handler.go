package handler

import (
	"Huangdu_HMC_Schedule/src/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	doctor     string
	hospital   string
	department string
}

func (h *Handler) GetDoctor(context *gin.Context) {
	var scheduleList []model.ScheduleInfo
	var res model.DoctorSchedule
	h.doctor = context.Query("doctor_name")
	h.hospital = context.Query("hospital_name")
	h.department = context.Query("department")
	if h.doctor == "" || h.department == "" || h.hospital == "" {
		res.Msg = "Bad Parameter"
		context.JSON(http.StatusBadRequest, res)
		return
	}

	list1 := h.handleDB()
	list2 := h.handleAPI()
	scheduleList = append(list1, list2...)
	if scheduleList == nil {
		res.Msg = "No Data"
		context.JSON(http.StatusNotFound, res)
		return
	}
	res.ScheduleList = scheduleList
	res.Name = h.doctor
	res.Department = h.department
	res.HospitalName = h.hospital
	res.Msg = "OK"
	context.JSON(http.StatusOK, res)
}
