package requests

// Request — команда от клиента.
type Request interface {
	Type() RequestType
}
