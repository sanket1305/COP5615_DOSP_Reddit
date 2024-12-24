# Reddit Engine

This project implements a Reddit-like engine using the actor model in the **Go programming language**. It was developed as part of **COP5615 - Distributed Operating System Principles** coursework and includes features such as user management, subreddit creation, post and comment handling, voting, and performance simulation. The engine also supports REST APIs and multi-client interactions.

![reddit](https://github.com/user-attachments/assets/8d6d50a4-d535-4e85-9df2-7a15d378ce4c)


Source: Reddit

---

## Table of Contents

1. [Description](#description)
2. [Features](#features)
3. [Architecture](#architecture)
5. [Getting Started](#getting-started)
   - [Prerequisites](#prerequisites)
   - [Installation](#installation)
6. [API Endpoints (Project 5)](#api-endpoints-project-5)
   - [User Management](#user-management)
   - [Subreddit Management](#subreddit-management)
   - [Post Management](#post-management)
   - [Commenting](#commenting)
   - [Voting](#voting)
   - [Messaging](#messaging)
7. [Simulation Testing](#simulation-testing)

---

## Description

This repository contains the code for the last two projects (Project 4 and Project 5) of **COP5615 - Distributed Operating System Principles**. 

The specific APIs supported by Reddit can be found at: [Reddit API Documentation](https://www.reddit.com/dev/api/).  
For an overview of Reddit and its functionality, visit: [What is Reddit?](https://www.oberlo.com/blog/what-is-reddit).

### Project Overview

- **Project 4**: Build a Reddit-like engine and implement a simulator to analyze the performance of the engine.
- **Project 5**: Extend the previously built engine by creating REST APIs and implementing multiple clients that can interact with the engine simultaneously using two separate processes.

---

## Features

The Reddit-like engine includes the following functionality:

1. **User Management**
   - Register an account.

2. **Subreddit Management**
   - Create and join subreddits.
   - Leave subreddits.

3. **Post Management**
   - Post in subreddits (text-only posts; no support for images or markdown).
   - Comment on posts (supports hierarchical comments, i.e., commenting on comments).

4. **Voting and Karma**
   - Upvote or downvote posts and comments.
   - Compute user karma based on votes.

5. **Feed and Messaging**
   - Retrieve a feed of posts from subscribed subreddits.
   - Send and receive direct messages.
   - Reply to direct messages.

6. **Simulator**
   - Simulate as many users as possible.
   - Simulate periods of live connection and disconnection for users.

---

## Architecture

- The project is designed with a separation between:
  1. **Client Processes**: Handle user actions such as posting, commenting, and subscribing.
  2. **Engine Process**: Distributes posts, tracks comments, handles votes, and manages user data.

- Multiple independent client processes simulate thousands of users interacting with a single-engine process.

---

## Getting Started

### Prerequisites

- Install [Go Programming Language](https://go.dev/).
- Familiarity with REST APIs is recommended.

### Installation

Check ReadMe of respective Project

## API Endpoints (Project 5)

### User Management
- `POST /register`: Register a new account.

### Subreddit Management
- `POST /subreddit/create`: Create a new subreddit.
- `POST /subreddit/join`: Join a subreddit.
- `POST /subreddit/leave`: Leave a subreddit.

### Post Management
- `POST /post`: Create a new post in a subreddit.
- `GET /feed`: Retrieve posts from subscribed subreddits.

### Commenting
- `POST /comment`: Add a comment to a post or another comment.

### Voting
- `POST /vote`: Upvote or downvote a post or comment.

### Messaging
- `GET /messages`: Retrieve direct messages.
- `POST /message/reply`: Reply to a direct message.

---

## Simulation Testing

The project includes a tester/simulator to validate functionality:

1. Simulate thousands of users performing actions like posting, commenting, voting, etc.
3. Test periods of live connection and disconnection for users.

---
