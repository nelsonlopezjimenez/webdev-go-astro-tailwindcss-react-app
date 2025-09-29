---
title: "Week 50 - Postcredentialing Steps" 
description: "Week50 - steps to do after credentialing a voced laptop"
week: 50
---
# Lesson 50: Post-credentialing Steps

## As Administrator

1. open cmd and gitbash
1. cd into /c/Users/sXXXXXX
1. mv AppData AppData.1.bak
1. mklink /D AppData c:\Users\Public\bin\golden25-AppData
1. logout

## As User sXXXXXX
1. Open Zeal
1. Select edit -> Preferences -> Browse set to c:\Users\Public\bin\Zeal\Zeal
1. copy/paste /c/Users/Public/Pictures/* to %userprofile%\Pictures\
1. cd ~/Pictures/current-old/Webdev and userprofile shortcuts copy/paste to Desktop
1. cd ~ or cd %userprofile%
1. mklink /D .vscode c:\Users\Public\bin\golden-25-root\.vscode
1. cd /c/Users/Public/bin/Desktop-icons/VSCode, chrome, cis-verdaccio, 7zFM.exe, bookmarks_9_24_24.html, Postman shortcuts copy/paste to Desktop
1. cd /c/Users/Public/my-express/public 
1. mklink /D youtube.nel c:\Users\Public\Videos\youtube.nel
1. cd /c/Users/Public/Downloads/2025/edu.gcfglobal.org-2025
1. run
```sh
mkdir -p en4 && for i in *.zip; do mkdir -p "en4/${i%.zip}" && unzip -q "$i" -d "en4/${i%.zip}" ; done
```
1. mkdir /c/websites
1. cd /c/websites
1. mklink /D edu.gcfglobal.org c:\Users\Public\Download\2025\edu.gcfglobal.org-2025\en4 (or not)
1. It depends of the final path as edu.gcfglobal.org or edu.gcfglobal.org/en

## Important Paths
1. Full-Stack React Projects by Shama Hoque: C:\Users\Public\CANVAS_FILE_CACHES\quarter1\
1. Networks, the Internet AIO_4_PPT and AIO_7_PPT at C:\Users\Public\CANVAS_FILE_CACHES\quarter1-in-base-image
1. Technology In Action : Using the Internet: C:\Users\Public\CANVAS_FILE_CACHES\quarter1-in-base-image\tia18e_Chp3
1. Duckett source code C:\Users\Public\Documents\duckett-textbook-src-code
1. RAW-VIDEOS\Q1BigRawVideos
1. Zen Garden : \Public\Videos\youtube.nel\public_html\zen

## Cheat Sheets
1. HTML/CSS : C:\Users\Public\CANVAS_FILE_CACHES\quarter1\a5.pdf
1. Git : C:\Users\Public\CANVAS_FILE_CACHES\quarter1-in-base-image\github-git-cheat-sheet
1. HTML5 cheat sheet: C:\Users\Public\Videos\RAW_VIDEOS\codingheroes-intro-cheat-sheets/html5-cheat-sheet.pdf
1. HTML5 cheat sheet: C:\Users\Public\Videos\RAW_VIDEOS\codingheroes-intro-cheat-sheets/WSU-HTML-Cheat-Sheet.pdf
1. 

# Lesson  50 Content Here

# Week 13: JavaScript Fundamentals

<!-- Solution 4: Use explicit HTML anchors (most reliable) -->
## 📚 Table of Contents
- [Variables and Data Types](#variables-and-data-types)
- [Functions](#functions)
- [Control Structures](#control-structures)
- [Arrays and Objects](#arrays-and-objects)
- [DOM Manipulation](#dom-manipulation)
- [Practice Exercises](#practice-exercises)



<!-- Use HTML anchors with proper IDs -->
<h2 id="variables-and-data-types">Variables and Data Types</h2>

In JavaScript, you can declare variables using `let`, `const`, or `var`...

```javascript
let name = "John";
const age = 25;
var city = "Seattle";
```

<h2 id="functions">Functions</h2>

Functions are reusable blocks of code...

```javascript
function greet(name) {
    return `Hello, ${name}!`;
}
```

<h2 id="control-structures">Control Structures</h2>

Control structures help you make decisions in your code...

```javascript
if (age >= 18) {
    console.log("You are an adult");
}
```

<h2 id="arrays-and-objects">Arrays and Objects</h2>

Arrays and objects are fundamental data structures...

```javascript
const fruits = ["apple", "banana", "orange"];
const person = { name: "John", age: 25 };
```

<h2 id="dom-manipulation">DOM Manipulation</h2>

The Document Object Model allows you to interact with web pages...

```javascript
document.getElementById("myButton").addEventListener("click", function() {
    alert("Button clicked!");
});
```

<h2 id="practice-exercises">Practice Exercises</h2>

Now let's put what you've learned into practice...

### Exercise 1: Variable Practice
Create variables for different data types and log them to the console.

### Exercise 2: Function Building
Build a calculator function that can add, subtract, multiply, and divide.



