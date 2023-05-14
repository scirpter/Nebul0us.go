package sigs

const (
	// might change after an update, but it's unlikely.
	// if incorrect, the client's IP address will be invalidated.
	KEEP_ALIVE_SIG = 0x5f71382d
	// necessary to calculate game object data relative to
	// the map size.
	MAP_SIZE_SIG = 844.1067172342606
	// app version
	APP_VERSION_SIG = 1217
)
