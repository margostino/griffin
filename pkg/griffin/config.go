package griffin

type CommandConfiguration struct {
	Id          string `yaml:"id"`
	Description string `yaml:"description"`
	Args        int    `yaml:"args"`
	Action      string `yaml:"action"`
	Pattern     string `yaml:"pattern"`
}

type CommandsConfiguration struct {
	CommandList []CommandConfiguration `yaml:"commands"`
}