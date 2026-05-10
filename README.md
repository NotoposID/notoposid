# NOTOPOS AI - Enterprise AI-Powered SaaS POS

![NOTOPOS AI Banner](https://images.unsplash.com/photo-1556742049-0cfed4f6a45d?q=80&w=2000&auto=format&fit=crop)

NOTOPOS AI is a next-generation, production-ready SaaS Point of Sale (POS) boilerplate designed for scalability, modularity, and AI-driven business intelligence. Built with **Next.js 14**, **Golang Fiber**, and **PostgreSQL**, it provides a rock-solid foundation for retail, cafes, restaurants, and UMKM.

## 🚀 Tech Stack

### Frontend
- **Framework**: Next.js 14 (App Router)
- **Styling**: Tailwind CSS & Shadcn UI
- **State Management**: Zustand & React Query
- **Animation**: Framer Motion
- **Validation**: Zod & React Hook Form
- **PWA**: Ready with offline support

### Backend
- **Language**: Golang (Fiber Framework)
- **Architecture**: Clean Architecture / Modular Monolith
- **Database**: PostgreSQL with `pgvector`
- **Caching**: Redis
- **Real-time**: WebSockets & Kafka ready
- **Auth**: JWT with RBAC & Multi-tenancy

### AI Stack
- **Engines**: OpenAI / OpenRouter
- **Vector Search**: pgvector for RAG and Recommendations
- **Features**: Sales Forecasting, AI Chatbot Assistant, Inventory Insights

## 🏗️ Architecture

The project follows a **Monorepo** structure managed by Turborepo:

```text
├── apps/
│   ├── web/          # Next.js Frontend
│   └── api/          # Golang Fiber Backend
├── packages/
│   └── config/       # Shared configurations
└── docker/           # Infrastructure & Docker setup
```

## 🛠️ Getting Started

### Prerequisites
- Node.js >= 20
- Go >= 1.21
- Docker & Docker Compose
- pnpm

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/your-org/notopos.git
   cd notopos
   ```

2. **Local Setup**
   Untuk menjalankan project di lokal, silakan baca [README_LOCAL.md](./README_LOCAL.md).

3. **Install dependencies**
   ```bash
   pnpm install
   ```

3. **Setup environment variables**
   ```bash
   cp apps/web/.env.example apps/web/.env.local
   cp apps/api/.env.example apps/api/.env
   ```

4. **Run with Docker Compose**
   ```bash
   docker-compose up -d
   ```

5. **Start development servers**
   ```bash
   pnpm dev
   ```

## 🔒 Multi-Tenancy

NOTOPOS AI uses a high-performance **column-based multi-tenancy** model. Every table is isolated by a `tenant_id`, and the backend middleware ensures that users can only access data belonging to their organization.

## 🤖 AI Capabilities

- **RAG Ready**: Easily feed your business data into AI models for custom insights.
- **Sales Forecasting**: Predict next week's revenue based on historical data.
- **AI Cashier Assistant**: Natural language commands for POS operations.

## 📄 License

Proprietary / Enterprise License.

---
Built with ❤️ by [NOTOPOS AI Team]
