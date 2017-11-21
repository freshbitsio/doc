package data

// Configuration for a plugin that resolves resources to content and metadata
type Resolver struct {

}

// User preference file, which is stored in ~/.doc/preferences.json
type UserPreferences struct {
	Fullname string `json:"fullname"'`
	Location string `json:"location"'`
	Email string `json:"email"'`
	Organization string `json:"organization"'`
	Resolvers []string `json:"resolvers"'`
	Credentials map[string]string `json:"credentials"'`
}
