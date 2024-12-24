# Reddit Engine
Our aim was to build the Reddit like backend engine and a simulator for testing.

## **Table of Contents**
- [Project Overview](#project-overview)
- [Features](#features)
  - [Core Engine](#core-engine)
  - [Client Tester/Simulator](#client-testersimulator)
- [System Architecture](#system-architecture)
- [Technologies Used](#technologies-used)
- [How to Run](#how-to-run)
  - [Prerequisites](#prerequisites)
  - [Steps to Run](#steps-to-run)

---

## **Project Overview**

This project is a Reddit-like engine paired with a client tester/simulator. It aims to replicate core functionalities of Reddit, such as account management, subreddit interactions, posting, commenting, and more.

---

## **Features**

### **Core Engine**
- **Account Management**:
  - Register new accounts.
- **Subreddit Management**:
  - Create subreddits.
  - Join and leave subreddits.
- **Posting**:
  - Post simple text content in subreddits.
- **Commenting**:
  - Add hierarchical comments (nested comments on posts or other comments).
- **Voting & Karma**:
  - Upvote/downvote posts and comments.
  - Compute user karma based on votes.
- **Feed Management**:
  - Retrieve a personalized feed of posts.
- **Direct Messaging**:
  - View direct messages.
  - Reply to direct messages.

### **Client Tester/Simulator**
- Simulate thousands of users interacting with the engine.
- Simulate live connection and disconnection periods for users.
- Apply a Zipf distribution for subreddit membership (popular subreddits have more members).
- Generate posts, comments, and re-posts based on activity levels.
- Measure performance metrics like latency, throughput, and resource usage.

---

## **System Architecture**

The system is designed with a separation of concerns:
1. **Engine Process**:
   - Manages backend logic (e.g., accounts, subreddits, posts, comments).
   - Distributes posts and tracks interactions like votes and comments.
2. **Client Processes**:
   - Simulate independent users performing actions such as posting, commenting, subscribing, and voting.

---

## **Technologies Used**
- Programming Language: Go lang and Proto Actors.

---

## **How to Run**

### Prerequisites
1. Install Go (Golang) on your system. You can download it from the official website: `https://golang.org/`.
2. Set up your Go workspace by configuring the `GOPATH` environment variable.

### Steps to Run
1. Clone this repository:
   ```
   git clone https://github.com/sanket1305/COP5615_DOSP_Reddit.git
   cd COP5615_DOSP_Reddit/project4_Reddit_engine
   ```

2. Start the engine and simulator process:
   ```
   go mod tidy
   go run .
   ```

---
