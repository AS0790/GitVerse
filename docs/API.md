# GitVerse API & Endpoint Specification

GitVerse exposes both serverless HTTP endpoints and content-negotiated formats.

---

## 1. Endpoints

### `GET /result?u={username}`
Returns the rendered HTML report or JSON representation for the given GitHub username.

* **Query Parameters:**
  * `u` (required): GitHub login handle (1–39 alphanumeric characters).
* **Headers:**
  * `Accept: application/json` -> Returns structured JSON report.
  * `Accept: text/markdown` -> Returns raw Markdown summary.
  * Default -> Returns full responsive HTML web application.

### `POST /api/private`
Calculates private repository achievements via single-use OAuth code exchange and revokes the token immediately before responding.

---

## 2. Response Codes

* `200 OK`: Successful analysis completed.
* `400 Bad Request`: Invalid GitHub username format.
* `404 Not Found`: Username does not exist on GitHub.
* `504 Gateway Timeout`: Account has too many contribution records to process within serverless limits.

