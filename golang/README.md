# Golang Developer Test

This test is designed to assess the skills and knowledge of Golang developers. Candidates are expected to demonstrate proficiency in Golang fundamentals, concurrency, web development, and testing.

### Instructions
Fork this repository.
Complete the test.
Ensure that your code is well-documented and follows best practices.
Push the changes to your forked repository.
Provide the repository link to the person conducting the assessment.

# Test Sections
## Section 1: Language Fundamentals

### Variables and Types:

Write a Golang function that swaps the values of two variables without using a temporary variable.

### Slices:

Create a function that takes a slice of integers as input and returns the sum of all even numbers in the slice.

### Interfaces:

Define an interface called Logger with a method Log(message string). Implement this interface in a struct called FileLogger that logs messages to a file and in a struct called ConsoleLogger that logs messages to the console.

## Section 2: Concurrency
### Goroutines:

Write a Golang program that uses goroutines to calculate the sum of elements in a large array concurrently. Ensure proper synchronization.

### Channels:

Implement a program that uses channels to simulate a simple producer-consumer scenario. The producer should generate random numbers and send them to the consumer, which calculates and prints their square.

## Section 3: Web Development
### HTTP Server:

Create a simple Golang HTTP server that listens on port 8080. It should respond with "Hello, World!" for any incoming request.

### REST API:

Design and implement a basic RESTful API for a todo list. Include operations to create, read, update, and delete tasks. Use in-memory storage for simplicity.

## Section 4: Testing
### Unit Testing:

Write unit tests for the functions created in question 2 (sum of even numbers in a slice). Ensure that the tests cover different scenarios, including edge cases.

### HTTP Testing:

Write a test for the HTTP server created in question 6. Use the standard Go testing framework to verify that the server responds correctly to different HTTP methods.

## Section 5: Code Review
### Code Review:
Using `snippet.go`, identify and describe at least one potential issues or improvements. Discuss what changes you would recommend.
