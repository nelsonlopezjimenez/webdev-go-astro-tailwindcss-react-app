---
title: "Lesson 2: HTML5 Fundamentals and Document Structure"
description: "Master HTML5 semantic elements and attributes, proper document structure, and create well-structured web pages"
week: 3
---

# Week 2: HTML5 Fundamentals and Document Structure

## Learning Objectives
- Master HTML5 semantic elements and attributes
- Understand proper document structure
- Learn HTML5 form basics
- Create well-structured web pages

## Daily Activities

### Day 1: HTML5 Semantic Elements

HTML5 introduced semantic elements that give meaning to your content structure:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Semantic HTML5 Structure</title>
</head>
<body>
    <!-- Header contains site navigation and branding -->
    <header>
        <nav>
            <ul>
                <li><a href="#home">Home</a></li>
                <li><a href="#about">About</a></li>
                <li><a href="#contact">Contact</a></li>
            </ul>
        </nav>
    </header>
    
    <!-- Main content area -->
    <main>
        <section id="home">
            <h1>Welcome to My Site</h1>
            <p>This is the main content area.</p>
        </section>
        
        <aside>
            <!-- Sidebar content like related links or ads -->
            <h3>Quick Links</h3>
            <ul>
                <li><a href="#">Link 1</a></li>
                <li><a href="#">Link 2</a></li>
            </ul>
        </aside>
    </main>
    
    <!-- Footer contains copyright and additional links -->
    <footer>
        <p>&copy; 2024 My Website. All rights reserved.</p>
    </footer>
</body>
</html>
```

### Day 2: Text Content and Formatting

Understanding the hierarchy and types of text content:

```html
<!-- Heading hierarchy is important for accessibility and SEO -->
<h1>Main Page Title</h1>
<h2>Section Title</h2>
<h3>Subsection Title</h3>

<!-- Different types of text content -->
<p>Regular paragraph text with <strong>important content</strong> and <em>emphasized text</em>.</p>

<!-- Lists for organized information -->
<ul>
    <li>Unordered list item 1</li>
    <li>Unordered list item 2</li>
</ul>

<ol>
    <li>Ordered list item 1</li>
    <li>Ordered list item 2</li>
</ol>

<!-- Definition lists for terms and descriptions -->
<dl>
    <dt>HTML</dt>
    <dd>HyperText Markup Language</dd>
    <dt>CSS</dt>
    <dd>Cascading Style Sheets</dd>
</dl>
```

### Day 3: Block vs Inline Elements

Understanding element behavior is crucial for layout:

```html
<!-- Block elements take full width and start new lines -->
<div>This is a generic block container</div>
<p>Paragraphs are block elements</p>
<section>Sections are block elements</section>

<!-- Inline elements only take necessary space -->
<span>This is an inline container</span>
<a href="#">Links are inline elements</a>
<strong>Strong text is inline</strong>
```

### Day 4: HTML Validation and Best Practices

**Key practices for clean HTML:**
- Use W3C Markup Validator to check your code
- Proper nesting of elements (no overlapping tags)
- Include required attributes and accessibility features
- Consistent code formatting and indentation
- Meaningful class and ID names

### Day 5: Document Metadata

The `<head>` section contains important metadata:

```html
<head>
    <!-- Essential meta tags -->
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="A brief description of your page for search engines">
    <meta name="keywords" content="web development, HTML, CSS">
    <meta name="author" content="Your Name">
    
    <title>Descriptive Page Title - Site Name</title>
    
    <!-- Link to external stylesheets -->
    <link rel="stylesheet" href="styles.css">
    
    <!-- Favicon -->
    <link rel="icon" href="favicon.ico" type="image/x-icon">
</head>
```

## Hands-On Exercise

Build a structured "My Favorite Movies" page using semantic HTML5 elements. Include:
- Proper document structure with header, main, and footer
- At least 3 different heading levels
- Multiple sections for different movie genres
- Lists of movies with descriptions
- Proper text formatting and emphasis

## Assignment Requirements

Create a complete HTML5 document that demonstrates:
1. Semantic structure using header, nav, main, section, aside, footer
2. Proper heading hierarchy (h1-h6)
3. Various list types (ul, ol, dl)
4. Text formatting elements (strong, em, mark, etc.)
5. Valid HTML5 markup (test with W3C validator)

## Resources
- [MDN HTML Element Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element)
- [W3Schools HTML5 Semantic Elements](https://www.w3schools.com/html/html5_semantic_elements.asp)
- [W3C Markup Validator](https://validator.w3.org/)

## Common Mistakes to Avoid
- Using `<div>` and `<span>` when semantic elements are more appropriate
- Skipping heading levels (going from h1 directly to h3)
- Forgetting closing tags or improperly nesting elements
- Not including required attributes like `alt` for images

---

**Next Week**: Advanced HTML5 and Site Structure