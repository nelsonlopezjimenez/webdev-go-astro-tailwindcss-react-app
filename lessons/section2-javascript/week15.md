---
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
----------------------------------------------------------
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