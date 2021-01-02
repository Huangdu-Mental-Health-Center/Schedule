package model

type HospitalSchedule struct {
	HospitalName   string `json:"hospital_name"`
	HospitalRegion string `json:"hospital_region"`
	Data           []struct {
		Name         string         `json:"name"`
		Department   string         `json:"department"`
		ScheduleList []ScheduleInfo `json:"schedule_list"`
	} `json:"data"`
}
type DoctorSchedule struct {
	Msg          string         `json:"msg"`
	HospitalName string         `json:"hospital_name"`
	Name         string         `json:"name"`
	Department   string         `json:"department"`
	ScheduleList []ScheduleInfo `json:"schedule_list"`
}
type ScheduleInfo struct {
	Date              string `json:"date"`
	TimeSlot          int    `json:"time_slot"`
	ProfessionalTitle string `json:"professional_title"`
	Price             int    `json:"price"`
}
