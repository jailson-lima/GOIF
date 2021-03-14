package types

type Properties struct {
	Host           string
	Port           int
	RegisterRoutes func()
}