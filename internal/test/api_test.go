package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

const baseURL = "http://localhost:8080/api"

var token string
var transactionID int // variável global para reusar nos testes

func TestAPI(t *testing.T) {
	t.Run("Register", testRegister)
	t.Run("Login", testLogin)
	t.Run("Create Transaction", testCreateTransaction)
	t.Run("List Transactions", testListTransactions)
	t.Run("Update Transaction", testUpdateTransaction)
	t.Run("Delete Transaction", testDeleteTransaction)
	t.Run("Get Balance Report", testGetBalance)
	t.Run("Get Expenses by Category", testGetExpensesByCategory)
}

func testRegister(t *testing.T) {
	body := []byte(`{"name":"Test User","email":"test@example.com","password":"123456"}`)
	resp, err := http.Post(baseURL+"/register", "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusConflict {
		t.Errorf("expected 201 or 409, got %d", resp.StatusCode)
	}
}

func testLogin(t *testing.T) {
	body := []byte(`{"email":"test@example.com","password":"123456"}`)
	resp, err := http.Post(baseURL+"/login", "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("login failed, status %d", resp.StatusCode)
	}

	var result map[string]string
	_ = json.NewDecoder(resp.Body).Decode(&result)
	token = result["token"]
	if token == "" {
		t.Fatal("token is empty")
	}
}

func authRequest(method, url string, body []byte) (*http.Response, error) {
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	return client.Do(req)
}

func testCreateTransaction(t *testing.T) {
	body := []byte(`{"amount":100.50,"type":"income","category":"Salary","description":"Test income"}`)
	resp, err := authRequest("POST", baseURL+"/transactions", body)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		data, _ := ioutil.ReadAll(resp.Body)
		t.Errorf("expected 201, got %d. body: %s", resp.StatusCode, string(data))
		return
	}

	// Captura o ID retornado
	var result map[string]interface{}
	_ = json.NewDecoder(resp.Body).Decode(&result)
	if id, ok := result["id"].(float64); ok {
		transactionID = int(id)
	} else {
		t.Fatal("transaction ID not returned")
	}
}

func testListTransactions(t *testing.T) {
	resp, err := authRequest("GET", baseURL+"/transactions", nil)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func testUpdateTransaction(t *testing.T) {
	body := []byte(`{"amount":200.75,"type":"expense","category":"Food","description":"Updated transaction"}`)
	url := baseURL + "/transactions/" + fmt.Sprint(transactionID)

	resp, err := authRequest("PUT", url, body)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func testDeleteTransaction(t *testing.T) {
	url := baseURL + "/transactions/" + fmt.Sprint(transactionID)

	resp, err := authRequest("DELETE", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func testGetBalance(t *testing.T) {
	resp, err := authRequest("GET", baseURL+"/reports/balance", nil)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func testGetExpensesByCategory(t *testing.T) {
 	resp, err := authRequest("GET", baseURL+"/reports/expenses-by-category", nil)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
 	}
}
