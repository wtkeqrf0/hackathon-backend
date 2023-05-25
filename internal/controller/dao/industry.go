package dao

type Industry struct {
	AvgWorkersNum    float64 `json:"avgWorkersNum,omitempty" sql:"avg_workers_num" validate:"required"`
	AvgWorkersNumCad float64 `json:"avgWorkersNumCad,omitempty" sql:"avg_workers_num_cad" validate:"required"`
	AvgSalary        float64 `json:"avgSalary,omitempty" sql:"avg_salary" validate:"required"`
	AvgSalaryCad     float64 `json:"avgSalaryCad,omitempty" sql:"avg_salary_cad" validate:"required"`
}
