const http = require("http");
const fs = require("fs");
const path = require("path");

const root = path.resolve(__dirname, "..", "site");
const resultTemplate = path.resolve(__dirname, "..", "web", "testdata", "result-public.html");
const types = {
  ".css": "text/css; charset=utf-8",
  ".html": "text/html; charset=utf-8",
  ".js": "application/javascript; charset=utf-8",
  ".json": "application/json; charset=utf-8",
  ".png": "image/png",
  ".svg": "image/svg+xml",
  ".woff2": "font/woff2",
  ".txt": "text/plain; charset=utf-8"
};

const port = 5173;

http.createServer((request, response) => {
  const url = new URL(request.url, `http://localhost:${port}`);
  const requestPath = decodeURIComponent(url.pathname);

  // Result view for profile check
  if (requestPath === "/result") {
    const user = url.searchParams.get("u") || "torvalds";
    if (fs.existsSync(resultTemplate)) {
      let content = fs.readFileSync(resultTemplate, "utf8");
      content = content.replace(/torvalds/g, user);
      response.writeHead(200, { "Content-Type": "text/html; charset=utf-8" });
      return response.end(content);
    }
  }

  let file = path.resolve(root, "." + (requestPath === "/" ? "/index.html" : requestPath));
  if (!file.startsWith(root)) return response.writeHead(403).end("Forbidden");

  if (fs.existsSync(file) && fs.statSync(file).isDirectory()) {
    file = path.join(file, "index.html");
  } else if (!fs.existsSync(file) && fs.existsSync(file + ".html")) {
    file = file + ".html";
  }

  fs.readFile(file, (error, data) => {
    if (error) {
      const notFound = path.join(root, "404.html");
      if (fs.existsSync(notFound)) {
        response.writeHead(404, { "Content-Type": "text/html; charset=utf-8" });
        return fs.createReadStream(notFound).pipe(response);
      }
      return response.writeHead(404).end("Not found");
    }
    response.writeHead(200, { "Content-Type": types[path.extname(file)] || "application/octet-stream" });
    response.end(data);
  });
}).listen(port, "127.0.0.1", () => console.log(`GitVerse preview: http://localhost:${port}`));
