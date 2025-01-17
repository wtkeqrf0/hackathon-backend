// Code generated by ent, DO NOT EDIT.

package history

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/while-act/hackathon-backend/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.History {
	return predicate.History(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.History {
	return predicate.History(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.History {
	return predicate.History(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.History {
	return predicate.History(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.History {
	return predicate.History(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.History {
	return predicate.History(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.History {
	return predicate.History(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.History {
	return predicate.History(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.History {
	return predicate.History(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.History {
	return predicate.History(sql.FieldEQ(FieldName, v))
}

// OrganizationalLegal applies equality check predicate on the "organizational_legal" field. It's identical to OrganizationalLegalEQ.
func OrganizationalLegal(v string) predicate.History {
	return predicate.History(sql.FieldEQ(FieldOrganizationalLegal, v))
}

// IndustryBranch applies equality check predicate on the "industry_branch" field. It's identical to IndustryBranchEQ.
func IndustryBranch(v string) predicate.History {
	return predicate.History(sql.FieldEQ(FieldIndustryBranch, v))
}

// FullTimeEmployees applies equality check predicate on the "full_time_employees" field. It's identical to FullTimeEmployeesEQ.
func FullTimeEmployees(v int) predicate.History {
	return predicate.History(sql.FieldEQ(FieldFullTimeEmployees, v))
}

// AvgSalary applies equality check predicate on the "avg_salary" field. It's identical to AvgSalaryEQ.
func AvgSalary(v float64) predicate.History {
	return predicate.History(sql.FieldEQ(FieldAvgSalary, v))
}

// DistrictTitle applies equality check predicate on the "district_title" field. It's identical to DistrictTitleEQ.
func DistrictTitle(v string) predicate.History {
	return predicate.History(sql.FieldEQ(FieldDistrictTitle, v))
}

// LandArea applies equality check predicate on the "land_area" field. It's identical to LandAreaEQ.
func LandArea(v float64) predicate.History {
	return predicate.History(sql.FieldEQ(FieldLandArea, v))
}

// IsBuy applies equality check predicate on the "is_buy" field. It's identical to IsBuyEQ.
func IsBuy(v bool) predicate.History {
	return predicate.History(sql.FieldEQ(FieldIsBuy, v))
}

// ConstructionFacilitiesArea applies equality check predicate on the "construction_facilities_area" field. It's identical to ConstructionFacilitiesAreaEQ.
func ConstructionFacilitiesArea(v float64) predicate.History {
	return predicate.History(sql.FieldEQ(FieldConstructionFacilitiesArea, v))
}

// BuildingType applies equality check predicate on the "building_type" field. It's identical to BuildingTypeEQ.
func BuildingType(v string) predicate.History {
	return predicate.History(sql.FieldEQ(FieldBuildingType, v))
}

// AccountingSupport applies equality check predicate on the "accounting_support" field. It's identical to AccountingSupportEQ.
func AccountingSupport(v bool) predicate.History {
	return predicate.History(sql.FieldEQ(FieldAccountingSupport, v))
}

// TaxationSystemOperations applies equality check predicate on the "taxation_system_operations" field. It's identical to TaxationSystemOperationsEQ.
func TaxationSystemOperations(v int) predicate.History {
	return predicate.History(sql.FieldEQ(FieldTaxationSystemOperations, v))
}

// OperationType applies equality check predicate on the "operation_type" field. It's identical to OperationTypeEQ.
func OperationType(v string) predicate.History {
	return predicate.History(sql.FieldEQ(FieldOperationType, v))
}

// PatentCalc applies equality check predicate on the "patent_calc" field. It's identical to PatentCalcEQ.
func PatentCalc(v bool) predicate.History {
	return predicate.History(sql.FieldEQ(FieldPatentCalc, v))
}

// BusinessActivityID applies equality check predicate on the "business_activity_id" field. It's identical to BusinessActivityIDEQ.
func BusinessActivityID(v int) predicate.History {
	return predicate.History(sql.FieldEQ(FieldBusinessActivityID, v))
}

// Other applies equality check predicate on the "other" field. It's identical to OtherEQ.
func Other(v string) predicate.History {
	return predicate.History(sql.FieldEQ(FieldOther, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.History {
	return predicate.History(sql.FieldEQ(FieldUserID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.History {
	return predicate.History(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.History {
	return predicate.History(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.History {
	return predicate.History(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.History {
	return predicate.History(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.History {
	return predicate.History(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.History {
	return predicate.History(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.History {
	return predicate.History(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.History {
	return predicate.History(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.History {
	return predicate.History(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.History {
	return predicate.History(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.History {
	return predicate.History(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.History {
	return predicate.History(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.History {
	return predicate.History(sql.FieldContainsFold(FieldName, v))
}

// OrganizationalLegalEQ applies the EQ predicate on the "organizational_legal" field.
func OrganizationalLegalEQ(v string) predicate.History {
	return predicate.History(sql.FieldEQ(FieldOrganizationalLegal, v))
}

// OrganizationalLegalNEQ applies the NEQ predicate on the "organizational_legal" field.
func OrganizationalLegalNEQ(v string) predicate.History {
	return predicate.History(sql.FieldNEQ(FieldOrganizationalLegal, v))
}

// OrganizationalLegalIn applies the In predicate on the "organizational_legal" field.
func OrganizationalLegalIn(vs ...string) predicate.History {
	return predicate.History(sql.FieldIn(FieldOrganizationalLegal, vs...))
}

// OrganizationalLegalNotIn applies the NotIn predicate on the "organizational_legal" field.
func OrganizationalLegalNotIn(vs ...string) predicate.History {
	return predicate.History(sql.FieldNotIn(FieldOrganizationalLegal, vs...))
}

// OrganizationalLegalGT applies the GT predicate on the "organizational_legal" field.
func OrganizationalLegalGT(v string) predicate.History {
	return predicate.History(sql.FieldGT(FieldOrganizationalLegal, v))
}

// OrganizationalLegalGTE applies the GTE predicate on the "organizational_legal" field.
func OrganizationalLegalGTE(v string) predicate.History {
	return predicate.History(sql.FieldGTE(FieldOrganizationalLegal, v))
}

// OrganizationalLegalLT applies the LT predicate on the "organizational_legal" field.
func OrganizationalLegalLT(v string) predicate.History {
	return predicate.History(sql.FieldLT(FieldOrganizationalLegal, v))
}

// OrganizationalLegalLTE applies the LTE predicate on the "organizational_legal" field.
func OrganizationalLegalLTE(v string) predicate.History {
	return predicate.History(sql.FieldLTE(FieldOrganizationalLegal, v))
}

// OrganizationalLegalContains applies the Contains predicate on the "organizational_legal" field.
func OrganizationalLegalContains(v string) predicate.History {
	return predicate.History(sql.FieldContains(FieldOrganizationalLegal, v))
}

// OrganizationalLegalHasPrefix applies the HasPrefix predicate on the "organizational_legal" field.
func OrganizationalLegalHasPrefix(v string) predicate.History {
	return predicate.History(sql.FieldHasPrefix(FieldOrganizationalLegal, v))
}

// OrganizationalLegalHasSuffix applies the HasSuffix predicate on the "organizational_legal" field.
func OrganizationalLegalHasSuffix(v string) predicate.History {
	return predicate.History(sql.FieldHasSuffix(FieldOrganizationalLegal, v))
}

// OrganizationalLegalEqualFold applies the EqualFold predicate on the "organizational_legal" field.
func OrganizationalLegalEqualFold(v string) predicate.History {
	return predicate.History(sql.FieldEqualFold(FieldOrganizationalLegal, v))
}

// OrganizationalLegalContainsFold applies the ContainsFold predicate on the "organizational_legal" field.
func OrganizationalLegalContainsFold(v string) predicate.History {
	return predicate.History(sql.FieldContainsFold(FieldOrganizationalLegal, v))
}

// IndustryBranchEQ applies the EQ predicate on the "industry_branch" field.
func IndustryBranchEQ(v string) predicate.History {
	return predicate.History(sql.FieldEQ(FieldIndustryBranch, v))
}

// IndustryBranchNEQ applies the NEQ predicate on the "industry_branch" field.
func IndustryBranchNEQ(v string) predicate.History {
	return predicate.History(sql.FieldNEQ(FieldIndustryBranch, v))
}

// IndustryBranchIn applies the In predicate on the "industry_branch" field.
func IndustryBranchIn(vs ...string) predicate.History {
	return predicate.History(sql.FieldIn(FieldIndustryBranch, vs...))
}

// IndustryBranchNotIn applies the NotIn predicate on the "industry_branch" field.
func IndustryBranchNotIn(vs ...string) predicate.History {
	return predicate.History(sql.FieldNotIn(FieldIndustryBranch, vs...))
}

// IndustryBranchGT applies the GT predicate on the "industry_branch" field.
func IndustryBranchGT(v string) predicate.History {
	return predicate.History(sql.FieldGT(FieldIndustryBranch, v))
}

// IndustryBranchGTE applies the GTE predicate on the "industry_branch" field.
func IndustryBranchGTE(v string) predicate.History {
	return predicate.History(sql.FieldGTE(FieldIndustryBranch, v))
}

// IndustryBranchLT applies the LT predicate on the "industry_branch" field.
func IndustryBranchLT(v string) predicate.History {
	return predicate.History(sql.FieldLT(FieldIndustryBranch, v))
}

// IndustryBranchLTE applies the LTE predicate on the "industry_branch" field.
func IndustryBranchLTE(v string) predicate.History {
	return predicate.History(sql.FieldLTE(FieldIndustryBranch, v))
}

// IndustryBranchContains applies the Contains predicate on the "industry_branch" field.
func IndustryBranchContains(v string) predicate.History {
	return predicate.History(sql.FieldContains(FieldIndustryBranch, v))
}

// IndustryBranchHasPrefix applies the HasPrefix predicate on the "industry_branch" field.
func IndustryBranchHasPrefix(v string) predicate.History {
	return predicate.History(sql.FieldHasPrefix(FieldIndustryBranch, v))
}

// IndustryBranchHasSuffix applies the HasSuffix predicate on the "industry_branch" field.
func IndustryBranchHasSuffix(v string) predicate.History {
	return predicate.History(sql.FieldHasSuffix(FieldIndustryBranch, v))
}

// IndustryBranchEqualFold applies the EqualFold predicate on the "industry_branch" field.
func IndustryBranchEqualFold(v string) predicate.History {
	return predicate.History(sql.FieldEqualFold(FieldIndustryBranch, v))
}

// IndustryBranchContainsFold applies the ContainsFold predicate on the "industry_branch" field.
func IndustryBranchContainsFold(v string) predicate.History {
	return predicate.History(sql.FieldContainsFold(FieldIndustryBranch, v))
}

// FullTimeEmployeesEQ applies the EQ predicate on the "full_time_employees" field.
func FullTimeEmployeesEQ(v int) predicate.History {
	return predicate.History(sql.FieldEQ(FieldFullTimeEmployees, v))
}

// FullTimeEmployeesNEQ applies the NEQ predicate on the "full_time_employees" field.
func FullTimeEmployeesNEQ(v int) predicate.History {
	return predicate.History(sql.FieldNEQ(FieldFullTimeEmployees, v))
}

// FullTimeEmployeesIn applies the In predicate on the "full_time_employees" field.
func FullTimeEmployeesIn(vs ...int) predicate.History {
	return predicate.History(sql.FieldIn(FieldFullTimeEmployees, vs...))
}

// FullTimeEmployeesNotIn applies the NotIn predicate on the "full_time_employees" field.
func FullTimeEmployeesNotIn(vs ...int) predicate.History {
	return predicate.History(sql.FieldNotIn(FieldFullTimeEmployees, vs...))
}

// FullTimeEmployeesGT applies the GT predicate on the "full_time_employees" field.
func FullTimeEmployeesGT(v int) predicate.History {
	return predicate.History(sql.FieldGT(FieldFullTimeEmployees, v))
}

// FullTimeEmployeesGTE applies the GTE predicate on the "full_time_employees" field.
func FullTimeEmployeesGTE(v int) predicate.History {
	return predicate.History(sql.FieldGTE(FieldFullTimeEmployees, v))
}

// FullTimeEmployeesLT applies the LT predicate on the "full_time_employees" field.
func FullTimeEmployeesLT(v int) predicate.History {
	return predicate.History(sql.FieldLT(FieldFullTimeEmployees, v))
}

// FullTimeEmployeesLTE applies the LTE predicate on the "full_time_employees" field.
func FullTimeEmployeesLTE(v int) predicate.History {
	return predicate.History(sql.FieldLTE(FieldFullTimeEmployees, v))
}

// AvgSalaryEQ applies the EQ predicate on the "avg_salary" field.
func AvgSalaryEQ(v float64) predicate.History {
	return predicate.History(sql.FieldEQ(FieldAvgSalary, v))
}

// AvgSalaryNEQ applies the NEQ predicate on the "avg_salary" field.
func AvgSalaryNEQ(v float64) predicate.History {
	return predicate.History(sql.FieldNEQ(FieldAvgSalary, v))
}

// AvgSalaryIn applies the In predicate on the "avg_salary" field.
func AvgSalaryIn(vs ...float64) predicate.History {
	return predicate.History(sql.FieldIn(FieldAvgSalary, vs...))
}

// AvgSalaryNotIn applies the NotIn predicate on the "avg_salary" field.
func AvgSalaryNotIn(vs ...float64) predicate.History {
	return predicate.History(sql.FieldNotIn(FieldAvgSalary, vs...))
}

// AvgSalaryGT applies the GT predicate on the "avg_salary" field.
func AvgSalaryGT(v float64) predicate.History {
	return predicate.History(sql.FieldGT(FieldAvgSalary, v))
}

// AvgSalaryGTE applies the GTE predicate on the "avg_salary" field.
func AvgSalaryGTE(v float64) predicate.History {
	return predicate.History(sql.FieldGTE(FieldAvgSalary, v))
}

// AvgSalaryLT applies the LT predicate on the "avg_salary" field.
func AvgSalaryLT(v float64) predicate.History {
	return predicate.History(sql.FieldLT(FieldAvgSalary, v))
}

// AvgSalaryLTE applies the LTE predicate on the "avg_salary" field.
func AvgSalaryLTE(v float64) predicate.History {
	return predicate.History(sql.FieldLTE(FieldAvgSalary, v))
}

// DistrictTitleEQ applies the EQ predicate on the "district_title" field.
func DistrictTitleEQ(v string) predicate.History {
	return predicate.History(sql.FieldEQ(FieldDistrictTitle, v))
}

// DistrictTitleNEQ applies the NEQ predicate on the "district_title" field.
func DistrictTitleNEQ(v string) predicate.History {
	return predicate.History(sql.FieldNEQ(FieldDistrictTitle, v))
}

// DistrictTitleIn applies the In predicate on the "district_title" field.
func DistrictTitleIn(vs ...string) predicate.History {
	return predicate.History(sql.FieldIn(FieldDistrictTitle, vs...))
}

// DistrictTitleNotIn applies the NotIn predicate on the "district_title" field.
func DistrictTitleNotIn(vs ...string) predicate.History {
	return predicate.History(sql.FieldNotIn(FieldDistrictTitle, vs...))
}

// DistrictTitleGT applies the GT predicate on the "district_title" field.
func DistrictTitleGT(v string) predicate.History {
	return predicate.History(sql.FieldGT(FieldDistrictTitle, v))
}

// DistrictTitleGTE applies the GTE predicate on the "district_title" field.
func DistrictTitleGTE(v string) predicate.History {
	return predicate.History(sql.FieldGTE(FieldDistrictTitle, v))
}

// DistrictTitleLT applies the LT predicate on the "district_title" field.
func DistrictTitleLT(v string) predicate.History {
	return predicate.History(sql.FieldLT(FieldDistrictTitle, v))
}

// DistrictTitleLTE applies the LTE predicate on the "district_title" field.
func DistrictTitleLTE(v string) predicate.History {
	return predicate.History(sql.FieldLTE(FieldDistrictTitle, v))
}

// DistrictTitleContains applies the Contains predicate on the "district_title" field.
func DistrictTitleContains(v string) predicate.History {
	return predicate.History(sql.FieldContains(FieldDistrictTitle, v))
}

// DistrictTitleHasPrefix applies the HasPrefix predicate on the "district_title" field.
func DistrictTitleHasPrefix(v string) predicate.History {
	return predicate.History(sql.FieldHasPrefix(FieldDistrictTitle, v))
}

// DistrictTitleHasSuffix applies the HasSuffix predicate on the "district_title" field.
func DistrictTitleHasSuffix(v string) predicate.History {
	return predicate.History(sql.FieldHasSuffix(FieldDistrictTitle, v))
}

// DistrictTitleEqualFold applies the EqualFold predicate on the "district_title" field.
func DistrictTitleEqualFold(v string) predicate.History {
	return predicate.History(sql.FieldEqualFold(FieldDistrictTitle, v))
}

// DistrictTitleContainsFold applies the ContainsFold predicate on the "district_title" field.
func DistrictTitleContainsFold(v string) predicate.History {
	return predicate.History(sql.FieldContainsFold(FieldDistrictTitle, v))
}

// LandAreaEQ applies the EQ predicate on the "land_area" field.
func LandAreaEQ(v float64) predicate.History {
	return predicate.History(sql.FieldEQ(FieldLandArea, v))
}

// LandAreaNEQ applies the NEQ predicate on the "land_area" field.
func LandAreaNEQ(v float64) predicate.History {
	return predicate.History(sql.FieldNEQ(FieldLandArea, v))
}

// LandAreaIn applies the In predicate on the "land_area" field.
func LandAreaIn(vs ...float64) predicate.History {
	return predicate.History(sql.FieldIn(FieldLandArea, vs...))
}

// LandAreaNotIn applies the NotIn predicate on the "land_area" field.
func LandAreaNotIn(vs ...float64) predicate.History {
	return predicate.History(sql.FieldNotIn(FieldLandArea, vs...))
}

// LandAreaGT applies the GT predicate on the "land_area" field.
func LandAreaGT(v float64) predicate.History {
	return predicate.History(sql.FieldGT(FieldLandArea, v))
}

// LandAreaGTE applies the GTE predicate on the "land_area" field.
func LandAreaGTE(v float64) predicate.History {
	return predicate.History(sql.FieldGTE(FieldLandArea, v))
}

// LandAreaLT applies the LT predicate on the "land_area" field.
func LandAreaLT(v float64) predicate.History {
	return predicate.History(sql.FieldLT(FieldLandArea, v))
}

// LandAreaLTE applies the LTE predicate on the "land_area" field.
func LandAreaLTE(v float64) predicate.History {
	return predicate.History(sql.FieldLTE(FieldLandArea, v))
}

// IsBuyEQ applies the EQ predicate on the "is_buy" field.
func IsBuyEQ(v bool) predicate.History {
	return predicate.History(sql.FieldEQ(FieldIsBuy, v))
}

// IsBuyNEQ applies the NEQ predicate on the "is_buy" field.
func IsBuyNEQ(v bool) predicate.History {
	return predicate.History(sql.FieldNEQ(FieldIsBuy, v))
}

// ConstructionFacilitiesAreaEQ applies the EQ predicate on the "construction_facilities_area" field.
func ConstructionFacilitiesAreaEQ(v float64) predicate.History {
	return predicate.History(sql.FieldEQ(FieldConstructionFacilitiesArea, v))
}

// ConstructionFacilitiesAreaNEQ applies the NEQ predicate on the "construction_facilities_area" field.
func ConstructionFacilitiesAreaNEQ(v float64) predicate.History {
	return predicate.History(sql.FieldNEQ(FieldConstructionFacilitiesArea, v))
}

// ConstructionFacilitiesAreaIn applies the In predicate on the "construction_facilities_area" field.
func ConstructionFacilitiesAreaIn(vs ...float64) predicate.History {
	return predicate.History(sql.FieldIn(FieldConstructionFacilitiesArea, vs...))
}

// ConstructionFacilitiesAreaNotIn applies the NotIn predicate on the "construction_facilities_area" field.
func ConstructionFacilitiesAreaNotIn(vs ...float64) predicate.History {
	return predicate.History(sql.FieldNotIn(FieldConstructionFacilitiesArea, vs...))
}

// ConstructionFacilitiesAreaGT applies the GT predicate on the "construction_facilities_area" field.
func ConstructionFacilitiesAreaGT(v float64) predicate.History {
	return predicate.History(sql.FieldGT(FieldConstructionFacilitiesArea, v))
}

// ConstructionFacilitiesAreaGTE applies the GTE predicate on the "construction_facilities_area" field.
func ConstructionFacilitiesAreaGTE(v float64) predicate.History {
	return predicate.History(sql.FieldGTE(FieldConstructionFacilitiesArea, v))
}

// ConstructionFacilitiesAreaLT applies the LT predicate on the "construction_facilities_area" field.
func ConstructionFacilitiesAreaLT(v float64) predicate.History {
	return predicate.History(sql.FieldLT(FieldConstructionFacilitiesArea, v))
}

// ConstructionFacilitiesAreaLTE applies the LTE predicate on the "construction_facilities_area" field.
func ConstructionFacilitiesAreaLTE(v float64) predicate.History {
	return predicate.History(sql.FieldLTE(FieldConstructionFacilitiesArea, v))
}

// BuildingTypeEQ applies the EQ predicate on the "building_type" field.
func BuildingTypeEQ(v string) predicate.History {
	return predicate.History(sql.FieldEQ(FieldBuildingType, v))
}

// BuildingTypeNEQ applies the NEQ predicate on the "building_type" field.
func BuildingTypeNEQ(v string) predicate.History {
	return predicate.History(sql.FieldNEQ(FieldBuildingType, v))
}

// BuildingTypeIn applies the In predicate on the "building_type" field.
func BuildingTypeIn(vs ...string) predicate.History {
	return predicate.History(sql.FieldIn(FieldBuildingType, vs...))
}

// BuildingTypeNotIn applies the NotIn predicate on the "building_type" field.
func BuildingTypeNotIn(vs ...string) predicate.History {
	return predicate.History(sql.FieldNotIn(FieldBuildingType, vs...))
}

// BuildingTypeGT applies the GT predicate on the "building_type" field.
func BuildingTypeGT(v string) predicate.History {
	return predicate.History(sql.FieldGT(FieldBuildingType, v))
}

// BuildingTypeGTE applies the GTE predicate on the "building_type" field.
func BuildingTypeGTE(v string) predicate.History {
	return predicate.History(sql.FieldGTE(FieldBuildingType, v))
}

// BuildingTypeLT applies the LT predicate on the "building_type" field.
func BuildingTypeLT(v string) predicate.History {
	return predicate.History(sql.FieldLT(FieldBuildingType, v))
}

// BuildingTypeLTE applies the LTE predicate on the "building_type" field.
func BuildingTypeLTE(v string) predicate.History {
	return predicate.History(sql.FieldLTE(FieldBuildingType, v))
}

// BuildingTypeContains applies the Contains predicate on the "building_type" field.
func BuildingTypeContains(v string) predicate.History {
	return predicate.History(sql.FieldContains(FieldBuildingType, v))
}

// BuildingTypeHasPrefix applies the HasPrefix predicate on the "building_type" field.
func BuildingTypeHasPrefix(v string) predicate.History {
	return predicate.History(sql.FieldHasPrefix(FieldBuildingType, v))
}

// BuildingTypeHasSuffix applies the HasSuffix predicate on the "building_type" field.
func BuildingTypeHasSuffix(v string) predicate.History {
	return predicate.History(sql.FieldHasSuffix(FieldBuildingType, v))
}

// BuildingTypeEqualFold applies the EqualFold predicate on the "building_type" field.
func BuildingTypeEqualFold(v string) predicate.History {
	return predicate.History(sql.FieldEqualFold(FieldBuildingType, v))
}

// BuildingTypeContainsFold applies the ContainsFold predicate on the "building_type" field.
func BuildingTypeContainsFold(v string) predicate.History {
	return predicate.History(sql.FieldContainsFold(FieldBuildingType, v))
}

// AccountingSupportEQ applies the EQ predicate on the "accounting_support" field.
func AccountingSupportEQ(v bool) predicate.History {
	return predicate.History(sql.FieldEQ(FieldAccountingSupport, v))
}

// AccountingSupportNEQ applies the NEQ predicate on the "accounting_support" field.
func AccountingSupportNEQ(v bool) predicate.History {
	return predicate.History(sql.FieldNEQ(FieldAccountingSupport, v))
}

// TaxationSystemOperationsEQ applies the EQ predicate on the "taxation_system_operations" field.
func TaxationSystemOperationsEQ(v int) predicate.History {
	return predicate.History(sql.FieldEQ(FieldTaxationSystemOperations, v))
}

// TaxationSystemOperationsNEQ applies the NEQ predicate on the "taxation_system_operations" field.
func TaxationSystemOperationsNEQ(v int) predicate.History {
	return predicate.History(sql.FieldNEQ(FieldTaxationSystemOperations, v))
}

// TaxationSystemOperationsIn applies the In predicate on the "taxation_system_operations" field.
func TaxationSystemOperationsIn(vs ...int) predicate.History {
	return predicate.History(sql.FieldIn(FieldTaxationSystemOperations, vs...))
}

// TaxationSystemOperationsNotIn applies the NotIn predicate on the "taxation_system_operations" field.
func TaxationSystemOperationsNotIn(vs ...int) predicate.History {
	return predicate.History(sql.FieldNotIn(FieldTaxationSystemOperations, vs...))
}

// TaxationSystemOperationsIsNil applies the IsNil predicate on the "taxation_system_operations" field.
func TaxationSystemOperationsIsNil() predicate.History {
	return predicate.History(sql.FieldIsNull(FieldTaxationSystemOperations))
}

// TaxationSystemOperationsNotNil applies the NotNil predicate on the "taxation_system_operations" field.
func TaxationSystemOperationsNotNil() predicate.History {
	return predicate.History(sql.FieldNotNull(FieldTaxationSystemOperations))
}

// OperationTypeEQ applies the EQ predicate on the "operation_type" field.
func OperationTypeEQ(v string) predicate.History {
	return predicate.History(sql.FieldEQ(FieldOperationType, v))
}

// OperationTypeNEQ applies the NEQ predicate on the "operation_type" field.
func OperationTypeNEQ(v string) predicate.History {
	return predicate.History(sql.FieldNEQ(FieldOperationType, v))
}

// OperationTypeIn applies the In predicate on the "operation_type" field.
func OperationTypeIn(vs ...string) predicate.History {
	return predicate.History(sql.FieldIn(FieldOperationType, vs...))
}

// OperationTypeNotIn applies the NotIn predicate on the "operation_type" field.
func OperationTypeNotIn(vs ...string) predicate.History {
	return predicate.History(sql.FieldNotIn(FieldOperationType, vs...))
}

// OperationTypeGT applies the GT predicate on the "operation_type" field.
func OperationTypeGT(v string) predicate.History {
	return predicate.History(sql.FieldGT(FieldOperationType, v))
}

// OperationTypeGTE applies the GTE predicate on the "operation_type" field.
func OperationTypeGTE(v string) predicate.History {
	return predicate.History(sql.FieldGTE(FieldOperationType, v))
}

// OperationTypeLT applies the LT predicate on the "operation_type" field.
func OperationTypeLT(v string) predicate.History {
	return predicate.History(sql.FieldLT(FieldOperationType, v))
}

// OperationTypeLTE applies the LTE predicate on the "operation_type" field.
func OperationTypeLTE(v string) predicate.History {
	return predicate.History(sql.FieldLTE(FieldOperationType, v))
}

// OperationTypeContains applies the Contains predicate on the "operation_type" field.
func OperationTypeContains(v string) predicate.History {
	return predicate.History(sql.FieldContains(FieldOperationType, v))
}

// OperationTypeHasPrefix applies the HasPrefix predicate on the "operation_type" field.
func OperationTypeHasPrefix(v string) predicate.History {
	return predicate.History(sql.FieldHasPrefix(FieldOperationType, v))
}

// OperationTypeHasSuffix applies the HasSuffix predicate on the "operation_type" field.
func OperationTypeHasSuffix(v string) predicate.History {
	return predicate.History(sql.FieldHasSuffix(FieldOperationType, v))
}

// OperationTypeIsNil applies the IsNil predicate on the "operation_type" field.
func OperationTypeIsNil() predicate.History {
	return predicate.History(sql.FieldIsNull(FieldOperationType))
}

// OperationTypeNotNil applies the NotNil predicate on the "operation_type" field.
func OperationTypeNotNil() predicate.History {
	return predicate.History(sql.FieldNotNull(FieldOperationType))
}

// OperationTypeEqualFold applies the EqualFold predicate on the "operation_type" field.
func OperationTypeEqualFold(v string) predicate.History {
	return predicate.History(sql.FieldEqualFold(FieldOperationType, v))
}

// OperationTypeContainsFold applies the ContainsFold predicate on the "operation_type" field.
func OperationTypeContainsFold(v string) predicate.History {
	return predicate.History(sql.FieldContainsFold(FieldOperationType, v))
}

// PatentCalcEQ applies the EQ predicate on the "patent_calc" field.
func PatentCalcEQ(v bool) predicate.History {
	return predicate.History(sql.FieldEQ(FieldPatentCalc, v))
}

// PatentCalcNEQ applies the NEQ predicate on the "patent_calc" field.
func PatentCalcNEQ(v bool) predicate.History {
	return predicate.History(sql.FieldNEQ(FieldPatentCalc, v))
}

// BusinessActivityIDEQ applies the EQ predicate on the "business_activity_id" field.
func BusinessActivityIDEQ(v int) predicate.History {
	return predicate.History(sql.FieldEQ(FieldBusinessActivityID, v))
}

// BusinessActivityIDNEQ applies the NEQ predicate on the "business_activity_id" field.
func BusinessActivityIDNEQ(v int) predicate.History {
	return predicate.History(sql.FieldNEQ(FieldBusinessActivityID, v))
}

// BusinessActivityIDIn applies the In predicate on the "business_activity_id" field.
func BusinessActivityIDIn(vs ...int) predicate.History {
	return predicate.History(sql.FieldIn(FieldBusinessActivityID, vs...))
}

// BusinessActivityIDNotIn applies the NotIn predicate on the "business_activity_id" field.
func BusinessActivityIDNotIn(vs ...int) predicate.History {
	return predicate.History(sql.FieldNotIn(FieldBusinessActivityID, vs...))
}

// BusinessActivityIDIsNil applies the IsNil predicate on the "business_activity_id" field.
func BusinessActivityIDIsNil() predicate.History {
	return predicate.History(sql.FieldIsNull(FieldBusinessActivityID))
}

// BusinessActivityIDNotNil applies the NotNil predicate on the "business_activity_id" field.
func BusinessActivityIDNotNil() predicate.History {
	return predicate.History(sql.FieldNotNull(FieldBusinessActivityID))
}

// OtherEQ applies the EQ predicate on the "other" field.
func OtherEQ(v string) predicate.History {
	return predicate.History(sql.FieldEQ(FieldOther, v))
}

// OtherNEQ applies the NEQ predicate on the "other" field.
func OtherNEQ(v string) predicate.History {
	return predicate.History(sql.FieldNEQ(FieldOther, v))
}

// OtherIn applies the In predicate on the "other" field.
func OtherIn(vs ...string) predicate.History {
	return predicate.History(sql.FieldIn(FieldOther, vs...))
}

// OtherNotIn applies the NotIn predicate on the "other" field.
func OtherNotIn(vs ...string) predicate.History {
	return predicate.History(sql.FieldNotIn(FieldOther, vs...))
}

// OtherGT applies the GT predicate on the "other" field.
func OtherGT(v string) predicate.History {
	return predicate.History(sql.FieldGT(FieldOther, v))
}

// OtherGTE applies the GTE predicate on the "other" field.
func OtherGTE(v string) predicate.History {
	return predicate.History(sql.FieldGTE(FieldOther, v))
}

// OtherLT applies the LT predicate on the "other" field.
func OtherLT(v string) predicate.History {
	return predicate.History(sql.FieldLT(FieldOther, v))
}

// OtherLTE applies the LTE predicate on the "other" field.
func OtherLTE(v string) predicate.History {
	return predicate.History(sql.FieldLTE(FieldOther, v))
}

// OtherContains applies the Contains predicate on the "other" field.
func OtherContains(v string) predicate.History {
	return predicate.History(sql.FieldContains(FieldOther, v))
}

// OtherHasPrefix applies the HasPrefix predicate on the "other" field.
func OtherHasPrefix(v string) predicate.History {
	return predicate.History(sql.FieldHasPrefix(FieldOther, v))
}

// OtherHasSuffix applies the HasSuffix predicate on the "other" field.
func OtherHasSuffix(v string) predicate.History {
	return predicate.History(sql.FieldHasSuffix(FieldOther, v))
}

// OtherIsNil applies the IsNil predicate on the "other" field.
func OtherIsNil() predicate.History {
	return predicate.History(sql.FieldIsNull(FieldOther))
}

// OtherNotNil applies the NotNil predicate on the "other" field.
func OtherNotNil() predicate.History {
	return predicate.History(sql.FieldNotNull(FieldOther))
}

// OtherEqualFold applies the EqualFold predicate on the "other" field.
func OtherEqualFold(v string) predicate.History {
	return predicate.History(sql.FieldEqualFold(FieldOther, v))
}

// OtherContainsFold applies the ContainsFold predicate on the "other" field.
func OtherContainsFold(v string) predicate.History {
	return predicate.History(sql.FieldContainsFold(FieldOther, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.History {
	return predicate.History(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.History {
	return predicate.History(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.History {
	return predicate.History(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.History {
	return predicate.History(sql.FieldNotIn(FieldUserID, vs...))
}

// HasIndustry applies the HasEdge predicate on the "industry" edge.
func HasIndustry() predicate.History {
	return predicate.History(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, IndustryTable, IndustryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasIndustryWith applies the HasEdge predicate on the "industry" edge with a given conditions (other predicates).
func HasIndustryWith(preds ...predicate.Industry) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		step := newIndustryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTaxationSystems applies the HasEdge predicate on the "taxation_systems" edge.
func HasTaxationSystems() predicate.History {
	return predicate.History(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TaxationSystemsTable, TaxationSystemsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTaxationSystemsWith applies the HasEdge predicate on the "taxation_systems" edge with a given conditions (other predicates).
func HasTaxationSystemsWith(preds ...predicate.TaxationSystem) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		step := newTaxationSystemsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBusinessActivity applies the HasEdge predicate on the "business_activity" edge.
func HasBusinessActivity() predicate.History {
	return predicate.History(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BusinessActivityTable, BusinessActivityColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessActivityWith applies the HasEdge predicate on the "business_activity" edge with a given conditions (other predicates).
func HasBusinessActivityWith(preds ...predicate.BusinessActivity) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		step := newBusinessActivityStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDistrict applies the HasEdge predicate on the "district" edge.
func HasDistrict() predicate.History {
	return predicate.History(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DistrictTable, DistrictColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDistrictWith applies the HasEdge predicate on the "district" edge with a given conditions (other predicates).
func HasDistrictWith(preds ...predicate.District) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		step := newDistrictStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.History {
	return predicate.History(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.History) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.History) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.History) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		p(s.Not())
	})
}
