package manager_request

type ManagerRequest struct {
		"email": "name@name.com", // should in email format
		"name": "", // string | minLength 4 | maxLength 52
		"userImageUri": "", // string | should be an URI
		"companyName": "", // string | minLength 4 | maxLength 52
		"companyImageUri": "", // string | should be an URI
}