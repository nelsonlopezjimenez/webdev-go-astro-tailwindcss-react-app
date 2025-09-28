---
title: "WEEK 60 - Wireframe creation" 
description: "Practical Web Design: chapter 1 about wireframe creation"
week: 60
---
# 3 Website-Designing Workflow
From Practical Web Design by Philippe Hong, Packt>, 2018,  ISBN 978-1-78839-503-8
 Before we really start digging into creating and implementing our first website, I want you
 to go through all the processes from start to launch. Web design is not just only designing
 aesthetic websites and beautiful layouts; web design is a whole process, especially when
 you want to implement your design to the real world afterward.

 In this chapter, we'll cover the following: 

- Goal identification: how to identify our goal
-  Scope definition: List out the scope 
- Wireframe creation: How to create wireframes
-  Designing: Framework to create a great design
-  Implementing, testing, and launching
## Our situation

 Let's imagine yourself working as a designer, and receiving client work. Your first project is
 to design a website for a Racing Club; here's the brief of the project:
 Racing Club is a club for racing fans that was founded in 2016. It started out with a bunch of
 friends with a love of cars, but it rapidly grew into a community keen to share their passion.
 So now, let's go through the entire process.

## Goal identification
 Within this first stage, you need to identify the website's end goal. By communicating with
 the client and asking them questions about their business and its goals.

 ## What is the purpose of the website?
 This is the right moment to identify the problem to be solved or to set a goal for the website.
 Sell tickets or increase ticket sales? See what your client wants and find the best solution for
 it. You also need to do your own investigation, dig into their website, and search for
 anything that needs to be fixed. With all the fundamentals we studied before, you're now
 able to see what is good and what is bad.
 ## Who is the website for?
 The best way to simplify the design process and the decision making is to know your target
 audience. There are plenty of ways to know your audience. You can either ask the client,
 track with different analytics, or check out previous reports on the same market that will
 help you in that phase. 
Is this useful for our audience?
 Your client should have some information about their customers, such as their ages,
 incomes, and so on. Use these pieces of information to create personas, and create user
 flows that fit the website. In our example, the Racing Club, we would create personas such
 as the following:
 George: 38 years old, father, garage worker, passionate about racing
 Paul: 28 years old, single, works in finance, has a love for cars and racing
 Both users, flows will work differently and you can already suppress any flows that don't
 concern our target users.
 What do they expect to find or do there?
 It is also important to know to define the Information Architecture (IA) of your website.
 Knowing what to show your users will set the design of your screens and plan the user
 experience.
 You'll have to create a sitemap and define every screen you need to do. Doing this first will
 help you greatly in designing your website, as you won't even have to think about it.
 [ 58 ]
Website-Designing Workflow
 Chapter 3
 Does the website need to follow a brand or have
 its own brand identity?
 Designing a website can be different when you need to follow a brand style guide. As the
 style guide will help keeping the consistency in the brand, the client will want you to follow
 it, even if it will restrict your creativity a bit.
 If the client doesn't have a brand identity, it's a good opportunity to create one for them.
 Are there any competitors? If there are, how is
 the website different than others?
 Knowing a client's competitors is also a good way to know what to do and what not to do.
 In your process of gathering information, you'll need to research the client's competitors. It's
 not just about doing something totally different, but doing what is good for the client. If
 some user experiences of your competitors are good, take inspiration from them, to make
 your client's website better. You often don't need to re-invent the wheel, but just improve it.
 So here's our project.
 We need a website with the following:
 Homepage
 Upcoming events page
 Past events page
 Event page details (view info of event and be able to buy a ticket)
 Blog page
 About Us page
 Contact page
 Login page (see the history of ticket purchased) 
The website needs to be responsive so people can access it on their mobile. The client
 doesn't have a brand identity and is willing to let us create one.
 The main goal of the website is to firstly show relevant information for the users, and then,
 if they want, enable them to purchase tickets online instead of going to the physical
 location.
 [ 59 ]
Website-Designing Workflow
 Chapter 3
 The following diagram is the sitemap of the website: 
Sitemap example
 Defining the scope
 This is often a tough part for designers: knowing and defining the scope of a project. It's
 usual for projects to last longer than expected, but this should not be a problem, as it leads
 to more work. But, sometimes, clients' expectations and your expectations are not the same,
 so it is best to set boundaries to prevent unexpected work and scope creep. Putting
 everything in a contract will help you. Here are some templates that you can use: https:/​/
 www.​smashingmagazine.​com/​2013/​04/​legal-​guide-​contract-​samples-​for-​designers/​.
 Creating wireframes
 Now that we have defined the goal of the project, we can start designing some wireframes.
 In this project example, we'll only do a couple of screens. Here's the wireframe that we'll use
 for the homepage, events page, and upcoming events page. Wireframes are not meant to be
 polished and designed, they are just shaped to get an idea of the layout and content. So just
 use simple rectangles with your favorite design application, or you can even sketch it by
 hand. 
[ 60 ]
Website-Designing Workflow
 Chapter 3
 Here's what we came up with: 
[ 61 ]
Website-Designing Workflow
 Chapter 3
 For the events page:
 [ 62 ]
Website-Designing Workflow Chapter 3
 [ 63 ]
 And this is what we came up with for the event page: 
Website-Designing Workflow
 Chapter 3
 Designing
 I always use the same framework when designing a project: 
1.
 2.
 3.
 Get inspiration
 Improve
 Invent
 Let me explain what I do for each step: 
Get inspiration
 I really think that inspiration is one of the main parts of design creation. Looking and
 gathering inspiration is crucial for me, as I need those inspirations to make my own design
 by taking a few pieces of design that I'll find cool or relevant for this type of project.
 The following are a few websites that I use to find inspiration:
 Awwwards
 CSSDesignawards
 FWA
 Dribbble
 Behance
 Then I'll use BukketApp to grab images and create a mood board for this specific project.
 Improve
 Once you have your mood board, you can start to tweak and improve the design. Like a
 puzzle, try to compose by remixing shapes, colors, and so on. This is the most fun and
 creative part, because it's up to your imagination and the amount of flexibility you have to
 create something unique and of high quality

## Remove entire folder from git history
## Remove folder with sensitive data
```sh
git filter-repo --path secrets/ --invert-paths
```

# Clean up and push
```sh
git push origin --force --all
git push origin --force --tags
```
## Remove Large Executable from History

bash# Using git filter-repo (recommended)
git filter-repo --path COMPUTER_BASICS-8.25.2025.exe --invert-paths

# Force push to update remote
git push origin --force --all


## Clean up local repository
    ```sh
git reflog expire --expire=now --all
git gc --prune=now --aggressive
```

# Force push to update remote (DANGER: affects all collaborators)
git push origin --force --all
git push origin --force --tags


 ![development branch](/week38/image-1.png)
<img src="/week38/image-1.png" alt="development branch" width="500" height="600">
![weeknavigation start branch](/week38/image.png)
<img src="/week38/image.png" alt="weeknavigation branch" width="500" height="600">
![secondev branch](/week38/image-2.png)
<img src="/week38/image-2.png" alt="secondev branch" width="500" height="600">
![added this images not rendered](/week38/image-3.png)
<img src="/week38/image-3.png" alt="alt pics rendering branch" width="500" height="600">
![with img tags](/week38/image-4.png)
<img src="/week38/image-4.png" alt="alt with img tags" width="500" height="600">