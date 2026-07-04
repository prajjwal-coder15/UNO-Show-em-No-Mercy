# 🎮 UNO Show 'Em No Mercy

A feature-rich implementation of **UNO Show 'Em No Mercy** built in **Go**, featuring multiplayer support, AI bots, REST APIs, WebSockets, customizable rules, and an interactive CLI.

> Designed with clean architecture and modular packages to demonstrate backend engineering, game logic, concurrency, and real-time communication.

![Go](https://img.shields.io/badge/Go-1.24+-00ADD8?logo=go)
![License](https://img.shields.io/badge/License-MIT-green)
![Status](https://img.shields.io/badge/Status-Active-success)

---

# ✨ Features

- 🎮 Complete UNO Show 'Em No Mercy gameplay
- 👥 Multiplayer game rooms
- 🤖 AI Bots with multiple difficulty levels
- 🌐 REST API
- ⚡ Real-time gameplay using WebSockets
- 💻 Interactive CLI
- 🏠 Room & lobby management
- 🃏 Complete deck generation
- 🎲 Turn management
- 🔥 Card stacking
- 🎨 Wild color selection
- 💀 Mercy Rule support
- ⚙️ Multiple rule presets
- ❤️ Health endpoint
- 🔄 Middleware support

---

# 🏗 Architecture

```text
                    Client
               /             \
          CLI Client      Web Client
               \             /
                REST / WebSocket API
                        │
          ┌─────────────┴─────────────┐
          │                           │
      Game Engine                Room Manager
          │                           │
          └─────────────┬─────────────┘
                        │
                   Rules Engine
                        │
                     AI Bots
```

---

# 📂 Project Structure

```text
uno-show-em-no-mercy/
│
├── api/          # REST API & WebSocket handlers
├── cmd/          # CLI & Server entry points
├── game/         # Core game engine
├── room/         # Room & lobby management
├── rules/        # Rule presets & configurations
├── runner/       # Application runner
├── ui/           # Terminal UI components
└── go.mod
```

---

# 📦 Package Overview

## `api/`

Handles all API-related functionality.

- REST endpoints
- WebSocket connections
- Game requests
- Room APIs
- Bot APIs
- Middleware
- Health checks

---

## `game/`

Contains the complete UNO game engine.

Includes:

- Card models
- Deck generation
- Turn handling
- Draw mechanics
- Discard pile
- Card effects
- Validation
- Winner detection
- Mercy Rule
- Stacking
- Multi-card play
- AI Bot logic

---

## `room/`

Responsible for multiplayer rooms.

Features:

- Lobby management
- Room creation
- Player joining
- Room codes
- Room state
- Room manager

---

## `rules/`

Supports different game configurations.

Available presets:

- Official UNO
- Show 'Em No Mercy
- House Rules
- Custom configurations

---

## `ui/`

Terminal-based interface.

Includes:

- Menus
- Player hand rendering
- Input handling
- Colored output
- Room list display

---

## `runner/`

Coordinates application startup and lifecycle.

---

## `cmd/`

Application entry points.

```text
cmd/
├── cli/
└── server/
```

---

# 🎮 Gameplay Features

- Classic UNO mechanics
- Show 'Em No Mercy rules
- Draw stacking
- Skip
- Reverse
- Wild cards
- Wild Draw Four
- Wild Color Roulette
- Mercy elimination
- Multiple winners detection
- Custom rule support

---

# 🤖 AI Bots

Supports intelligent AI players.

Features:

- Easy
- Medium
- Hard
- Card prioritization
- Color strategy
- Smart decision making
- Risk evaluation

---

# 🌐 API Features

- Room management
- Game state
- Player actions
- Move validation
- WebSocket events
- Bot management
- Rules endpoint
- Health endpoint

---

# 🛠 Tech Stack

- **Language:** Go
- **Communication:** REST API, WebSockets
- **CLI:** Go Terminal
- **Architecture:** Clean Architecture
- **Concurrency:** Goroutines & Channels

---

# 🚀 Getting Started

## Clone Repository

```bash
git clone https://github.com/your-username/uno-show-em-no-mercy.git

cd uno-show-em-no-mercy
```

## Install Dependencies

```bash
go mod tidy
```

---

# ▶️ Run CLI

```bash
go run ./cmd/cli
```

---

# ▶️ Run Server

```bash
go run ./cmd/server
```

---

# 🧪 Run Tests

```bash
go test ./...
```

---

# 📚 Concepts Covered

- Clean Architecture
- Package Organization
- REST API Development
- WebSockets
- Game Engine Design
- State Management
- Concurrency
- Middleware
- Strategy Pattern
- AI Decision Making
- Validation
- Modular Design

---

# 🚀 Future Improvements

- 🌍 React Frontend
- 🗄 PostgreSQL persistence
- 👤 User authentication
- 🏆 Leaderboard
- 🎯 Match history
- 📈 Player statistics
- 🎵 Sound effects
- 🎨 Better animations
- 🐳 Docker support
- ☁️ Cloud deployment

---

# 📜 License

This project is licensed under the **MIT License**.

---

# 👨‍💻 Author

**Prajwal**

Backend Engineer • Go Developer • System Design Enthusiast

GitHub: **https://github.com/prajwal-coder15**

---

## ⭐ Support

If you found this project useful, consider giving it a **⭐ Star** on GitHub!
