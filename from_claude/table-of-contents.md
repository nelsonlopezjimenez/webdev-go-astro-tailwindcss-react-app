There are several easy ways to create a table of contents with local anchor links in Markdown. Here are the most common approaches:## The Easiest Methods:

### **1. Simple Bullet List (Most Common)**
```markdown
## Table of Contents
- [Variables and Data Types](#variables-and-data-types)
- [Functions](#functions)
- [DOM Manipulation](#dom-manipulation)
```

### **2. Auto-Generated Anchors**
Markdown automatically creates anchors from headers:
- `## Variables and Data Types` → `#variables-and-data-types`
- `## DOM Manipulation` → `#dom-manipulation`

### **Rules for Auto-Generated Anchors:**
- Lowercase everything
- Replace spaces with hyphens
- Remove special characters
- `## Part 1: JavaScript Basics` → `#part-1-javascript-basics`

### **3. For Your Course Lessons:**
Since you're using Astro, I'd recommend this format for your week files:

```markdown
---
title: "Week 13: JavaScript Fundamentals"
description: "Learn JavaScript basics including variables, functions, and DOM manipulation"
week: 13
---

# Week 13: JavaScript Fundamentals

## 📚 Quick Navigation
- [Day 1: Variables and Data Types](#day-1-variables-and-data-types)
- [Day 2: Functions](#day-2-functions)
- [Day 3: Control Structures](#day-3-control-structures)
- [Day 4: Arrays and Objects](#day-4-arrays-and-objects)
- [Day 5: DOM Manipulation](#day-5-dom-manipulation)

---

## Day 1: Variables and Data Types
Content here...

## Day 2: Functions
Content here...
```

**Which style would work best for your lesson structure?** The simple bullet list is easiest to maintain and works everywhere.

# Week 13: JavaScript Fundamentals

## Table of Contents
- [Variables and Data Types](#variables-and-data-types)
- [Functions](#functions)
- [Control Structures](#control-structures)
- [Arrays and Objects](#arrays-and-objects)
- [DOM Manipulation](#dom-manipulation)
- [Practice Exercises](#practice-exercises)

---

## Variables and Data Types

Content about variables...

```javascript
let name = "John";
const age = 25;
```

## Functions

Content about functions...

```javascript
function greet(name) {
    return `Hello, ${name}!`;
}
```

## Control Structures

Content about if statements, loops...

```javascript
if (age >= 18) {
    console.log("Adult");
}
```

## Arrays and Objects

Content about arrays and objects...

```javascript
const fruits = ["apple", "banana", "orange"];
const person = { name: "John", age: 25 };
```

## DOM Manipulation

Content about working with the DOM...

```javascript
document.getElementById("myButton").addEventListener("click", function() {
    alert("Button clicked!");
});
```

## Practice Exercises

1. Create a calculator function
2. Build a simple to-do list
3. Make an interactive button


# Week 13: JavaScript Fundamentals

<!-- Method 1: Simple bullet list -->
## 📚 Table of Contents
- [🔢 Variables and Data Types](#variables-and-data-types)
- [⚡ Functions](#functions)
- [🔄 Control Structures](#control-structures)
- [📦 Arrays and Objects](#arrays-and-objects)
- [🌐 DOM Manipulation](#dom-manipulation)
- [💻 Practice Exercises](#practice-exercises)

---

<!-- Method 2: Numbered list with estimated time -->
## 📋 Learning Path
1. [Variables and Data Types](#variables-and-data-types) *(15 minutes)*
2. [Functions](#functions) *(20 minutes)*
3. [Control Structures](#control-structures) *(25 minutes)*
4. [Arrays and Objects](#arrays-and-objects) *(30 minutes)*
5. [DOM Manipulation](#dom-manipulation) *(20 minutes)*
6. [Practice Exercises](#practice-exercises) *(45 minutes)*

**Total Estimated Time: 2.5 hours**

---

<!-- Method 3: Categorized sections -->
## 🗂️ Quick Navigation

### Core Concepts
- [Variables and Data Types](#variables-and-data-types)
- [Functions](#functions)

### Programming Logic
- [Control Structures](#control-structures)
- [Arrays and Objects](#arrays-and-objects)

### Web Development
- [DOM Manipulation](#dom-manipulation)
- [Practice Exercises](#practice-exercises)

---

<!-- Method 4: With descriptions -->
## 📖 What You'll Learn

| Topic | Description | Link |
|-------|-------------|------|
| Variables | Learn about let, const, var | [Go to Variables](#variables-and-data-types) |
| Functions | Create reusable code blocks | [Go to Functions](#functions) |
| Control Flow | If statements, loops | [Go to Control](#control-structures) |
| Data Structures | Arrays and objects | [Go to Arrays](#arrays-and-objects) |
| DOM | Interact with web pages | [Go to DOM](#dom-manipulation) |
| Practice | Hands-on exercises | [Go to Exercises](#practice-exercises) |

---

<!-- The actual content sections -->
## Variables and Data Types

### Learning Objectives
- Understand the difference between let, const, and var
- Learn about primitive data types
- Practice variable declaration and assignment

### Content
JavaScript has several ways to declare variables...

## Functions

### Learning Objectives
- Create function declarations and expressions
- Understand parameters and return values
- Learn about arrow functions

### Content
Functions are reusable blocks of code...

## Control Structures

### Learning Objectives
- Use if/else statements for decision making
- Implement for and while loops
- Understand break and continue

### Content
Control structures help you make decisions...

## Arrays and Objects

### Learning Objectives
- Create and manipulate arrays
- Work with object properties and methods
- Understand array methods like map, filter, forEach

### Content
Arrays and objects are fundamental data structures...

## DOM Manipulation

### Learning Objectives
- Select elements with querySelector
- Modify element content and attributes
- Handle user events

### Content
The Document Object Model (DOM) allows you to...

## Practice Exercises

### Exercise 1: Variable Practice
Create variables for different data types...

### Exercise 2: Function Building
Build a calculator function...

### Exercise 3: DOM Interaction
Create an interactive webpage...

# Week 13: JavaScript Fundamentals

<!-- Method 5: Custom anchor names (if automatic ones don't work) -->
## Table of Contents
- [Introduction](#intro)
- [Part 1: Basics](#part1)
- [Part 2: Advanced](#part2)
- [Troubleshooting](#troubleshoot)
- [Next Steps](#next)

---

<!-- Using custom anchor IDs -->
<a id="intro"></a>
## Introduction

Welcome to JavaScript fundamentals...

<a id="part1"></a>
## Part 1: The Basics

Let's start with variables...

<a id="part2"></a>
## Part 2: Advanced Concepts

Now we'll explore more complex topics...

<a id="troubleshoot"></a>
## Common Issues and Troubleshooting

If you encounter problems...

<a id="next"></a>
## What's Next?

In the next lesson, we'll cover...

---

<!-- Method 6: For Astro/React components (if you need HTML) -->
## Interactive Table of Contents

<div class="toc-container">
  <h3>📚 Jump to Section</h3>
  <div class="toc-grid">
    <a href="#variables" class="toc-item">
      <strong>Variables</strong>
      <small>Data storage basics</small>
    </a>
    <a href="#functions" class="toc-item">
      <strong>Functions</strong>
      <small>Reusable code blocks</small>
    </a>
    <a href="#dom" class="toc-item">
      <strong>DOM</strong>
      <small>Web page interaction</small>
    </a>
  </div>
</div>

---

<!-- Method 7: Auto-generated TOC (if supported by your processor) -->
<!-- Some Markdown processors support this: -->
[TOC]

<!-- Or this: -->
[[TOC]]

---

<!-- Method 8: Back to top links -->
## Variables and Data Types

Content here...

[↑ Back to Top](#table-of-contents)

## Functions

More content...

[↑ Back to Top](#table-of-contents)

---

<!-- Method 9: Collapsible sections (GitHub-style) -->
<details>
<summary><strong>📋 Detailed Table of Contents</strong></summary>

- **Introduction**
  - [What is JavaScript?](#what-is-javascript)
  - [Setting up your environment](#setting-up-environment)
  
- **Variables and Data Types**
  - [Variable declarations](#variable-declarations)
  - [Primitive types](#primitive-types)
  - [Type conversion](#type-conversion)
  
- **Functions**
  - [Function declarations](#function-declarations)
  - [Arrow functions](#arrow-functions)
  - [Scope and closures](#scope-and-closures)

</details>

---

<!-- Method 10: Quick reference card style -->
## 🎯 Quick Reference

| Section | Key Concepts | Time |
|---------|--------------|------|
| [Variables](#variables-and-data-types) | `let`, `const`, `var`, types | 15 min |
| [Functions](#functions) | declarations, expressions, arrows | 20 min |
| [Control](#control-structures) | `if`, loops, `switch` | 25 min |
| [Arrays](#arrays-and-objects) | `[]`, methods, iteration | 30 min |
| [DOM](#dom-manipulation) | selectors, events, manipulation | 20 min |

**🏁 Total: ~2 hours**

