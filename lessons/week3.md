---
title: "Data Types and Variables WEEK 03"
description: "Learn about different data types in Python, how to work with variables, and handle user input effectively."
week: 2
---

# Data Types and Variables WEEK 03

## Learning Objectives

- Understand Python's basic data types
- Learn how to work with variables effectively
- Handle user input and output
- Perform basic operations on different data types

## Python Data Types

### 1. Integers (int)
Whole numbers without decimal points.

```python
age = 25
year = 2024
temperature = -5
```

### 2. Floating Point Numbers (float)
Numbers with decimal points.

```python
price = 19.99
pi = 3.14159
temperature = 98.6
```

### 3. Strings (str)
Text data enclosed in quotes.

```python
name = "Alice"
message = 'Hello, World!'
quote = """This is a
multi-line string"""
```

### 4. Booleans (bool)
True or False values.

```python
is_student = True
has_license = False
is_raining = True
```

## Variable Naming Rules

### Valid Names
```python
user_name = "Alice"
age2 = 25
_private_var = "hidden"
firstName = "Bob"
```

### Invalid Names
```python
# These will cause errors:
# 2age = 25        # Can't start with number
# user-name = "Alice"  # Can't use hyphens
# class = "Math"   # Can't use reserved words
```

## User Input and Output

### Getting Input
```python
name = input("What's your name? ")
age = input("How old are you? ")

print(f"Hello {name}, you are {age} years old!")
```

### Converting Input Types
```python
# input() always returns a string
age_str = input("Enter your age: ")
age = int(age_str)  # Convert to integer

# Or do it in one line
age = int(input("Enter your age: "))
```

## String Operations

### Concatenation
```python
first_name = "John"
last_name = "Doe"

# Using +
full_name = first_name + " " + last_name

# Using f-strings (recommended)
full_name = f"{first_name} {last_name}"

print(full_name)  # Output: John Doe
```

### Useful String Methods
```python
text = "Hello, World!"

print(text.upper())      # HELLO, WORLD!
print(text.lower())      # hello, world!
print(text.replace("World", "Python"))  # Hello, Python!
print(len(text))         # 13
```

## Numeric Operations

### Basic Arithmetic
```python
a = 10
b = 3

print(a + b)    # Addition: 13
print(a - b)    # Subtraction: 7
print(a * b)    # Multiplication: 30
print(a / b)    # Division: 3.333...
print(a // b)   # Floor division: 3
print(a % b)    # Modulus (remainder): 1
print(a ** b)   # Exponentiation: 1000
```

### Order of Operations
```python
result = 2 + 3 * 4    # Result: 14 (not 20)
result = (2 + 3) * 4  # Result: 20

# PEMDAS: Parentheses, Exponents, Multiplication/Division, Addition/Subtraction
```

## Type Checking and Conversion

### Checking Types
```python
age = 25
name = "Alice"
height = 5.8

print(type(age))     # <class 'int'>
print(type(name))    # <class 'str'>
print(type(height))  # <class 'float'>
```

### Type Conversion
```python
# String to number
age_str = "25"
age = int(age_str)

# Number to string
score = 95
score_str = str(score)

# String to float
price_str = "19.99"
price = float(price_str)
```

## Practical Examples

### Example 1: Simple Calculator
```python
print("Simple Calculator")
print("=" * 20)

num1 = float(input("Enter first number: "))
num2 = float(input("Enter second number: "))

print(f"\n{num1} + {num2} = {num1 + num2}")
print(f"{num1} - {num2} = {num1 - num2}")
print(f"{num1} * {num2} = {num1 * num2}")

if num2 != 0:
    print(f"{num1} / {num2} = {num1 / num2}")
else:
    print("Cannot divide by zero!")
```

### Example 2: Personal Information
```python
# Collect user information
name = input("What's your name? ")
age = int(input("What's your age? "))
city = input("What city do you live in? ")

# Calculate birth year (approximately)
current_year = 2024
birth_year = current_year - age

# Display information
print("\n" + "=" * 30)
print("PERSONAL INFORMATION")
print("=" * 30)
print(f"Name: {name}")
print(f"Age: {age} years old")
print(f"City: {city}")
print(f"Approximate birth year: {birth_year}")
```

## Common Errors and Solutions

### 1. ValueError
```python
# This will cause an error:
age = int(input("Enter your age: "))  # User enters "twenty"

# Solution: Add error handling (we'll learn this later)
```

### 2. TypeError
```python
# This will cause an error:
name = "Alice"
age = 25
message = "Hello " + name + ", you are " + age + " years old"

# Solution: Convert age to string
message = "Hello " + name + ", you are " + str(age) + " years old"
# Or use f-strings
message = f"Hello {name}, you are {age} years old"
```

## Practice Exercises

### Exercise 1: Temperature Converter
Create a program that converts Celsius to Fahrenheit:

Formula: F = (C × 9/5) + 32

```python
# Your code here
celsius = float(input("Enter temperature in Celsius: "))
# Convert to Fahrenheit and display
```

### Exercise 2: Circle Calculator
Calculate the area and circumference of a circle:

```python
import math

radius = float(input("Enter the radius: "))
# Calculate area: π × r²
# Calculate circumference: 2 × π × r
# Display results
```

### Exercise 3: Shopping Receipt
Create a simple shopping receipt:

```python
# Get item information
item1_name = input("First item name: ")
item1_price = float(input("First item price: "))
# Continue for more items...

# Calculate total, tax, and final amount
```

## Key Takeaways

1. **Choose the right data type** for your variables
2. **Always convert user input** to the appropriate type
3. **Use f-strings** for clean string formatting
4. **Be careful with type mixing** - Python is strict about types
5. **Follow naming conventions** for readable code

## Next Week Preview

In Week 3, we'll explore:
- Conditional statements (if, elif, else)
- Comparison operators
- Logical operators
- Decision-making in programs

## Homework

1. Complete all three practice exercises
2. Experiment with different data types
3. Try breaking your programs and fix the errors
4. Create a program that uses all four data types we learned

---

**Pro Tip**: Use the `type()` function whenever you're unsure about a variable's data type!