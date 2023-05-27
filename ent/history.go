// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/while-act/hackathon-backend/ent/district"
	"github.com/while-act/hackathon-backend/ent/equipment"
	"github.com/while-act/hackathon-backend/ent/history"
	"github.com/while-act/hackathon-backend/ent/industry"
	"github.com/while-act/hackathon-backend/ent/user"
)

// History is the model entity for the History schema.
type History struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CompanyName holds the value of the "company_name" field.
	CompanyName string `json:"company_name,omitempty"`
	// IndustryBranch holds the value of the "industry_branch" field.
	IndustryBranch string `json:"industry_branch,omitempty"`
	// FullTimeEmployees holds the value of the "full_time_employees" field.
	FullTimeEmployees int `json:"full_time_employees,omitempty"`
	// DistrictTitle holds the value of the "district_title" field.
	DistrictTitle string `json:"district_title,omitempty"`
	// LandArea holds the value of the "land_area" field.
	LandArea float64 `json:"land_area,omitempty"`
	// ConstructionFacilitiesArea holds the value of the "construction_facilities_area" field.
	ConstructionFacilitiesArea float64 `json:"construction_facilities_area,omitempty"`
	// EquipmentType holds the value of the "equipment_type" field.
	EquipmentType string `json:"equipment_type,omitempty"`
	// OrganizationType holds the value of the "organization_type" field.
	OrganizationType string `json:"organization_type,omitempty"`
	// FacilityType holds the value of the "facility_type" field.
	FacilityType string `json:"facility_type,omitempty"`
	// AccountingServices holds the value of the "accounting_services" field.
	AccountingServices bool `json:"accounting_services,omitempty"`
	// Patent holds the value of the "patent" field.
	Patent bool `json:"patent,omitempty"`
	// Other holds the value of the "other" field.
	Other string `json:"other,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HistoryQuery when eager-loading is set.
	Edges        HistoryEdges `json:"edges"`
	selectValues sql.SelectValues
}

// HistoryEdges holds the relations/edges for other nodes in the graph.
type HistoryEdges struct {
	// Industry holds the value of the industry edge.
	Industry *Industry `json:"industry,omitempty"`
	// District holds the value of the district edge.
	District *District `json:"district,omitempty"`
	// Equipment holds the value of the equipment edge.
	Equipment *Equipment `json:"equipment,omitempty"`
	// Users holds the value of the users edge.
	Users *User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// IndustryOrErr returns the Industry value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HistoryEdges) IndustryOrErr() (*Industry, error) {
	if e.loadedTypes[0] {
		if e.Industry == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: industry.Label}
		}
		return e.Industry, nil
	}
	return nil, &NotLoadedError{edge: "industry"}
}

// DistrictOrErr returns the District value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HistoryEdges) DistrictOrErr() (*District, error) {
	if e.loadedTypes[1] {
		if e.District == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: district.Label}
		}
		return e.District, nil
	}
	return nil, &NotLoadedError{edge: "district"}
}

// EquipmentOrErr returns the Equipment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HistoryEdges) EquipmentOrErr() (*Equipment, error) {
	if e.loadedTypes[2] {
		if e.Equipment == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: equipment.Label}
		}
		return e.Equipment, nil
	}
	return nil, &NotLoadedError{edge: "equipment"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HistoryEdges) UsersOrErr() (*User, error) {
	if e.loadedTypes[3] {
		if e.Users == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*History) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case history.FieldAccountingServices, history.FieldPatent:
			values[i] = new(sql.NullBool)
		case history.FieldLandArea, history.FieldConstructionFacilitiesArea:
			values[i] = new(sql.NullFloat64)
		case history.FieldID, history.FieldFullTimeEmployees, history.FieldUserID:
			values[i] = new(sql.NullInt64)
		case history.FieldCompanyName, history.FieldIndustryBranch, history.FieldDistrictTitle, history.FieldEquipmentType, history.FieldOrganizationType, history.FieldFacilityType, history.FieldOther:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the History fields.
func (h *History) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case history.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			h.ID = int(value.Int64)
		case history.FieldCompanyName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field company_name", values[i])
			} else if value.Valid {
				h.CompanyName = value.String
			}
		case history.FieldIndustryBranch:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field industry_branch", values[i])
			} else if value.Valid {
				h.IndustryBranch = value.String
			}
		case history.FieldFullTimeEmployees:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field full_time_employees", values[i])
			} else if value.Valid {
				h.FullTimeEmployees = int(value.Int64)
			}
		case history.FieldDistrictTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field district_title", values[i])
			} else if value.Valid {
				h.DistrictTitle = value.String
			}
		case history.FieldLandArea:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field land_area", values[i])
			} else if value.Valid {
				h.LandArea = value.Float64
			}
		case history.FieldConstructionFacilitiesArea:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field construction_facilities_area", values[i])
			} else if value.Valid {
				h.ConstructionFacilitiesArea = value.Float64
			}
		case history.FieldEquipmentType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field equipment_type", values[i])
			} else if value.Valid {
				h.EquipmentType = value.String
			}
		case history.FieldOrganizationType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field organization_type", values[i])
			} else if value.Valid {
				h.OrganizationType = value.String
			}
		case history.FieldFacilityType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field facility_type", values[i])
			} else if value.Valid {
				h.FacilityType = value.String
			}
		case history.FieldAccountingServices:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field accounting_services", values[i])
			} else if value.Valid {
				h.AccountingServices = value.Bool
			}
		case history.FieldPatent:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field patent", values[i])
			} else if value.Valid {
				h.Patent = value.Bool
			}
		case history.FieldOther:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field other", values[i])
			} else if value.Valid {
				h.Other = value.String
			}
		case history.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				h.UserID = int(value.Int64)
			}
		default:
			h.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the History.
// This includes values selected through modifiers, order, etc.
func (h *History) Value(name string) (ent.Value, error) {
	return h.selectValues.Get(name)
}

// QueryIndustry queries the "industry" edge of the History entity.
func (h *History) QueryIndustry() *IndustryQuery {
	return NewHistoryClient(h.config).QueryIndustry(h)
}

// QueryDistrict queries the "district" edge of the History entity.
func (h *History) QueryDistrict() *DistrictQuery {
	return NewHistoryClient(h.config).QueryDistrict(h)
}

// QueryEquipment queries the "equipment" edge of the History entity.
func (h *History) QueryEquipment() *EquipmentQuery {
	return NewHistoryClient(h.config).QueryEquipment(h)
}

// QueryUsers queries the "users" edge of the History entity.
func (h *History) QueryUsers() *UserQuery {
	return NewHistoryClient(h.config).QueryUsers(h)
}

// Update returns a builder for updating this History.
// Note that you need to call History.Unwrap() before calling this method if this History
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *History) Update() *HistoryUpdateOne {
	return NewHistoryClient(h.config).UpdateOne(h)
}

// Unwrap unwraps the History entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (h *History) Unwrap() *History {
	_tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("ent: History is not a transactional entity")
	}
	h.config.driver = _tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *History) String() string {
	var builder strings.Builder
	builder.WriteString("History(")
	builder.WriteString(fmt.Sprintf("id=%v, ", h.ID))
	builder.WriteString("company_name=")
	builder.WriteString(h.CompanyName)
	builder.WriteString(", ")
	builder.WriteString("industry_branch=")
	builder.WriteString(h.IndustryBranch)
	builder.WriteString(", ")
	builder.WriteString("full_time_employees=")
	builder.WriteString(fmt.Sprintf("%v", h.FullTimeEmployees))
	builder.WriteString(", ")
	builder.WriteString("district_title=")
	builder.WriteString(h.DistrictTitle)
	builder.WriteString(", ")
	builder.WriteString("land_area=")
	builder.WriteString(fmt.Sprintf("%v", h.LandArea))
	builder.WriteString(", ")
	builder.WriteString("construction_facilities_area=")
	builder.WriteString(fmt.Sprintf("%v", h.ConstructionFacilitiesArea))
	builder.WriteString(", ")
	builder.WriteString("equipment_type=")
	builder.WriteString(h.EquipmentType)
	builder.WriteString(", ")
	builder.WriteString("organization_type=")
	builder.WriteString(h.OrganizationType)
	builder.WriteString(", ")
	builder.WriteString("facility_type=")
	builder.WriteString(h.FacilityType)
	builder.WriteString(", ")
	builder.WriteString("accounting_services=")
	builder.WriteString(fmt.Sprintf("%v", h.AccountingServices))
	builder.WriteString(", ")
	builder.WriteString("patent=")
	builder.WriteString(fmt.Sprintf("%v", h.Patent))
	builder.WriteString(", ")
	builder.WriteString("other=")
	builder.WriteString(h.Other)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", h.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// Histories is a parsable slice of History.
type Histories []*History
