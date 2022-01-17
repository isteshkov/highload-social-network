package dto

type Response struct {
	Status      string      `json:"status"`
	ErrorCode   string      `json:"errorCode"`
	Payload     interface{} `json:"payload"`
	Description string      `json:"description"`
}

type ResponsePayloadProfile struct {
}

type ResponsePayloadProfiles struct {
}

type ResponsePayloadFriendship struct {
}

type ResponsePayloadSignIn struct {
}

type ResponsePayloadSignUp struct {
}

type ResponsePayloadSignUpCheck struct {
}

type ResponsePayloadSignOut struct {
}
