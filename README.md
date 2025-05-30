# go-key-pg-template
This repository provides a foundational template for building Go backend applications with goa dsl, pre-configured with Keycloak for robust authentication and authorization, and integrated with a PostgreSQL database.

# Template Backend (Goa + PostgreSQL + OAuth2 - KeyCloak)

This project is a backend written in **Go** using **Goa DSL**, **OAuth2**, and **PostgreSQL** with support for **migrations** via `golang-migrate`.

---

## Requirements

- Go == 1.23.7 
- PostgreSQL == 15
- [golang-migrate](https://github.com/golang-migrate/migrate) CLI installed (optional)
- [Docker](https://www.docker.com/)
- [goa](goa.design/goa/v3/cmd/goa@v3.20.0) install by: go install 

---

## Project Setup

### 1. Clone the repository
```bash
git clone https://github.com/Rejrak/go-key-pg-template.git
cd go-key-pg-template
```

### 2. Create the `.env` file

Copy `.env.example` to `.env` and edit:

```bash
cp .env.example .env
```

Edit the following values:

```env
DB_HOST="postgres"
DB_NAME="ld_db"
DB_USER="ld_user"
DB_PASS="ld_pass"
DB_PORT="5432"

KC_DB_NAME="keycloak"
KC_ADMIN_USER="admin"
KC_ADMIN_PASS="adminpassword"

KC_CLIENT_ID="be-client"
KC_CLIENT_SECRET="be-client-secret"
KC_REALM="LastingDynamics"
```

> `RS256PK` is your JWT public key in base64 if using RS256

---

### 3. Start the database (via Docker)

```bash
docker compose up -d
```

> This will start a PostgreSQL container ready for use.

---

## Goa Code Generation

Every time you modify DSL files in `design/`, regenerate code:

runs:
```bash
goa gen be/design
```

---

## ▶ Run the Backend

```bash
docker compose up --build
```

Or using [air](https://github.com/cosmtrek/air) for live development:

```bash
air
```

Backend will run at:
```
http://localhost:9090
```

KeyCloak will run at:
```
http://localhost:8080
```
---

## Keycloak Configuration

To make Keycloak work with this backend and Swagger UI, follow these steps:

### 1. Access the Keycloak Admin Panel

Open your browser and go to [http://localhost:8080](http://localhost:8080).  
Log in using the credentials defined in your `.env` file:

```env
KC_ADMIN_USER=admin
KC_ADMIN_PASS=adminpassword
```

---

### 2. Create a Realm (if it doesn't exist)

- Go to **Realm Settings**
- Click **Create Realm**
- Name it:

```
LastingDynamics
```

> This must match `KC_REALM` in your `.env`.

---

### 3. Create a Client for the Backend

- Go to **Clients > Create**
- Fill the following:
  - **Client ID**: `be-client`
  - **Client Protocol**: `openid-connect`
- Click **Next**

#### Update Client Settings:

- **Root URL**:  
  ```
  http://localhost:9090
  ```

- Set:
  - **Access Type**: `confidential`
  - Enable:
    - ✅ Standard Flow
    - ✅ Direct Access Grants

- **Valid Redirect URIs**:
  ```
  http://localhost:9090/docs/oauth2-redirect
  ```

- Save the client.

---

### 4. Retrieve the Client Secret

- After saving, go to the **Credentials** tab.
- Copy the `Secret` value.
- Paste it into your `.env` file:

```env
KC_CLIENT_SECRET=your-client-secret-here
```

---

### 5. Create a Test User

- Go to **Users > Add User**
- Fill in a username, email, etc. → Save
- Go to the **Credentials** tab
  - Set a password
  - Disable "Temporary"
- Save the password.

---

### 6. Copy the Realm Public Key

- Go to **Realm Settings > Keys > RS256**
- Click **Public Key**
- Copy it and save it in `.env` as a single-line PEM:

```env
KC_RSA_PUBLIC_KEY="-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkq...\n-----END PUBLIC KEY-----"
```

> Replace line breaks with `\n` or load the key from a `.pem` file at runtime.

---



## 🔍 Swagger UI (interactive docs)

Swagger UI is available at:
```
http://localhost:9090/docs
```

Redoc UI is available at:
```
http://localhost:9090/redoc
```

Click **"Authorize"** to log in via OAuth2 using your Identity Provider (e.g. Keycloak).

>  This uses **Authorization Code Flow with PKCE**

---

## ⚖️ Project Structure

```
be/
├── cmd/                 # Entrypoint
├── design/              # Goa DSL
├── gen/                 # Generated code (do not modify)
├── internal/
│   ├── config/          # Bootstrap and init
│   ├── database/        # Repositories + models + migrations
│       ├── migrations
│   ├── services/        # Business logic implementations
│   └── middleware/      # JWT middleware
└── migrations/          # SQL migrations for golang-migrate
```

---

## Troubleshooting

### Port already in use
> Check if PostgreSQL or another service is using port `5432` or `9090`

### OAuth2 Error: `OAuth2Auth not implemented`
> Ensure each Goa service implements `OAuth2Auth(...)`

### Invalid token / missing public key
> Verify `RS256PK` in your `.env`

---

## Useful Resources
- [Goa DSL Docs](https://goa.design/)
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [Keycloak Docs](https://www.keycloak.org/docs/)

---

**Rejrak** - Roberto Lucchetti

