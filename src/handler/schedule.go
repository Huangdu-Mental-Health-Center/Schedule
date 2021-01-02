package handler

import (
	"Huangdu_HMC_Schedule/src/driver"
	"Huangdu_HMC_Schedule/src/logger"
	"database/sql"

	"Huangdu_HMC_Schedule/src/model"
)

var nameDB = [...]string{"医院A", "医院C"}
var nameFile = [...]string{"医院B"}

func (h *Handler) queryFromDB(pDB *sql.DB) []model.ScheduleInfo {
	var resInfo []model.ScheduleInfo
	rows, err := pDB.Query(
		"SELECT doctor_date,time_slot,professional_title,price FROM schedule  WHERE doctor_name=? and department=? ",
		h.doctor, h.department)
	if err != nil {
		logger.Error.Println(err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var tmpInfo model.ScheduleInfo
		err1 := rows.Scan(&tmpInfo.Date, &tmpInfo.TimeSlot, &tmpInfo.ProfessionalTitle, &tmpInfo.Price)
		if err1 != nil {
			logger.Error.Println(err1)
			continue
		}
		//Error.Println(tmpInfo.Date)
		resInfo = append(resInfo, tmpInfo)
	}
	err = rows.Err()
	if err != nil {
		logger.Error.Println(err)
	}
	return resInfo
}

func (h *Handler) handleDB() []model.ScheduleInfo {
	isHospital := false
	var dbName string
	for _, hospital := range nameDB {
		if hospital == h.hospital {
			isHospital = true
			dbName = hospital
		}
	}
	if !isHospital {
		return nil
	}
	isOpen, pDB := driver.ConnectDB(dbName)
	if !isOpen {
		logger.Error.Printf("Failed to connect %s database\n", dbName)
		return nil
	}
	list := h.queryFromDB(pDB)
	err := pDB.Close()
	if err != nil {
		logger.Error.Println(err)
	}
	return list
}
func (h *Handler) handleAPI() []model.ScheduleInfo {
	isHospital := false
	var fileName string
	for _, hospital := range nameFile {
		if hospital == h.hospital {
			isHospital = true
			fileName = hospital
		}
	}
	if !isHospital {
		return nil
	}
	v := model.HospitalSchedule{}
	driver.LoadFile(fileName, &v)
	for _, doctorSchedule := range v.Data {
		if doctorSchedule.Name == h.doctor &&
			doctorSchedule.Department == h.department {
			return doctorSchedule.ScheduleList
		}
	}
	return nil
}
