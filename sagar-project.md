Scope of Work

1. Frontend Development (React)
Develop interactive, modular UI using React, HTML/CSS, TailwindCSS or Bootstrap
Build dynamic, collaborative UI elements including:
Gantt charts, drag-and-drop workflows
Collaborative editing, commenting, and task interactions
Real-time notifications, event-based updates, and chatbot UI
Enable email-based notifications and responses (ticket-like threads)
Integrate AI-powered UX assistants to automate and streamline workflows (e.g., auto-suggest, summarization, chat)


2. Backend Development (Node.js Microservices)
Design and implement Node.js + Express-based microservices
ORM-based data access with full support for:
Data hierarchy modelling (e.g., project > program > task structures)
Schema migration, rollback, and versioning
Role-based access control down to field-level permissions
Real-time sync: Notify frontend of database changes via event-driven updates (WebSockets/Event Bus)


3. Admin & Support Panel
Lightweight internal tools (e.g., using Retool) for admin dashboards
Manage tickets, user roles, issue statuses, and internal approvals


4. Testing, DevOps & Automation
100% unit and integration test coverage for frontend and backend
Scalable CI/CD pipelines using GitHub Actions
Auto-run on PR merges
Includes static code analysis (e.g., SonarQube), linting, and security checks
Auto-deploy with rollback hooks
Proper source control hygiene, semantic versioning, and release planning


5. Security & Performance
Frontend + Backend security best practices:
SQL injection protection
URL tampering prevention
Granular RBAC (endpoint & data-level)
Caching strategies (Redis or similar) for performance optimization
Secure key management, encrypted data storage, and token-based session handling


6. Developer Experience & External Integration
OAuth2 Provider Support (for client credentials + authorization code flows)
API Authentication using industry standards (JWT/Bearer Tokens, API Keys)
Secure 3rd-party access to APIs (role- and scope-based)
Comprehensive API Documentation Portal (e.g., Swagger UI, Redoc)
Authenticated Developer Login Console to manage keys, scopes, sandbox environments
Access Controls: Fine-grained ACLs for endpoints based on user roles and org contexts
API Access Logging and usage telemetry for auditing and throttling
Webhook system to support event-based integrations with external tools
Integration with other tools like Slack, Asana, etc.
Desired Capabilities
AI-assisted UX (e.g., LLM-based summaries, chat support, dynamic recommendations)
Support for real-time team collaboration (multi-user live interaction)
Email notification + reply support for workflows and issues
Preferred Stack
Frontend: React.js, Tailwind CSS, Lucide-React, Jest, Storybook
Backend: Node.js, Express.js, Prisma/Sequelize ORM
Database: Azure SQL, Cosmos DB (hierarchical modeling), Redis (caching)
Notifications: WebSockets, Azure Event Grid/Event Hub
DevOps: GitHub Actions, SonarQube, Azure Monitor, Terraform (IaC)
Security: Azure Key Vault, OWASP best practices, Azure AD for IAM
Admin Panel: Retool (or equivalent embedded tooling)
