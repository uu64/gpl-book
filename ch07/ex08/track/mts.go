package track

import (
	"reflect"
)

type MultiTierSort struct {
	Records      []*Track
	PrimaryKey   string
	SecondaryKey string
}

func (mts MultiTierSort) Len() int { return len(mts.Records) }

func (mts MultiTierSort) Swap(i, j int) {
	mts.Records[i], mts.Records[j] = mts.Records[j], mts.Records[i]
}

func (mts MultiTierSort) Less(i, j int) bool {
	ri := reflect.Indirect(reflect.ValueOf(mts.Records[i]))
	rj := reflect.Indirect(reflect.ValueOf(mts.Records[j]))

	if mts.PrimaryKey != "" {
		c1 := compareByFieldName(ri, rj, mts.PrimaryKey)
		if c1 != 0 {
			return c1 == -1
		}
	}

	if mts.SecondaryKey != "" {
		c2 := compareByFieldName(ri, rj, mts.SecondaryKey)
		if c2 != 0 {
			return c2 == -1
		}
	}

	return false
}

func compareByFieldName(ri, rj reflect.Value, key string) int {
	fi := ri.FieldByName(key)
	fj := rj.FieldByName(key)
	switch fi.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if fi.Int() == fj.Int() {
			return 0
		} else if fi.Int() > fj.Int() {
			return 1
		} else {
			return -1
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if fi.Uint() == fj.Uint() {
			return 0
		} else if fi.Uint() > fj.Uint() {
			return 1
		} else {
			return -1
		}
	case reflect.Float32, reflect.Float64:
		if fi.Float() == fj.Float() {
			return 0
		} else if fi.Float() > fj.Float() {
			return 1
		} else {
			return -1
		}
	case reflect.String:
		if fi.String() == fj.String() {
			return 0
		} else if fi.String() > fj.String() {
			return 1
		} else {
			return -1
		}
	default:
		if fi.String() == fj.String() {
			return 0
		} else if fi.String() > fj.String() {
			return 1
		} else {
			return -1
		}
	}
}
