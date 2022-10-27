Website: tinyurl.com/coffeeleo

# Context

My friends know that when I do my work, there's a good chance I'm sitting in a coffee shop. In fact, it's always the same coffee shop.
Sometimes, my friends want to pop by and do some work or simply say "hi".
So, they send me a text message to make sure that I am, in fact, at the coffee shop. A problem arises, though, when I miss the text. 
If I don't confirm my location, I risk them not showing up when I'm at the coffee shop or, even worse, showing up when I'm not there. 
Thus, my solution was to create this project to allow them to see whether or not I am at the coffee shop without the need for texting.

# Goal

The goal wasn't only to give my friends status updates. 
In fact, the actual goal was to practice creating a web application with a focus on Docker and golang.

# Project Pages

### Main Page

The main page simply shows the current status of whether I am at the coffee shop or not.

### Admin Login Page

This page allows me to login with an administrator account.

### Admin Control Page

The control page allows me to actually change the values status values displayed on the main page.

# Tech

***Docker*** was used to separate the frontend, backend, and database into their own separate containers.

***HTML***, ***CSS***, and ***JavaScript*** were used in the frontend side of things to create each page, make them responsive, and send HTTP requests to the backend.

***Go***, mainly the ***Gin*** package, was used to handle the backend which included receiving and handling HTTP requests, handling password and session authentication, and manipulating the database.

***PostgreSQL*** was used to hold all the data which includes the current status, admin account, and session values.   

***AWS*** used to host the project (ElasticBeanstalk and Relational Database Service). 

***Github Workflows/Actions*** used to automatically deploy to AWS.
