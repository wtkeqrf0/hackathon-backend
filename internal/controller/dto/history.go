package dto

type History struct {
	CompanyName                string `json:"companyName,omitempty" validate:"required,gte=2,lte=150"`
	IndustryBranch             string `json:"industryBranch,omitempty" validate:"required"`
	FullTimeEmployees          int    `json:"fullTimeEmployees,omitempty" validate:"required,lte=0"`
	DistrictTitle              string `json:"districtTitle,omitempty" validate:"required"`
	LandArea                   int    `json:"landArea,omitempty" validate:"required,lte=0"`
	ConstructionFacilitiesArea int    `json:"constructionFacilitiesArea,omitempty" validate:"required,lte=0"`
	EquipmentType              string `json:"equipmentType,omitempty" validate:"required"`
	FacilityType               string `json:"facilityType,omitempty" validate:"required"`
	AccountingServices         bool   `json:"accountingServices,omitempty" validate:"required"`
	Patent                     bool   `json:"patent,omitempty" validate:"required"`
	Other                      string `json:"other,omitempty" validate:"required"`
}

type CompanyName struct {
	CompanyName string `json:"companyName,omitempty" validate:"required,gte=2,lte=150"`
}
