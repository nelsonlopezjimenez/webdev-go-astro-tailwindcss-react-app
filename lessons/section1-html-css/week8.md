---
title: "Week 7: HTML Tables and Data Presentation"
description: "Master HTML tables for displaying structured data with proper semantics, accessibility, and responsive design"
week: 8
---

# Week 7: HTML Tables and Data Presentation

## Learning Objectives
- Create semantic HTML tables for structured data
- Implement accessible table design
- Style tables with advanced CSS techniques
- Make tables responsive for mobile devices
- Present data effectively and professionally

## When to Use Tables

Tables should be used for **tabular data** - information that has a logical relationship between rows and columns. Don't use tables for layout purposes; use CSS Grid or Flexbox instead.

**Good uses for tables:**
- Financial reports and pricing
- Comparison charts
- Scientific data
- Schedules and timetables
- Contact directories

## Daily Activities

### Day 1: Basic Table Structure

Understanding the semantic structure of HTML tables:

```html
<!-- Simple table with proper structure -->
<table>
    <!-- Table caption for accessibility -->
    <caption>Monthly Sales Report - Q1 2024</caption>
    
    <!-- Table header -->
    <thead>
        <tr>
            <th scope="col">Month</th>
            <th scope="col">Sales ($)</th>
            <th scope="col">Growth (%)</th>
            <th scope="col">Target ($)</th>
        </tr>
    </thead>
    
    <!-- Table body -->
    <tbody>
        <tr>
            <td>January</td>
            <td>$25,000</td>
            <td>+5.2%</td>
            <td>$24,000</td>
        </tr>
        <tr>
            <td>February</td>
            <td>$28,500</td>
            <td>+14.0%</td>
            <td>$25,000</td>
        </tr>
        <tr>
            <td>March</td>
            <td>$31,200</td>
            <td>+9.5%</td>
            <td>$26,000</td>
        </tr>
    </tbody>
    
    <!-- Table footer (optional) -->
    <tfoot>
        <tr>
            <th scope="row">Total</th>
            <td>$84,700</td>
            <td>+9.6%</td>
            <td>$75,000</td>
        </tr>
    </tfoot>
</table>
```

**Key Elements:**
- `<table>`: Container for the entire table
- `<caption>`: Describes the table's content (important for accessibility)
- `<thead>`: Groups header content
- `<tbody>`: Groups body content
- `<tfoot>`: Groups footer content
- `<th>`: Header cells with `scope` attribute
- `<td>`: Data cells

### Day 2: Advanced Table Features

Complex tables with spanning cells and relationships:

```html
<!-- Complex table with spanning cells -->
<table>
    <caption>Product Comparison Chart</caption>
    <thead>
        <tr>
            <th rowspan="2" scope="col">Product</th>
            <th colspan="2" scope="colgroup">Specifications</th>
            <th colspan="2" scope="colgroup">Pricing</th>
            <th rowspan="2" scope="col">Rating</th>
        </tr>
        <tr>
            <th scope="col">Weight</th>
            <th scope="col">Dimensions</th>
            <th scope="col">MSRP</th>
            <th scope="col">Sale Price</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <th scope="row">Laptop A</th>
            <td>2.1 kg</td>
            <td>35 × 24 × 2 cm</td>
            <td>$999</td>
            <td>$799</td>
            <td>4.5/5</td>
        </tr>
        <tr>
            <th scope="row">Laptop B</th>
            <td>1.8 kg</td>
            <td>33 × 23 × 1.8 cm</td>
            <td>$1,299</td>
            <td>$1,099</td>
            <td>4.8/5</td>
        </tr>
        <tr>
            <th scope="row">Laptop C</th>
            <td>2.3 kg</td>
            <td>36 × 25 × 2.2 cm</td>
            <td>$1,499</td>
            <td>$1,299</td>
            <td>4.7/5</td>
        </tr>
    </tbody>
</table>

<!-- Table with grouped data -->
<table>
    <caption>Employee Directory by Department</caption>
    <thead>
        <tr>
            <th scope="col">Department</th>
            <th scope="col">Name</th>
            <th scope="col">Position</th>
            <th scope="col">Email</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <th scope="rowgroup" rowspan="3">Development</th>
            <td>Sarah Johnson</td>
            <td>Lead Developer</td>
            <td>sarah@company.com</td>
        </tr>
        <tr>
            <td>Mike Chen</td>
            <td>Frontend Developer</td>
            <td>mike@company.com</td>
        </tr>
        <tr>
            <td>Lisa Rodriguez</td>
            <td>Backend Developer</td>
            <td>lisa@company.com</td>
        </tr>
    </tbody>
</table>
```

**Important attributes:**
- `scope="col"`: Header applies to column
- `scope="row"`: Header applies to row  
- `scope="colgroup"`: Header applies to group of columns
- `scope="rowgroup"`: Header applies to group of rows
- `colspan="n"`: Cell spans n columns
- `rowspan="n"`: Cell spans n rows

### Day 3: Table Styling with CSS

Making tables visually appealing and readable:

```css
/* Basic table styling */
table {
    width: 100%;
    border-collapse: collapse; /* Removes double borders */
    margin: 2rem 0;
    font-size: 0.9rem;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    border-radius: 8px;
    overflow: hidden;
}

/* Table caption styling */
caption {
    font-size: 1.2rem;
    font-weight: 600;
    margin-bottom: 1rem;
    color: #2c3e50;
    text-align: left;
}

/* Header styling */
thead th {
    background: linear-gradient(135deg, #3498db, #2980b9);
    color: white;
    padding: 1rem;
    text-align: left;
    font-weight: 600;
    text-transform: uppercase;
    letter-spacing: 0.5px;
}

/* Body cell styling */
tbody td,
tbody th {
    padding: 0.75rem 1rem;
    border-bottom: 1px solid #ecf0f1;
    vertical-align: top;
}

/* Row header styling */
tbody th {
    background-color: #f8f9fa;
    font-weight: 600;
    color: #2c3e50;
}

/* Zebra striping for better readability */
tbody tr:nth-child(even) {
    background-color: #f8f9fa;
}

tbody tr:nth-child(odd) {
    background-color: white;
}

/* Hover effect */
tbody tr:hover {
    background-color: #e3f2fd;
    transform: scale(1.01);
    transition: all 0.3s ease;
}

/* Footer styling */
tfoot td,
tfoot th {
    background-color: #34495e;
    color: white;
    font-weight: 600;
    padding: 1rem;
    border-top: 3px solid #2c3e50;
}

/* Numeric data alignment */
.numeric {
    text-align: right;
    font-family: 'Courier New', monospace;
}

/* Status indicators */
.status {
    padding: 0.25rem 0.75rem;
    border-radius: 20px;
    font-size: 0.875rem;
    font-weight: 500;
    text-transform: uppercase;
}

.status.active {
    background-color: #d4edda;
    color: #155724;
}

.status.inactive {
    background-color: #f8d7da;
    color: #721c24;
}

.status.pending {
    background-color: #fff3cd;
    color: #856404;
}
```

### Day 4: Responsive Tables

Making tables work on mobile devices:

```css
/* Responsive table wrapper */
.table-wrapper {
    overflow-x: auto;
    margin: 1rem 0;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

/* Horizontal scroll for wide tables */
@media screen and (max-width: 768px) {
    .table-wrapper {
        -webkit-overflow-scrolling: touch;
    }
    
    table {
        min-width: 600px; /* Ensures table doesn't shrink too much */
    }
}

/* Alternative: Stacked table for mobile */
@media screen and (max-width: 600px) {
    .responsive-table thead {
        display: none;
    }
    
    .responsive-table tbody,
    .responsive-table tr,
    .responsive-table td {
        display: block;
    }
    
    .responsive-table tr {
        border: 2px solid #ddd;
        margin-bottom: 1rem;
        padding: 1rem;
        border-radius: 8px;
        background: white;
    }
    
    .responsive-table td {
        border: none;
        padding: 0.5rem 0;
        position: relative;
        padding-left: 120px;
    }
    
    .responsive-table td:before {
        content: attr(data-label) ": ";
        position: absolute;
        left: 0;
        width: 110px;
        font-weight: 600;
        color: #2c3e50;
    }
}
```

```html
<!-- Table with data labels for mobile -->
<table class="responsive-table">
    <thead>
        <tr>
            <th>Product</th>
            <th>Price</th>
            <th>Stock</th>
            <th>Status</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td data-label="Product">Laptop Pro</td>
            <td data-label="Price">$1,299</td>
            <td data-label="Stock">15</td>
            <td data-label="Status">In Stock</td>
        </tr>
        <tr>
            <td data-label="Product">Tablet Air</td>
            <td data-label="Price">$599</td>
            <td data-label="Stock">0</td>
            <td data-label="Status">Out of Stock</td>
        </tr>
    </tbody>
</table>
```

### Day 5: Advanced Table Features

Adding interactivity and enhanced functionality:

```html
<!-- Table with sorting and filtering -->
<div class="table-controls">
    <div class="search-box">
        <input type="search" id="table-search" placeholder="Search products...">
    </div>
    <div class="filter-controls">
        <select id="category-filter">
            <option value="">All Categories</option>
            <option value="laptops">Laptops</option>
            <option value="tablets">Tablets</option>
            <option value="phones">Phones</option>
        </select>
        <select id="sort-column">
            <option value="">Sort by...</option>
            <option value="name">Name A-Z</option>
            <option value="price-low">Price Low-High</option>
            <option value="price-high">Price High-Low</option>
            <option value="rating">Rating</option>
        </select>
    </div>
</div>

<table id="product-table" class="sortable-table">
    <thead>
        <tr>
            <th data-sort="name" class="sortable">
                Product Name 
                <span class="sort-indicator">↕</span>
            </th>
            <th data-sort="category" class="sortable">
                Category
                <span class="sort-indicator">↕</span>
            </th>
            <th data-sort="price" class="sortable numeric">
                Price
                <span class="sort-indicator">↕</span>
            </th>
            <th data-sort="rating" class="sortable numeric">
                Rating
                <span class="sort-indicator">↕</span>
            </th>
            <th>Actions</th>
        </tr>
    </thead>
    <tbody>
        <tr data-category="laptops">
            <td>MacBook Pro 16"</td>
            <td>Laptops</td>
            <td class="numeric">$2,399</td>
            <td class="numeric">4.8/5</td>
            <td>
                <button class="btn-small">View</button>
                <button class="btn-small">Edit</button>
            </td>
        </tr>
        <!-- More rows... -->
    </tbody>
</table>
```

```css
/* Table controls styling */
.table-controls {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 1rem;
    gap: 1rem;
    flex-wrap: wrap;
}

.search-box input {
    padding: 0.5rem;
    border: 2px solid #ddd;
    border-radius: 4px;
    width: 250px;
}

.filter-controls {
    display: flex;
    gap: 1rem;
}

.filter-controls select {
    padding: 0.5rem;
    border: 2px solid #ddd;
    border-radius: 4px;
}

/* Sortable headers */
.sortable {
    cursor: pointer;
    position: relative;
    user-select: none;
}

.sortable:hover {
    background-color: rgba(255, 255, 255, 0.1);
}

.sort-indicator {
    margin-left: 0.5rem;
    opacity: 0.5;
}

.sortable.asc .sort-indicator::after {
    content: "↑";
    opacity: 1;
}

.sortable.desc .sort-indicator::after {
    content: "↓";
    opacity: 1;
}

/* Small buttons for table actions */
.btn-small {
    padding: 0.25rem 0.75rem;
    margin-right: 0.5rem;
    border: 1px solid #ddd;
    background: white;
    border-radius: 3px;
    cursor: pointer;
    font-size: 0.875rem;
}

.btn-small:hover {
    background: #f8f9fa;
}
```

## Table Accessibility Best Practices

### Essential Features
1. **Always use `<caption>`** to describe table content
2. **Use `scope` attributes** on header cells
3. **Provide `summary` attribute** for complex tables
4. **Ensure sufficient color contrast**
5. **Make tables keyboard navigable**

### Complex Table Example
```html
<table role="table" aria-labelledby="financial-caption">
    <caption id="financial-caption">
        Quarterly Financial Results - Revenue and Expenses by Region
    </caption>
    <thead>
        <tr>
            <th scope="col" id="region">Region</th>
            <th scope="col" id="q1-rev">Q1 Revenue</th>
            <th scope="col" id="q1-exp">Q1 Expenses</th>
            <th scope="col" id="q2-rev">Q2 Revenue</th>
            <th scope="col" id="q2-exp">Q2 Expenses</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <th scope="row" id="north">North America</th>
            <td headers="north q1-rev">$2.1M</td>
            <td headers="north q1-exp">$1.8M</td>
            <td headers="north q2-rev">$2.4M</td>
            <td headers="north q2-exp">$2.0M</td>
        </tr>
    </tbody>
</table>
```

## Common Table Patterns

### Pricing Table
```html
<table class="pricing-table">
    <thead>
        <tr>
            <th scope="col">Feature</th>
            <th scope="col">Basic</th>
            <th scope="col">Pro</th>
            <th scope="col">Enterprise</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <th scope="row">Price</th>
            <td>Free</td>
            <td>$19/month</td>
            <td>$49/month</td>
        </tr>
        <tr>
            <th scope="row">Storage</th>
            <td>5GB</td>
            <td>100GB</td>
            <td>1TB</td>
        </tr>
        <tr>
            <th scope="row">Users</th>
            <td>1</td>
            <td>5</td>
            <td>Unlimited</td>
        </tr>
    </tbody>
</table>
```

### Schedule Table
```html
<table class="schedule-table">
    <caption>Weekly Course Schedule</caption>
    <thead>
        <tr>
            <th scope="col">Time</th>
            <th scope="col">Monday</th>
            <th scope="col">Tuesday</th>
            <th scope="col">Wednesday</th>
            <th scope="col">Thursday</th>
            <th scope="col">Friday</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <th scope="row">9:00 AM</th>
            <td>HTML Basics</td>
            <td>CSS Styling</td>
            <td>JavaScript</td>
            <td>React</td>
            <td>Project Work</td>
        </tr>
    </tbody>
</table>
```

## Hands-On Exercise

Create a comprehensive data table for your final project:

**Project Ideas:**
- **Portfolio**: Project comparison table
- **Business Site**: Service pricing table or team directory
- **Educational Site**: Course schedule or comparison chart

**Requirements:**
1. Use proper semantic markup with caption, thead, tbody
2. Include appropriate scope attributes
3. Style with CSS for visual appeal
4. Make it responsive for mobile devices
5. Test with screen readers or accessibility tools

## Resources
- [MDN HTML Tables](https://developer.mozilla.org/en-US/docs/Learn/HTML/Tables)
- [W3Schools HTML Tables](https://www.w3schools.com/html/html_tables.asp)
- [WebAIM Table Accessibility](https://webaim.org/techniques/tables/)

## Assignment Checklist
- [ ] Create semantic table structure
- [ ] Add proper accessibility attributes
- [ ] Style table with CSS for readability
- [ ] Implement responsive design
- [ ] Test with different screen sizes
- [ ] Validate HTML and check accessibility

---

**Next Week**: HTML Forms and User Input