---
title: "Arrays and Modern JavaScript Methods"
description: "Master array manipulation, iteration methods, and modern JavaScript features for data processing"
week: 18
section: 2
prerequisites: ["Objects", "Functions", "Loops"]
objectives:
  - "Master array creation, manipulation, and transformation methods"
  - "Use modern array methods (map, filter, reduce, forEach)"
  - "Implement functional programming concepts with arrays"
  - "Handle complex data structures and nested arrays"
  - "Apply array methods for real-world data processing"
---

# Week 18: Arrays and Modern JavaScript Methods

## Learning Objectives

By the end of this week, students will be able to:
- Create and manipulate arrays using modern JavaScript methods
- Apply functional programming principles with array methods
- Transform and filter data efficiently using built-in methods
- Work with complex data structures and nested arrays
- Implement data processing pipelines using method chaining
- Choose appropriate array methods for different scenarios

## Daily Activities

### Day 1: Array Fundamentals and Basic Methods
**Complexity: Beginner**

**Morning Session: Array Creation and Basic Operations**
- Array creation methods
- Adding and removing elements
- Array indexing and length

**Code Example:**
```javascript
// Array creation methods
let fruits = ["apple", "banana", "orange"];           // Array literal
let numbers = new Array(1, 2, 3, 4, 5);             // Array constructor
let empty = [];                                       // Empty array
let mixed = ["text", 42, true, {name: "John"}];     // Mixed data types

// Creating arrays with specific length
let tenItems = new Array(10);                        // Array with 10 undefined elements
let filledArray = Array(5).fill(0);                 // [0, 0, 0, 0, 0]
let rangeArray = Array.from({length: 5}, (_, i) => i + 1); // [1, 2, 3, 4, 5]

console.log("Fruits:", fruits);
console.log("Range array:", rangeArray);

// Basic array properties
console.log("Array length:", fruits.length);
console.log("First fruit:", fruits[0]);
console.log("Last fruit:", fruits[fruits.length - 1]);

// Adding elements
fruits.push("grape");                    // Add to end: ["apple", "banana", "orange", "grape"]
fruits.unshift("strawberry");            // Add to beginning: ["strawberry", "apple", "banana", "orange", "grape"]
console.log("After adding:", fruits);

// Removing elements
let lastFruit = fruits.pop();            // Remove from end: "grape"
let firstFruit = fruits.shift();         // Remove from beginning: "strawberry"
console.log("Removed:", lastFruit, firstFruit);
console.log("After removing:", fruits);

// Splice method - add/remove at specific position
let vegetables = ["carrot", "broccoli", "spinach", "kale"];
vegetables.splice(2, 1, "lettuce", "tomato"); // Remove 1 at index 2, add "lettuce" and "tomato"
console.log("After splice:", vegetables); // ["carrot", "broccoli", "lettuce", "tomato", "kale"]

// Slice method - extract portion (non-destructive)
let citrus = fruits.slice(1, 3);         // Extract from index 1 to 3 (not including 3)
console.log("Citrus fruits:", citrus);
console.log("Original fruits unchanged:", fruits);

// Finding elements
let inventory = ["laptop", "mouse", "keyboard", "monitor", "laptop"];
console.log("Index of mouse:", inventory.indexOf("mouse"));           // 1
console.log("Last index of laptop:", inventory.lastIndexOf("laptop")); // 4
console.log("Includes keyboard:", inventory.includes("keyboard"));     // true

// Array concatenation and joining
let array1 = [1, 2, 3];
let array2 = [4, 5, 6];
let combined = array1.concat(array2);    // [1, 2, 3, 4, 5, 6]
let withSpread = [...array1, ...array2]; // Modern way using spread operator

console.log("Combined arrays:", combined);
console.log("With spread operator:", withSpread);

// Converting array to string
let words = ["Hello", "world", "from", "JavaScript"];
let sentence = words.join(" ");          // "Hello world from JavaScript"
let csvFormat = words.join(",");         // "Hello,world,from,JavaScript"
console.log("Joined as sentence:", sentence);

// Reversing and sorting
let letters = ["c", "a", "b", "d"];
let reversed = [...letters].reverse();   // Create copy then reverse
let sorted = [...letters].sort();       // Create copy then sort
console.log("Original:", letters);
console.log("Reversed copy:", reversed);
console.log("Sorted copy:", sorted);

// Working with arrays of objects
let students = [
    { name: "Alice", grade: 85, subject: "Math" },
    { name: "Bob", grade: 92, subject: "Science" },
    { name: "Carol", grade: 78, subject: "English" }
];

// Adding new student
students.push({ name: "David", grade: 88, subject: "History" });

// Finding student by name
let bobIndex = students.findIndex(student => student.name === "Bob");
console.log("Bob's index:", bobIndex);

// Updating student grade
if (bobIndex !== -1) {
    students[bobIndex].grade = 95;
}

console.log("Updated students:", students);

// Array destructuring
let colors = ["red", "green", "blue", "yellow"];
let [primary1, primary2, primary3, ...others] = colors;
console.log("Primary colors:", primary1, primary2, primary3);
console.log("Other colors:", others);

// Swapping variables using destructuring
let a = 10, b = 20;
[a, b] = [b, a];
console.log("After swap - a:", a, "b:", b);
```

**Activity:** Create an inventory management system using basic array methods to add, remove, and organize products.

**Evening Session: Array Searching and Basic Iteration**
- Array search methods
- Basic iteration patterns
- Array copying techniques

### Day 2: forEach, map, and Array Transformation
**Complexity: Beginner to Intermediate**

**Morning Session: Modern Array Iteration**
- forEach method for iteration
- map method for transformation
- Differences between forEach and map

**Code Example:**
```javascript
// forEach - execute function for each array element
let numbers = [1, 2, 3, 4, 5];

// Basic forEach usage
console.log("Using forEach:");
numbers.forEach(function(number, index) {
    console.log(`Index ${index}: ${number}`);
});

// forEach with arrow function (modern syntax)
numbers.forEach((number, index) => {
    console.log(`Number: ${number}, Square: ${number * number}`);
});

// forEach with objects
let products = [
    { name: "Laptop", price: 999, category: "Electronics" },
    { name: "Book", price: 15, category: "Education" },
    { name: "Headphones", price: 199, category: "Electronics" }
];

products.forEach(product => {
    console.log(`${product.name}: ${product.price} (${product.category})`);
});

// map - transform each element and return new array
let doubled = numbers.map(num => num * 2);
console.log("Original numbers:", numbers);      // [1, 2, 3, 4, 5]
console.log("Doubled numbers:", doubled);       // [2, 4, 6, 8, 10]

// map with more complex transformations
let fahrenheitTemps = [32, 68, 86, 104, 122];
let celsiusTemps = fahrenheitTemps.map(temp => Math.round((temp - 32) * 5/9));
console.log("Fahrenheit:", fahrenheitTemps);
console.log("Celsius:", celsiusTemps);

// map with objects - extracting properties
let productNames = products.map(product => product.name);
let productPrices = products.map(product => product.price);
console.log("Product names:", productNames);
console.log("Product prices:", productPrices);

// map to create new objects
let discountedProducts = products.map(product => ({
    name: product.name,
    originalPrice: product.price,
    salePrice: Math.round(product.price * 0.8), // 20% discount
    savings: Math.round(product.price * 0.2)
}));

console.log("Discounted products:", discountedProducts);

// map vs forEach - important difference!
let forEachResult = numbers.forEach(num => num * 2); // forEach returns undefined
let mapResult = numbers.map(num => num * 2);         // map returns new array

console.log("forEach result:", forEachResult);       // undefined
console.log("map result:", mapResult);               // [2, 4, 6, 8, 10]

// Practical example: Processing user data
let users = [
    { firstName: "John", lastName: "Doe", age: 30, email: "john@example.com" },
    { firstName: "Jane", lastName: "Smith", age: 25, email: "jane@example.com" },
    { firstName: "Bob", lastName: "Johnson", age: 35, email: "bob@example.com" }
];

// Create display names
let displayUsers = users.map(user => ({
    id: `user_${user.firstName.toLowerCase()}_${user.lastName.toLowerCase()}`,
    fullName: `${user.firstName} ${user.lastName}`,
    age: user.age,
    email: user.email,
    isAdult: user.age >= 18,
    initials: `${user.firstName[0]}${user.lastName[0]}`
}));

console.log("Processed users:", displayUsers);

// Chaining array methods
let processedNumbers = numbers
    .map(num => num * 2)        // Double each number
    .map(num => num + 1)        // Add 1 to each
    .map(num => `#${num}`);     // Convert to string with #

console.log("Chained processing:", processedNumbers);

// Advanced map usage with nested data
let orders = [
    {
        id: 1,
        customer: "Alice",
        items: [
            { product: "Laptop", quantity: 1, price: 999 },
            { product: "Mouse", quantity: 2, price: 25 }
        ]
    },
    {
        id: 2,
        customer: "Bob",
        items: [
            { product: "Keyboard", quantity: 1, price: 75 },
            { product: "Monitor", quantity: 1, price: 300 }
        ]
    }
];

// Calculate order totals
let orderSummaries = orders.map(order => {
    let total = order.items.reduce((sum, item) => sum + (item.quantity * item.price), 0);
    let itemCount = order.items.reduce((count, item) => count + item.quantity, 0);
    
    return {
        orderId: order.id,
        customer: order.customer,
        itemCount: itemCount,
        total: total,
        formattedTotal: `${total.toFixed(2)}`
    };
});

console.log("Order summaries:", orderSummaries);

// Real-world example: Processing API response data
let apiResponse = [
    { id: 1, title: "JavaScript Basics", author: "John Smith", publishedDate: "2023-01-15", tags: ["javascript", "programming"] },
    { id: 2, title: "React Guide", author: "Jane Doe", publishedDate: "2023-02-20", tags: ["react", "frontend"] },
    { id: 3, title: "Node.js Server", author: "Bob Wilson", publishedDate: "2023-03-10", tags: ["nodejs", "backend"] }
];

// Transform API data for frontend display
let blogPosts = apiResponse.map(post => ({
    id: post.id,
    title: post.title,
    author: post.author,
    publishedDate: new Date(post.publishedDate).toLocaleDateString(),
    daysSincePublished: Math.floor((new Date() - new Date(post.publishedDate)) / (1000 * 60 * 60 * 24)),
    tags: post.tags.map(tag => `#${tag}`).join(" "),
    slug: post.title.toLowerCase().replace(/\s+/g, "-").replace(/[^\w-]/g, ""),
    excerpt: `${post.title} by ${post.author}...`
}));

console.log("Blog posts for display:", blogPosts);

// Performance tip: map creates new array, use forEach for side effects only
let sideEffectData = [];

// Good use of forEach (side effects)
numbers.forEach(num => {
    if (num % 2 === 0) {
        sideEffectData.push(num);
    }
});

// Better use of filter and map for same result
let functionalApproach = numbers
    .filter(num => num % 2 === 0)
    .map(num => num);

console.log("Side effect approach:", sideEffectData);
console.log("Functional approach:", functionalApproach);
```

**Activity:** Build a data transformation pipeline that processes student records, calculating grades, formatting names, and creating summary reports.

**Evening Session: Method Chaining and Performance**
- Chaining multiple array methods
- Performance considerations
- When to use each method

### Day 3: filter, find, and Array Searching
**Complexity: Intermediate**

**Morning Session: Filtering and Finding Data**
- filter method for data selection
- find and findIndex methods
- some and every methods for testing

**Code Example:**
```javascript
// filter - create new array with elements that pass a test
let numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

let evenNumbers = numbers.filter(num => num % 2 === 0);
let greaterThanFive = numbers.filter(num => num > 5);
let oddAndLarge = numbers.filter(num => num % 2 === 1 && num > 5);

console.log("Even numbers:", evenNumbers);           // [2, 4, 6, 8, 10]
console.log("Greater than 5:", greaterThanFive);     // [6, 7, 8, 9, 10]
console.log("Odd and large:", oddAndLarge);          // [7, 9]

// filter with objects
let employees = [
    { name: "Alice", department: "Engineering", salary: 80000, experience: 5 },
    { name: "Bob", department: "Sales", salary: 60000, experience: 3 },
    { name: "Carol", department: "Engineering", salary: 95000, experience: 8 },
    { name: "David", department: "Marketing", salary: 55000, experience: 2 },
    { name: "Eva", department: "Engineering", salary: 105000, experience: 10 }
];

// Filter by department
let engineers = employees.filter(emp => emp.department === "Engineering");
console.log("Engineers:", engineers);

// Filter by salary range
let highEarners = employees.filter(emp => emp.salary >= 80000);
console.log("High earners:", highEarners);

// Complex filtering
let seniorEngineers = employees.filter(emp => 
    emp.department === "Engineering" && emp.experience >= 8
);
console.log("Senior engineers:", seniorEngineers);

// find - return first element that matches condition
let firstEngineer = employees.find(emp => emp.department === "Engineering");
let highestPaid = employees.find(emp => emp.salary > 100000);
let nonExistent = employees.find(emp => emp.department === "HR");

console.log("First engineer:", firstEngineer);
console.log("Highest paid:", highestPaid);
console.log("Non-existent department:", nonExistent); // undefined

// findIndex - return index of first matching element
let aliceIndex = employees.findIndex(emp => emp.name === "Alice");
let marketingIndex = employees.findIndex(emp => emp.department === "Marketing");

console.log("Alice's index:", aliceIndex);
console.log("First marketing employee index:", marketingIndex);

// some - test if at least one element passes condition
let hasHighEarner = employees.some(emp => emp.salary > 100000);
let hasIntern = employees.some(emp => emp.experience < 1);
let hasEngineers = employees.some(emp => emp.department === "Engineering");

console.log("Has high earner:", hasHighEarner);     // true
console.log("Has intern:", hasIntern);             // false
console.log("Has engineers:", hasEngineers);       // true

// every - test if all elements pass condition
let allExperienced = employees.every(emp => emp.experience > 0);
let allHighPaid = employees.every(emp => emp.salary > 70000);
let allEmployed = employees.every(emp => emp.name && emp.department);

console.log("All experienced:", allExperienced);   // true
console.log("All high paid:", allHighPaid);       // false
console.log("All have required fields:", allEmployed); // true

// Practical example: E-commerce product filtering
let products = [
    { id: 1, name: "Laptop", category: "Electronics", price: 999, rating: 4.5, inStock: true },
    { id: 2, name: "Smartphone", category: "Electronics", price: 699, rating: 4.7, inStock: true },
    { id: 3, name: "Book", category: "Education", price: 25, rating: 4.2, inStock: false },
    { id: 4, name: "Headphones", category: "Electronics", price: 199, rating: 4.3, inStock: true },
    { id: 5, name: "Tablet", category: "Electronics", price: 399, rating: 4.1, inStock: true },
    { id: 6, name: "Desk", category: "Furniture", price: 299, rating: 3.9, inStock: true }
];

// Product search and filtering functions
function searchProducts(products, filters) {
    return products.filter(product => {
        // Category filter
        if (filters.category && product.category !== filters.category) {
            return false;
        }
        
        // Price range filter
        if (filters.minPrice && product.price < filters.minPrice) {
            return false;
        }
        
        if (filters.maxPrice && product.price > filters.maxPrice) {
            return false;
        }
        
        // Rating filter
        if (filters.minRating && product.rating < filters.minRating) {
            return false;
        }
        
        // Stock filter
        if (filters.inStock && !product.inStock) {
            return false;
        }
        
        // Name search
        if (filters.searchTerm) {
            return product.name.toLowerCase().includes(filters.searchTerm.toLowerCase());
        }
        
        return true;
    });
}

// Example searches
let electronicsUnder500 = searchProducts(products, {
    category: "Electronics",
    maxPrice: 500
});

let highRatedInStock = searchProducts(products, {
    minRating: 4.3,
    inStock: true
});

let laptopSearch = searchProducts(products, {
    searchTerm: "laptop"
});

console.log("Electronics under $500:", electronicsUnder500);
console.log("High rated in stock:", highRatedInStock);
console.log("Laptop search:", laptopSearch);

// Advanced filtering with multiple conditions
let premiumProducts = products
    .filter(product => product.price > 300)
    .filter(product => product.rating >= 4.0)
    .filter(product => product.inStock);

console.log("Premium products:", premiumProducts);

// Using filter with map for data transformation
let productSummary = products
    .filter(product => product.inStock)
    .map(product => ({
        name: product.name,
        price: `${product.price}`,
        rating: `${product.rating}/5 stars`,
        category: product.category
    }));

console.log("In-stock product summary:", productSummary);

// Real-world example: User permission system
let users = [
    { id: 1, name: "Admin User", role: "admin", permissions: ["read", "write", "delete"], isActive: true },
    { id: 2, name: "Editor", role: "editor", permissions: ["read", "write"], isActive: true },
    { id: 3, name: "Viewer", role: "viewer", permissions: ["read"], isActive: true },
    { id: 4, name: "Inactive User", role: "editor", permissions: ["read", "write"], isActive: false }
];

// Find users with specific permissions
let canDelete = users.filter(user => 
    user.permissions.includes("delete") && user.isActive
);

let canWrite = users.filter(user => 
    user.permissions.includes("write") && user.isActive
);

// Check if any user has admin privileges
let hasAdmin = users.some(user => user.role === "admin" && user.isActive);

// Check if all active users can read
let allCanRead = users
    .filter(user => user.isActive)
    .every(user => user.permissions.includes("read"));

console.log("Users who can delete:", canDelete);
console.log("Users who can write:", canWrite);
console.log("Has active admin:", hasAdmin);
console.log("All active users can read:", allCanRead);

// Combining filter with other array methods
function analyzeEmployeeData(employees) {
    return {
        totalEmployees: employees.length,
        engineerCount: employees.filter(emp => emp.department === "Engineering").length,
        averageSalary: employees.reduce((sum, emp) => sum + emp.salary, 0) / employees.length,
        seniorEmployees: employees.filter(emp => emp.experience >= 5),
        departments: [...new Set(employees.map(emp => emp.department))],
        hasHighEarners: employees.some(emp => emp.salary > 100000)
    };
}

let employeeAnalysis = analyzeEmployeeData(employees);
console.log("Employee analysis:", employeeAnalysis);
```

**Activity:** Create a comprehensive search and filtering system for a movie database with multiple criteria and advanced search features.

**Evening Session: Advanced Filtering Patterns**
- Complex filter conditions
- Combining filter with other methods
- Performance optimization for large datasets

### Day 4: reduce and Advanced Array Operations
**Complexity: Intermediate to Advanced**

**Morning Session: The reduce Method**
- Understanding reduce fundamentals
- Common reduce patterns
- Using reduce for calculations and transformations

**Code Example:**
```javascript
// reduce - powerful method for transforming array into single value
let numbers = [1, 2, 3, 4, 5];

// Basic reduce - sum all numbers
let sum = numbers.reduce((accumulator, currentValue) => {
    console.log(`Acc: ${accumulator}, Current: ${currentValue}`);
    return accumulator + currentValue;
}, 0); // 0 is initial value

console.log("Sum:", sum); // 15

// reduce without initial value (uses first element as initial)
let product = numbers.reduce((acc, curr) => acc * curr);
console.log("Product:", product); // 120

// Find maximum value using reduce
let max = numbers.reduce((max, current) => current > max ? current : max);
console.log("Maximum:", max); // 5

// More complex example: processing sales data
let sales = [
    { product: "Laptop", amount: 999, quantity: 2, date: "2023-01-15" },
    { product: "Mouse", amount: 25, quantity: 5, date: "2023-01-16" },
    { product: "
    
```