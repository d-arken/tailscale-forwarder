package config

type ConnectionMapping struct {
	SourcePort int
	TargetAddr string
	TargetPort int
}

type config struct {
	TSHostname   string `env:"TS_HOSTNAME,required"`
	TSAuthKey    string `env:"TS_AUTHKEY,required"`
	TSControlURL string `env:"TS_CONTROL_URL"`
	TSStateDir   string `env:"TS_STATE_DIR"`
	TSEphemeral  bool   `env:"TS_EPHEMERAL" envDefault:"true"`
	LocalListen  bool   `env:"LOCAL_LISTEN" envDefault:"false"`

	ConnectionMappings []ConnectionMapping
}
