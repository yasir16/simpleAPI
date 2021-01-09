package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestGetUsers(t *testing.T) {
	ts := httptest.NewServer(setupServer())

	defer ts.Close()

	client := &http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/users", ts.URL), nil)
	req.Header.Add("Authorization", "initestyasir")
	resp, err := client.Do(req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]
	fmt.Println(val)
	// Assert that the "content-type" header is actually set
	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	// Assert that it was set as expected
	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}
}

func TestCreateUser(t *testing.T) {
	ts := httptest.NewServer(setupServer())

	defer ts.Close()

	data := url.Values{}
	data.Set("fullName", "Muhammad Yasir Abdulazis")
	data.Set("userName", "namaBelakang")
	data.Set("phone", "081703078960")
	data.Set("address", "semoloaraawe")

	client := &http.Client{}
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/user", ts.URL), strings.NewReader(data.Encode()))
	req.Header.Add("Authorization", "initestyasir")
	resp, err := client.Do(req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

}
func TestGetUser(t *testing.T) {
	ts := httptest.NewServer(setupServer())

	defer ts.Close()

	client := &http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/user/1", ts.URL), nil)
	req.Header.Add("Authorization", "initestyasir")
	resp, err := client.Do(req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]
	fmt.Println(val)
	// Assert that the "content-type" header is actually set
	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	// Assert that it was set as expected
	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}
}
func TestUpdateUser(t *testing.T) {
	ts := httptest.NewServer(setupServer())

	defer ts.Close()

	data := url.Values{}
	data.Set("fullName", "Muhammad Yasir Abdulazis")
	data.Set("userName", "namaBelakang")
	data.Set("phone", "081703078960")
	data.Set("address", "semoloaraawe")

	client := &http.Client{}
	req, _ := http.NewRequest("PUT", fmt.Sprintf("%s/user/1", ts.URL), strings.NewReader(data.Encode()))
	req.Header.Add("Authorization", "initestyasir")
	resp, err := client.Do(req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

}

func TestDeleteUser(t *testing.T) {
	ts := httptest.NewServer(setupServer())

	defer ts.Close()

	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/user/1", ts.URL), nil)
	req.Header.Add("Authorization", "initestyasir")
	resp, err := client.Do(req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

}
