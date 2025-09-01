---
title: "Introduction to Programming"
description: "Get started with programming concepts, setting up your development environment, and writing your first program."
week: 1
---

# Welcome to Programming Fundamentals!

## Learning Objectives

By the end of this lesson, you will be able to:

- Understand what programming is and why it's important
- Set up your development environment
- Write and run your first program
- Understand basic programming terminology

![alt text](image.png)

## What is Programming?

Programming is the process of creating instructions for computers to follow. These instructions, called **code**, tell the computer exactly what to do step by step.

Think of it like writing a recipe:
- A recipe tells a person how to cook a dish
- A program tells a computer how to solve a problem

## Key Programming Concepts

### 1. Algorithm
An algorithm is a set of step-by-step instructions to solve a problem.

**Example**: Making a sandwich
1. Get two slices of bread
2. Spread butter on one slice
3. Add your favorite filling
4. Place the second slice on top

### 2. Variables
Variables are containers that store data values.

```python
name = "Alice"
age = 25
is_student = True
```

### 3. Functions
Functions are reusable blocks of code that perform specific tasks.

```python
def greet(name):
    return f"Hello, {name}!"

message = greet("Alice")
print(message)  # Output: Hello, Alice!
```

## Setting Up Your Environment

### Required Software

1. **Text Editor or IDE**
   - VS Code (recommended)
   - Sublime Text
   - Atom

2. **Programming Language**
   - Python 3.8 or higher
   - Download from [python.org](https://python.org)

3. **Terminal/Command Prompt**
   - Built into your operating system

### Installation Steps

1. Download and install Python
2. Download and install VS Code
3. Install the Python extension in VS Code
4. Open terminal and type `python --version` to verify installation

## Your First Program

Let's write the classic "Hello, World!" program:

```python
# This is your first Python program!
print("Hello, World!")
print("Welcome to Programming Fundamentals!")

# You can also use variables
course_name = "Programming Fundamentals"
week = 1

print(f"You are in {course_name}, Week {week}")
```

### Running Your Program

1. Save the code in a file called `hello.py`
2. Open terminal in the same directory
3. Type: `python hello.py`
4. Press Enter

You should see:
```
Hello, World!
Welcome to Programming Fundamentals!
You are in Programming Fundamentals, Week 1
```

## Practice Exercises

### Exercise 1: Personal Introduction
Write a program that prints information about yourself:

```python
# Your code here
name = "Your Name"
favorite_color = "Your Favorite Color"
hobby = "Your Hobby"

# Print statements to introduce yourself
```

### Exercise 2: Simple Calculator
Create a program that adds two numbers:

```python
# Your code here
number1 = 10
number2 = 5

# Calculate and print the sum
```

## Common Mistakes to Avoid

1. **Spelling Errors**: `print` not `Print`
2. **Missing Quotes**: `"Hello"` not `Hello`
3. **Incorrect Indentation**: Python is sensitive to spaces
4. **Forgetting Colons**: After `if`, `for`, `def`, etc.

## Next Week Preview

In Week 2, we'll dive deeper into:
- Data types (numbers, strings, booleans)
- User input and output
- Basic operations and expressions
- Conditional statements

## Resources

- [Python Official Tutorial](https://docs.python.org/3/tutorial/)
- [Python.org Beginner's Guide](https://wiki.python.org/moin/BeginnersGuide)
- [VS Code Python Tutorial](https://code.visualstudio.com/docs/python/python-tutorial)

## Homework

1. Complete both practice exercises
2. Experiment with different print statements
3. Try changing the values in variables and see what happens
4. Read Chapter 1 of your textbook

---

**Remember**: Programming is a skill that improves with practice. Don't worry if everything doesn't make sense immediately – that's completely normal!