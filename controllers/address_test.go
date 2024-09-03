package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jash707/ecommerce-go/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mock UserCollection to use in the tests
var mockUserCollection *mongo.Collection

// stringPtr is a helper function that returns a pointer to a string.
func stringPtr(s string) *string {
	return &s
}

func TestAddAddress(t *testing.T) {
	// Initialize Gin in test mode
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Set up the route
	router.POST("/addaddress", AddAddress())

	// Create a mock address to add
	mockAddress := models.Address{
		Address_ID: primitive.NewObjectID(),
		House:      stringPtr("123 Main St"),
		Street:     stringPtr("Test City"),
		City:       stringPtr("Test State"),
		Pincode:    stringPtr("12345"),
	}

	// Convert mock address to JSON
	mockAddressJSON, _ := json.Marshal(mockAddress)

	// Create a mock ObjectID for user
	mockUserID := primitive.NewObjectID().Hex()

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, "/addaddress?userID="+mockUserID, bytes.NewBuffer(mockAddressJSON))
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Mock MongoDB aggregation result
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	// Insert mock data to the mockUserCollection
	mockUserCollection = setupMockCollection(t)
	defer mockUserCollection.Drop(ctx)

	// Mock aggregation to return a predefined value
	// mockData := bson.M{
	// 	"_id":   mockAddress.Address_ID,
	// 	"count": int32(1), // Set count less than 2 to simulate the addition of address
	// }
	// mockCursor := mockMongoCursor(ctx, mockData)
	// UserCollection = mockUserCollection // Set the global variable to use the mock collection

	// Run the handler
	router.ServeHTTP(recorder, req)

	// Verify that the response status code is 200 OK
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Verify the response body (you can add more checks depending on your needs)
	expectedResponse := "\"Successfully added the address\""
	assert.Equal(t, expectedResponse, recorder.Body.String())
}

func TestEditHomeAddress(t *testing.T) {
	// Initialize Gin in test mode
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Set up the route for the EditHomeAddress handler
	router.PUT("/edithomeaddress", EditHomeAddress())

	// Create a mock initial address and user document
	mockUserID := primitive.NewObjectID() // Generate a mock user ID
	mockInitialAddress := models.Address{
		Address_ID: mockUserID, // Using user ID as the address ID for simplicity
		House:      stringPtr("123 Initial St"),
		Street:     stringPtr("Initial City"),
		City:       stringPtr("Initial State"),
		Pincode:    stringPtr("11111"),
	}

	// Create a mock user document with an address array
	mockUser := models.User{
		ID:              mockUserID,
		User_ID:         mockUserID.Hex(),
		Address_Details: []models.Address{mockInitialAddress},
	}

	// Insert mock user into the mockUserCollection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mockUserCollection = setupMockCollection(t)
	defer mockUserCollection.Drop(ctx)

	_, err := mockUserCollection.InsertOne(ctx, mockUser)
	if err != nil {
		t.Fatalf("could not insert mock user: %v", err)
	}

	// Create a mock address to edit (the new address data)
	mockEditAddress := models.Address{
		House:   stringPtr("456 New Ave"),
		Street:  stringPtr("New City"),
		City:    stringPtr("New State"),
		Pincode: stringPtr("67890"),
	}

	// Convert mock address to JSON
	mockEditAddressJSON, _ := json.Marshal(mockEditAddress)

	// Create a new HTTP request with PUT method
	req, err := http.NewRequest(http.MethodPut, "/edithomeaddress?userID="+mockUser.User_ID, bytes.NewBuffer(mockEditAddressJSON))
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Run the handler
	router.ServeHTTP(recorder, req)

	// Verify that the response status code is 200 OK
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Print the actual response body for debugging
	fmt.Println("Actual Response Body:", recorder.Body.String())

	// Verify the response body matches the expected plain JSON string
	expectedResponse := "\"Successfully updated the home address\""
	assert.Equal(t, expectedResponse, recorder.Body.String())
}

// setupMockCollection sets up a mock MongoDB collection
func setupMockCollection(t *testing.T) *mongo.Collection {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("could not create Mongo client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		t.Fatalf("could not connect to Mongo: %v", err)
	}

	return client.Database("testdb").Collection("users")
}

// mockMongoCursor returns a mock cursor for MongoDB aggregation
// func mockMongoCursor(ctx context.Context, data bson.M) *mongo.Cursor {
// 	cursor, _ := mongo.NewCursor(ctx, nil, nil)
// 	// Assuming we're setting the cursor to return predefined data
// 	cursor.SetBatchSize(1)
// 	return cursor
// }
