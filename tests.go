package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

// TestLoadConfig tests the loadConfig function
func TestLoadConfig(t *testing.T) {
	// Create a temporary config file for testing
	tempConfig := `{
		"port": ":8080",
		"domain": "localhost",
		"static_dir": "html",
		"tls_config": {
			"cert_file": "server.crt",
			"key_file": "server.key"
		},
		"repo_config": {
			"version": "1.2.0",
			"author": "recon (contact@mail.recon.best)",
			"product": "Gamma Web Server",
			"repository": "https://github.com/gamma-gws/gws"
		}
	}`

	// Write the temporary config file
	err := ioutil.WriteFile("temp_config.json", []byte(tempConfig), 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		// Cleanup: Remove the temporary config file
		err := os.Remove("temp_config.json")
		if err != nil {
			t.Fatal(err)
		}
	}()

	// Call the loadConfig function
	config := loadConfig("temp_config.json")

	// Perform assertions to check if the loaded config matches the expected values
	if config.Port != ":8080" {
		t.Errorf("Unexpected Port. Expected: \":8080\", Got: %s", config.Port)
	}
	if config.Domain != "localhost" {
		t.Errorf("Unexpected Domain. Expected: \"localhost\", Got: %s", config.Domain)
	}
}

// TestPrintHeader tests the printHeader function
func TestPrintHeader(t *testing.T) {
	// Create a Config instance for testing
	config := Config{
		Port:   ":8080",
		Domain: "localhost",
		// ... Set other fields accordingly
	}

	// Call the printHeader function and capture the output
	out := captureOutput(func() {
		printHeader(config)
	})

	// Perform assertions to check if the output matches the expected format
	expectedOutput := "Hello, World! | Gamma Web Server v1.2.0 | Created by recon (contact@mail.recon.best)\n"
	if out != expectedOutput {
		t.Errorf("Unexpected output. Expected: %s, Got: %s", expectedOutput, out)
	}
}

// Helper function to capture stdout for testing
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}
