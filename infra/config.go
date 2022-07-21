package infra

import (
	"encoding"
	"log"
	"reflect"

	"github.com/mitchellh/mapstructure"

	"github.com/spf13/viper"
)

//Config The configuration struc has all the configurations needed by the appConfig
type Config struct {
	APP  appConfig  `mapstructure:",squash"`
	GRPC grpcConfig `mapstructure:",squash"`
}

type appConfig struct {
	LogLevel string `mapstructure:"log_level"`
	Name     string `mapstructure:"app_name"`
	Debug    bool   `mapstructure:"debug"`
}

type grpcConfig struct {
	Port string `mapstructure:"grpc_port"`
}

//NewConfig Creates a new config struct with current env var values
func NewConfig() (*Config, error) {
	viper.SetDefault("GRPC_PORT", "80")

	viper.SetDefault("DEBUG", "true")
	viper.SetDefault("LOG_LEVEL", "debug")

	viper.SetDefault("APP_NAME", "event_broker")

	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("config: .env file not found")
	}

	cfg := &Config{}
	if err := viper.Unmarshal(cfg, viper.DecodeHook(
		mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			unmarshalerHook(),
		),
	)); err != nil {
		return nil, err
	}

	return cfg, nil
}

func unmarshalerHook() mapstructure.DecodeHookFunc {
	return func(from reflect.Value, to reflect.Value) (interface{}, error) {
		if to.CanAddr() {
			to = to.Addr()
		}

		// If the destination implements the unmarshalling interface
		u, ok := to.Interface().(encoding.TextUnmarshaler)
		if !ok {
			return from.Interface(), nil
		}

		// If it is nil and a pointer, create and assign the target value first
		if to.IsNil() && to.Type().Kind() == reflect.Ptr {
			to.Set(reflect.New(to.Type().Elem()))
			u = to.Interface().(encoding.TextUnmarshaler)
		}

		var text []byte
		switch v := from.Interface().(type) {
		case string:
			text = []byte(v)
		case []byte:
			text = v
		default:
			return v, nil
		}

		if err := u.UnmarshalText(text); err != nil {
			return to.Interface(), err
		}
		return to.Interface(), nil
	}
}
