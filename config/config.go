package config

import (
  "os"
  "strconv"
  "strings"
)

type DbConfig struct {
  DB_HOST string;
  DB_PORT string;
  DB_USER string;
  DB_PASSWORD string;
  DB_DATABASE string;
}

type DirectusConfig struct {
  DIRECTUS_HOST string;
  ADMIN_API_KEY string;
}

type Config struct {
  DbConfig DbConfig;
}

func New() *Config {
  return &Config {
    DbConfig: DbConfig{
      DB_HOST: getEnv("DB_HOST", ""),
      DB_PORT: getEnv("DB_PORT", ""),
      DB_USER: getEnv("DB_USER", ""),
      DB_PASSWORD: getEnv("DB_PASSWORD", ""),
      DB_DATABASE: getEnv("DB_DATABASE", ""),
    },
  };
}

func getEnv(key string, defaultVal string) string {
  if value, exists := os.LookupEnv(key); exists {
return value
  }

  return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
  valueStr := getEnv(name, "")
  if value, err := strconv.Atoi(valueStr); err == nil {
return value
  }

  return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
  valStr := getEnv(name, "")
  if val, err := strconv.ParseBool(valStr); err == nil {
return val
  }

  return defaultVal
}

func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
  valStr := getEnv(name, "")

  if valStr == "" {
return defaultVal
  }

  val := strings.Split(valStr, sep)

  return val
}