package podtypes

// Scaleway
//
// Configuration for interacting with Scaleway
// +kubebuilder:object:generate=true
type Scaleway struct {
	// Configuration for authenticating a Pod with Scaleway
	//
	//+kubebuilder:validation:Optional
	Auth *ScalewayAuth `json:"auth,omitempty"`
}

// ScalewayAuth
//
// Configuration for authenticating a Pod with Scaleway
type ScalewayAuth struct {
	// Name of the Kubernetes Secret (in the application's namespace) containing the
	// Scaleway API access key and secret key.
	//
	//+kubebuilder:validation:Required
	SecretName string `json:"secretName"`
}
