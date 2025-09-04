---
title: "Week 5: Advanced CSS and Responsive Design"
description: "Master CSS positioning, flexbox, grid, and responsive design principles for modern web layouts"
week: 5
---

# Week 5: Advanced CSS and Responsive Design

## Learning Objectives
- Master CSS positioning and layout techniques
- Understand responsive design principles
- Implement CSS Grid and Flexbox
- Create mobile-first designs

## The Mobile-First Era

With over 50% of web traffic coming from mobile devices, responsive design isn't optional—it's essential. This week, we'll learn to create websites that work beautifully on any screen size.

## Daily Activities

### Day 1: CSS Positioning

Understanding how elements are positioned on the page:

```css
/* Static positioning (default) */
.static-element {
    position: static; /* Normal document flow */
}

/* Relative positioning - relative to its normal position */
.relative-element {
    position: relative;
    top: 20px;    /* Moves 20px down from normal position */
    left: 10px;   /* Moves 10px right from normal position */
}

/* Absolute positioning - relative to nearest positioned ancestor */
.absolute-element {
    position: absolute;
    top: 50px;    /* 50px from top of positioned parent */
    right: 20px;  /* 20px from right of positioned parent */
}

/* Fixed positioning - relative to viewport */
.fixed-header {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    background-color: white;
    z-index: 1000; /* Ensures it stays on top */
    box-shadow: 0 2px 5px rgba(0,0,0,0.1);
}

/* Sticky positioning - switches between relative and fixed */
.sticky-nav {
    position: sticky;
    top: 0; /* Sticks when scrolling reaches top of viewport */
    background-color: #333;
    padding: 1rem;
}
```

**Practical Example: Card with Badge**
```css
.card {
    position: relative;
    padding: 2rem;
    background: white;
    border-radius: 8px;
    box-shadow: 0 4px 6px rgba(0,0,0,0.1);
}

.card-badge {
    position: absolute;
    top: -10px;
    right: -10px;
    background: #e74c3c;
    color: white;
    padding: 0.5rem;
    border-radius: 50%;
    font-size: 0.875rem;
}
```

### Day 2: Flexbox Layout

Flexbox makes it easy to create flexible, responsive layouts:

```css
/* Flex container */
.flex-container {
    display: flex;
    flex-direction: row; /* row | column | row-reverse | column-reverse */
    justify-content: center; /* flex-start | flex-end | center | space-between | space-around | space-evenly */
    align-items: center; /* flex-start | flex-end | center | stretch | baseline */
    flex-wrap: wrap; /* nowrap | wrap | wrap-reverse */
    gap: 20px; /* Space between items */
}

/* Flex items */
.flex-item {
    flex: 1; /* flex-grow: 1, flex-shrink: 1, flex-basis: 0 */
    min-width: 250px; /* Prevents items from getting too small */
}

/* Responsive navigation with flexbox */
.nav-flex {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1rem;
    background: #2c3e50;
}

.nav-brand {
    font-size: 1.5rem;
    font-weight: bold;
    color: white;
}

.nav-links {
    display: flex;
    list-style: none;
    gap: 2rem;
    margin: 0;
    padding: 0;
}

.nav-links a {
    color: white;
    text-decoration: none;
    transition: color 0.3s;
}

.nav-links a:hover {
    color: #3498db;
}

/* Centering with flexbox */
.center-content {
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 100vh;
}
```

### Day 3: CSS Grid Layout

CSS Grid provides powerful two-dimensional layout capabilities:

```css
/* Grid container */
.grid-container {
    display: grid;
    grid-template-columns: repeat(3, 1fr); /* 3 equal columns */
    grid-template-rows: auto; /* Rows size to content */
    gap: 20px;
    padding: 20px;
}

/* Responsive grid */
.responsive-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    gap: 2rem;
}

/* Named grid areas for layout */
.layout-grid {
    display: grid;
    grid-template-areas: 
        "header header header"
        "sidebar main main"
        "footer footer footer";
    grid-template-columns: 200px 1fr 1fr;
    grid-template-rows: auto 1fr auto;
    gap: 20px;
    min-height: 100vh;
}

.header { grid-area: header; }
.sidebar { grid-area: sidebar; }
.main { grid-area: main; }
.footer { grid-area: footer; }

/* Grid item positioning */
.featured-item {
    grid-column: 1 / 3; /* Spans from column 1 to 3 */
    grid-row: 1 / 3;    /* Spans from row 1 to 3 */
}
```

### Day 4: Responsive Design Fundamentals

Creating websites that adapt to different screen sizes:

```css
/* Mobile-first approach */
/* Base styles for mobile (320px and up) */
.container {
    padding: 1rem;
    max-width: 100%;
}

.grid {
    display: grid;
    grid-template-columns: 1fr;
    gap: 1rem;
}

.nav-menu {
    display: none; /* Hidden by default on mobile */
}

.nav-toggle {
    display: block;
    background: none;
    border: none;
    color: white;
    font-size: 1.5rem;
}

/* Tablet styles (768px and up) */
@media screen and (min-width: 768px) {
    .container {
        padding: 2rem;
        max-width: 750px;
        margin: 0 auto;
    }
    
    .grid {
        grid-template-columns: repeat(2, 1fr);
        gap: 2rem;
    }
    
    .nav-menu {
        display: flex;
    }
    
    .nav-toggle {
        display: none;
    }
}

/* Desktop styles (1024px and up) */
@media screen and (min-width: 1024px) {
    .container {
        max-width: 1200px;
        padding: 3rem;
    }
    
    .grid {
        grid-template-columns: repeat(3, 1fr);
        gap: 3rem;
    }
}

/* Large desktop (1440px and up) */
@media screen and (min-width: 1440px) {
    .container {
        max-width: 1400px;
    }
    
    .grid {
        grid-template-columns: repeat(4, 1fr);
    }
}

/* Responsive images */
img {
    max-width: 100%;
    height: auto;
}

/* Responsive video */
.video-wrapper {
    position: relative;
    padding-bottom: 56.25%; /* 16:9 aspect ratio */
    height: 0;
    overflow: hidden;
}

.video-wrapper iframe {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
}
```

### Day 5: Advanced Responsive Techniques

Fine-tuning responsive design with advanced techniques:

```css
/* Container queries (modern browsers) */
@container (min-width: 400px) {
    .card {
        display: flex;
        align-items: center;
    }
}

/* Responsive typography */
html {
    font-size: 16px; /* Base size for mobile */
}

@media screen and (min-width: 768px) {
    html {
        font-size: 18px;
    }
}

@media screen and (min-width: 1024px) {
    html {
        font-size: 20px;
    }
}

/* Fluid typography with clamp() */
h1 {
    font-size: clamp(2rem, 5vw, 4rem);
}

/* Responsive spacing */
.section {
    padding: clamp(2rem, 5vw, 4rem) 0;
}

/* Print styles */
@media print {
    .no-print {
        display: none;
    }
    
    body {
        font-size: 12pt;
        line-height: 1.4;
        color: black;
        background: white;
    }
    
    a::after {
        content: " (" attr(href) ")";
    }
}

/* Dark mode support */
@media (prefers-color-scheme: dark) {
    :root {
        --background-color: #1a1a1a;
        --text-color: #ffffff;
        --card-background: #2d2d2d;
    }
}

/* Reduced motion preference */
@media (prefers-reduced-motion: reduce) {
    * {
        animation-duration: 0.01ms !important;
        animation-iteration-count: 1 !important;
        transition-duration: 0.01ms !important;
    }
}
```

## Responsive Design Strategy

### 1. Mobile-First Approach
Start with mobile styles and progressively enhance for larger screens:

```css
/* Mobile first (default) */
.element {
    width: 100%;
    padding: 1rem;
}

/* Then add larger screen styles */
@media (min-width: 768px) {
    .element {
        width: 50%;
        padding: 2rem;
    }
}
```

### 2. Common Breakpoints
```css
/* Extra small devices (phones, 600px and down) */
@media only screen and (max-width: 600px) {}

/* Small devices (portrait tablets and large phones, 600px and up) */
@media only screen and (min-width: 600px) {}

/* Medium devices (landscape tablets, 768px and up) */
@media only screen and (min-width: 768px) {}

/* Large devices (laptops/desktops, 992px and up) */
@media only screen and (min-width: 992px) {}

/* Extra large devices (large laptops and desktops, 1200px and up) */
@media only screen and (min-width: 1200px) {}
```

## Hands-On Exercise

Transform your website into a fully responsive design:

1. **Audit Current Layout**: Test your site on different screen sizes
2. **Implement Mobile-First CSS**: Start with mobile styles, then enhance
3. **Add Flexbox Navigation**: Create a responsive navigation menu
4. **Create Grid Layouts**: Use CSS Grid for complex layouts
5. **Test Across Devices**: Ensure consistent experience on all screens

## Testing Responsive Design

### Browser Developer Tools
- Chrome DevTools Device Mode
- Firefox Responsive Design Mode
- Safari Web Inspector

### Physical Device Testing
- Test on actual phones and tablets
- Use tools like BrowserStack or Sauce Labs
- Check performance on slower devices

### Checklist
- [ ] Navigation works on small screens
- [ ] Text is readable without zooming
- [ ] Touch targets are at least 44px
- [ ] Images scale appropriately
- [ ] Forms are easy to use on mobile
- [ ] Page loads quickly on mobile networks

## Resources
- [MDN Responsive Design](https://developer.mozilla.org/en-US/docs/Learn/CSS/CSS_layout/Responsive_Design)
- [CSS Grid Guide](https://css-tricks.com/snippets/css/complete-guide-grid/)
- [Flexbox Guide](https://css-tricks.com/snippets/css/a-guide-to-flexbox/)

## Assignment Checklist
- [ ] Implement mobile-first responsive design
- [ ] Use Flexbox for navigation and card layouts
- [ ] Create at least one CSS Grid layout
- [ ] Test on multiple device sizes
- [ ] Optimize images for different screen densities
- [ ] Ensure touch-friendly interface elements

---

**Next Week**: Links, Navigation, and User Experience