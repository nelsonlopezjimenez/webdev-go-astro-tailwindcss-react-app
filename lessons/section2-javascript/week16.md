---
title: "Functions - Organizing and Reusing Code"
description: "Master function declaration, parameters, return values, and scope for modular, maintainable JavaScript code"
week: 16
section: 2
prerequisites: ["Variables and data types", "Control flow", "Loops and arrays"]
objectives:
  - "Declare and invoke functions with proper syntax"
  - "Use parameters and arguments for flexible function behavior"
  - "Return values from functions and handle return data"
  - "Understand function scope and variable accessibility"
  - "Create modular, reusable code using function best practices"
---

# Week 16: Functions - Organizing and Reusing Code

## Learning Objectives

By the end of this week, students will be able to:
- Write function declarations and expressions correctly
- Pass data to functions using parameters and arguments
- Return calculated results and processed data from functions
- Understand and apply function scope principles
- Create helper functions and organize code modularly
- Debug function-related issues and trace execution flow

## Daily Activities

### Day 1: Function Declaration and Basic Usage
**Complexity: Beginner**

**Morning Session: Function Fundamentals**
- What functions are and why they're essential
- Function declaration syntax
- Calling (invoking) functions
- Function naming conventions

**Code Example:**
```javascript
// Basic function declaration
function sayHello() {
    console.log("Hello, World!"); // Beginner: Code inside function runs when called
}

// Calling (invoking) the function
sayHello(); // Output: Hello, World!

// Function with multiple statements
function showWelcomeMessage() {
    console.log("Welcome to our website!");
    console.log("We're glad you're here!");
    console.log("Enjoy browsing our content!");
}

showWelcomeMessage(); // Runs all three console.log statements

// Functions for calculations
function calculateCircleArea() {
    let radius = 5;
    let area = Math.PI * radius * radius;
    console.log(`The area of a circle with radius ${radius} is ${area.toFixed(2)}`);
}

calculateCircleArea(); // Output: The area of a circle with radius 5 is 78.54

// Functions for repetitive tasks
function drawLine() {
    console.log("=" * 50); // Creates a line of equal signs
}

function displayHeader() {
    drawLine();
    console.log("        DAILY REPORT");
    console.log("        " + new Date().toDateString());
    drawLine();
}

displayHeader(); // Creates a formatted header

// Function hoisting example (advanced concept)
// You can call functions before they're declared
greetUser(); // This works because of hoisting

function greetUser() {
    console.log("Thanks for visiting!");
}

// Best practice: Group related functions
function showLoadingMessage() {
    console.log("Loading...");
}

function showSuccessMessage() {
    console.log("Operation completed successfully!");
}

function showErrorMessage() {
    console.log("An error occurred. Please try again.");
}

// Using functions in sequence
console.log("Starting application...");
showLoadingMessage();
// Simulate some work happening
setTimeout(() => { // Advanced: setTimeout runs code after a delay
    showSuccessMessage();
}, 1000);
```

**Activity:** Create a collection of utility functions for a simple text-based game (show menu, display rules, show game over screen).

**Evening Session: Function Organization**
- Grouping related functions
- Function naming best practices
- Planning function structure

### Day 2: Parameters and Arguments
**Complexity: Beginner to Intermediate**

**Morning Session: Making Functions Flexible**
- Function parameters vs arguments
- Passing data to functions
- Multiple parameters and parameter order

**Code Example:**
```javascript
// Functions with parameters
function greetPerson(name) {
    console.log(`Hello, ${name}!`); // Beginner: name is a parameter
}

greetPerson("Alice");  // "Alice" is an argument
greetPerson("Bob");    // Same function, different argument
greetPerson("Carol");  // Reusable and flexible!

// Multiple parameters
function calculateRectangleArea(length, width) {
    let area = length * width;
    console.log(`A rectangle ${length} x ${width} has area ${area}`);
}

calculateRectangleArea(5, 3);  // Arguments: 5 and 3
calculateRectangleArea(10, 7); // Arguments: 10 and 7

// Parameter order matters
function introduceStudent(name, age, major) {
    console.log(`${name} is ${age} years old and studies ${major}.`);
}

introduceStudent("Emma", 20, "Computer Science");
// introduceStudent(20, "Emma", "Computer Science"); // Wrong order!

// Default parameters (newer JavaScript feature)
function createUserProfile(name, age = 18, country = "USA") {
    console.log(`User: ${name}, Age: ${age}, Country: ${country}`);
}

createUserProfile("John");                    // Uses defaults: John, 18, USA
createUserProfile("Maria", 25);               // Uses default country: Maria, 25, USA
createUserProfile("Yuki", 22, "Japan");       // All values provided

// Functions with different parameter types
function displayProductInfo(name, price, inStock, categories) {
    console.log(`Product: ${name}`);
    console.log(`Price: ${price}`);
    console.log(`Available: ${inStock ? "Yes" : "No"}`);
    console.log(`Categories: ${categories.join(", ")}`); // Arrays as parameters
}

displayProductInfo("Laptop", 999, true, ["Electronics", "Computers", "Portable"]);

// Advanced: Rest parameters for variable number of arguments
function calculateSum(...numbers) { // ...numbers collects all arguments into an array
    let total = 0;
    for (let i = 0; i < numbers.length; i++) {
        total += numbers[i];
    }
    console.log(`Sum of ${numbers.join(" + ")} = ${total}`);
    return total;
}

calculateSum(1, 2, 3);       // Works with 3 numbers
calculateSum(5, 10, 15, 20); // Works with 4 numbers
calculateSum(100);           // Works with 1 number

// Practical example: User validation
function validateUser(username, email, age) {
    let errors = [];
    
    if (username.length < 3) {
        errors.push("Username must be at least 3 characters");
    }
    
    if (!email.includes("@")) {
        errors.push("Email must contain @ symbol");
    }
    
    if (age < 13 || age > 120) {
        errors.push("Age must be between 13 and 120");
    }
    
    if (errors.length > 0) {
        console.log("Validation errors:");
        for (let i = 0; i < errors.length; i++) {
            console.log(`- ${errors[i]}`);
        }
    } else {
        console.log("User validation passed!");
    }
}

validateUser("jo", "invalid-email", 150); // Shows multiple errors
validateUser("john_doe", "john@example.com", 25); // Passes validation
```

**Activity:** Build a personal information manager with functions that accept different types and numbers of parameters.

**Evening Session: Parameter Validation**
- Checking parameter types and values
- Handling missing or invalid parameters
- Error prevention strategies

### Day 3: Return Values and Function Output
**Complexity: Beginner to Intermediate**

**Morning Session: Getting Data Back from Functions**
- The return statement
- Using returned values
- Functions that transform data

**Code Example:**
```javascript
// Basic return statement
function addTwoNumbers(a, b) {
    return a + b; // Beginner: return sends value back to where function was called
}

let result = addTwoNumbers(5, 3); // result now contains 8
console.log(`5 + 3 = ${result}`);

// Functions can return different data types
function getGreeting(name) {
    return `Hello, ${name}! Welcome back.`; // Returns a string
}

function isEven(number) {
    return number % 2 === 0; // Returns a boolean
}

function getRandomNumber(min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min; // Returns a number
}

// Using returned values
let greeting = getGreeting("Sarah");
console.log(greeting);

if (isEven(10)) {
    console.log("10 is even");
}

let dice = getRandomNumber(1, 6);
console.log(`You rolled a ${dice}`);

// Returning objects for complex data
function createStudent(name, grade, subjects) {
    return {
        name: name,
        grade: grade,
        subjects: subjects,
        isActive: true,
        enrollmentDate: new Date()
    };
}

let newStudent = createStudent("Alex", 10, ["Math", "Science", "English"]);
console.log("New student created:", newStudent);

// Returning arrays
function getTopScores(scores, count) {
    let sortedScores = scores.slice(); // Create a copy
    sortedScores.sort((a, b) => b - a); // Sort descending
    return sortedScores.slice(0, count); // Return top 'count' scores
}

let gameScores = [85, 92, 78, 96, 88, 91, 83];
let topThree = getTopScores(gameScores, 3);
console.log("Top 3 scores:", topThree);

// Early returns for efficiency
function findFirstMatch(array, target) {
    for (let i = 0; i < array.length; i++) {
        if (array[i] === target) {
            return i; // Return immediately when found - no need to continue
        }
    }
    return -1; // Return -1 if not found (common convention)
}

let numbers = [10, 25, 3, 47, 12];
let position = findFirstMatch(numbers, 47);
if (position !== -1) {
    console.log(`Found 47 at position ${position}`);
} else {
    console.log("47 not found in array");
}

// Multiple return points for different conditions
function categorizeGrade(score) {
    if (score >= 90) {
        return "A - Excellent";
    } else if (score >= 80) {
        return "B - Good";
    } else if (score >= 70) {
        return "C - Satisfactory";
    } else if (score >= 60) {
        return "D - Needs Improvement";
    } else {
        return "F - Failing";
    }
    // Code after this point never runs
}

let studentGrade = categorizeGrade(87);
console.log(`Grade: ${studentGrade}`);

// Advanced: Returning functions (higher-order functions preview)
function createMultiplier(factor) {
    return function(number) {
        return number * factor;
    };
}

let double = createMultiplier(2);
let triple = createMultiplier(3);

console.log(double(5)); // 10
console.log(triple(4)); // 12

// Practical example: Data processing pipeline
function processOrderData(rawOrder) {
    // Validate the order first
    if (!rawOrder.items || rawOrder.items.length === 0) {
        return { error: "Order must contain items" };
    }
    
    if (!rawOrder.customerEmail || !rawOrder.customerEmail.includes("@")) {
        return { error: "Valid customer email required" };
    }
    
    // Calculate totals
    let subtotal = 0;
    for (let i = 0; i < rawOrder.items.length; i++) {
        subtotal += rawOrder.items[i].price * rawOrder.items[i].quantity;
    }
    
    let tax = subtotal * 0.08; // 8% tax
    let shipping = subtotal > 50 ? 0 : 9.99; // Free shipping over $50
    let total = subtotal + tax + shipping;
    
    // Return processed order
    return {
        orderId: Date.now(), // Simple ID generation
        customer: rawOrder.customerEmail,
        items: rawOrder.items,
        subtotal: Math.round(subtotal * 100) / 100, // Round to 2 decimal places
        tax: Math.round(tax * 100) / 100,
        shipping: shipping,
        total: Math.round(total * 100) / 100,
        status: "pending"
    };
}

// Test the function
let rawOrder = {
    customerEmail: "customer@example.com",
    items: [
        { name: "Widget", price: 19.99, quantity: 2 },
        { name: "Gadget", price: 34.50, quantity: 1 }
    ]
};

let processedOrder = processOrderData(rawOrder);
if (processedOrder.error) {
    console.log("Order error:", processedOrder.error);
} else {
    console.log("Processed order:", processedOrder);
}
```

**Activity:** Create a calculator library with functions that return computed values for various mathematical operations.

**Evening Session: Return Value Patterns**
- When to return vs when to log
- Returning multiple values using objects/arrays
- Chaining function calls

### Day 4: Function Scope and Variable Access
**Complexity: Intermediate**

**Morning Session: Understanding Scope**
- Global vs local scope
- Block scope with let and const
- Variable shadowing and scope chain

**Code Example:**
```javascript
// Global scope - accessible everywhere
let globalMessage = "I'm accessible everywhere!";
var oldGlobalVar = "I'm also global (legacy)";

function demonstrateScope() {
    console.log(globalMessage); // Can access global variables
    console.log(oldGlobalVar);  // Can access global var
}

demonstrateScope();

// Function scope - variables declared inside functions
function createLocalVariables() {
    let localMessage = "I only exist inside this function"; // Beginner: local scope
    var functionScoped = "I'm also local to this function";
    
    console.log("Inside function:", localMessage);
    console.log("Inside function:", functionScoped);
}

createLocalVariables();
// console.log(localMessage); // Error! localMessage doesn't exist here

// Block scope with let and const
function demonstrateBlockScope() {
    let outerVariable = "I'm in the function scope";
    
    if (true) {
        let blockVariable = "I'm in the block scope"; // Only exists in this { }
        const alsoBlockScoped = "Me too!";
        var functionScoped = "I escape the block"; // var ignores block scope
        
        console.log("Inside block:", outerVariable);     // Works
        console.log("Inside block:", blockVariable);     // Works
        console.log("Inside block:", alsoBlockScoped);   // Works
    }
    
    console.log("Outside block:", outerVariable);       // Works
    console.log("Outside block:", functionScoped);      // Works - var escaped!
    // console.log("Outside block:", blockVariable);    // Error! Block scoped
}

demonstrateBlockScope();

// Variable shadowing - local variables "hide" global ones
let userName = "Global User";

function loginUser() {
    let userName = "Local User"; // Shadows the global userName
    console.log("Inside function:", userName); // Prints "Local User"
}

loginUser();
console.log("Outside function:", userName); // Prints "Global User"

// Parameters create local variables
function processUserData(userName, userAge) { // userName parameter shadows global
    let userStatus = userName ? "active" : "inactive"; // Local variable
    console.log(`Processing: ${userName}, age ${userAge}, status: ${userStatus}`);
}

processUserData("Alice", 25);
// userName, userAge, and userStatus don't exist outside the function

// Practical example: Avoiding global pollution
function createShoppingCart() {
    let items = []; // Private to this function
    let total = 0;  // Private to this function
    
    function addItem(name, price, quantity) {
        let item = { name, price, quantity }; // Local to addItem
        items.push(item);
        total += price * quantity;
        console.log(`Added ${quantity} ${name}(s) for ${price * quantity}`);
    }
    
    function removeItem(name) {
        for (let i = 0; i < items.length; i++) { // i is block scoped
            if (items[i].name === name) {
                total -= items[i].price * items[i].quantity;
                items.splice(i, 1);
                console.log(`Removed ${name} from cart`);
                return;
            }
        }
        console.log(`${name} not found in cart`);
    }
    
    function getCartSummary() {
        return {
            items: items.slice(), // Return a copy, not the original
            total: Math.round(total * 100) / 100,
            itemCount: items.length
        };
    }
    
    // Return an object with the functions (closure pattern)
    return {
        add: addItem,
        remove: removeItem,
        summary: getCartSummary
    };
}

// Usage - items and total are protected/private
let cart = createShoppingCart();
cart.add("Laptop", 999, 1);
cart.add("Mouse", 29, 2);
console.log("Cart summary:", cart.summary());
cart.remove("Mouse");
console.log("Updated summary:", cart.summary());

// Advanced: Closure example
function createCounter(startValue) {
    let count = startValue; // This variable is "closed over"
    
    return function() {
        count++; // Can access and modify count even after createCounter finishes
        return count;
    };
}

let counter1 = createCounter(0);
let counter2 = createCounter(100);

console.log(counter1()); // 1
console.log(counter1()); // 2
console.log(counter2()); // 101
console.log(counter1()); // 3 - each counter has its own count variable
```

**Activity:** Build a simple game where different functions need access to shared game state while keeping some data private.

**Evening Session: Best Practices for Scope**
- Minimizing global variables
- When to use different scope levels
- Debugging scope-related issues

### Day 5: Advanced Functions and Best Practices
**Complexity: Intermediate to Advanced**

**Morning Session: Function Expressions and Advanced Patterns**
- Function expressions vs declarations
- Anonymous functions and callbacks
- Arrow functions (ES6) introduction

**Code Example:**
```javascript
// Function expressions - storing functions in variables
let sayGoodbye = function(name) {
    return `Goodbye, ${name}! See you later.`;
};

console.log(sayGoodbye("Maria")); // Works just like regular functions

// Anonymous functions - functions without names
let numbers = [1, 2, 3, 4, 5];

// Using anonymous function with array methods
let doubled = numbers.map(function(num) {  // Anonymous function as parameter
    return num * 2;
});
console.log("Doubled:", doubled);

// Arrow functions - shorter syntax for simple functions
let tripled = numbers.map(num => num * 3); // Beginner: => creates arrow function
console.log("Tripled:", tripled);

// Arrow functions with multiple parameters
let addNumbers = (a, b) => a + b;
let result = addNumbers(10, 5);
console.log("Sum:", result);

// Arrow functions with block body
let processOrder = (order) => {
    let tax = order.total * 0.08;
    let finalTotal = order.total + tax;
    return {
        ...order,
        tax: tax,
        finalTotal: finalTotal
    };
};

// Callback functions - functions passed to other functions
function processArray(array, callback) {
    let results = [];
    for (let i = 0; i < array.length; i++) {
        results.push(callback(array[i], i)); // Call the callback for each item
    }
    return results;
}

// Using different callbacks with the same function
let squares = processArray([1, 2, 3, 4], (num) => num * num);
let withIndex = processArray(['a', 'b', 'c'], (letter, index) => `${index}: ${letter}`);

console.log("Squares:", squares);
console.log("With index:", withIndex);

// Higher-order functions - functions that work with other functions
function createValidator(rule) {
    return function(value) {
        return rule(value);
    };
}

// Create specific validators
let isPositive = createValidator(num => num > 0);
let isValidEmail = createValidator(email => email.includes("@") && email.includes("."));
let isLongEnough = createValidator(str => str.length >= 8);

console.log(isPositive(5));        // true
console.log(isValidEmail("test@example.com")); // true
console.log(isLongEnough("password123"));      // true

// Function composition - combining functions
function addOne(x) { return x + 1; }
function double(x) { return x * 2; }
function square(x) { return x * x; }

// Manual composition
let result1 = square(double(addOne(3))); // ((3 + 1) * 2)² = (8)² = 64

// Compose function for cleaner composition
function compose(...functions) {
    return function(value) {
        return functions.reduceRight((acc, fn) => fn(acc), value);
    };
}

let transform = compose(square, double, addOne);
let result2 = transform(3); // Same as above: 64

// Practical example: Event handling system
function createEventHandler() {
    let listeners = {};
    
    function addEventListener(event, callback) {
        if (!listeners[event]) {
            listeners[event] = [];
        }
        listeners[event].push(callback);
    }
    
    function removeEventListener(event, callback) {
        if (listeners[event]) {
            listeners[event] = listeners[event].filter(cb => cb !== callback);
        }
    }
    
    function triggerEvent(event, data) {
        if (listeners[event]) {
            listeners[event].forEach(callback => {
                try {
                    callback(data);
                } catch (error) {
                    console.error(`Error in event handler for ${event}:`, error);
                }
            });
        }
    }
    
    return {
        on: addEventListener,
        off: removeEventListener,
        trigger: triggerEvent
    };
}

// Usage example
let eventSystem = createEventHandler();

// Define event handlers
let userLoginHandler = (userData) => {
    console.log(`User ${userData.name} logged in at ${new Date()}`);
};

let analyticsHandler = (userData) => {
    console.log(`Tracking login for user ID: ${userData.id}`);
};

// Register handlers
eventSystem.on('userLogin', userLoginHandler);
eventSystem.on('userLogin', analyticsHandler);

// Trigger event
eventSystem.trigger('userLogin', { id: 123, name: 'Alice' });

// Advanced: Function memoization for performance
function memoize(fn) {
    let cache = {};
    return function(...args) {
        let key = JSON.stringify(args);
        if (key in cache) {
            console.log('Cache hit!');
            return cache[key];
        }
        console.log('Computing...');
        let result = fn.apply(this, args);
        cache[key] = result;
        return result;
    };
}

// Expensive function to memoize
function fibonacci(n) {
    if (n <= 1) return n;
    return fibonacci(n - 1) + fibonacci(n - 2);
}

let memoizedFib = memoize(fibonacci);

console.log(memoizedFib(10)); // Computes
console.log(memoizedFib(10)); // Cache hit!
console.log(memoizedFib(11)); // Computes 11, uses cached 10

// Best practices example: Modular code organization
const UserUtils = {
    validateEmail: (email) => {
        return email.includes('@') && email.includes('.');
    },
    
    formatName: (firstName, lastName) => {
        return `${firstName.charAt(0).toUpperCase()}${firstName.slice(1)} ${lastName.charAt(0).toUpperCase()}${lastName.slice(1)}`;
    },
    
    generateUsername: (firstName, lastName) => {
        return `${firstName.toLowerCase()}.${lastName.toLowerCase()}`;
    },
    
    createUserProfile: function(userData) {
        if (!this.validateEmail(userData.email)) {
            throw new Error('Invalid email address');
        }
        
        return {
            id: Date.now(),
            name: this.formatName(userData.firstName, userData.lastName),
            username: this.generateUsername(userData.firstName, userData.lastName),
            email: userData.email,
            createdAt: new Date(),
            isActive: true
        };
    }
};

// Usage of modular functions
try {
    let newUser = UserUtils.createUserProfile({
        firstName: 'john',
        lastName: 'doe',
        email: 'john.doe@example.com'
    });
    console.log('Created user:', newUser);
} catch (error) {
    console.error('User creation failed:', error.message);
}
```

**Activity:** Build a comprehensive utility library with various types of functions that work together to solve complex problems.

**Evening Session: Function Design Principles**
- Single responsibility principle
- Function naming and documentation
- Testing and debugging functions

## Hands-on Exercises

### Exercise 1: Personal Finance Calculator
Create a comprehensive financial planning tool that uses multiple functions to:
- Calculate loan payments with different parameters
- Process budget data and return analysis
- Generate financial reports using helper functions
- Validate financial inputs and handle edge cases

### Exercise 2: Text Processing Library
Build a text analysis system featuring:
- Functions that transform and analyze text data
- Modular functions for different text operations
- Higher-order functions for flexible text processing
- Callback-based functions for custom transformations

### Exercise 3: Interactive Web Application
Develop a dynamic web page that:
- Uses functions to generate different types of content
- Implements event handling with function callbacks
- Organizes code into logical function modules
- Demonstrates proper scope management and data flow

## Resources

**MDN Documentation:**
- [Functions](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Functions)
- [Function declarations](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/function)
- [Function expressions](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/function)
- [Arrow functions](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Functions/Arrow_functions)
- [Closures](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Closures)

**W3Schools References:**
- [JavaScript Functions](https://www.w3schools.com/js/js_functions.asp)
- [Function Parameters](https://www.w3schools.com/js/js_function_parameters.asp)
- [Function Return](https://www.w3schools.com/js/js_function_return.asp)
- [Function Scope](https://www.w3schools.com/js/js_scope.asp)

## Assignment Checklist

**Basic Requirements:**
- [ ] Created function declarations with proper syntax
- [ ] Used parameters and arguments effectively
- [ ] Implemented return statements and used returned values
- [ ] Demonstrated understanding of function scope
- [ ] Organized code using multiple related functions

**Intermediate Challenges:**
- [ ] Implemented functions with complex parameter handling
- [ ] Used function expressions and arrow functions appropriately
- [ ] Created higher-order functions and callbacks
- [ ] Applied proper scope management and avoided global pollution
- [ ] Built reusable function libraries

**Advanced Extensions:**
- [ ] Implemented advanced function patterns (closures, memoization)
- [ ] Created modular code organization systems
- [ ] Used functions for complex data processing pipelines
- [ ] Implemented error handling and validation in functions
- [ ] Demonstrated mastery of function composition and design patterns

// Practical example: Data processing pipeline
function processOrderData(rawOrder) {
    // Validate the order first
    if (!rawOrder.items || rawOrder.items.length === 0) {
        return { error: "Order must contain items" };
    }
    
    if (!rawOrder.customerEmail || !rawOrder.customerEmail.includes("@")) {
        return { error: "Valid customer email required" };
    }
    
    // Calculate totals
    let subtotal = 0;
    for (let i = 0; i < rawOrder.items.length; i++) {
        subtotal += rawOrder.items[i].price * rawOrder.items[i].quantity;
    }
    
    let tax = subtotal * 0.08; // 8% tax
    let shipping = subtotal > 50 ? 0 : 9.99; // Free shipping over $50
    let total = subtotal + tax + shipping;
    
    // Return processed order
    return {
        orderId: Date.now(), // Simple ID generation
        customer: rawOrder.customerEmail,
        items: rawOrder.items,
        subtotal: Math.round(subtotal * 100) / 100, // Round to 2 decimal places
        tax: Math.round(tax * 100) / 100,
        shipping: shipping,
        total: Math.round(total * 100) / 100,
        status: "pending"
    };
}

// Test the function
let rawOrder = {
    customerEmail: "customer@example.com",
    items: [
        { name: "Widget", price: 19.99, quantity: 2 },
        { name: "Gadget", price: 34.50, quantity: 1 }
    ]
};

let processedOrder = processOrderData(rawOrder);
if (processedOrder.error) {
    console.log("Order error:", processedOrder.error);
} else {
    console.log("Processed order:", processedOrder);
}
```

**Activity:** Create a calculator library with functions that return computed values for various mathematical operations.

**Evening Session: Return Value Patterns**
- When to return vs when to log
- Returning multiple values using objects/arrays
- Chaining function calls

### Day 4: Function Scope and Variable Access
**Complexity: Intermediate**

**Morning Session: Understanding Scope**
- Global vs local scope
- Block scope with let and const
- Variable shadowing and scope chain

**Code Example:**
```javascript
// Global scope - accessible everywhere
let globalMessage = "I'm accessible everywhere!";
var oldGlobalVar = "I'm also global (legacy)";

function demonstrateScope() {
    console.log(globalMessage); // Can access global variables
    console.log(oldGlobalVar);  // Can access global var
}

demonstrateScope();

// Function scope - variables declared inside functions
function createLocalVariables() {
    let localMessage = "I only exist inside this function"; // Beginner: local scope
    var functionScoped = "I'm also local to this function";
    
    console.log("Inside function:", localMessage);
    console.log("Inside function:", functionScoped);
}

createLocalVariables();
// console.log(localMessage); // Error! localMessage doesn't exist here

// Block scope with let and const
function demonstrateBlockScope() {
    let outerVariable = "I'm in the function scope";
    
    if (true) {
        let blockVariable = "I'm in the block scope"; // Only exists in this { }
        const alsoBlockScoped = "Me too!";
        var functionScoped = "I escape the block"; // var ignores block scope
        
        console.log("Inside block:", outerVariable);     // Works
        console.log("Inside block:", blockVariable);     // Works
        console.log("Inside block:", alsoBlockScoped);   // Works
    }
    
    console.log("Outside block:", outerVariable);       // Works
    console.log("Outside block:", functionScoped);      // Works - var escaped!
    // console.log("Outside block:", blockVariable);    // Error! Block scoped
}

demonstrateBlockScope();

// Variable shadowing - local variables "hide" global ones
let userName = "Global User";

function loginUser() {
    let userName = "Local User"; // Shadows the global userName
    console.log("Inside function:", userName); // Prints "Local User"
}

loginUser();
console.log("Outside function:", userName); // Prints "Global User"

// Parameters create local variables
function processUserData(userName, userAge) { // userName parameter shadows global
    let userStatus = userName ? "active" : "inactive"; // Local variable
    console.log(`Processing: ${userName}, age ${userAge}, status: ${userStatus}`);
}

processUserData("Alice", 25);
// userName, userAge, and userStatus don't exist outside the function

// Practical example: Avoiding global pollution
function createShoppingCart() {
    let items = []; // Private to this function
    let total = 0;  // Private to this function
    
    function addItem(name, price, quantity) {
        let item = { name, price, quantity }; // Local to addItem
        items.push(item);
        total += price * quantity;
        console.log(`Added ${quantity} ${name}(s) for ${price * quantity}`);
    }
    
    function removeItem(name) {
        for (let i = 0; i < items.length; i++) { // i is block scoped
            if (items[i].name === name) {
                total -= items[i].price * items[i].quantity;
                items.splice(i, 1);
                console.log(`Removed ${name} from cart`);
                return;
            }
        }
        console.log(`${name} not found in cart`);
    }
    
    function getCartSummary() {
        return {
            items: items.slice(), // Return a copy, not the original
            total: Math.round(total * 100) / 100,
            itemCount: items.length
        };
    }
    
    // Return an object with the functions (closure pattern)
    return {
        add: addItem,
        remove: removeItem,
        summary: getCartSummary
    };
}

// Usage - items and total are protected/private
let cart = createShoppingCart();
cart.add("Laptop", 999, 1);
cart.add("Mouse", 29, 2);
console.log("Cart summary:", cart.summary());
cart.remove("Mouse");
console.log("Updated summary:", cart.summary());

// Advanced: Closure example
function createCounter(startValue) {
    let count = startValue; // This variable is "closed over"
    
    return function() {
        count++; // Can access and modify count even after createCounter finishes
        return count;
    };
}

let counter1 = createCounter(0);
let counter2 = createCounter(100);

console.log(counter1()); // 1
console.log(counter1()); // 2
console.log(counter2()); // 101
console.log(counter1()); // 3 - each counter has its own count variable
```

**Activity:** Build a simple game where different functions need access to shared game state while keeping some data private.

**Evening Session: Best Practices for Scope**
- Minimizing global variables
- When to use different scope levels
- Debugging scope-related issues

### Day 5: Advanced Functions and Best Practices
**Complexity: Intermediate to Advanced**

**Morning Session: Function Expressions and Advanced Patterns**
- Function expressions vs declarations
- Anonymous functions and callbacks
- Arrow functions (ES6) introduction

**Code Example:**
```javascript
// Function expressions - storing functions in variables
let sayGoodbye = function(name) {
    return `Goodbye, ${name}! See you later.`;
};

console.log(sayGoodbye("Maria")); // Works just like regular functions

// Anonymous functions - functions without names
let numbers = [1, 2, 3, 4, 5];

// Using anonymous function with array methods
let doubled = numbers.map(function(num) {  // Anonymous function as parameter
    return num * 2;
});
console.log("Doubled:", doubled);

// Arrow functions - shorter syntax for simple functions
let tripled = numbers.map(num => num * 3); // Beginner: => creates arrow function
console.log("Tripled:", tripled);

// Arrow functions with multiple parameters
let addNumbers = (a, b) => a + b;
let result = addNumbers(10, 5);
console.log("Sum:", result);

// Arrow functions with block body
let processOrder = (order) => {
    let tax = order.total * 0.08;
    let finalTotal = order.total + tax;
    return {
        ...order,
        tax: tax,
        finalTotal: finalTotal
    };
};

// Callback functions - functions passed to other functions
function processArray(array, callback) {
    let results = [];
    for (let i = 0; i < array.length; i++) {
        results.push(callback(array[i], i)); // Call the callback for each item
    }
    return results;
}

// Using different callbacks with the same function
let squares = processArray([1, 2, 3, 4], (num) => num * num);
let withIndex = processArray(['a', 'b', 'c'], (letter, index) => `${index}: ${letter}`);

console.log("Squares:", squares);
console.log("With index:", withIndex);

// Higher-order functions - functions that work with other functions
function createValidator(rule) {
    return function(value) {
        return rule(value);
    };
}

// Create specific validators
let isPositive = createValidator(num => num > 0);
let isValidEmail = createValidator(email => email.includes("@") && email.includes("."));
let isLongEnough = createValidator(str => str.length >= 8);

console.log(isPositive(5));        // true
console.log(isValidEmail("test@example.com")); // true
console.log(isLongEnough("password123"));      // true

// Function composition - combining functions
function addOne(x) { return x + 1; }
function double(x) { return x * 2; }
function square(x) { return x * x; }

// Manual composition
let result1 = square(double(addOne(3))); // ((3 + 1) * 2)² = (8)² = 64

// Compose function for cleaner composition
function compose(...functions) {
    return function(value) {
        return functions.reduceRight((acc, fn) => fn(acc), value);
    };
}

let transform = compose(square, double, addOne);
let result2 = transform(3); // Same as above: 64

// Practical example: Event handling system
function createEventHandler() {
    let listeners = {};
    
    function addEventListener(event, callback) {
        if (!listeners[event]) {
            listeners[event] = [];
        }
        listeners[event].push(callback);
    }
    
    function removeEventListener(event, callback) {
        if (listeners[event]) {
            listeners[event] = listeners[event].filter(cb => cb !== callback);
        }
    }
    
    function triggerEvent(event, data) {
        if (listeners[event]) {
            listeners[event].forEach(callback => {
                try {
                    callback(data);
                } catch (error) {
                    console.error(`Error in event handler for ${event}:`, error);
                }
            });
        }
    }
    
    return {
        on: addEventListener,
        off: removeEventListener,
        trigger: triggerEvent
    };
}

// Usage example
let eventSystem = createEventHandler();

// Define event handlers
let userLoginHandler = (userData) => {
    console.log(`User ${userData.name} logged in at ${new Date()}`);
};

let analyticsHandler = (userData) => {
    console.log(`Tracking login for user ID: ${userData.id}`);
};

// Register handlers
eventSystem.on('userLogin', userLoginHandler);
eventSystem.on('userLogin', analyticsHandler);

// Trigger event
eventSystem.trigger('userLogin', { id: 123, name: 'Alice' });

// Advanced: Function memoization for performance
function memoize(fn) {
    let cache = {};
    return function(...args) {
        let key = JSON.stringify(args);
        if (key in cache) {
            console.log('Cache hit!');
            return cache[key];
        }
        console.log('Computing...');
        let result = fn.apply(this, args);
        cache[key] = result;
        return result;
    };
}

// Expensive function to memoize
function fibonacci(n) {
    if (n <= 1) return n;
    return fibonacci(n - 1) + fibonacci(n - 2);
}

let memoizedFib = memoize(fibonacci);

console.log(memoizedFib(10)); // Computes
console.log(memoizedFib(10)); // Cache hit!
console.log(memoizedFib(11)); // Computes 11, uses cached 10

// Best practices example: Modular code organization
const UserUtils = {
    validateEmail: (email) => {
        return email.includes('@') && email.includes('.');
    },
    
    formatName: (firstName, lastName) => {
        return `${firstName.charAt(0).toUpperCase()}${firstName.slice(1)} ${lastName.charAt(0).toUpperCase()}${lastName.slice(1)}`;
    },
    
    generateUsername: (firstName, lastName) => {
        return `${firstName.toLowerCase()}.${lastName.toLowerCase()}`;
    },
    
    createUserProfile: function(userData) {
        if (!this.validateEmail(userData.email)) {
            throw new Error('Invalid email address');
        }
        
        return {
            id: Date.now(),
            name: this.formatName(userData.firstName, userData.lastName),
            username: this.generateUsername(userData.firstName, userData.lastName),
            email: userData.email,
            createdAt: new Date(),
            isActive: true
        };
    }
};

// Usage of modular functions
try {
    let newUser = UserUtils.createUserProfile({
        firstName: 'john',
        lastName: 'doe',
        email: 'john.doe@example.com'
    });
    console.log('Created user:', newUser);
} catch (error) {
    console.error('User creation failed:', error.message);
}
```

**Activity:** Build a comprehensive utility library with various types of functions that work together to solve complex problems.

**Evening Session: Function Design Principles**
- Single responsibility principle
- Function naming and documentation
- Testing and debugging functions

## Hands-on Exercises

### Exercise 1: Personal Finance Calculator
Create a comprehensive financial planning tool that uses multiple functions to:
- Calculate loan payments with different parameters
- Process budget data and return analysis
- Generate financial reports using helper functions
- Validate financial inputs and handle edge cases

### Exercise 2: Text Processing Library
Build a text analysis system featuring:
- Functions that transform and analyze text data
- Modular functions for different text operations
- Higher-order functions for flexible text processing
- Callback-based functions for custom transformations

### Exercise 3: Interactive Web Application
Develop a dynamic web page that:
- Uses functions to generate different types of content
- Implements event handling with function callbacks
- Organizes code into logical function modules
- Demonstrates proper scope management and data flow

## Resources

**MDN Documentation:**
- [Functions](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Functions)
- [Function declarations](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/function)
- [Function expressions](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/function)
- [Arrow functions](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Functions/Arrow_functions)
- [Closures](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Closures)

**W3Schools References:**
- [JavaScript Functions](https://www.w3schools.com/js/js_functions.asp)
- [Function Parameters](https://www.w3schools.com/js/js_function_parameters.asp)
- [Function Return](https://www.w3schools.com/js/js_function_return.asp)
- [Function Scope](https://www.w3schools.com/js/js_scope.asp)

## Assignment Checklist

**Basic Requirements:**
- [ ] Created function declarations with proper syntax
- [ ] Used parameters and arguments effectively
- [ ] Implemented return statements and used returned values
- [ ] Demonstrated understanding of function scope
- [ ] Organized code using multiple related functions

**Intermediate Challenges:**
- [ ] Implemented functions with complex parameter handling
- [ ] Used function expressions and arrow functions appropriately
- [ ] Created higher-order functions and callbacks
- [ ] Applied proper scope management and avoided global pollution
- [ ] Built reusable function libraries

**Advanced Extensions:**
- [ ] Implemented advanced function patterns (closures, memoization)
- [ ] Created modular code organization systems
- [ ] Used functions for complex data processing pipelines
- [ ] Implemented error handling and validation in functions
- [ ] Demonstrated mastery of function composition and design patterns# Week 13: JavaScript Fundamentals - Variables, Data Types, and Basic Operations
