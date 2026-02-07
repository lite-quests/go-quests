package exit_quest

import (
	"os"
	"os/exec"
	"testing"
)

const helperEnv = "GO_HELPER_PROCESS"

func TestHelperProcess(t *testing.T) {
	if os.Getenv(helperEnv) != "1" {
		return
	}
	host := os.Getenv("TEST_HOST")
	port := os.Getenv("TEST_PORT")
	creds := os.Getenv("TEST_CREDS")

	portNum := 0

	if port == "valid" {
		portNum = 5432
	} else if port == "invalid" {
		portNum = 70000
	}

	credentialsValid := creds == "true"

	if host == "empty" {
		host = ""
	} else {
		host = "localhost"
	}

	ConnectDB(host, portNum, credentialsValid)
}

func runCase(host, port, creds string) int {
	cmd := exec.Command(os.Args[0], "-test.run=TestHelperProcess")
	cmd.Env = append(os.Environ(),
		helperEnv+"=1",
		"TEST_HOST="+host,
		"TEST_PORT="+port,
		"TEST_CREDS="+creds,
	)

	err := cmd.Run()
	if err == nil {
		return 0
	}

	if exitErr, ok := err.(*exec.ExitError); ok {
		return exitErr.ExitCode()
	}

	return -1
}

func TestConnectDB(t *testing.T) {
	tests := []struct {
		name  string
		host  string
		port  string
		creds string
		want  int
	}{
		{
			name:  "success",
			host:  "valid",
			port:  "valid",
			creds: "true",
			want:  0,
		},
		{
			name:  "empty host",
			host:  "empty",
			port:  "valid",
			creds: "true",
			want:  1,
		},
		{
			name:  "invalid port",
			host:  "valid",
			port:  "invalid",
			creds: "true",
			want:  2,
		},
		{
			name:  "invalid credentials",
			host:  "valid",
			port:  "valid",
			creds: "false",
			want:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := runCase(tt.host, tt.port, tt.creds)

			if got != tt.want {
				t.Fatalf("expected exit code %d, got %d", tt.want, got)
			}
		})
	}
}
