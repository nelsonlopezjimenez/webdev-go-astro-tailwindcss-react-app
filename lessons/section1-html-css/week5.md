---
title: "Week 4: Introduction to CSS3"
description: "Learn CSS syntax, selectors, box model, colors, fonts, and basic styling techniques"
week: 5
---

# Week 4: Introduction to CSS3

## Learning Objectives
- Understand CSS syntax and selectors
- Apply basic styling to HTML elements
- Learn the box model concept
- Implement colors, fonts, and basic layouts

## What is CSS?

**CSS (Cascading Style Sheets)** separates presentation from content. While HTML provides structure and meaning, CSS controls how that content looks and is laid out on the page.

## Daily Activities

### Day 1: CSS Syntax and Selectors

Understanding how to target and style HTML elements:

```css
/* Element selector - targets all p elements */
p {
    color: #333;
    font-size: 16px;
    line-height: 1.5;
}

/* Class selector - targets elements with class="highlight" */
.highlight {
    background-color: yellow;
    padding: 10px;
}

/* ID selector - targets element with id="header" */
#header {
    background-color: #2c3e50;
    color: white;
    text-align: center;
}

/* Descendant selector - targets spans inside paragraphs */
p span {
    font-weight: bold;
}

/* Multiple selectors */
h1, h2, h3 {
    color: #2c3e50;
    margin-bottom: 1rem;
}
```

**Linking CSS to HTML:**

```html
<!-- External stylesheet (recommended) -->
<link rel="stylesheet" href="styles.css">

<!-- Internal stylesheet -->
<style>
    body { font-family: Arial, sans-serif; }
</style>

<!-- Inline styles (use sparingly) -->
<p style="color: red;">This text is red</p>
```

### Day 2: The CSS Box Model

Every HTML element is a rectangular box with four parts:

```css
.box {
    /* Content area */
    width: 300px;
    height: 200px;
    
    /* Padding - space inside the border */
    padding: 20px;
    
    /* Border - outline around the padding */
    border: 2px solid #333;
    
    /* Margin - space outside the border */
    margin: 10px;
    
    /* Background only covers content + padding */
    background-color: lightblue;
}

/* Box-sizing changes how width/height are calculated */
.better-box {
    box-sizing: border-box; /* Includes padding and border in width/height */
    width: 300px;
    padding: 20px;
    border: 2px solid #333;
}
```

**Visual representation:**
```
┌─────────────────────────────┐
│         MARGIN              │
│  ┌─────────────────────┐    │
│  │       BORDER        │    │
│  │  ┌───────────────┐  │    │
│  │  │   PADDING     │  │    │
│  │  │  ┌─────────┐  │  │    │
│  │  │  │CONTENT  │  │  │    │
│  │  │  └─────────┘  │  │    │
│  │  └───────────────┘  │    │
│  └─────────────────────┘    │
└─────────────────────────────┘
```

### Day 3: Typography and Colors

Making text readable and visually appealing:

```css
body {
    /* Font families with fallbacks */
    font-family: 'Arial', 'Helvetica', sans-serif;
    font-size: 16px;
    line-height: 1.6;
    color: #333;
}

h1 {
    font-size: 2.5em; /* Relative to parent element */
    font-weight: 700;
    color: #2c3e50;
    text-align: center;
    margin-bottom: 1rem;
}

/* Different color formats */
.color-examples {
    color: red;                    /* Named color */
    color: #ff0000;               /* Hexadecimal */
    color: rgb(255, 0, 0);        /* RGB */
    color: rgba(255, 0, 0, 0.7);  /* RGB with alpha (transparency) */
    color: hsl(0, 100%, 50%);     /* HSL (Hue, Saturation, Lightness) */
}

/* Text properties */
.text-styling {
    font-style: italic;
    font-weight: bold;
    text-decoration: underline;
    text-transform: uppercase;
    letter-spacing: 2px;
    word-spacing: 5px;
}
```

### Day 4: Basic Layout with CSS

Creating simple layouts and positioning:

```css
/* Simple layout container */
.container {
    max-width: 1200px;
    margin: 0 auto; /* Centers the container */
    padding: 0 20px;
}

/* Basic flexbox layout */
.flex-container {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 20px;
}

.flex-item {
    flex: 1; /* Takes equal space */
    padding: 20px;
    background-color: #f4f4f4;
    border-radius: 5px;
}

/* Float layout (older method, but still useful) */
.float-left {
    float: left;
    width: 48%;
    margin-right: 2%;
}

.float-right {
    float: right;
    width: 48%;
}

.clearfix::after {
    content: "";
    display: table;
    clear: both;
}
```

### Day 5: CSS Organization and Best Practices

Writing maintainable and organized CSS:

```css
/* CSS Reset or Normalize */
* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

/* CSS Custom Properties (Variables) */
:root {
    --primary-color: #3498db;
    --secondary-color: #2ecc71;
    --text-color: #333;
    --background-color: #f8f9fa;
    --max-width: 1200px;
    --border-radius: 5px;
}

/* Using variables */
.button {
    background-color: var(--primary-color);
    color: white;
    padding: 12px 24px;
    border: none;
    border-radius: var(--border-radius);
    cursor: pointer;
    transition: background-color 0.3s ease;
}

.button:hover {
    background-color: var(--secondary-color);
}

/* BEM methodology for naming */
.card { /* Block */ }
.card__title { /* Element */ }
.card__title--large { /* Modifier */ }
```

## Hands-On Exercise

Style your multi-page website from Week 3 using CSS:

1. Create a separate `styles.css` file
2. Link it to all your HTML pages
3. Apply consistent typography across all pages
4. Style your navigation menu
5. Create a simple layout with header, main content, and footer
6. Use the box model to add spacing and borders
7. Implement a cohesive color scheme

## CSS File Organization

```css
/* styles.css structure */

/* 1. CSS Reset/Normalize */
/* 2. CSS Custom Properties */
/* 3. Base styles (html, body, headings) */
/* 4. Layout (containers, grids) */
/* 5. Components (buttons, cards, forms) */
/* 6. Utilities (text-center, margin helpers) */
/* 7. Media queries (responsive design) */
```

## Common CSS Properties Reference

**Text and Fonts:**
- `font-family`, `font-size`, `font-weight`, `font-style`
- `color`, `text-align`, `text-decoration`, `line-height`

**Box Model:**
- `width`, `height`, `padding`, `margin`, `border`
- `box-sizing`

**Background:**
- `background-color`, `background-image`, `background-size`

**Display and Position:**
- `display`, `position`, `float`, `clear`

## Resources
- [MDN CSS Basics](https://developer.mozilla.org/en-US/docs/Learn/Getting_started_with_the_web/CSS_basics)
- [W3Schools CSS Tutorial](https://www.w3schools.com/css/)
- [CSS Tricks Almanac](https://css-tricks.com/almanac/)

## Assignment Checklist
- [ ] Create external CSS file linked to all pages
- [ ] Apply consistent typography using web-safe fonts
- [ ] Implement color scheme using CSS variables
- [ ] Style navigation menu with hover effects
- [ ] Use box model properties for spacing and layout
- [ ] Test your site in different browsers

---

**Next Week**: Advanced CSS and Responsive Design