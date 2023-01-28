package GeneratePassword

type Hash struct {
	UserText      string
	ArgonText     string
	CustomSetting CustomOptions
}
type CustomOptions struct {
	Salt       bool
	SpecialKey bool
}
