package dto

type History struct {
	CompanyName                string `json:"companyName,omitempty" validate:"required,gte=2,lte=150" example:"ООО 'Парк'"`
	IndustryBranch             string `json:"industryBranch,omitempty" validate:"required" example:"Авиационная промышленность"`
	FullTimeEmployees          int    `json:"fullTimeEmployees,omitempty" validate:"required,lte=0" example:"50"`
	DistrictTitle              string `json:"districtTitle,omitempty" validate:"required" example:"ВАО"`
	LandArea                   int    `json:"landArea,omitempty" validate:"required,lte=0" example:"120"`
	ConstructionFacilitiesArea int    `json:"constructionFacilitiesArea,omitempty" validate:"required,lte=0" example:"50"`
	EquipmentType              string `json:"equipmentType,omitempty" validate:"required" example:"Токарные станки"`
	FacilityType               string `json:"facilityType,omitempty" validate:"required" example:"idk"` //TODO
	AccountingServices         bool   `json:"accountingServices,omitempty" validate:"required" example:"true"`
	Patent                     bool   `json:"patent,omitempty" validate:"required" example:"true"`
	Other                      string `json:"other,omitempty" validate:"required" example:"I want some cookies"`
}

type CompanyName struct {
	CompanyName string `json:"companyName,omitempty" validate:"required,gte=2,lte=150" example:"ООО 'Парк'"`
}
