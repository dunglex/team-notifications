require("dotenv").config();
const {
  CloudAdapter,
  ConfigurationServiceClientCredentialFactory,
  MemoryStorage,
  ConversationState,
  UserState,
} = require("botbuilder");
const http = require("http");
const { TeamsBot } = require("./team-bot");

console.log("Starting application...");

// Check for required environment variables
if (
  !process.env.MicrosoftAppId ||
  !process.env.MicrosoftAppPassword ||
  !process.env.MicrosoftAppType ||
  !process.env.MicrosoftAppTenantId
) {
  console.error("Required environment variables are missing");
  process.exit(1);
}
console.log("MicrosoftAppId: " + process.env.MicrosoftAppId);
console.log("MicrosoftAppPassword: " + process.env.MicrosoftAppPassword);
console.log("MicrosoftAppType: " + process.env.MicrosoftAppType);
console.log("MicrosoftAppTenantId: " + process.env.MicrosoftAppTenantId);
console.log("Environment variables loaded successfully");

// Create adapter
const adapter = new CloudAdapter(
  new ConfigurationServiceClientCredentialFactory({
    MicrosoftAppId: process.env.MicrosoftAppId,
    MicrosoftAppPassword: process.env.MicrosoftAppPassword,
    MicrosoftAppType: process.env.MicrosoftAppType,
    MicrosoftAppTenantId: process.env.MicrosoftAppTenantId,
  })
);
console.log("CloudAdapter created successfully");

// Create conversation and user state
const memoryStorage = new MemoryStorage();
const conversationState = new ConversationState(memoryStorage);
const userState = new UserState(memoryStorage);
console.log("ConversationState and UserState created successfully");

// Create bot
const bot = new TeamsBot(conversationState, userState);
console.log("TeamsBot created successfully");

// Function to handle /api/messages endpoint
const handleApiMessages = (req, res) => {
  const startTime = Date.now();
  const { method, url, headers } = req;
  const ip = req.socket.remoteAddress;
  let body = "";

  req.on("data", (chunk) => {
    body += chunk.toString();
  });

  req.on("end", () => {
    const duration = Date.now() - startTime;
    console.log(
      `Request: ${method} ${url} | Time: ${new Date(
        startTime
      ).toISOString()} | Origin: ${
        headers.origin || "N/A"
      } | IP: ${ip} | Duration: ${duration}ms | Body: ${body}`
    );

    req.body = JSON.parse(body);
    adapter
      .processActivity(req, res, async (context) => {
        await bot.run(context);
      })
      .catch((err) => {
        console.error(err);
        res.statusCode = 500;
        res.end("Internal Server Error");
      });
  });
};
console.log("API message handler created successfully");

// Create an HTTP server
const server = http.createServer((req, res) => {
  if (req.method === "POST" && req.url === "/api/messages") {
    handleApiMessages(req, res);
  } else {
    res.statusCode = 404;
    res.end("Not Found");
  }
});
console.log("HTTP server created successfully");

// Start the server
server.listen(process.env.PORT || 3978, () => {
  console.log(`Server started on port ${server.address().port}`);
});
