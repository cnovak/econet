package econet

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	accountID = "ffqk2yixk1jip4syvvq9e839wkxcjwv"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

func setup() func() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	return func() {
		server.Close()
	}
}

func testData(path string) string {
	b, err := ioutil.ReadFile("testdata/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func TestEconetGetLocations(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/auth", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, testData("auth.json"))
	})
	mux.HandleFunc("/code/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, testData("get_location.json"))
	})

	baseURL = server.URL
	client, err := New("user@example.com", "1111")
	if err != nil {
		t.Errorf("got err: %s", err)
	}

	locations, err := client.getLocations()
	if err != nil {
		t.Fatalf("getLocations() returned error %s", err)
	}

	locationList := locations.Results.Locations
	if locationList == nil {
		t.Fatal("Location list is nil")
	}

	if len(locationList) != 1 {
		t.Fatalf("Location list should have 1 item, got %v", len(locationList))
	}

}

func TestEconetAuth(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/auth", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, testData("auth.json"))
	})

	baseURL = server.URL
	client, err := New("user@example.com", "1111")
	if err != nil {
		t.Errorf("got err: %s", err)
	}

	if client == nil {
		t.Fatal("client is nil")
	}

	// validate attributes
	if client.UserID != "ffqk2yixk1jip4syvvq9e839wkxcjwv" {
		t.Fatalf("UserID not set %v", client.UserID)
	}

	if client.Options.AccountID != "37507646828174499081" {
		t.Fatal("AccountID not set")
	}

	if client.UserToken == "" {
		t.Fatal("UserToken not set")
	}

}
