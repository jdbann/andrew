// Defaults
GenerateSeeds: bool | *false

// Generate seeds in development and ephemeral environments (review apps)
if #Meta.Environment.Type == "development" || #Meta.Environment.Type == "ephemeral" {
    GenerateSeeds: true
}
