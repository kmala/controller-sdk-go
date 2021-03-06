package api

// AppSettings is the structure of an app's settings.
type AppSettings struct {
	// Owner is the app owner. It cannot be updated with AppSettings.Set(). See app.Transfer().
	Owner string `json:"owner,omitempty"`
	// App is the app name. It cannot be updated at all right now.
	App string `json:"app,omitempty"`
	// Created is the time that the application settings was created and cannot be updated.
	Created string `json:"created,omitempty"`
	// Updated is the last time the application settings was changed and cannot be updated.
	Updated string `json:"updated,omitempty"`
	// UUID is a unique string reflecting the application settings in its current state.
	// It changes every time the application settings is changed and cannot be updated.
	UUID string `json:"uuid,omitempty"`
	// Maintenance determines if the application is taken down for maintenance or not.
	Maintenance *bool `json:"maintenance,omitempty"`
	// Routable determines if the application should be exposed by the router.
	Routable  *bool    `json:"routable,omitempty"`
	Whitelist []string `json:"whitelist,omitempty"`
}

// NewRoutable returns a default value for the AppSettings.Routable field.
func NewRoutable() *bool {
	b := true
	return &b
}

// Whitelist is the structure of POST /v2/app/<app id>/whitelist/.
type Whitelist struct {
	Addresses []string `json:"addresses"`
}
