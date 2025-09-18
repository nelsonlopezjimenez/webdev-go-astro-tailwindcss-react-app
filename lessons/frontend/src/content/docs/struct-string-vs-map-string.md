---
title: "struct string vs map array of objects"
description: "Brief description file1"
publishDate: 2025-09-17
tags: ["struct", "map", "documentation"]
draft: false
---
The error indicates that your struct field `Resources` is still defined as `[]string`, but you're trying to assign `[]map[string]interface{}`. You need to update the struct definition.

## Option 1: Update the Struct Field Type

Change your struct definition from:

```go
type YourStruct struct {
    // ... other fields
    Resources []string  // <- Change this
}
```

To:

```go
type YourStruct struct {
    // ... other fields
    Resources []map[string]interface{} `json:"resources"`
}
```

## Option 2: Use a Custom Resource Struct (Recommended)

Define a Resource struct and update your main struct:

```go
type Resource struct {
    Text string `json:"text"`
    URL  string `json:"url"`
    Type string `json:"type"`
}

type YourStruct struct {
    // ... other fields
    Resources []Resource `json:"resources"`
}
```

Then use it like:

```go
Resources: []Resource{
    {
        Text: "MDN Web Docs - JavaScript Reference",
        URL:  "https://developer.mozilla.org/en-US/docs/Web/JavaScript",
        Type: "documentation",
    },
    {
        Text: "JavaScript.info - Modern JavaScript tutorial",
        URL:  "https://javascript.info/",
        Type: "tutorial",
    },
    // ... rest of your resources
},
```

## Option 3: Keep it as []string (Quick Fix)

If you want to keep the current structure and just add URLs temporarily, you can use strings with HTML:

```go
Resources: []string{
    `<a href="https://developer.mozilla.org/en-US/docs/Web/JavaScript" target="_blank">MDN Web Docs - JavaScript Reference</a>`,
    `<a href="https://javascript.info/" target="_blank">JavaScript.info - Modern JavaScript tutorial</a>`,
    `<a href="https://www.w3schools.com/js/" target="_blank">W3Schools - JavaScript tutorials and examples</a>`,
    `<a href="https://codepen.io/" target="_blank">CodePen - JavaScript code playground</a>`,
    `<a href="https://developer.chrome.com/docs/devtools/" target="_blank">Chrome DevTools - Debugging and testing</a>`,
},
```

## Recommended Solution

I recommend **Option 2** because it provides:
- Type safety
- Better JSON serialization
- Cleaner separation of data and presentation
- Easier to extend with additional fields

You'll need to find where your struct is defined (likely around line 633 or earlier) and update the `Resources` field type.