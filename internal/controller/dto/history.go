package dto

type History struct {
	Name                       string            `json:"name,omitempty" validate:"required,gte=2,lte=200" example:"ООО 'Парк'"`
	OrganizationLegal          string            `json:"organizationLegal,omitempty" validate:"required" example:"ИП"`
	IndustryBranch             string            `json:"industryBranch,omitempty" validate:"required" example:"Авиационная_промышленность"`
	FullTimeEmployees          int               `json:"fullTimeEmployees,omitempty" validate:"required,lte=0" example:"50"`
	AvgSalary                  float64           `json:"avgSalary,omitempty" validate:"required,lte=0" example:"3058.12"`
	DistrictTitle              string            `json:"districtTitle,omitempty" validate:"required" example:"ВАО"`
	LandArea                   float64           `json:"landArea,omitempty" validate:"required,lte=0" example:"120"`
	IsBuy                      bool              `json:"isBuy,omitempty" validate:"required" example:"true"`
	ConstructionFacilitiesArea float64           `json:"constructionFacilitiesArea,omitempty" validate:"required,lte=0" example:"50"`
	BuildingType               string            `json:"buildingType,omitempty" validate:"required" example:"Энергетическое"`
	Equipment                  []Equipment       `json:"equipment,omitempty" validate:"omitempty"`
	AccountingSupport          bool              `json:"accountingSupport,omitempty" validate:"required" example:"true"`
	TaxationSystemOperations   *int              `json:"taxationSystemOperations,omitempty" validate:"omitempty"`
	OperationsNum              *int              `json:"operationsNum,omitempty" validate:"omitempty,lte=0"`
	PatentCalc                 bool              `json:"patentCalc,omitempty" validate:"required" example:"true"`
	BusinessActivity           *BusinessActivity `json:"business_activity,omitempty" validate:"omitempty"`
	Other                      *string           `json:"other,omitempty" validate:"omitempty" example:"I want some cookies"`
}

type HistoryId struct {
	Id string `json:"id,omitempty" validate:"required,number" example:"12"`
}
