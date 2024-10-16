require('dotenv').config();
const http = require('http');
const url = require('url');
const axios = require('axios'); // Use axios for HTTP requests
const { Size } = require('adaptivecards');

const powerAutomateUrl = process.env.POWER_AUTOMATE_URL; // Set your Power Automate URL in environment variables

const server = http.createServer((req, res) => {
  const parsedUrl = url.parse(req.url, true);

  if (req.method === 'POST' && parsedUrl.pathname === '/api/messages') {
    let body = '';

    req.on('data', chunk => {
      body += chunk.toString();
    });

    req.on('end', async () => {
      try {
        const message = JSON.parse(body);
        const teamMessage = {
          type: "AdaptiveCard",
          attachments: [
            {
              contentType: "application/vnd.microsoft.card.adaptive",
              content: {
                $schema: "http://adaptivecards.io/schemas/adaptive-card.json",
                type: "AdaptiveCard",
                version: "1.0",
                body: [
                  {
                    type: "TextBlock",
                    text: "For Samples and Templates, see [https://adaptivecards.io/samples](https://adaptivecards.io/samples)",
                    size: "large"
                  }
                ]
              }
            }
          ]
        };

        console.log(JSON.stringify(teamMessage, null, 2));
        // Send message to Power Automate
        await axios.post(powerAutomateUrl, teamMessage);

        res.statusCode = 200;
        res.end('Message received and processed');
      } catch (error) {
        console.error('Error processing message:', error);
        res.statusCode = 500;
        res.end('Internal Server Error');
      }
    });
  } else {
    res.statusCode = 404;
    res.end('Not Found');
  }
});

console.log('HTTP server created successfully');

// Start the server
server.listen(process.env.PORT || 3978, () => {
  console.log(`Server started on port ${server.address().port}`);
});