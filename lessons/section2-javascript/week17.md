---
title: "Objects and Object-Oriented Programming"
description: "Master JavaScript objects, properties, methods, and object-oriented programming concepts"
week: 17
section: 2
prerequisites: ["Functions", "Arrays", "Basic data types"]
objectives:
  - "Create and manipulate JavaScript objects effectively"
  - "Understand object properties, methods, and this keyword"
  - "Implement object-oriented programming patterns"
  - "Use constructor functions and ES6 classes"
  - "Apply inheritance and polymorphism concepts"
---

# Week 17: Objects and Object-Oriented Programming

## Learning Objectives

By the end of this week, students will be able to:
- Create objects using literal notation and constructor functions
- Access and modify object properties and methods
- Understand and use the `this` keyword correctly
- Implement object-oriented programming principles
- Use ES6 classes and inheritance
- Build complex applications using object-oriented design

## Daily Activities

### Day 1: Object Basics and Property Management
**Complexity: Beginner**

**Morning Session: Object Fundamentals**
- Object literal notation
- Property access and modification
- Dynamic property creation

**Code Example:**
```javascript
// Object literal notation - the most common way to create objects
let student = {
    name: "Alice Johnson",
    age: 20,
    major: "Computer Science",
    gpa: 3.8,
    isEnrolled: true
};

// Accessing object properties
console.log(student.name);        // Dot notation: "Alice Johnson"
console.log(student["age"]);      // Bracket notation: 20
console.log(student.major);       // "Computer Science"

// Modifying existing properties
student.age = 21;                 // Update age
student["gpa"] = 3.9;            // Update GPA using bracket notation
console.log("Updated student:", student);

// Adding new properties dynamically
student.email = "alice@university.edu";
student.courses = ["CS101", "MATH201", "PHYS101"];
student.address = {              // Nested object
    street: "123 Campus Ave",
    city: "College Town",
    state: "CA",
    zipCode: "90210"
};

// Beginner: Objects can contain any data type
let product = {
    id: 1001,
    name: "Laptop Pro",
    price: 1299.99,
    inStock: true,
    categories: ["Electronics", "Computers"],
    specifications: {
        processor: "Intel i7",
        memory: "16GB",
        storage: "512GB SSD"
    },
    reviews: [
        { rating: 5, comment: "Excellent laptop!" },
        { rating: 4, comment: "Good performance" }
    ]
};

// Working with nested objects
console.log(product.specifications.processor);  // "Intel i7"
console.log(product.reviews[0].rating);         // 5

// Dynamic property access
let propertyName = "price";
console.log(product[propertyName]);             // 1299.99

// Checking if properties exist
if ("email" in student) {
    console.log("Student has email:", student.email);
}

if (student.phone !== undefined) {
    console.log("Student phone:", student.phone);
} else {
    console.log("Student phone not provided");
}

// Deleting properties
delete student.isEnrolled;
console.log("After deletion:", student);

// Object.keys(), Object.values(), Object.entries()
console.log("Property names:", Object.keys(product));
console.log("Property values:", Object.values(product));
console.log("Key-value pairs:", Object.entries(product));
```

**Activity:** Create a personal library system using objects to represent books, authors, and library information.

**Evening Session: Object Iteration and Manipulation**
- Looping through object properties
- Object copying and cloning
- Comparing objects

### Day 2: Methods and the `this` Keyword
**Complexity: Beginner to Intermediate**

**Morning Session: Object Methods**
- Adding methods to objects
- Understanding `this` context
- Method invocation patterns

**Code Example:**
```javascript
// Objects with methods - functions inside objects
let calculator = {
    value: 0,
    
    // Method using function declaration
    add: function(number) {
        this.value += number;        // Beginner: 'this' refers to the calculator object
        return this;                 // Return this for method chaining
    },
    
    // Method using ES6 shorthand
    subtract(number) {
        this.value -= number;
        return this;
    },
    
    multiply(number) {
        this.value *= number;
        return this;
    },
    
    divide(number) {
        if (number !== 0) {
            this.value /= number;
        } else {
            console.log("Cannot divide by zero!");
        }
        return this;
    },
    
    clear() {
        this.value = 0;
        return this;
    },
    
    getResult() {
        return this.value;
    },
    
    display() {
        console.log(`Current value: ${this.value}`);
        return this;
    }
};

// Using object methods
calculator.add(10).multiply(2).subtract(5).display(); // Method chaining
console.log("Final result:", calculator.getResult());

// Bank account example with more complex methods
let bankAccount = {
    accountNumber: "12345678",
    accountHolder: "John Doe",
    balance: 1000,
    transactions: [],
    
    deposit(amount) {
        if (amount > 0) {
            this.balance += amount;
            this.transactions.push({
                type: "deposit",
                amount: amount,
                date: new Date(),
                balance: this.balance
            });
            console.log(`Deposited $${amount}. New balance: $${this.balance}`);
        } else {
            console.log("Deposit amount must be positive");
        }
        return this;
    },
    
    withdraw(amount) {
        if (amount > 0 && amount <= this.balance) {
            this.balance -= amount;
            this.transactions.push({
                type: "withdrawal",
                amount: amount,
                date: new Date(),
                balance: this.balance
            });
            console.log(`Withdrew $${amount}. New balance: $${this.balance}`);
        } else if (amount > this.balance) {
            console.log("Insufficient funds");
        } else {
            console.log("Withdrawal amount must be positive");
        }
        return this;
    },
    
    getBalance() {
        return this.balance;
    },
    
    getTransactionHistory() {
        return this.transactions.slice(); // Return a copy
    },
    
    displayStatement() {
        console.log(`\n=== Account Statement ===`);
        console.log(`Account: ${this.accountNumber}`);
        console.log(`Holder: ${this.accountHolder}`);
        console.log(`Current Balance: $${this.balance}`);
        console.log(`\nTransaction History:`);
        
        this.transactions.forEach((transaction, index) => {
            console.log(`${index + 1}. ${transaction.type.toUpperCase()}: $${transaction.amount} - Balance: $${transaction.balance}`);
        });
        
        return this;
    }
};

// Using the bank account
bankAccount
    .deposit(500)
    .withdraw(200)
    .deposit(100)
    .displayStatement();

// Arrow functions and 'this' - Important difference!
let userProfile = {
    name: "Sarah",
    age: 25,
    
    // Regular function - 'this' refers to userProfile
    greet: function() {
        console.log(`Hello, I'm ${this.name} and I'm ${this.age} years old`);
    },
    
    // Arrow function - 'this' refers to global/window object (avoid in object methods)
    greetArrow: () => {
        console.log(`Hello, I'm ${this.name}`); // 'this' is NOT userProfile!
    },
    
    // Method that uses arrow function inside (common pattern)
    listHobbies: function() {
        let hobbies = ["reading", "coding", "hiking"];
        
        hobbies.forEach((hobby) => {
            // Arrow function preserves 'this' from enclosing scope
            console.log(`${this.name} enjoys ${hobby}`);
        });
    }
};

userProfile.greet();        // Works correctly
userProfile.greetArrow();   // 'this' doesn't work as expected
userProfile.listHobbies();  // Arrow function preserves 'this'

// Advanced: Dynamic method assignment
let gameCharacter = {
    name: "Hero",
    health: 100,
    level: 1
};

// Adding methods dynamically
gameCharacter.attack = function(target) {
    console.log(`${this.name} attacks ${target}!`);
};

gameCharacter.heal = function(amount) {
    this.health += amount;
    console.log(`${this.name} heals for ${amount} HP. Health: ${this.health}`);
};

gameCharacter.levelUp = function() {
    this.level++;
    this.health += 20;
    console.log(`${this.name} reached level ${this.level}! Health: ${this.health}`);
};

gameCharacter.attack("Goblin");
gameCharacter.heal(25);
gameCharacter.levelUp();
```

**Activity:** Build a task management system with objects that have methods for adding, completing, and organizing tasks.

**Evening Session: Method Context and Binding**
- Call, apply, and bind methods
- Method borrowing
- Context preservation

### Day 3: Constructor Functions and Prototypes
**Complexity: Intermediate**

**Morning Session: Constructor Functions**
- Creating objects with constructor functions
- The `new` keyword
- Instance vs prototype properties

**Code Example:**
```javascript
// Constructor function - template for creating objects
function Person(name, age, occupation) {
    // 'this' refers to the new object being created
    this.name = name;
    this.age = age;
    this.occupation = occupation;
    this.isAlive = true;
    
    // Method defined in constructor (not recommended for performance)
    this.introduce = function() {
        console.log(`Hi, I'm ${this.name}, a ${this.age}-year-old ${this.occupation}`);
    };
}

// Creating instances using 'new' keyword
let person1 = new Person("Alice", 30, "Engineer");
let person2 = new Person("Bob", 25, "Designer");

person1.introduce(); // "Hi, I'm Alice, a 30-year-old Engineer"
person2.introduce(); // "Hi, I'm Bob, a 25-year-old Designer"

// Better approach: Methods on prototype (shared across instances)
function Animal(species, name) {
    this.species = species;
    this.name = name;
    this.energy = 100;
}

// Adding methods to prototype - shared by all instances
Animal.prototype.eat = function(food) {
    this.energy += 10;
    console.log(`${this.name} the ${this.species} eats ${food}. Energy: ${this.energy}`);
};

Animal.prototype.sleep = function() {
    this.energy += 25;
    console.log(`${this.name} sleeps. Energy: ${this.energy}`);
};

Animal.prototype.play = function() {
    this.energy -= 15;
    console.log(`${this.name} plays. Energy: ${this.energy}`);
};

Animal.prototype.getInfo = function() {
    return `${this.name} is a ${this.species} with ${this.energy} energy`;
};

// Creating animal instances
let dog = new Animal("Dog", "Buddy");
let cat = new Animal("Cat", "Whiskers");

dog.eat("bone");
dog.play();
dog.sleep();
console.log(dog.getInfo());

cat.eat("fish");
cat.play();
console.log(cat.getInfo());

// More complex constructor example: Car
function Car(make, model, year, color) {
    this.make = make;
    this.model = model;
    this.year = year;
    this.color = color;
    this.mileage = 0;
    this.isRunning = false;
    this.fuel = 100;
}

Car.prototype.start = function() {
    if (!this.isRunning && this.fuel > 0) {
        this.isRunning = true;
        console.log(`${this.make} ${this.model} started`);
    } else if (this.fuel <= 0) {
        console.log("Cannot start: No fuel");
    } else {
        console.log("Car is already running");
    }
};

Car.prototype.stop = function() {
    if (this.isRunning) {
        this.isRunning = false;
        console.log(`${this.make} ${this.model} stopped`);
    } else {
        console.log("Car is already stopped");
    }
};

Car.prototype.drive = function(miles) {
    if (this.isRunning && this.fuel > 0) {
        this.mileage += miles;
        this.fuel -= miles * 0.1; // 0.1 fuel per mile
        console.log(`Drove ${miles} miles. Total mileage: ${this.mileage}`);
        
        if (this.fuel <= 0) {
            this.fuel = 0;
            this.isRunning = false;
            console.log("Ran out of fuel!");
        }
    } else if (!this.isRunning) {
        console.log("Start the car first");
    } else {
        console.log("No fuel");
    }
};

Car.prototype.refuel = function(amount) {
    this.fuel = Math.min(this.fuel + amount, 100);
    console.log(`Refueled. Current fuel: ${this.fuel}%`);
};

Car.prototype.getDetails = function() {
    return {
        vehicle: `${this.year} ${this.make} ${this.model}`,
        color: this.color,
        mileage: this.mileage,
        fuel: this.fuel,
        isRunning: this.isRunning
    };
};

// Using the Car constructor
let myCar = new Car("Toyota", "Camry", 2022, "Blue");
console.log("Car details:", myCar.getDetails());

myCar.start();
myCar.drive(50);
myCar.drive(30);
myCar.refuel(20);
myCar.stop();

// Advanced: Constructor with validation
function BankAccount(accountHolder, initialDeposit) {
    // Validation
    if (!accountHolder || typeof accountHolder !== 'string') {
        throw new Error("Account holder name is required");
    }
    
    if (initialDeposit < 0) {
        throw new Error("Initial deposit cannot be negative");
    }
    
    this.accountNumber = Math.random().toString(36).substr(2, 9).toUpperCase();
    this.accountHolder = accountHolder;
    this.balance = initialDeposit || 0;
    this.transactions = [];
    this.isActive = true;
    this.createdDate = new Date();
}

BankAccount.prototype.deposit = function(amount) {
    if (!this.isActive) {
        console.log("Account is closed");
        return false;
    }
    
    if (amount <= 0) {
        console.log("Deposit amount must be positive");
        return false;
    }
    
    this.balance += amount;
    this.transactions.push({
        type: "deposit",
        amount: amount,
        date: new Date(),
        balance: this.balance
    });
    
    console.log(`Deposited $${amount}. New balance: $${this.balance}`);
    return true;
};

BankAccount.prototype.withdraw = function(amount) {
    if (!this.isActive) {
        console.log("Account is closed");
        return false;
    }
    
    if (amount <= 0) {
        console.log("Withdrawal amount must be positive");
        return false;
    }
    
    if (amount > this.balance) {
        console.log("Insufficient funds");
        return false;
    }
    
    this.balance -= amount;
    this.transactions.push({
        type: "withdrawal",
        amount: amount,
        date: new Date(),
        balance: this.balance
    });
    
    console.log(`Withdrew $${amount}. New balance: $${this.balance}`);
    return true;
};

BankAccount.prototype.closeAccount = function() {
    this.isActive = false;
    console.log(`Account ${this.accountNumber} has been closed`);
};

// Testing the BankAccount constructor
try {
    let account = new BankAccount("John Doe", 1000);
    account.deposit(500);
    account.withdraw(200);
    console.log("Account info:", {
        number: account.accountNumber,
        holder: account.accountHolder,
        balance: account.balance,
        transactions: account.transactions.length
    });
} catch (error) {
    console.error("Error creating account:", error.message);
}
```

**Activity:** Create a library management system using constructor functions for books, members, and library operations.

**Evening Session: Prototype Chain and Inheritance**
- Understanding the prototype chain
- Prototype inheritance
- Object.create() method

### Day 4: ES6 Classes and Modern OOP
**Complexity: Intermediate to Advanced**

**Morning Session: ES6 Class Syntax**
- Class declarations and expressions
- Constructor methods
- Class methods and static methods

**Code Example:**
```javascript
// ES6 Class syntax - modern way to create object templates
class Rectangle {
    // Constructor method - runs when new instance is created
    constructor(width, height) {
        this.width = width;
        this.height = height;
    }
    
    // Instance methods
    getArea() {
        return this.width * this.height;
    }
    
    getPerimeter() {
        return 2 * (this.width + this.height);
    }
    
    isSquare() {
        return this.width === this.height;
    }
    
    // Getter method
    get dimensions() {
        return `${this.width}x${this.height}`;
    }
    
    // Setter method
    set dimensions(value) {
        const [width, height] = value.split('x').map(Number);
        this.width = width;
        this.height = height;
    }
    
    // Static method - belongs to class, not instances
    static createSquare(side) {
        return new Rectangle(side, side);
    }
    
    static compareAreas(rect1, rect2) {
        const area1 = rect1.getArea();
        const area2 = rect2.getArea();
        
        if (area1 > area2) return `First rectangle is larger (${area1} vs ${area2})`;
        if (area2 > area1) return `Second rectangle is larger (${area2} vs ${area1})`;
        return `Both rectangles have equal area (${area1})`;
    }
    
    toString() {
        return `Rectangle(${this.width}, ${this.height}) - Area: ${this.getArea()}`;
    }
}

// Using the Rectangle class
let rect1 = new Rectangle(10, 5);
let rect2 = new Rectangle(8, 8);

console.log(rect1.toString());
console.log(rect2.toString());
console.log("Is rect2 a square?", rect2.isSquare());

// Using getter and setter
console.log("Rect1 dimensions:", rect1.dimensions);
rect1.dimensions = "12x6";
console.log("Updated rect1:", rect1.toString());

// Using static methods
let square = Rectangle.createSquare(7);
console.log("Created square:", square.toString());
console.log(Rectangle.compareAreas(rect1, square));

// More complex class example: User Management
class User {
    constructor(username, email, role = 'user') {
        this.username = username;
        this.email = email;
        this.role = role;
        this.isActive = true;
        this.loginCount = 0;
        this.lastLogin = null;
        this.createdAt = new Date();
    }
    
    login() {
        if (!this.isActive) {
            console.log("Account is deactivated");
            return false;
        }
        
        this.loginCount++;
        this.lastLogin = new Date();
        console.log(`${this.username} logged in successfully`);
        return true;
    }
    
    logout() {
        console.log(`${this.username} logged out`);
    }
    
    updateEmail(newEmail) {
        if (this.isValidEmail(newEmail)) {
            this.email = newEmail;
            console.log(`Email updated to ${newEmail}`);
            return true;
        } else {
            console.log("Invalid email format");
            return false;
        }
    }
    
    deactivate() {
        this.isActive = false;
        console.log(`User ${this.username} has been deactivated`);
    }
    
    reactivate() {
        this.isActive = true;
        console.log(`User ${this.username} has been reactivated`);
    }
    
    isValidEmail(email) {
        return email.includes('@') && email.includes('.');
    }
    
    // Getter for user info
    get info() {
        return {
            username: this.username,
            email: this.email,
            role: this.role,
            isActive: this.isActive,
            loginCount: this.loginCount,
            lastLogin: this.lastLogin,
            accountAge: Math.floor((new Date() - this.createdAt) / (1000 * 60 * 60 * 24))
        };
    }
    
    // Static method for user validation
    static validateUserData(userData) {
        const errors = [];
        
        if (!userData.username || userData.username.length < 3) {
            errors.push("Username must be at least 3 characters");
        }
        
        if (!userData.email || !userData.email.includes('@')) {
            errors.push("Valid email is required");
        }
        
        return {
            isValid: errors.length === 0,
            errors: errors
        };
    }
    
    static createAdmin(username, email) {
        return new User(username, email, 'admin');
    }
}

// Using the User class
let user1 = new User("john_doe", "john@example.com");
let admin = User.createAdmin("admin_user", "admin@example.com");

user1.login();
user1.updateEmail("john.doe@example.com");
console.log("User info:", user1.info);

// Validation example
let newUserData = { username: "ab", email: "invalid" };
let validation = User.validateUserData(newUserData);
if (!validation.isValid) {
    console.log("Validation errors:", validation.errors);
}

// Advanced class with private fields (newer JavaScript feature)
class BankAccount {
    // Private fields (prefix with #)
    #balance = 0;
    #accountNumber;
    #transactions = [];
    
    constructor(accountHolder, initialDeposit = 0) {
        this.accountHolder = accountHolder;
        this.#accountNumber = this.#generateAccountNumber();
        this.#balance = initialDeposit;
        this.isActive = true;
        
        if (initialDeposit > 0) {
            this.#transactions.push({
                type: 'initial deposit',
                amount: initialDeposit,
                date: new Date(),
                balance: this.#balance
            });
        }
    }
    
    // Private method
    #generateAccountNumber() {
        return 'ACC' + Math.random().toString(36).substr(2, 9).toUpperCase();
    }
    
    #recordTransaction(type, amount) {
        this.#transactions.push({
            type: type,
            amount: amount,
            date: new Date(),
            balance: this.#balance
        });
    }
    
    // Public methods
    deposit(amount) {
        if (!this.isActive) throw new Error("Account is closed");
        if (amount <= 0) throw new Error("Amount must be positive");
        
        this.#balance += amount;
        this.#recordTransaction('deposit', amount);
        console.log(`Deposited $${amount}. New balance: $${this.#balance}`);
    }
    
    withdraw(amount) {
        if (!this.isActive) throw new Error("Account is closed");
        if (amount <= 0) throw new Error("Amount must be positive");
        if (amount > this.#balance) throw new Error("Insufficient funds");
        
        this.#balance -= amount;
        this.#recordTransaction('withdrawal', amount);
        console.log(`Withdrew $${amount}. New balance: $${this.#balance}`);
    }
    
    // Getter for balance (read-only access to private field)
    get balance() {
        return this.#balance;
    }
    
    get accountNumber() {
        return this.#accountNumber;
    }
    
    get transactionHistory() {
        return [...this.#transactions]; // Return copy of private array
    }
    
    closeAccount() {
        this.isActive = false;
        console.log(`Account ${this.#accountNumber} has been closed`);
    }
}

// Using the BankAccount with private fields
let account = new BankAccount("Jane Smith", 1000);
console.log("Account number:", account.accountNumber);
console.log("Initial balance:", account.balance);

account.deposit(500);
account.withdraw(200);
console.log("Transaction history:", account.transactionHistory);

// Cannot access private fields directly
// console.log(account.#balance); // This would cause an error
```

**Activity:** Build a complete e-commerce product catalog system using ES6 classes for products, categories, and shopping cart functionality.

**Evening Session: Advanced Class Features**
- Private fields and methods
- Getters and setters
- Static methods and properties

### Day 5: Inheritance and Polymorphism
**Complexity: Advanced**

**Morning Session: Class Inheritance**
- Extending classes with `extends`
- The `super` keyword
- Method overriding

**Code Example:**
```javascript
// Base class (parent/superclass)
class Vehicle {
    constructor(make, model, year) {
        this.make = make;
        this.model = model;
        this.year = year;
        this.isRunning = false;
        this.mileage = 0;
    }
    
    start() {
        if (!this.isRunning) {
            this.isRunning = true;
            console.log(`${this.make} ${this.model} started`);
        } else {
            console.log("Vehicle is already running");
        }
    }
    
    stop() {
        if (this.isRunning) {
            this.isRunning = false;
            console.log(`${this.make} ${this.model} stopped`);
        } else {
            console.log("Vehicle is already stopped");
        }
    }
    
    getInfo() {
        return `${this.year} ${this.make} ${this.model} - ${this.mileage} miles`;
    }
    
    // Method to be overridden by subclasses
    getDescription() {
        return `This is a ${this.year} ${this.make} ${this.model}`;
    }
}

// Derived class (child/subclass) - Car
class Car extends Vehicle {
    constructor(make, model, year, doors, fuelType) {
        super(make, model, year); // Call parent constructor
        this.doors = doors;
        this.fuelType = fuelType;
        this.fuel = 100;
    }
    
    // Override parent method
    start() {
        if (this.fuel <= 0) {
            console.log("Cannot start: No fuel");
            return;
        }
        super.start(); // Call parent method
    }
    
    drive(miles) {
        if (!this.isRunning) {
            console.log("Start the car first");
            return;
        }
        
        if (this.fuel <= 0) {
            console.log("No fuel to drive");
            return;
        }
        
        this.mileage += miles;
        this.fuel -= miles * 0.1; // Consume fuel
        console.log(`Drove ${miles} miles. Fuel: ${this.fuel.toFixed(1)}%`);
        
        if (this.fuel <= 0) {
            this.fuel = 0;
            this.stop();
            console.log("Ran out of fuel!");
        }
    }
    
    refuel() {
        this.fuel = 100;
        console.log("Car refueled to 100%");
    }
    
    // Override getDescription method
    getDescription() {
        return `${super.getDescription()} with ${this.doors} doors, runs on ${this.fuelType}`;
    }
    
    // Car-specific method
    honk() {
        console.log(`${this.make} ${this.model} goes BEEP BEEP!`);
    }
}

// Another derived class - Motorcycle
class Motorcycle extends Vehicle {
    constructor(make, model, year, engineSize) {
        super(make, model, year);
        this.engineSize = engineSize;
        this.hasWindshield = false;
    }
    
    // Override start method
    start() {
        console.log("Motorcycle starting sequence...");
        super.start();
        console.log("Ready to ride!");
    }
    
    wheelie() {
        if (this.isRunning) {
            console.log(`${this.make} ${this.model} pops a wheelie!`);
        } else {
            console.log("Start the motorcycle first");
        }
    }
    
    // Override getDescription
    getDescription() {
        return `${super.getDescription()} with ${this.engineSize}cc engine`;
    }
}

// Using inheritance
let myCar = new Car("Toyota", "Camry", 2022, 4, "gasoline");
let myBike = new Motorcycle("Yamaha", "R1", 2023, 1000);

console.log(myCar.getDescription());
console.log(myBike.getDescription());

myCar.start();
myCar.drive(50);
myCar.honk();

myBike.start();
myBike.wheelie();

// More complex inheritance example: Employee hierarchy
class Employee {
    constructor(name, id, department, salary) {
        this.name = name;
        this.id = id;
        this.department = department;
        this.salary = salary;
        this.hireDate = new Date();
        this.isActive = true;
    }
    
    work() {
        console.log(`${this.name} is working`);
    }
    
    takeBreak() {
        console.log(`${this.name} is taking a break`);
    }
    
    getAnnualSalary() {
        return this.salary * 12;
    }
    
    getInfo() {
        return {
            name: this.name,
            id: this.id,
            department: this.department,
            salary: this.salary,
            isActive: this.isActive
        };
    }
    
    // Template method pattern
    doWork() {
        this.clockIn();
        this.work();
        this.clockOut();
    }
    
    clockIn() {
        console.log(`${this.name} clocked in`);
    }
    
    clockOut() {
        console.log(`${this.name} clocked out`);
    }
}

class Developer extends Employee {
    constructor(name, id, salary, programmingLanguages) {
        super(name, id, "Engineering", salary);
        this.programmingLanguages = programmingLanguages;
        this.projects = [];
    }
    
    // Override work method
    work() {
        console.log(`${this.name} is coding in ${this.programmingLanguages.join(", ")}`);
    }
    
    code(project) {
        console.log(`${this.name} is working on ${project}`);
        if (!this.projects.includes(project)) {
            this.projects.push(project);
        }
    }
    
    debug() {
        console.log(`${this.name} is debugging code`);
    }
    
    // Override doWork with developer-specific workflow
    doWork() {
        this.clockIn();
        this.code("Daily Development Tasks");
        this.debug();
        this.takeBreak();
        this.code("Code Review");
        this.clockOut();
    }
}

class Manager extends Employee {
    constructor(name, id, salary, teamSize) {
        super(name, id, "Management", salary);
        this.teamSize = teamSize;
        this.meetings = [];
    }
    
    work() {
        console.log(`${this.name} is managing a team of ${this.teamSize} people`);
    }
    
    scheduleMeeting(topic, duration) {
        this.meetings.push({ topic, duration, date: new Date() });
        console.log(`${this.name} scheduled a meeting about ${topic}`);
    }
    
    reviewPerformance(employee) {
        console.log(`${this.name} is reviewing ${employee}'s performance`);
    }
    
    doWork() {
        this.clockIn();
        this.scheduleMeeting("Daily Standup", 30);
        this.work();
        this.reviewPerformance("team member");
        this.clockOut();
    }
}

// Using the employee hierarchy
let dev = new Developer("Alice Johnson", "DEV001", 8000, ["JavaScript", "Python", "React"]);
let manager = new Manager("Bob Smith", "MGR001", 12000, 8);

console.log("=== Developer Day ===");
dev.doWork();
dev.code("E-commerce Website");

console.log("\n=== Manager Day ===");
manager.doWork();
manager.scheduleMeeting("Sprint Planning", 60);

// Polymorphism - same method call, different behavior
let employees = [dev, manager];

console.log("\n=== Polymorphism Example ===");
employees.forEach(employee => {
    employee.work(); // Each class implements work() differently
});

// Advanced inheritance with multiple levels
class Animal {
    constructor(name, species) {
        this.name = name;
        this.species = species;
        this.energy = 100;
    }
    
    eat() {
        this.energy += 10;
        console.log(`${this.name} is eating. Energy: ${this.energy}`);
    }
    
    sleep() {
        this.energy += 25;
        console.log(`${this.name} is sleeping. Energy: ${this.energy}`);
    }
    
    makeSound() {
        console.log(`${this.name} makes a sound`);
    }
}

class Mammal extends Animal {
    constructor(name, species, furColor) {
        super(name, species);
        this.furColor = furColor;
        this.bodyTemperature = "warm";
    }
    
    giveBirth() {
        console.log(`${this.name} gives birth to live young`);
    }
}

class Dog extends Mammal {
    constructor(name, breed, furColor) {
        super(name, "Canine", furColor);
        this.breed = breed;
        this.loyalty = 100;
    }
    
    makeSound() {
        console.log(`${this.name} barks: Woof! Woof!`);
    }
    
    fetch() {
        this.energy -= 15;
        console.log(`${this.name} fetches the ball. Energy: ${this.energy}`);
    }
    
    wagTail() {
        console.log(`${this.name} wags tail happily`);
    }
}

class Cat extends Mammal {
    constructor(name, breed, furColor) {
        super(name, "Feline", furColor);
        this.breed = breed;
        this.independence = 90;
    }
    
    makeSound() {
        console.log(`${this.name} meows: Meow!`);
    }
    
    purr() {
        console.log(`${this.name} purrs contentedly`);
    }
    
    hunt() {
        this.energy -= 20;
        console.log(`${this.name} hunts. Energy: ${this.energy}`);
    }
}

// Using multi-level inheritance
let dog = new Dog("Buddy", "Golden Retriever", "golden");
let cat = new Cat("Whiskers", "Persian", "white");

console.log("\n=== Animal Behaviors ===");
dog.makeSound();
dog.fetch();
dog.wagTail();

cat.makeSound();
cat.purr();
cat.hunt();

// Polymorphism with animals
let animals = [dog, cat];
console.log("\n=== Animal Sounds ===");
animals.forEach(animal => {
    animal.makeSound(); // Each animal makes different sound
});
```

**Activity:** Design and implement a complete game character system with base classes and specialized character types (warriors, mages, archers) using inheritance.

**Evening Session: Advanced OOP Patterns**
- Abstract classes and interfaces (JavaScript patterns)
- Composition vs inheritance
- Design patterns introduction

## Hands-on Exercises

### Exercise 1: Library Management System
Create a comprehensive library system using OOP principles:
- Base classes for library items (books, magazines, DVDs)
- User classes with different permission levels
- Inheritance for specialized item types
- Methods for borrowing, returning, and managing inventory

### Exercise 2: E-commerce Platform
Build an object-oriented e-commerce system:
- Product hierarchy with categories and variations
- User account management with different user types
- Shopping cart and order processing classes
- Payment processing using polymorphism

### Exercise 3: Game Development Framework
Design a mini game engine using advanced OOP:
- Base game object classes
- Character inheritance hierarchies
- Game state management
- Event system using observer pattern

## Resources

**MDN Documentation:**
- [Objects](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Working_with_Objects)
- [Classes](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Classes)
- [Inheritance](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Inheritance_and_the_prototype_chain)
- [this keyword](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/this)

**W3Schools References:**
- [JavaScript Objects](https://www.w3schools.com/js/js_objects.asp)
- [JavaScript Classes](https://www.w3schools.com/js/js_classes.asp)
- [Object Methods](https://www.w3schools.com/js/js_object_methods.asp)

## Assignment Checklist

**Basic Requirements:**
- [ ] Created objects using literal notation and constructor functions
- [ ] Implemented object methods and used `this` keyword correctly
- [ ] Built constructor functions with prototype methods
- [ ] Created ES6 classes with constructors and methods
- [ ] Implemented basic inheritance using `extends` and `super`

**Intermediate Challenges:**
- [ ] Used getters, setters, and static methods effectively
- [ ] Implemented proper encapsulation and data privacy
- [ ] Created complex inheritance hierarchies
- [ ] Applied polymorphism principles
- [ ] Built reusable object-oriented components

**Advanced Extensions:**
- [ ] Used private fields and methods (modern JavaScript)
- [ ] Implemented design patterns (factory, observer, etc.)
- [ ] Created abstract base classes and interfaces
- [ ] Built complex object relationships and composition
- [ ] Designed scalable object-oriented architectures


