package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWithBothEnv(t *testing.T) {
	err := os.Setenv("HOST", "192.168.1.1")
	if err != nil {
		t.Fatal("cannot set env `HOST`")
	}

	err = os.Setenv("PORT", "1111")
	if err != nil {
		t.Fatal("cannot set env `PORT`")
	}

	testConfig := &Config{
		"192.168.1.1",
		"1111",
	}

	gotConfig := New()

	assert.Equal(t, *testConfig, *gotConfig)
}

func TestNewWithHostOnly(t *testing.T) {
	err := os.Setenv("HOST", "192.168.1.1")
	if err != nil {
		t.Fatal("cannot set env `HOST`")
	}

	err = os.Unsetenv("PORT")
	if err != nil {
		t.Fatal("cannot unset env `PORT`")
	}

	testConfig := &Config{
		"192.168.1.1",
		"8080",
	}

	gotConfig := New()

	assert.Equal(t, *testConfig, *gotConfig)
}

func TestNewWithPortOnly(t *testing.T) {
	err := os.Setenv("PORT", "1111")
	if err != nil {
		t.Fatal("cannot set env `PORT`")
	}

	err = os.Unsetenv("HOST")
	if err != nil {
		t.Fatal("cannot unset env `HOST`")
	}

	testConfig := &Config{
		"localhost",
		"1111",
	}

	gotConfig := New()

	assert.Equal(t, *testConfig, *gotConfig)
}

func TestNewWithNoEnv(t *testing.T) {
	err := os.Unsetenv("PORT")
	if err != nil {
		t.Fatal("cannot unset env `PORT`")
	}

	err = os.Unsetenv("HOST")
	if err != nil {
		t.Fatal("cannot unset env `HOST`")
	}

	testConfig := &Config{
		"localhost",
		"8080",
	}

	gotConfig := New()

	assert.Equal(t, *testConfig, *gotConfig)
}

func TestCheckHost(t *testing.T) {
	tests := []struct {
		name    string
		host    string
		wantErr bool
	}{
		{
			"Good Host 1",
			"192.168.1.1",
			false,
		},
		{
			"Good Host 2",
			"10.10.10.10",
			false,
		},
		{
			"Good Host 3",
			"localhost",
			false,
		},
		{
			"Good Host 4",
			"0.0.0.0",
			false,
		},
		{
			"Good Host 5",
			"4.250.20.3",
			false,
		},
		{
			"Bad Host 1",
			"01.0.0",
			true,
		},
		{
			"Bad Host 2",
			"192.168.1.300",
			true,
		},
		{
			"Bad Host 3",
			"aboba",
			true,
		},
		{
			"Bad Host 4",
			"1921.68.1.1",
			true,
		},
		{
			"Empty host",
			"",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkHost(tt.host)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCheckPort(t *testing.T) {
	tests := []struct {
		name    string
		port    string
		wantErr bool
	}{
		{
			"Good port 1",
			"1000",
			false,
		},
		{
			"Good port 2",
			"65534",
			false,
		},
		{
			"Good port 3",
			"1",
			false,
		},
		{
			"Bad port 1",
			"0",
			true,
		},
		{
			"Bad port 2",
			"65536",
			true,
		},
		{
			"Bad port 3",
			"aboba",
			true,
		},
		{
			"Bad port 4",
			"-1",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkPort(tt.port)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
