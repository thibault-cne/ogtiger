package ttype

type TypeID int

const (
	Int TypeID = iota
	String
	Record
	AnyRecord
	Array
	NoReturn
)

type RecordField struct {
	Name string
	Type *TigerType
}

type TigerType struct {
	ID TypeID

	// For records
	Fields []*RecordField

	// For arrays
	ElementType *TigerType
}

func (t *TigerType) Equals(other *TigerType) bool {
	if t.ID != other.ID {
		return false
	}

	switch t.ID {
	case Record:
		// If either type is AnyRecord, they are equal
		if t.ID == AnyRecord || other.ID == AnyRecord {
			return true
		}

		if other.ID != Record {
			return false
		}

		// Check that the number of fields is the same
		if len(t.Fields) != len(other.Fields) {
			return false
		}

		// It doesn't matter which order the fields are in
		for i := range t.Fields {
			found := false
			for j := range other.Fields {
				if t.Fields[i].Name == other.Fields[j].Name && !t.Fields[i].Type.Equals(other.Fields[j].Type) {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	case Array:
		return t.ElementType.Equals(other.ElementType)
	}

	return true
}

func NewTigerType(id TypeID) TigerType {
	return TigerType{
		ID: id,
	}
}

func NewRecordType(fields []*RecordField) TigerType {
	return TigerType{
		ID:     Record,
		Fields: fields,
	}
}

func NewArrayType(elementType *TigerType) TigerType {
	return TigerType{
		ID:          Array,
		ElementType: elementType,
	}
}
