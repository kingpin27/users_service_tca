package server

import "github.com/kingpin27/users_service_tca/model"

type InputFields struct {
	Name    *string
	Phone   *uint64
	City    *string
	Height  *uint64
	Married *bool
}

func ParseOptionalStringField(in InputFields, inId *uint32) *model.User {
	name := ""
	if in.Name != nil {
		name = *in.Name
	}
	city := ""
	if in.City != nil {
		city = *in.City
	}
	var phone uint64 = 0
	if in.Phone != nil {
		phone = uint64(*in.Phone)
	}
	var height uint64 = 0
	if in.Height != nil {
		height = uint64(*in.Height)
	}
	var married int8 = model.MARITAL_STATUS_NOT_SET
	if in.Married != nil {
		if *in.Married {
			married = model.MARITAL_STATUS_TRUE
		} else {
			married = model.MARITAL_STATUS_FALSE

		}
	}
	var id uint32 = 0
	if inId != nil {
		id = *inId
	}
	return &model.User{
		Id:      id,
		FName:   name,
		City:    city,
		Phone:   phone,
		Height:  uint8(height), // Convert uint64 to uint8
		Married: int8(married),
	}
}
