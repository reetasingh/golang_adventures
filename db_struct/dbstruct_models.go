package dbstruct

type DynamicStruct struct {
	Field1 int    `db:"field1,omitempty"`
	Field2 string `db:"field2,omitempty"`
}
