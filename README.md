# Go Echo AI Boilerplate

A production-ready, microservices-friendly boilerplate for Golang using the Echo framework. 

This boilerplate is specifically designed with an **AI-first, Architecture-Driven + Process-Driven Development** philosophy. It enforces **Spec-Driven Development** compatibility to ensure seamless collaboration between human developers and AI agents.

## 🚀 Features

- **Framework**: [Labstack Echo v4](https://echo.labstack.com/)
- **Database**: PostgreSQL with [GORM](https://gorm.io/)
- **Architecture**: Hybrid Vertical Slice (Domain-Driven Design)
- **AI-First**: Built-in 3-Layer Agent Architecture, compatible with **Antigravity**, **Claude Code**, and **OpenHands**.
- **Process-Driven**: Integrated planning and verification workflows.
- **Spec-Driven**: Mandatory data-first and plan-first implementation protocols.

### 🧠 AI Agent Intelligence

This repository is optimized for AI-assisted development. All agents follow the core protocols defined in the documentation.

- **System Architecture**: [ARCHITECTURE.md](.agent/ARCHITECTURE.md)
- **Agent Protocols**: [AGENTS.md](.agent/AGENTS.md)
- **Plan Structure**: [PLAN-STRUCTURE.md](.agent/PLAN-STRUCTURE.md)
- **Workflows**: `.agent/workflows/` (use `/create`, `/enhance`, `/plan`)

## 📂 Project Structure

```bash
/.agent         # AI Agent Intelligence (Rules, Agents, Skills)
/cmd
  /api          # Main HTTP API entrypoint
  /cron         # Cron job entrypoint
  /migrate      # Database migration entrypoint
  /worker       # Background worker entrypoint
/internal
  /api_handler  # HTTP Handlers grouped by Domain (Vertical Slices)
  /domain       # Global domain definitions (Errors, Context)
  /model        # Shared Database Entities (GORM)
  /assets       # Static assets (HTML, CSS, JS) for Landing Page

/pkg            # Shared Infrastructure (Logger, Database, Config)
/plans          # Project task planning and tracking
```

## 🛠️ Getting Started

### Prerequisites
- Go 1.23+
- PostgreSQL
- Redis (optional)

### Installation

1. **Clone the repo**
   ```bash
   git clone https://github.com/PohLee/go-echo-ai-boilerplate.git
   cd go-echo-ai-boilerplate
   ```

2. **Setup Environment**
   ```bash
   cp .env.example .env
   # Update .env with your DB credentials
   ```

3. **Run Migrations**
   ```bash
   go run cmd/migrate/main.go
   ```

4. **Start Server**
   ```bash
   go run cmd/api/main.go
   ```


## ✅ Definition of Done

This project uses a strict definition of done enforced by script.
Before committing, run:

```bash
python .tmp/checklist.py
```

This ensures:
- Code is formatted and vetted
- Tests pass (Unit + Integration)
- Documentation is synchronized
- Success criteria meet the original spec

