package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/patrickmn/go-cache"
)

var geoCache = cache.New(24*time.Hour, 10*time.Minute)

func StartAreaSearchServer() {
	fmt.Println("Starting server on :8080...")

	http.HandleFunc("/search", searchHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Missing 'q' parameter", http.StatusBadRequest)
		return
	}

	// Call nominatim and get results
	results, err := callNominatim(query)
	if err != nil {
		http.Error(w, "Failed to fetch location data", http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode the result and send it
	json.NewEncoder(w).Encode(results)
}

func callNominatim(address string) ([]map[string]interface{}, error) {
	fmt.Println("Checking cache for:", address)

	if x, found := geoCache.Get(address); found {
		fmt.Println("Cache hit!")
		results := x.([]map[string]interface{})
		return results, nil
	}

	fmt.Println("Cache miss! Calling Nominatim...")

	baseURL := "https://nominatim.openstreetmap.org/search"

	params := url.Values{}
	params.Add("q", address)
	params.Add("format", "json")

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("User-Agent", "ChariotX-AreaSearch/1.0 (contact: sdtp2244@gmail.com)")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body) // use io instead of ioutil
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var results []map[string]interface{}
	if err := json.Unmarshal(body, &results); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil, err
	}

	// Cache the results
	geoCache.Set(address, results, cache.DefaultExpiration)

	return results, nil
}
