package amjwt

var (
	// Version defines the version of amjwt
	Version string
)

func init() {
	// If version, commit, or build time are not set, make that clear.
	const unknown = "unknown"
	if Version == "" {
		Version = unknown
	}
}
