//go:build e2e

package e2e

import (
	"DoramaSet/internal/handler/apiserver"
	"context"
	"flag"
	"fmt"
	"github.com/pressly/goose"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"testing"
	"text/template"
)

func setupTestDatabase() (testcontainers.Container, string, int, error) {
	containerReq := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
		Env: map[string]string{
			"POSTGRES_DB":       "testdb",
			"POSTGRES_PASSWORD": "postgres",
			"POSTGRES_USER":     "postgres",
		},
	}

	dbContainer, _ := testcontainers.GenericContainer(
		context.Background(),
		testcontainers.GenericContainerRequest{
			ContainerRequest: containerReq,
			Started:          true,
		})

	host, _ := dbContainer.Host(context.Background())
	port, _ := dbContainer.MappedPort(context.Background(), "5432")

	dsn := fmt.Sprintf("host=%s port=%d user=postgres password=postgres dbname=testdb sslmode=disable", host, port.Int())
	pureDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, "", 0, fmt.Errorf("gorm open: %w", err)
	}

	text, err := os.ReadFile("../../deployments/init/postgres/01_create.sql")
	if err != nil {
		return nil, "", 0, fmt.Errorf("read file: %w", err)
	}
	if err := pureDB.Exec(string(text)).Error; err != nil {
		return nil, "", 0, fmt.Errorf("exec: %w", err)
	}

	text, err = os.ReadFile("../../deployments/init/postgres/02_constraints.sql")
	if err != nil {
		return nil, "", 0, fmt.Errorf("read file: %w", err)
	}
	if err := pureDB.Exec(string(text)).Error; err != nil {
		return nil, "", 0, fmt.Errorf("exec: %w", err)
	}

	text, err = os.ReadFile("./03_insert.sql")
	if err != nil {
		return nil, "", 0, fmt.Errorf("read file: %w", err)
	}
	if err := pureDB.Exec(string(text)).Error; err != nil {
		return nil, "", 0, fmt.Errorf("exec: %w", err)
	}

	sqlDB, err := pureDB.DB()
	if err != nil {
		return nil, "", 0, fmt.Errorf("get db: %w", err)
	}

	if err = goose.Up(sqlDB, "../../deployments/migration"); err != nil {
		return nil, "", 0, fmt.Errorf("up migrations: %w", err)
	}

	return dbContainer, host, port.Int(), nil
}

func TestMain(m *testing.M) {
	rc := 0
	defer func() {
		os.Exit(rc)
	}()

	tmplFile := flag.String("tmpl", "../../configs/test_config_temp.yml", "tmpl file name")

	err := flag.Set("config", "../../configs/test_config.yml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "set flag value: %s", err)
		rc = 1
		return
	}
	configPath := flag.Lookup("config").Value.(flag.Getter).Get().(string)

	flag.Parse()

	dbContainer, host, port, err := setupTestDatabase()
	if err != nil {
		fmt.Fprintf(os.Stderr, "setup test database: %s", err)
		rc = 1
		return
	}
	defer dbContainer.Terminate(context.Background())

	tmpl, err := template.ParseFiles(*tmplFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse template: %s", err)
		rc = 1
		return
	}

	cfg, err := os.Create(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create cfg file: %s", err)
		rc = 1
		return
	}
	v := struct {
		Host string
		Port int
	}{host, port}

	err = tmpl.Execute(cfg, v)
	if err != nil {
		fmt.Fprintf(os.Stderr, "execute tmpl: %s", err)
		rc = 1
		return
	}

	err = cfg.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "close file: %s", err)
		rc = 1
		return
	}

	_, err = apiserver.Init()
}
