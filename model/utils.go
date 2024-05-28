package model

func GetMarrigeStatusBoolToEnum(married bool) int8 {
	if married {
		return MARITAL_STATUS_TRUE
	}
	return MARITAL_STATUS_FALSE
}

func GetMarriedStatusEnumToBool(married int8) bool {
	return married == MARITAL_STATUS_TRUE
}
