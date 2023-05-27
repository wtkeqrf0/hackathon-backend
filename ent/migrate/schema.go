// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CompaniesColumns holds the columns for the "companies" table.
	CompaniesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "inn", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Nullable: true, Size: 150},
		{Name: "website", Type: field.TypeString, Nullable: true},
	}
	// CompaniesTable holds the schema information for the "companies" table.
	CompaniesTable = &schema.Table{
		Name:       "companies",
		Columns:    CompaniesColumns,
		PrimaryKey: []*schema.Column{CompaniesColumns[0]},
	}
	// DistrictsColumns holds the columns for the "districts" table.
	DistrictsColumns = []*schema.Column{
		{Name: "title", Type: field.TypeString, Unique: true},
		{Name: "avg_cadastral_val", Type: field.TypeFloat64},
	}
	// DistrictsTable holds the schema information for the "districts" table.
	DistrictsTable = &schema.Table{
		Name:       "districts",
		Columns:    DistrictsColumns,
		PrimaryKey: []*schema.Column{DistrictsColumns[0]},
	}
	// EntrepreneurshipsColumns holds the columns for the "entrepreneurships" table.
	EntrepreneurshipsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeString, Unique: true},
	}
	// EntrepreneurshipsTable holds the schema information for the "entrepreneurships" table.
	EntrepreneurshipsTable = &schema.Table{
		Name:       "entrepreneurships",
		Columns:    EntrepreneurshipsColumns,
		PrimaryKey: []*schema.Column{EntrepreneurshipsColumns[0]},
	}
	// EquipmentColumns holds the columns for the "equipment" table.
	EquipmentColumns = []*schema.Column{
		{Name: "type", Type: field.TypeString, Unique: true},
		{Name: "avg_price_dol", Type: field.TypeFloat64},
		{Name: "avg_price_rub", Type: field.TypeFloat64},
	}
	// EquipmentTable holds the schema information for the "equipment" table.
	EquipmentTable = &schema.Table{
		Name:       "equipment",
		Columns:    EquipmentColumns,
		PrimaryKey: []*schema.Column{EquipmentColumns[0]},
	}
	// HistoriesColumns holds the columns for the "histories" table.
	HistoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "company_name", Type: field.TypeString, Size: 150},
		{Name: "full_time_employees", Type: field.TypeInt},
		{Name: "land_area", Type: field.TypeFloat64},
		{Name: "construction_facilities_area", Type: field.TypeFloat64},
		{Name: "organization_type", Type: field.TypeString},
		{Name: "facility_type", Type: field.TypeString},
		{Name: "accounting_services", Type: field.TypeBool},
		{Name: "patent", Type: field.TypeBool},
		{Name: "other", Type: field.TypeString, Size: 2147483647},
		{Name: "district_title", Type: field.TypeString},
		{Name: "equipment_type", Type: field.TypeString},
		{Name: "industry_branch", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeInt},
	}
	// HistoriesTable holds the schema information for the "histories" table.
	HistoriesTable = &schema.Table{
		Name:       "histories",
		Columns:    HistoriesColumns,
		PrimaryKey: []*schema.Column{HistoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "histories_districts_histories",
				Columns:    []*schema.Column{HistoriesColumns[10]},
				RefColumns: []*schema.Column{DistrictsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "histories_equipment_histories",
				Columns:    []*schema.Column{HistoriesColumns[11]},
				RefColumns: []*schema.Column{EquipmentColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "histories_industries_histories",
				Columns:    []*schema.Column{HistoriesColumns[12]},
				RefColumns: []*schema.Column{IndustriesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "histories_users_histories",
				Columns:    []*schema.Column{HistoriesColumns[13]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// IndustriesColumns holds the columns for the "industries" table.
	IndustriesColumns = []*schema.Column{
		{Name: "branch", Type: field.TypeString, Unique: true},
		{Name: "avg_workers_num", Type: field.TypeFloat64},
		{Name: "avg_workers_num_cad", Type: field.TypeFloat64},
		{Name: "avg_salary", Type: field.TypeFloat64},
		{Name: "avg_salary_cad", Type: field.TypeFloat64},
	}
	// IndustriesTable holds the schema information for the "industries" table.
	IndustriesTable = &schema.Table{
		Name:       "industries",
		Columns:    IndustriesColumns,
		PrimaryKey: []*schema.Column{IndustriesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "role", Type: field.TypeString, Default: "USER"},
		{Name: "name", Type: field.TypeString, Unique: true, Default: schema.Expr("'user' || setval(pg_get_serial_sequence('users','id'),nextval(pg_get_serial_sequence('users','id'))-1)")},
		{Name: "password_hash", Type: field.TypeBytes},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "first_name", Type: field.TypeString, Size: 30},
		{Name: "last_name", Type: field.TypeString, Size: 30},
		{Name: "father_name", Type: field.TypeString, Nullable: true, Size: 30},
		{Name: "position", Type: field.TypeString, Nullable: true, Size: 50},
		{Name: "country", Type: field.TypeString, Nullable: true},
		{Name: "city", Type: field.TypeString, Nullable: true},
		{Name: "biography", Type: field.TypeString, Nullable: true, Size: 1024},
		{Name: "company_id", Type: field.TypeInt, Unique: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_companies_users",
				Columns:    []*schema.Column{UsersColumns[14]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CompaniesTable,
		DistrictsTable,
		EntrepreneurshipsTable,
		EquipmentTable,
		HistoriesTable,
		IndustriesTable,
		UsersTable,
	}
)

func init() {
	HistoriesTable.ForeignKeys[0].RefTable = DistrictsTable
	HistoriesTable.ForeignKeys[1].RefTable = EquipmentTable
	HistoriesTable.ForeignKeys[2].RefTable = IndustriesTable
	HistoriesTable.ForeignKeys[3].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = CompaniesTable
}
