package utils

import "regexp"

// IsValidUUID validate v4 uuid
func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

// WrapJSONResponse wrap json response with Data key
func WrapJSONResponse(data interface{}) map[string]interface{} {
	response := make(map[string]interface{})
	response["Data"] = data
	return response
}
