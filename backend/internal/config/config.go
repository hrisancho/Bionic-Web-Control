package config

import (
	"log"
	"os"
	"reflect"
	"time"

	"Bionic-Web-Control/internal/validator"
	"go.uber.org/zap/zapcore"

	"github.com/spf13/viper"
)

const (
	ProtocolHTTP  = "http"
	ProtocolHTTPS = "https"
)

const (
	DefaultTimeFormat = "02.01.2006 15:04:05"
	FilterTimeFormat  = "01/02/2006"
)

var (
	DebugLevel = zapcore.DebugLevel
)

// Config используется для хранения конфигурации сервера.
type Config struct {
	// ServerAddr - адрес HTTP сервера (по умолчанию: 127.0.0.1:8080)
	ServerAddr string `mapstructure:"SERVER_ADDR"`
	// SiteURL - URL адрес сайта (обязательное значение)
	// TODO стоит ли добавлять
	//SiteURL string `mapstructure:"SITE_URL" validate:"required"`
	// TimeFormat - формат времени (по умолчанию: 02.01.2006 15:04:05)
	TimeFormat string `mapstructure:"TIME_FORMAT"`

	// AllowOriginsCors - CORS разрешенные источники (по умолчанию: *)
	AllowOriginsCors string `mapstructure:"ALLOW_ORIGINS_CORS"`

	// MqttBrokerAddr - MQTT broker full addr (пример: tcp://192.168.1.108:1883)
	MqttBrokerAddr string `mapstructure:"MQTT_BROKER_ADDR" validate:"required"`
	MqttUsername   string `mapstructure:"MQTT_USERNAME"`
	MqttPassword   string `mapstructure:"MQTT_PASSWORD"`
	// MqttKeepAlive - MQTT client keep alive
	MqttKeepAlive time.Duration `mapstructure:"MQTT_KEEP_ALIVE"`
	// MqttPingTimeout - MQTT client ping timeout
	MqttPingTimeout time.Duration `mapstructure:"MQTT_PING_TIMEOUT"`
	// MqttQOS - MQTT client QoS
	MqttQOS byte `mapstructure:"MQTT_QOS"`
}

func initDefaultConfig() (v *viper.Viper) {
	v = viper.New()

	v.SetDefault("SERVER_ADDR", "127.0.0.1:8080")
	v.SetDefault("TIME_FORMAT", DefaultTimeFormat)
	v.SetDefault("ALLOW_ORIGINS_CORS", "*")
	v.SetDefault("MQTT_BROKER_ADDR", "mqtt://127.0.0.1:1883")
	v.SetDefault("MQTT_USERNAME", "backend-client")
	v.SetDefault("MQTT_PASSWORD", "public")
	v.SetDefault("MQTT_KEEP_ALIVE", 60*time.Second)
	v.SetDefault("MQTT_PING_TIMEOUT", 10*time.Second)
	v.SetDefault("MQTT_QOS", 2)

	return
}

func loadConfigFile(v *viper.Viper, path string) (config Config, err error) {
	v.AddConfigPath(path)
	v.SetConfigName("main")
	v.SetConfigType("env")

	v.AutomaticEnv()

	err = v.ReadInConfig()
	if err != nil {
		return
	}

	configReflectType := reflect.ValueOf(&config).Elem()
	configFieldsCount := configReflectType.NumField()

	err = v.Unmarshal(&config)
	if err != nil {
		return
	}
	time.Now().Unix()

	for fieldInd := 0; fieldInd < configFieldsCount; fieldInd++ {
		configField := configReflectType.Field(fieldInd)

		if configField.Kind() != reflect.Struct {
			continue
		}

		err = v.Unmarshal(configField.Addr().Interface())
		if err != nil {
			return
		}
	}

	return
}

func loadConfigEnv(v *viper.Viper) (config Config, err error) {
	envNameList := envNameListByConfig(reflect.TypeOf(config))
	for _, envName := range envNameList {
		err = v.BindEnv(envName, envName)
		if err != nil {
			return
		}
	}

	err = v.Unmarshal(&config)
	return
}

func envNameListByConfig(configType reflect.Type) (envNameList []string) {
	configFieldsCount := configType.NumField()
	envNameList = make([]string, 0, configFieldsCount)

	for fieldInd := 0; fieldInd < configFieldsCount; fieldInd++ {
		configField := configType.Field(fieldInd)

		if configField.Type.Kind() == reflect.Struct {
			envNameList = append(envNameList, envNameListByConfig(configField.Type)...)
		}

		envNameList = append(envNameList, configField.Tag.Get("mapstructure"))
	}
	return
}

func LoadConfig(appValidator *validator.AppValidator) (config Config, err error) {
	v := initDefaultConfig()

	if _, err = os.Stat("../main.env"); err == nil {
		config, err = loadConfigFile(v, "../")
	} else {
		log.Println("Loading config from env...")
		config, err = loadConfigEnv(v)
	}

	if err = appValidator.Struct(&config); err != nil {
		err = appValidator.ErrorTranslated(err)
		return
	}

	return config, err
}
