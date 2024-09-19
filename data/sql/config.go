package sql

type Config struct {
	ConnectionString    string `env:"SQL_CONNECTION_STRING,unset"`
	IsLoggingStatements bool   `env:"SQL_ENABLE_LOGGING" envDefault:"false"`
}

type ConfigTransactionFactory struct {
	IsolationLevel int  `env:"SQL_ISOLATION_LEVEL" envDefault:"0"`
	ReadOnly       bool `env:"SQL_READ_ONLY" envDefault:"false"`
}
