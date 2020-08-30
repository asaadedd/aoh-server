package status

type Status string

const(
	CONNECTED    Status = "CONNECTED"
	CONNECTING          = "CONNECTING"
	DISCONNECTED        = "DISCONNECTED"
)
