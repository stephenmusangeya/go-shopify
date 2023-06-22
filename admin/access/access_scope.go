package access

type AccessScopeResponse struct {
	AccessScopes []AccessScope `json:"access_scopes"`
}

type AccessScope struct {
	Scope APIScope `json:"handle"`
}
