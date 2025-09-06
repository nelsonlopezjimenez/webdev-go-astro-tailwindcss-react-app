```

**Activity:** Create a simple guessing game where the computer picks a random number and the user (simulated) tries to guess it using a while loop.

**Evening Session: Infinite Loop Prevention**
- Common causes of infinite loops
- Debugging loop conditions
- Safe practices for condition-based loops

### Day 3: Arrays and Loop Integration
**Complexity: Beginner to Intermediate**

**Morning Session: Processing Arrays with Loops**
- Introduction to arrays and indexing
- Using loops to access array elements
- Common array processing patterns

**Code Example:**
```javascript
// Introduction to arrays - collections of data
let fruits = ["apple", "banana", "orange", "grape", "kiwi"];
let numbers = [10, 25, 3, 47, 12, 8, 33];
let mixedData = ["John", 25, true, "engineer"];

// Basic array access
console.log(`First fruit: ${fruits[0]}`);     // Beginner: Arrays start at index 0
console.log(`Array length: ${fruits.length}`); // Get total number of items

// Processing arrays with for loops
console.log("All fruits:");
for (let i = 0; i < fruits.length; i++) {
    console.log(`${i + 1}. ${fruits[i]}`); // Display numbered list
}

// Finding maximum value in array
let maxNumber = numbers[0]; // Start with first number
for (let i = 1; i < numbers.length; i++) {
    if (numbers[i] > maxNumber) {
        maxNumber = numbers[i];
    }
}
console.log(`Maximum number: ${maxNumber}`);

// Counting specific items
let longWords = 0;
for (let i = 0; i < fruits.length; i++) {
    if (fruits[i].length > 5) {  // Words longer than 5 characters
        longWords++;
    }
}
console.log(`Number of long fruit names: ${longWords}`);

// Building new arrays based on conditions
let evenNumbers = [];
for (let i = 0; i < numbers.length; i++) {
    if (numbers[i] % 2 === 0) {  // Check if number is even
        evenNumbers.push(numbers[i]); // Beginner: push() adds item to end of array
    }
}
console.log("Even numbers:", evenNumbers);

// Advanced: Finding and modifying array data
let students = [
    { name: "Alice", grade: 85, subject: "Math" },
    { name: "Bob", grade: 92, subject: "Science" },
    { name: "Carol", grade: 78, subject: "Math" },
    { name: "David", grade: 88, subject: "Science" }
];

// Calculate average grade for Math students
let mathGrades = [];
for (let i = 0; i < students.length; i++) {
    if (students[i].subject === "Math") {
        mathGrades.push(students[i].grade);
    }
}

let mathTotal = 0;
for (let i = 0; i < mathGrades.length; i++) {
    mathTotal += mathGrades[i];
}
let mathAverage = mathTotal / mathGrades.length;
console.log(`Average Math grade: ${mathAverage.toFixed(1)}`);

// String processing with loops
let sentence = "JavaScript is awesome for web development";
let words = sentence.split(" "); // Split sentence into array of words
let capitalizedWords = [];

for (let i = 0; i < words.length; i++) {
    // Capitalize first letter of each word
    let capitalizedWord = words[i][0].toUpperCase() + words[i].slice(1);
    capitalizedWords.push(capitalizedWord);
}

let result = capitalizedWords.join(" "); // Join array back into string
console.log(result); // "JavaScript Is Awesome For Web Development"
```

**Activity:** Create a grade book system that processes student data, calculates averages, and identifies top performers using loops.

**Evening Session: Array Methods Preview**
- Introduction to built-in array methods
- When to use loops vs built-in methods
- Performance considerations

### Day 4: Break, Continue, and Loop Control
**Complexity: Intermediate**

**Morning Session: Controlling Loop Execution**
- Using break to exit loops early
- Using continue to skip iterations
- Nested loops and labeled breaks

**Code Example:**
```javascript
// Break statement - exits loop immediately
console.log("Finding first even number:");
let numbersToCheck = [1, 3, 7, 8, 11, 14, 19];

for (let i = 0; i < numbersToCheck.length; i++) {
    console.log(`Checking: ${numbersToCheck[i]}`);
    
    if (numbersToCheck[i] % 2 === 0) {
        console.log(`Found first even number: ${numbersToCheck[i]}`);
        break; // Beginner: break stops the loop immediately
    }
}
console.log("Search complete!");

// Continue statement - skips current iteration
console.log("\nProcessing positive numbers only:");
let allNumbers = [-2, 5, -1, 8, -3, 12, 0, 7];

for (let i = 0; i < allNumbers.length; i++) {
    if (allNumbers[i] <= 0) {
        continue; // Beginner: continue skips to next iteration
    }
    
    console.log(`Processing positive number: ${allNumbers[i]}`);
    // Any code here only runs for positive numbers
}

// Practical example: User input validation
let responses = ["yes", "", "maybe", "no", "", "yes", "invalid"];
let validResponses = [];

for (let i = 0; i < responses.length; i++) {
    // Skip empty responses
    if (responses[i] === "") {
        console.log(`Skipping empty response at position ${i}`);
        continue;
    }
    
    // Stop if we encounter invalid data
    if (responses[i] === "invalid") {
        console.log("Invalid data found, stopping processing");
        break;
    }
    
    validResponses.push(responses[i]);
    console.log(`Added valid response: ${responses[i]}`);
}

console.log("Valid responses:", validResponses);

// Advanced: Nested loops with break and continue
console.log("\nSearching 2D data structure:");
let matrix = [
    [1, 2, 3, 4],
    [5, 6, 7, 8],
    [9, 10, 11, 12]
];

let searchTarget = 7;
let found = false;

// Label for breaking out of nested loops
outerLoop: for (let row = 0; row < matrix.length; row++) {
    for (let col = 0; col < matrix[row].length; col++) {
        console.log(`Checking position [${row}][${col}]: ${matrix[row][col]}`);
        
        if (matrix[row][col] === searchTarget) {
            console.log(`Found ${searchTarget} at position [${row}][${col}]`);
            found = true;
            break outerLoop; // Advanced: labeled break exits both loops
        }
    }
}

if (!found) {
    console.log(`${searchTarget} not found in matrix`);
}

// Practical application: Processing shopping cart
let shoppingCart = [
    { item: "laptop", price: 999, quantity: 1 },
    { item: "mouse", price: 0, quantity: 2 },    // Invalid price
    { item: "keyboard", price: 79, quantity: 1 },
    { item: "monitor", price: 299, quantity: 1 },
    { item: "headphones", price: -50, quantity: 1 } // Invalid price
];

let total = 0;
let itemCount = 0;

for (let i = 0; i < shoppingCart.length; i++) {
    let item = shoppingCart[i];
    
    // Skip items with invalid prices
    if (item.price <= 0) {
        console.log(`Skipping ${item.item} - invalid price: ${item.price}`);
        continue;
    }
    
    // Stop processing if we find extremely expensive items (business rule)
    if (item.price > 2000) {
        console.log(`Stopping - found expensive item: ${item.item} (${item.price})`);
        break;
    }
    
    let itemTotal = item.price * item.quantity;
    total += itemTotal;
    itemCount += item.quantity;
    
    console.log(`${item.item}: ${item.price} x ${item.quantity} = ${itemTotal}`);
}

console.log(`\nCart summary: ${itemCount} items, Total: ${total.toFixed(2)}`);
```

**Activity:** Build a data validation system that processes a list of user records, skipping invalid entries and stopping on critical errors.

**Evening Session: Loop Performance and Optimization**
- Efficient loop patterns
- Avoiding unnecessary iterations
- When to cache array length

### Day 5: Dynamic HTML Generation with Loops
**Complexity: Intermediate to Advanced**

**Morning Session: Creating HTML with JavaScript**
- Building HTML strings with loops
- Introduction to DOM manipulation
- Dynamic list and table generation

**Code Example:**
```javascript
// Creating dynamic HTML content with loops
// This connects JavaScript to your HTML pages!

// Building HTML lists dynamically
let todoItems = [
    { task: "Complete JavaScript homework", priority: "high", done: false },
    { task: "Buy groceries", priority: "medium", done: true },
    { task: "Call dentist", priority: "low", done: false },
    { task: "Finish project proposal", priority: "high", done: false }
];

// Method 1: Building HTML string
let todoHTML = "<ul class='todo-list'>";

for (let i = 0; i < todoItems.length; i++) {
    let item = todoItems[i];
    let cssClass = item.done ? "completed" : "pending";
    let priorityClass = `priority-${item.priority}`;
    
    todoHTML += `
        <li class="${cssClass} ${priorityClass}">
            <input type="checkbox" ${item.done ? "checked" : ""}>
            <span class="task-text">${item.task}</span>
            <span class="priority-badge">${item.priority}</span>
        </li>
    `;
}

todoHTML += "</ul>";
console.log("Generated HTML:", todoHTML);

// Method 2: Direct DOM manipulation (more advanced)
// This would go in your HTML page inside <script> tags
function createTodoList(items, containerId) {
    // Get the container element from your HTML
    let container = document.getElementById(containerId);
    
    // Clear existing content
    container.innerHTML = "";
    
    // Create the list element
    let list = document.createElement("ul");
    list.className = "todo-list";
    
    for (let i = 0; i < items.length; i++) {
        let item = items[i];
        
        // Create list item
        let listItem = document.createElement("li");
        listItem.className = `${item.done ? "completed" : "pending"} priority-${item.priority}`;
        
        // Create checkbox
        let checkbox = document.createElement("input");
        checkbox.type = "checkbox";
        checkbox.checked = item.done;
        
        // Create task text
        let taskText = document.createElement("span");
        taskText.className = "task-text";
        taskText.textContent = item.task;
        
        // Create priority badge
        let priorityBadge = document.createElement("span");
        priorityBadge.className = "priority-badge";
        priorityBadge.textContent = item.priority;
        
        // Assemble the list item
        listItem.appendChild(checkbox);
        listItem.appendChild(taskText);
        listItem.appendChild(priorityBadge);
        
        // Add to list
        list.appendChild(listItem);
    }
    
    // Add list to container
    container.appendChild(list);
}

// Creating dynamic tables
let students = [
    { name: "Alice Johnson", grade: 92, subject: "Mathematics", email: "alice@school.edu" },
    { name: "Bob Smith", grade: 87, subject: "Science", email: "bob@school.edu" },
    { name: "Carol Davis", grade: 95, subject: "English", email: "carol@school.edu" },
    { name: "David Wilson", grade: 89, subject: "History", email: "david@school.edu" }
];

function generateStudentTable(studentData) {
    let tableHTML = `
        <table class="student-table">
            <thead>
                <tr>
                    <th>Name</th>
                    <th>Subject</th>
                    <th>Grade</th>
                    <th>Status</th>
                    <th>Contact</th>
                </tr>
            </thead>
            <tbody>
    `;
    
    for (let i = 0; i < studentData.length; i++) {
        let student = studentData[i];
        let status = student.grade >= 90 ? "Excellent" : 
                    student.grade >= 80 ? "Good" : 
                    student.grade >= 70 ? "Satisfactory" : "Needs Improvement";
        
        let statusClass = student.grade >= 90 ? "status-excellent" : 
                         student.grade >= 80 ? "status-good" : 
                         student.grade >= 70 ? "status-satisfactory" : "status-poor";
        
        tableHTML += `
            <tr>
                <td>${student.name}</td>
                <td>${student.subject}</td>
                <td>${student.grade}%</td>
                <td class="${statusClass}">${status}</td>
                <td><a href="mailto:${student.email}">${student.email}</a></td>
            </tr>
        `;
    }
    
    tableHTML += `
            </tbody>
        </table>
    `;
    
    return tableHTML;
}

// Advanced: Interactive elements with loops
function createInteractiveMenu(menuItems) {
    let menuHTML = '<nav class="dynamic-menu">';
    
    for (let i = 0; i < menuItems.length; i++) {
        let item = menuItems[i];
        
        menuHTML += `
            <div class="menu-item" onclick="handleMenuClick('${item.id}')">
                <i class="icon ${item.icon}"></i>
                <span class="menu-text">${item.text}</span>
                ${item.badge ? `<span class="badge">${item.badge}</span>` : ''}
            </div>
        `;
    }
    
    menuHTML += '</nav>';
    return menuHTML;
}

// Example menu data
let navigationItems = [
    { id: 'home', text: 'Home', icon: 'fa-home', badge: null },
    { id: 'messages', text: 'Messages', icon: 'fa-envelope', badge: '3' },
    { id: 'profile', text: 'Profile', icon: 'fa-user', badge: null },
    { id: 'settings', text: 'Settings', icon: 'fa-cog', badge: '!' }
];

console.log("Generated menu:", createInteractiveMenu(navigationItems));

// Practical example: Building a product catalog
let products = [
    { id: 1, name: "Laptop Pro", price: 1299, category: "Electronics", inStock: true, rating: 4.5 },
    { id: 2, name: "Wireless Mouse", price: 29, category: "Electronics", inStock: true, rating: 4.2 },
    { id: 3, name: "Coffee Mug", price: 15, category: "Kitchen", inStock: false, rating: 4.8 },
    { id: 4, name: "Desk Lamp", price: 89, category: "Furniture", inStock: true, rating: 4.1 }
];

function createProductGrid(products) {
    let gridHTML = '<div class="product-grid">';
    
    for (let i = 0; i < products.length; i++) {
        let product = products[i];
        let stockClass = product.inStock ? "in-stock" : "out-of-stock";
        let stockText = product.inStock ? "In Stock" : "Out of Stock";
        
        // Generate star rating
        let stars = "";
        for (let j = 1; j <= 5; j++) {
            stars += j <= Math.floor(product.rating) ? "★" : "☆";
        }
        
        gridHTML += `
            <div class="product-card ${stockClass}">
                <div class="product-header">
                    <h3>${product.name}</h3>
                    <span class="category">${product.category}</span>
                </div>
                <div class="product-price">${product.price}</div>
                <div class="product-rating">${stars} (${product.rating})</div>
                <div class="stock-status">${stockText}</div>
                <button class="add-to-cart" ${!product.inStock ? 'disabled' : ''}>
                    ${product.inStock ? 'Add to Cart' : 'Notify When Available'}
                </button>
            </div>
        `;
    }
    
    gridHTML += '</div>';
    return gridHTML;
}
```

**Activity:** Create a dynamic dashboard that generates different types of content (charts, lists, cards) based on data arrays using loops.

**Evening Session: Integration with HTML/CSS**
- Connecting JavaScript loops to webpage elements
- CSS classes for dynamic styling
- Event handling for dynamic content

## Hands-on Exercises

### Exercise 1: Data Processing Dashboard
Create a comprehensive data analysis tool that:
- Processes multiple datasets using different loop types
- Generates summary statistics and reports
- Creates visual representations using HTML generation
- Implements filtering and search functionality

### Exercise 2: Interactive Game Engine
Build a simple game system featuring:
- Game loop using while loops for continuous play
- Grid-based game board using nested for loops
- Dynamic scoring and level progression
- User input processing and validation

### Exercise 3: Content Management System
Develop a basic CMS that:
- Processes and displays articles using loops
- Implements pagination and filtering
- Generates navigation menus dynamically
- Handles user interactions and updates

## Resources

**MDN Documentation:**
- [for statement](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/for)
- [while statement](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/while)
- [break statement](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/break)
- [continue statement](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/continue)
- [Arrays](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array)

**W3Schools References:**
- [JavaScript For Loop](https://www.w3schools.com/js/js_loop_for.asp)
- [JavaScript While Loop](https://www.w3schools.com/js/js_loop_while.asp)
- [JavaScript Arrays](https://www.w3schools.com/js/js_arrays.asp)
- [JavaScript Break and Continue](https://www.w3schools.com/js/js_break.asp)

## Assignment Checklist

**Basic Requirements:**
- [ ] Implemented for loops with proper initialization, condition, and increment
- [ ] Used while loops for condition-based iteration
- [ ] Processed arrays using loops effectively
- [ ] Applied break and continue statements appropriately
- [ ] Generated dynamic HTML content using loops

**Intermediate Challenges:**
- [ ] Chose appropriate loop types for different scenarios
- [ ] Implemented nested loops for complex data processing
- [ ] Created reusable functions that incorporate loops
- [ ] Handled edge cases and prevented infinite loops
- [ ] Integrated loops with conditional logic effectively

**Advanced Extensions:**
- [ ] Optimized loop performance for large datasets
- [ ] Implemented complex data processing algorithms
- [ ] Created interactive web elements using dynamic generation
- [ ] Used labeled breaks and advanced loop control
- [ ] Built complete applications that rely heavily on iteration

---

# Week 16: Functions - Organizing and Reusing Code

```yaml
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
```

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

```yaml
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
```

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

---

# Week 14: Control Flow - Conditionals and Decision Making

```yaml
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
```

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

# Week 15: Loops and Iteration - Repeating Code Efficiently

```yaml
title: "Loops and Iteration - Repeating Code Efficiently"
description: "Master for loops, while loops, and iteration techniques for processing data and creating dynamic content"
week: 15
section: 2
prerequisites: ["Conditional statements", "Variables", "Basic operators"]
objectives:
  - "Implement for loops for counting and iteration"
  - "Use while and do-while loops for condition-based repetition"
  - "Process arrays and strings using loops"
  - "Control loop execution with break and continue"
  - "Create dynamic HTML content using JavaScript loops"
```

## Learning Objectives

By the end of this week, students will be able to:
- Write efficient for loops for various iteration patterns
- Choose appropriate loop types for different scenarios
- Process collections of data using iteration
- Control loop flow with break and continue statements
- Generate dynamic content and HTML using loops
- Avoid infinite loops and debug loop-related issues

## Daily Activities

### Day 1: For Loops Fundamentals
**Complexity: Beginner**

**Morning Session: Basic For Loop Structure**
- Understanding when and why to use loops
- For loop syntax and components
- Loop counter variables and iteration

**Code Example:**
```javascript
// Basic for loop structure
// for (initialization; condition; increment/decrement)

// Simple counting loop
for (let i = 1; i <= 5; i++) {
    console.log(`Count: ${i}`); // Beginner: Runs 5 times, i goes from 1 to 5
}
// Output: Count: 1, Count: 2, Count: 3, Count: 4, Count: 5

// Counting backwards
for (let i = 10; i >= 1; i--) {
    console.log(`Countdown: ${i}`);
}
// Outputs countdown from 10 to 1

// Different step sizes
for (let i = 0; i <= 20; i += 2) {
    console.log(`Even number: ${i}`); // Beginner: i increases by 2 each time
}
// Outputs: 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20

// Processing strings with loops
let message = "Hello";
for (let i = 0; i < message.length; i++) {
    console.log(`Character ${i}: ${message[i]}`); // Beginner: Access each character
}
// H-e-l-l-o each on separate lines

// Building strings with loops
let stars = "";
for (let i = 1; i <= 5; i++) {
    stars += "*";  // Add a star each iteration
    console.log(stars);
}
// Outputs: *, **, ***, ****, *****

// Advanced: Nested loops for patterns
console.log("Creating a multiplication table:");
for (let row = 1; row <= 5; row++) {
    let line = "";
    for (let col = 1; col <= 5; col++) {
        line += `${row * col}\t`;  // \t adds a tab space
    }
    console.log(line);
}
```

**Activity:** Create a program that generates different visual patterns using nested for loops (stars, numbers, pyramids).

**Evening Session: Loop Planning and Logic**
- Breaking down problems into loop components
- Determining start, end, and step values
- Common loop patterns and applications

### Day 2: While Loops and Condition-Based Iteration
**Complexity: Beginner to Intermediate**

**Morning Session: While and Do-While Loops**
- When to use while vs for loops
- While loop syntax and flow
- Do-while loops for guaranteed execution

**Code Example:**
```javascript
// Basic while loop - runs as long as condition is true
let count = 1;
while (count <= 5) {
    console.log(`While count: ${count}`);
    count++; // Beginner: IMPORTANT - don't forget to update the condition variable!
}

// While loop with user input simulation
let password = "";
let attempts = 0;
let maxAttempts = 3;

while (password !== "secret123" && attempts < maxAttempts) {
    // In real app, this would get user input
    password = attempts === 0 ? "wrong1" : attempts === 1 ? "wrong2" : "secret123";
    attempts++;
    
    if (password === "secret123") {
        console