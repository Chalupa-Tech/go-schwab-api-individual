Here is a summary of the authentication process based on the Schwab Developer Portal documentation.

---

# Schwab API Authentication (OAuth 2.0)

The Schwab Trader API uses the **OAuth 2.0 "Three-Legged Flow"** to secure access. This protocol replaces username/password credentials with encrypted tokens.

## Key Components & Lifespans

- **Access Token:** Used to authenticate API calls (e.g., placing orders, streaming data).
  - **Lifespan:** **30 minutes**.
- **Refresh Token:** Used to generate new Access Tokens without user intervention.
  - **Lifespan:** **7 days**.
- **Client ID & Secret:** Unique credentials for your app found in the Developer Portal. The Secret must remain confidential.
- **Callback URL:** A required HTTPS URL where the user is redirected after logging in (e.g., `https://127.0.0.1` for local testing).

---

## The Authentication Flow

### Step 1: User Authorization (Browser)

Your application must direct the user to the Schwab Login Micro Site (LMS) to grant consent.

- **URL Pattern:**
  ```text
  https://api.schwabapi.com/v1/oauth/authorize?client_id={APP_KEY}&redirect_uri={CALLBACK_URL}
  ```
- **Result:** Upon success, the browser redirects to your Callback URL with a `code` appended (e.g., `https://127.0.0.1/?code={AUTHORIZATION_CODE}...`).

### Step 2: Create Initial Tokens (Server)

Exchange the Authorization Code for your first set of tokens.

- **Endpoint:** `POST https://api.schwabapi.com/v1/oauth/token`
- **Headers:**
  - `Authorization`: `Basic {BASE64_ENCODED_Client_ID:Client_Secret}`
  - `Content-Type`: `application/x-www-form-urlencoded`
- **Body:**
  ```text
  grant_type=authorization_code&code={AUTHORIZATION_CODE}&redirect_uri={CALLBACK_URL}
  ```
- **Note:** The `code` must be URL decoded (e.g., ending in `@` not `%40`) before sending.

### Step 3: Making API Calls

Include the Access Token in the header of all subsequent API requests.

- **Header Format:** `Authorization: Bearer {ACCESS_TOKEN}`.

### Step 4: Refreshing Tokens

Because the Access Token expires every 30 minutes, you must use the Refresh Token to get a new one. This can be done indefinitely as long as the Refresh Token itself (valid for 7 days) has not expired.

- **Endpoint:** `POST https://api.schwabapi.com/v1/oauth/token`
- **Headers:** Same as Step 2 (Basic Auth with Client ID/Secret).
- **Body:**
  ```text
  grant_type=refresh_token&refresh_token={REFRESH_TOKEN}
  ```
- **Logic:** If the Refresh Token expires (after 7 days), you must restart the flow at **Step 1**.
