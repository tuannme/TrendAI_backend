package models

type AuthenticationToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthenticationResponse struct {
	User  UserResponse        `json:"user"`
	Token AuthenticationToken `json:"token"`
}

func (u *User) ToAuthenticationResponse(token AuthenticationToken) AuthenticationResponse {
	return AuthenticationResponse{
		User:  u.ToResponse(),
		Token: token,
	}
}
