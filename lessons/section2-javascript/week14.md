---
title: "Control Flow - Conditionals and Decision Making"
description: "Master if/else statements, switch cases, and ternary operators for dynamic program behavior"
week: 14
section: 2
prerequisites: ["JavaScript variables", "Data types", "Comparison operators"]
objectives:
  - "Implement conditional logic using if/else statements"
  - "Use switch statements for multiple condition scenarios"
  - "Apply ternary operators for concise conditional expressions"
  - "Combine logical operators for complex conditions"
  - "Handle user input and create interactive experiences"
---

# Week 14: Control Flow - Conditionals and Decision Making

## Learning Objectives

By the end of this week, students will be able to:
- Write conditional statements to control program flow
- Choose appropriate conditional structures for different scenarios
- Create interactive web pages that respond to user actions
- Validate user input and provide appropriate feedback
- Combine multiple conditions using logical operators
- Debug conditional logic and fix common mistakes

## Daily Activities

### Day 1: If/Else Statements Foundation
**Complexity: Beginner**

**Morning Session: Basic If/Else Structure**
- Understanding program flow control
- If statement syntax and logic
- If/else and if/else if chains

**Code Example:**
```javascript
// Basic if statement structure
let userAge = 18;

if (userAge >= 18) {
    console.log("You can vote!"); // Beginner: Code inside {} runs only if condition is true
}

// If/else for two possibilities
let temperature = 75;

if (temperature > 80) {
    console.log("It's hot outside!");
} else {
    console.log("It's not too hot today.");
}

// If/else if/else for multiple conditions
let score = 85;

if (score >= 90) {
    console.log("Grade: A - Excellent work!");
} else if (score >= 80) {
    console.log("Grade: B - Good job!");     // This condition matches our score
} else if (score >= 70) {
    console.log("Grade: C - Satisfactory");
} else if (score >= 60) {
    console.log("Grade: D - Needs improvement");
} else {
    console.log("Grade: F - Please see instructor");
}

// Advanced: Nested conditions
let weather = "sunny";
let hasUmbrella = false;

if (weather === "rainy") {
    if (hasUmbrella) {
        console.log("Perfect! You're prepared for the rain.");
    } else {
        console.log("You might get wet - consider bringing an umbrella.");
    }
} else {
    console.log("Enjoy the nice weather!");
}
```

**Activity:** Create a simple age verification system that provides different messages for different age ranges.

**Evening Session: Common Conditional Patterns**
- Checking for null/undefined values
- Validating user input
- Setting default values

### Day 2: Complex Conditions and Logical Operators
**Complexity: Beginner to Intermediate**

**Morning Session: Combining Conditions**
- Using && (AND) in conditionals
- Using || (OR) in conditionals  
- Using ! (NOT) in conditionals
- Order of operations in logical expressions

**Code Example:**
```javascript
// Combining conditions with logical operators
let userAge = 25;
let hasLicense = true;
let hasInsurance = true;
let hasVehicle = false;

// AND operator - all conditions must be true
if (userAge >= 18 && hasLicense && hasInsurance) {
    console.log("You meet the basic driving requirements!");
}

// OR operator - at least one condition must be true
if (hasVehicle || hasLicense) {
    console.log("You have some form of transportation option.");
}

// Complex combinations
if ((userAge >= 18 && hasLicense) || (userAge >= 16 && hasParentalPermission)) {
    console.log("You can drive under current conditions.");
}

// NOT operator examples
let isLoggedIn = false;
let isGuest = true;

if (!isLoggedIn) {
    console.log("Please log in to continue.");  // Beginner: !isLoggedIn means "if NOT logged in"
}

if (!isGuest && isLoggedIn) {
    console.log("Welcome back, registered user!");
}

// Real-world example: Form validation
let email = "user@example.com";
let password = "securePass123";
let agreeToTerms = true;

if (email.includes("@") && password.length >= 8 && agreeToTerms) {
    console.log("Registration successful!");
} else {
    // We'll learn how to give specific error messages next
    console.log("Please check your registration details.");
}

// Advanced: Short-circuit evaluation
let user = null;
// This safely checks if user exists before accessing properties
if (user && user.isActive) {  
    console.log("Active user found");
}
```

**Activity:** Build a movie recommendation system that considers age, genre preferences, and viewing time.

**Evening Session: Debugging Conditional Logic**
- Common mistakes with logical operators
- Using parentheses for clarity
- Testing edge cases

### Day 3: Switch Statements and Pattern Matching
**Complexity: Beginner to Intermediate**

**Morning Session: Switch Statement Basics**
- When to use switch vs if/else
- Switch syntax and break statements
- Default cases and fall-through behavior

**Code Example:**
```javascript
// Switch statement for multiple specific values
let dayOfWeek = "Monday";

switch (dayOfWeek) {
    case "Monday":
        console.log("Start of the work week!");
        break;  // Beginner: break prevents code from continuing to next case
    case "Tuesday":
        console.log("Tuesday motivation!");
        break;
    case "Wednesday":
        console.log("Hump day!");
        break;
    case "Thursday":
        console.log("Almost Friday!");
        break;
    case "Friday":
        console.log("TGIF!");
        break;
    case "Saturday":
    case "Sunday":  // Beginner: Multiple cases can share the same code
        console.log("Weekend vibes!");
        break;
    default:  // Beginner: default runs if no cases match
        console.log("That's not a valid day.");
}

// Switch with numbers
let month = 3;
let season;

switch (month) {
    case 12:
    case 1:
    case 2:
        season = "Winter";
        break;
    case 3:
    case 4:
    case 5:
        season = "Spring";  // March falls here
        break;
    case 6:
    case 7:
    case 8:
        season = "Summer";
        break;
    case 9:
    case 10:
    case 11:
        season = "Fall";
        break;
    default:
        season = "Unknown";
}

console.log(`Month ${month} is in ${season}`);

// Advanced: Switch with expressions
let operation = "multiply";
let a = 10;
let b = 5;
let result;

switch (operation) {
    case "add":
        result = a + b;
        break;
    case "subtract":
        result = a - b;
        break;
    case "multiply":
        result = a * b;  // This case executes
        break;
    case "divide":
        result = b !== 0 ? a / b : "Cannot divide by zero";
        break;
    default:
        result = "Unknown operation";
}

console.log(`${a} ${operation} ${b} = ${result}`);
```

**Activity:** Create a restaurant menu system that shows different options based on meal time (breakfast, lunch, dinner).

**Evening Session: Switch vs If/Else Decision Making**
- Performance considerations
- Readability and maintenance
- When each approach is preferred

### Day 4: Ternary Operators and Concise Conditionals
**Complexity: Intermediate**

**Morning Session: Ternary Operator Basics**
- Syntax: condition ? valueIfTrue : valueIfFalse
- When to use ternary vs if/else
- Chaining ternary operators

**Code Example:**
```javascript
// Basic ternary operator syntax
let age = 20;
let canVote = age >= 18 ? "Yes" : "No";  // Beginner: Short way to write if/else
console.log(`Can vote: ${canVote}`);

// Ternary vs if/else comparison
// Using if/else:
let temperature = 75;
let clothing;
if (temperature > 80) {
    clothing = "shorts";
} else {
    clothing = "pants";
}

// Same logic with ternary:
let clothing2 = temperature > 80 ? "shorts" : "pants";  // Much shorter!

// Ternary in template literals
let user = { name: "Alice", isLoggedIn: true };
let greeting = `Hello, ${user.isLoggedIn ? user.name : "Guest"}!`;

// Multiple ternary operators (chaining)
let score = 85;
let grade = score >= 90 ? "A" : 
           score >= 80 ? "B" : 
           score >= 70 ? "C" : 
           score >= 60 ? "D" : "F";

// Advanced: Ternary with function calls
function showWelcomeMessage() {
    return "Welcome to our site!";
}

function showLoginPrompt() {
    return "Please log in to continue.";
}

let isAuthenticated = false;
let message = isAuthenticated ? showWelcomeMessage() : showLoginPrompt();

// Ternary for setting CSS classes (preview of DOM manipulation)
let isActive = true;
let buttonClass = isActive ? "btn-active" : "btn-inactive";

// Complex example: Shopping cart logic
let items = 3;
let subtotal = 45.99;
let hasCoupon = true;

let shipping = items > 0 ? (subtotal > 50 ? 0 : 5.99) : 0;
let discount = hasCoupon ? (subtotal * 0.1) : 0;
let total = subtotal + shipping - discount;

console.log(`Items: ${items}`);
console.log(`Subtotal: $${subtotal}`);
console.log(`Shipping: $${shipping}`);
console.log(`Discount: $${discount}`);
console.log(`Total: $${total.toFixed(2)}`);
```

**Activity:** Build a dynamic pricing calculator that applies different rules based on customer type, order size, and promotions.

**Evening Session: Readability and Best Practices**
- When ternary operators improve code
- When they make code harder to read
- Formatting complex ternary expressions

### Day 5: Interactive User Input and Validation
**Complexity: Intermediate to Advanced**

**Morning Session: Handling User Input**
- Getting input from HTML forms
- Using prompt() for simple input (development/testing)
- Validating and sanitizing user input

**Code Example:**
```javascript
// Simple input validation functions
function validateEmail(email) {
    // Beginner: Check if email contains @ symbol and has reasonable length
    if (email.includes("@") && email.length > 5) {
        return true;
    }
    return false;
}

function validatePassword(password) {
    // Check multiple password requirements
    if (password.length < 8) {
        return { valid: false, message: "Password must be at least 8 characters" };
    }
    if (!/[A-Z]/.test(password)) {  // Advanced: Regular expression for uppercase
        return { valid: false, message: "Password must contain an uppercase letter" };
    }
    if (!/[0-9]/.test(password)) {  // Advanced: Regular expression for numbers
        return { valid: false, message: "Password must contain a number" };
    }
    return { valid: true, message: "Password is strong" };
}

// User registration validation
function validateRegistration(userData) {
    let errors = [];
    
    // Name validation
    if (!userData.name || userData.name.trim().length < 2) {
        errors.push("Name must be at least 2 characters long");
    }
    
    // Age validation
    if (userData.age < 13 || userData.age > 120) {
        errors.push("Age must be between 13 and 120");
    }
    
    // Email validation
    if (!validateEmail(userData.email)) {
        errors.push("Please enter a valid email address");
    }
    
    // Password validation
    let passwordCheck = validatePassword(userData.password);
    if (!passwordCheck.valid) {
        errors.push(passwordCheck.message);
    }
    
    // Return validation result
    if (errors.length === 0) {
        return { success: true, message: "Registration successful!" };
    } else {
        return { success: false, errors: errors };
    }
}

// Example usage
let newUser = {
    name: "John Doe",
    age: 25,
    email: "john@example.com",
    password: "SecurePass123"
};

let validationResult = validateRegistration(newUser);

if (validationResult.success) {
    console.log(validationResult.message);
} else {
    console.log("Registration failed:");
    validationResult.errors.forEach(error => {
        console.log(`- ${error}`);
    });
}

// Advanced: Dynamic form validation
function createUserProfile(formData) {
    // Set defaults for optional fields
    let profile = {
        name: formData.name || "Anonymous",
        theme: formData.theme || "light",
        notifications: formData.notifications !== undefined ? formData.notifications : true,
        level: formData.experience || "beginner"
    };
    
    // Conditional profile setup based on user choices
    if (profile.level === "advanced") {
        profile.features = ["dark-mode", "advanced-tools", "beta-features"];
    } else if (profile.level === "intermediate") {
        profile.features = ["dark-mode", "intermediate-tools"];
    } else {
        profile.features = ["guided-tour", "help-tooltips"];
    }
    
    return profile;
}
```

**Activity:** Create a comprehensive user registration system with multiple validation rules and helpful error messages.

**Evening Session: Error Handling and User Experience**
- Providing helpful error messages
- Progressive validation
- Accessibility considerations

## Hands-on Exercises

### Exercise 1: Smart Home Control System
Create a control panel simulation that:
- Uses if/else statements for device status
- Implements switch statements for different room controls
- Includes safety checks and validation
- Provides user-friendly feedback

### Exercise 2: Grade Calculator with Multiple Conditions
Build a grade calculator that:
- Takes multiple assignment scores
- Applies different grading scales based on course type
- Handles extra credit and late penalties
- Uses ternary operators for concise logic

### Exercise 3: Interactive Quiz Application
Develop a quiz system featuring:
- Multiple question types (true/false, multiple choice)
- Complex scoring logic with conditionals
- User input validation
- Dynamic feedback based on performance

## Resources

**MDN Documentation:**
- [if...else](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/if...else)
- [switch](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/switch)
- [Conditional (ternary) operator](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Conditional_Operator)
- [Logical operators](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators#binary_logical_operators)

**W3Schools References:**
- [JavaScript if else](https://www.w3schools.com/js/js_if_else.asp)
- [JavaScript switch](https://www.w3schools.com/js/js_switch.asp)
- [JavaScript comparisons](https://www.w3schools.com/js/js_comparisons.asp)

## Assignment Checklist

**Basic Requirements:**
- [ ] Implemented if/else statements correctly
- [ ] Used switch statements for appropriate scenarios
- [ ] Applied logical operators (&&, ||, !) in conditions
- [ ] Created ternary operators for simple conditionals
- [ ] Handled user input with basic validation

**Intermediate Challenges:**
- [ ] Combined multiple conditions effectively
- [ ] Chose appropriate conditional structures for different problems
- [ ] Implemented comprehensive input validation
- [ ] Used nested conditionals appropriately
- [ ] Created helpful error messages

**Advanced Extensions:**
- [ ] Implemented complex validation logic
- [ ] Used short-circuit evaluation effectively
- [ ] Created reusable validation functions
- [ ] Handled edge cases and error scenarios
- [ ] Optimized conditional logic for performance

---

