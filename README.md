<a name="readme-top"></a>

![example workflow](https://github.com/YourUsername/chat-app/actions/workflows/build.yml/badge.svg)
[![Issues][issues-shield]][issues-url]
[![License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

<!-- PROJECT LOGO -->
<br />
<div align="center">

<h3 align="center">Real-time Chat Application</h3>

  <p align="center">
    A robust real-time chat application with features for managing chat rooms and private conversations using WebSocket technology.
    <br />
    <a href="https://github.com/YourUsername/chat-app/tree/main/api/specs"><strong>Explore API specs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/YourUsername/chat-app/issues">Report Bug</a>
    ·
    <a href="https://github.com/YourUsername/chat-app/issues">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#key-features">Key Features</a></li>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

## About The Project

This is a real-time chat application backend that enables users to create chat rooms, manage friendships, and engage in real-time conversations. The project utilizes WebSocket technology for real-time communication, MongoDB for persistent data storage, and Redis for managing real-time message delivery.

### Key Features

* **Chat Room Management**
    * Create new chat rooms
    * Join existing chat rooms
    * Leave chat rooms
    * Real-time updates on room activities

* **Friend System**
    * Add new friends
    * Manage friend list
    * Real-time friend status updates

* **Real-time Messaging**
    * Send and receive messages instantly
    * Private messaging between friends
    * Message history persistence
    * Real-time delivery status

### Built With

* [![Go][Go.dev]][Go-url]
* [![MongoDB][MongoDB]][MongoDB-url]
* [![Redis][Redis]][Redis-url]
* [![WebSocket][WebSocket]][WebSocket-url]

## Getting Started

### Prerequisites

Before running this project, make sure you have the following installed:
* Go 1.21 or higher
* MongoDB 6.0 or higher
* Redis 7.0 or higher

### Installation

1. Clone the repository
   ```sh
   git clone https://github.com/YourUsername/chat-app.git
   ```

2. Install Go dependencies
   ```sh
   go mod download
   ```

3. Set up environment variables
   ```sh
   cp .env.example .env
   ```

4. Configure your MongoDB connection
   ```sh
   # In your .env file
   MONGODB_URI=mongodb://localhost:27017/chatapp
   ```

5. Configure Redis connection
   ```sh
   # In your .env file
   REDIS_URL=redis://localhost:6379
   ```

6. Run the application
   ```sh
   go run main.go
   ```

## Usage

The application exposes both REST API endpoints and WebSocket connections for different functionalities:

### REST API Endpoints

```
POST /api/v1/rooms           # Create a new chat room
GET /api/v1/rooms           # List all available rooms
POST /api/v1/friends/add    # Add a new friend
GET /api/v1/friends         # Get friend list
```

### WebSocket Endpoints

```
ws://localhost:8080/ws/chat      # Main WebSocket endpoint for chat
```

## Contributing

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact
>
| <a href="https://github.com/samuelgamito" target="_blank">**Samuel Gamito**</a> |
| :---: |
| [<img src="https://avatars2.githubusercontent.com/u/12644639?s=460&u=4a0475c4309b27a91bb87f3adb13745ea76a917e" width="250">](https://github.com/samuelgamito)  |
| <a href="https://github.com/samuelgamito" target="_blank">`github.com/samuelgamito`</a> |



<!-- MARKDOWN LINKS & IMAGES -->
[issues-shield]: https://img.shields.io/github/issues/YourUsername/chat-app.svg?style=for-the-badge
[issues-url]: https://github.com/YourUsername/chat-app/issues
[license-shield]: https://img.shields.io/github/license/YourUsername/chat-app.svg?style=for-the-badge
[license-url]: https://github.com/YourUsername/chat-app/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white
[linkedin-url]: https://linkedin.com/in/YourUsername

<!-- BUILT WITH VARIABLES -->
[Go.dev]: https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://go.dev/
[MongoDB]: https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white
[MongoDB-url]: https://www.mongodb.com/
[Redis]: https://img.shields.io/badge/redis-%23DD0031.svg?style=for-the-badge&logo=redis&logoColor=white
[Redis-url]: https://redis.io/
[WebSocket]: https://img.shields.io/badge/websocket-%23000000.svg?style=for-the-badge&logo=socket.io&logoColor=white
[WebSocket-url]: https://developer.mozilla.org/en-US/docs/Web/API/WebSocket
