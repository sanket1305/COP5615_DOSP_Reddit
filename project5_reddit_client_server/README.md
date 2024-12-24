# Reddit Engine

This project extends the previously built Reddit-like engine by adding **REST APIs** using the **Go programming language**. It allows multiple clients to interact with the engine simultaneously, perform various actions such as posting, commenting, and voting, and also enables users to send and reply to direct messages. The functionality of the code is verified through rigorous testing.

---

## Table of Contents

1. [Description](#description)
2. [Features](#features)
3. [REST APIs](#rest-apis)
   - [User Management](#user-management)
   - [Subreddit Management](#subreddit-management)
   - [Post Management](#post-management)
   - [Commenting](#commenting)
   - [Voting](#voting)
   - [Messaging](#messaging)
4. [System Demo](#system-demo)
5. [Getting Started](#getting-started)
   - [Prerequisites](#prerequisites)
   - [Installation](#installation)
   - [Running the Application](#running-the-application)

---

## Description

The project focuses on creating REST APIs for a Reddit-like engine and enabling multiple clients to interact with the system simultaneously. The key objectives include:

- Implementing REST APIs for user management, subreddit management, post creation, commenting, voting, and messaging.
- Allowing multiple clients to send concurrent requests to the engine.
- Supporting direct messaging between users.
- Verifying the functionality of all implemented features.

---

## Features

1. **User Management**
   - Register new accounts.

2. **Subreddit Management**
   - Create, join, and leave subreddits.

3. **Post Management**
   - Create text-based posts in subreddits.
   - Retrieve a feed of posts from subscribed subreddits.

4. **Commenting**
   - Add hierarchical comments (comments on posts or other comments).

5. **Voting**
   - Upvote or downvote posts and comments.
   - Compute user karma based on votes.

6. **Messaging**
   - Send and retrieve direct messages.
   - Reply to direct messages.

7. **Concurrency**
   - Support multiple clients interacting with the engine simultaneously.

---

## REST APIs

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

## System Demo

A demo of the system is available on YouTube:  
[System Demo Video](youtube url)

---

## Getting Started

### Prerequisites

- Install [Go Programming Language](https://go.dev/).
- Familiarity with REST APIs is recommended.
- A tool like Postman or cURL for testing API endpoints is helpful.

### Installation

1. Clone this repository:
   ```
   https://github.com/sanket1305/COP5615_DOSP_Reddit.git
   cd COP5615_DOSP_Reddit/project5_reddit_client_server
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

### Running the Application

[Check youtube demo video](https://youtu.be/h1urknlePyE)

---
