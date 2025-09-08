---
title: "JavaScript Fundamentals - Variables, Data Types, and Basic Operations"
description: "Introduction to JavaScript syntax, variables, primitive data types, and basic operations"
week: 13
section: 2
prerequisites: ["HTML structure", "CSS styling", "Basic understanding of web pages"]
objectives:
  - "Understand what JavaScript is and how it enhances web pages"
  - "Declare and use variables with let, const, and var"
  - "Work with primitive data types: strings, numbers, booleans"
  - "Perform basic operations and type conversion"
  - "Debug JavaScript using browser developer tools"
---
# Week 13: JavaScript Fundamentals - Variables, Data Types, and Basic Operations
[TOC]

<!-- Or this: -->
[[TOC]]

## Learning Objectives

By the end of this week, students will be able to:
- Explain JavaScript's role in web development and the DOM
- Declare variables using appropriate keywords (let, const, var)
- Identify and work with JavaScript's primitive data types
- Perform arithmetic, string, and comparison operations
- Use browser developer tools to debug and test JavaScript code
- Write clean, commented JavaScript code following best practices

## Daily Activities

### Day 1: Introduction to JavaScript and Setting Up
**Complexity: Beginner**

**Morning Session: What is JavaScript?**
- JavaScript's role in web development (behavior layer)
- Difference between client-side and server-side JavaScript
- How JavaScript interacts with HTML and CSS
- Setting up development environment

**Code Example:**
```javascript
// Your first JavaScript program
// This code adds interactivity to a webpage

// Display a welcome message
console.log("Welcome to JavaScript!"); // Beginner: console.log prints messages to browser console

// Change webpage content (we'll learn more about this later)
document.getElementById("welcome").textContent = "Hello from JavaScript!";
/* 
Advanced concept preview: document.getElementById() finds HTML elements
This is called DOM manipulation - we'll cover this in detail in Week 15
*/
```

**Activity:** Create a simple HTML page with a `<p>` element, add a `<script>` tag, and make JavaScript change the paragraph text.

**Evening Session: Browser Developer Tools**
- Opening and navigating Chrome/Firefox DevTools
- Using the Console tab to run JavaScript
- Understanding error messages

### Day 2: Variables and Declaration Keywords
**Complexity: Beginner to Intermediate**

**Morning Session: Understanding Variables**
- What variables are and why we need them
- Naming conventions and best practices
- Introduction to let, const, and var

**Code Example:**
```javascript
// Variable declarations - storing information for later use

// let - for values that can change
let userName = "Alex";           // Beginner: let creates a variable that can be updated
let userAge = 25;
userName = "Jordan";             // This is allowed - we can change let variables

// const - for values that should never change  
const PI = 3.14159;             // Beginner: const creates a variable that cannot be changed
const siteName = "My Web App";
// PI = 3.14;                   // This would cause an error!

// var - older way (we'll mostly use let and const)
var oldVariable = "legacy code"; // Advanced: var has different scoping rules, prefer let/const

// Good variable naming examples
let firstName = "Emma";          // camelCase for JavaScript
let isLoggedIn = true;          // Boolean variables often start with "is" or "has"
let totalItemCount = 0;         // Descriptive names help others understand your code
```

**Activity:** Create variables for a user profile (name, age, email, isActive) and practice changing appropriate ones.

**Evening Session: Variable Scope Preview**
- Brief introduction to block scope
- Why we prefer let and const over var

### Day 3: Data Types - Strings
**Complexity: Beginner to Intermediate**

**Morning Session: Working with Strings**
- Creating strings with single, double, and template literals
- Common string methods and properties
- String concatenation vs template literals

**Code Example:**
```javascript
// Different ways to create strings
let singleQuote = 'Hello World';      // Beginner: Strings can use single quotes
let doubleQuote = "Hello World";      // Or double quotes (choose one style and be consistent)
let templateLiteral = `Hello World`;  // Template literals use backticks

// String properties and methods
let message = "JavaScript is awesome!";
console.log(message.length);          // Beginner: .length tells us how many characters
console.log(message.toUpperCase());   // Beginner: .toUpperCase() makes all letters capital
console.log(message.toLowerCase());   // Beginner: .toLowerCase() makes all letters lowercase

// Template literals for dynamic strings
let name = "Sarah";
let age = 28;
let greeting = `Hello, my name is ${name} and I'm ${age} years old.`;
// Beginner: ${} lets us put variables inside strings with template literals
console.log(greeting);

// String methods for manipulation
let email = "  user@example.com  ";
let cleanEmail = email.trim();        // Advanced: .trim() removes extra spaces
let domain = email.split("@")[1];     // Advanced: .split() breaks string into array
```

**Activity:** Create a personal introduction using template literals that combines multiple variables.

**Evening Session: String Exercises**
- Practice with string methods
- Building dynamic messages
- Introduction to escape characters

### Day 4: Data Types - Numbers and Booleans
**Complexity: Beginner to Intermediate**

**Morning Session: Working with Numbers**
- Number types in JavaScript
- Arithmetic operations
- Math object basics

**Code Example:**
```javascript
// Numbers in JavaScript
let wholeNumber = 42;              // Beginner: Integers (whole numbers)
let decimal = 3.14159;             // Beginner: Floating-point numbers (decimals)
let negative = -10;                // Negative numbers

// Arithmetic operations
let sum = 10 + 5;                  // Addition: 15
let difference = 10 - 5;           // Subtraction: 5  
let product = 10 * 5;              // Multiplication: 50
let quotient = 10 / 5;             // Division: 2
let remainder = 10 % 3;            // Modulus (remainder): 1

// Math object for complex operations
let randomNumber = Math.random();   // Advanced: Random number between 0 and 1
let rounded = Math.round(3.7);      // Advanced: Rounds to nearest integer: 4
let maximum = Math.max(5, 10, 3);   // Advanced: Finds largest number: 10

// Boolean values
let isComplete = true;              // Beginner: true or false values
let isVisible = false;
let hasPermission = true;

// Boolean operations
let canEdit = isComplete && hasPermission;  // AND: both must be true
let showWarning = !isVisible;               // NOT: opposite of isVisible
let needsAttention = !isComplete || hasErrors; // OR: either can be true
```

**Activity:** Create a simple calculator that performs basic arithmetic operations and displays results.

**Evening Session: Type Coercion Preview**
- How JavaScript converts between types
- Common pitfalls with automatic conversion

### Day 5: Operations and Type Conversion
**Complexity: Intermediate**

**Morning Session: Comparison and Logical Operations**
- Comparison operators
- Strict vs loose equality
- Logical operators in depth

**Code Example:**
```javascript
// Comparison operators
let a = 10;
let b = "10";
let c = 20;

// Equality comparisons
console.log(a == b);    // true - loose equality (converts types)
console.log(a === b);   // false - strict equality (no type conversion)
console.log(a != b);    // false - loose inequality  
console.log(a !== b);   // true - strict inequality

// Beginner: Always use === and !== to avoid surprises!

// Other comparisons
console.log(a < c);     // true
console.log(a >= 10);   // true
console.log(b <= "5");  // false (string comparison)

// Logical operators in detail
let user = {
  isLoggedIn: true,
  hasSubscription: false,
  credits: 5
};

// AND (&&) - all conditions must be true
let canAccessPremium = user.isLoggedIn && user.hasSubscription;

// OR (||) - at least one condition must be true  
let canViewContent = user.isLoggedIn || user.credits > 0;

// NOT (!) - reverses the boolean value
let needsLogin = !user.isLoggedIn;

// Type conversion examples
let numberFromString = Number("123");      // Convert string to number: 123
let stringFromNumber = String(456);        // Convert number to string: "456"
let booleanFromString = Boolean("hello");  // Convert to boolean: true (non-empty strings are true)

// Advanced: Falsy values in JavaScript
let falsyExamples = [false, 0, "", null, undefined, NaN];
// These all convert to false in boolean contexts
```

**Activity:** Build a user permission checker that uses logical operations to determine what a user can access.

**Evening Session: Debugging Practice**
- Using console.log strategically
- Understanding common error messages
- Reading stack traces

## Hands-on Exercises

### Exercise 1: Personal Profile Creator
Create a webpage that uses JavaScript to display dynamic user information:
- Declare variables for name, age, occupation, and hobbies
- Use template literals to create formatted output
- Include at least 3 different data types

### Exercise 2: Simple Calculator Interface  
Build on your HTML/CSS knowledge to create:
- An HTML form with number inputs
- JavaScript that performs calculations
- Display results dynamically on the page

### Exercise 3: Type Experiment Lab
Create experiments to understand JavaScript types:
- Test different variable declarations
- Explore type conversion scenarios
- Document surprising results with comments

## Resources

**MDN Documentation:**
- [JavaScript Basics](https://developer.mozilla.org/en-US/docs/Learn/Getting_started_with_the_web/JavaScript_basics)
- [Variables](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Grammar_and_types#Variables)
- [Data Types](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Data_structures)
- [Operators](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Expressions_and_Operators)

**W3Schools References:**
- [JavaScript Variables](https://www.w3schools.com/js/js_variables.asp)
- [JavaScript Data Types](https://www.w3schools.com/js/js_datatypes.asp)
- [JavaScript Operators](https://www.w3schools.com/js/js_operators.asp)

## Assignment Checklist

**Basic Requirements:**
- [ ] Created variables using let, const, and var appropriately
- [ ] Demonstrated understanding of strings, numbers, and booleans
- [ ] Used template literals for dynamic string creation
- [ ] Performed arithmetic and logical operations
- [ ] Used console.log for debugging and output

**Intermediate Challenges:**
- [ ] Implemented proper variable naming conventions
- [ ] Used strict equality (===) consistently
- [ ] Handled type conversion explicitly
- [ ] Created meaningful comments explaining code purpose

**Advanced Extensions:**
- [ ] Explored scope differences between let, const, and var
- [ ] Implemented complex logical operations
- [ ] Used Math object methods effectively
- [ ] Debugged code using browser developer tools


