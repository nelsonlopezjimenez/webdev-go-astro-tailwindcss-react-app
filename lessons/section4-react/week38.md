---
title: "Week 1: Web Development IV: Introduction to JavaScript Frameworks" 
description: "Quarter 4. Frontend development using react"
week: 38
---

<h1 id="day1">Welcome to Web Development IV!</h1>

## Daily Activities

### Day 1: Quarter 4 Introduction

- Review syllabus and course expectations
- Introduction to quarter 4 
- Set up folder structure for course projects
- Chapter 05 from Shama Hoque's React book
- nvs: node version switch
- Additional Resources.

## Full-Stack React Projects: Learn MERN stack development by building modern web apps using MongoDB, Express, React, and Node.js 2nd ed. Edition
### by Shama Hoque (Author)

![shama_hoque](/week38/shama_hoque.png)

### Chapter 05: Building a Backend with MongoDB, Express, and Node

- Overview of the skeleton application
- Backend code setup
- User model with Mongoose
- User CRUD API endpoints with Express
- User Auth with JSON Web Tokens
- Running backend code and checking APIs

#### Feature Breakdown

- Sign up: Users can register by creating a new account using an email address.
- User list: Any visitor can see a list of all registered users.
- Authentication: Registered users can sign-in and sign-out.
- Protected user profile: Only registered users can view individual user details after signing in.
- Authorized user edit and delete: Only a registered and authenticated user can edit or remove their own user account details.

#### User model

| Field name | Type |
| ------ | --------:|
| name | String |
| email | String |
| password | String |
| created | Date |
| updated | Date |

#### API endopoints for user CRUD

| Operation | API route | HTTP method |
| --------- | --------- | ----------- | 
| Create a user | /api/users | POST |
| List all users | /api/users | GET |
| Fetch a user | /api/users/:userId | GET |
| Update a user | /api/users/:userId | PUT |
| Delete a user | /api/users/:userId | DELETE  |
| User sign-in | /auth/signin | POST |
| User signout (optional) | /auth/signout | GET |

#### Setting up the skeleton backend

```
| mern_skeleton/
|-- config/
|--- config.js
|-- server/
|    |--- controllers/
|    |        |---- auth.controller.js
|    |        |---- user.controller.js
|    |--- helpers/
|    |       | ---- dbErrorHandler.js
|    |--- models/
|    |       |---- user.model.js
|    |--- routes/
|    |       |---- auth.routes.js
|    |       |---- user.routes.js
|    |--- express.js
|    |--- server.js
|-- .babelrc
|-- nodemon.json
|-- package.json
|-- webpack.config.server.js
|-- yarn.lock
```
#### Setting up the app
![bethany-griggs](/week38/bethany-griggs.png)
1. Page 154: 
```sh
mkdir express-unique-app
cd express-unique-app
npm init -y
npm install express
touch app.js
mkdir routes public
touch routes/index.js 
```
1. edit app.js file


#### Steps to run mern social app

 1. open cmd
 1. run nvs
 1. select 'm' or node v13.14.0
 1. cd Documents/_QUARTER3/chapter-five-mern-social-yes-nm
 1. run npm install or tar xvfz nm.tar.gz
 1. npm run development