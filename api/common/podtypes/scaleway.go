package podtypes

type scaleway struct {
	Auth *Auth `json:",auth,omitempty"`
}

type ScalewayAuth struct {
	SecretName string `json:"secretName"`
}
