package dao

type Industry struct {
	AvgWorkersNum    float64 `json:"avgWorkersNum,omitempty" sql:"avg_workers_num" validate:"required" example:"1.208"`
	AvgWorkersNumCad float64 `json:"avgWorkersNumCad,omitempty" sql:"avg_workers_num_cad" validate:"required" example:"1243.0"`
	AvgSalary        float64 `json:"avgSalary,omitempty" sql:"avg_salary" validate:"required" example:"72.7825875"`
	AvgSalaryCad     float64 `json:"avgSalaryCad,omitempty" sql:"avg_salary_cad" validate:"required" example:"95.54196489"`
}
