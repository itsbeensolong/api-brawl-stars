package utils

type Username struct {
	User     string
	Password string
	Database string
}

type Super struct {
	Name        string
	Description string
	Damage      string
	Scope       string
}
type Brawler struct {
	Name     string
	Image    string
	Type     string
	Category string
	Stellar  [2]string
	Gadget   [2]string
	Super
}
