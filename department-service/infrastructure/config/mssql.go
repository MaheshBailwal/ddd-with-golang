package config

//server=azuresql-mb.database.windows.net;user id=mahesh;password=password@123;port=1433;database=devServer_28.04-patched;"

type SMSSql struct {
	Server   string `validate:"required"`
	Port     string `validate:"required"`
	User     string `validate:"required"`
	Password string `validate:"required"`
	Database string `validate:"required"`
}
