require('dotenv').config();
const http = require('http');
const url = require('url');

// Function to handle the POST request when a pull request is created
function onPrCreated(req, res) {
  let reqBody = '';

  req.on('data', chunk => {
    reqBody += chunk.toString();
  });

  req.on('error', (err) => {
    console.error('Error receiving data:', err);
    res.statusCode = 400;
    res.end('Bad Request');
  });

  req.on('end', async () => {
    try {
      const message = JSON.parse(reqBody);
      let url = `${message.resource.repository.webUrl }/pullrequest/${message.resource.pullRequestId}`;
      let jiraUrl = (index = (description = message.resource.description).indexOf("https://sd.homecredit.vn")) != -1 ? description.substring(index) : null;
      let srcBranch = message.resource.sourceRefName.replace('refs/heads/', '');
      let targetBranch = message.resource.targetRefName.replace('refs/heads/', '');
      let repository = message.resource.repository.name;
      let author = message.resource.createdBy.displayName;
      const webhookRequestBody = JSON.stringify({
        "type": "message",
        "attachments": [
          {
            "contentType": "application/vnd.microsoft.card.adaptive",
            "contentUrl": null,
            "content": {
              "$schema": "http://adaptivecards.io/schemas/adaptive-card.json",
              "type": "AdaptiveCard",
              "version": "1.2",
              "body": [
                {
                    "type": "TextBlock",
                    "title": `${message.resource.title}`,
                    "text": `${message.resource.description}`,
                    "url": `${url}`,
                    "jiraUrl": `${jiraUrl}`,
                    "srcBranch": `${srcBranch}`,
                    "targetBranch": `${targetBranch}`,
                    "repository": `${repository}`,
                    "author": `${author}`
                }
              ]
            }
          }
        ]
      });
      console.log('Sending message:', webhookRequestBody);
      await fetch(process.env.WEBHOOK_URL, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: webhookRequestBody
      });

      res.statusCode = 200;
      res.end('Message received and processed');
    } catch (error) {
      console.error('Error processing message:', error);
      res.statusCode = 500;
      res.end('Internal Server Error');
    }
  });
}

// Create the web server
const server = http.createServer((req, res) => {
  const parsedUrl = url.parse(req.url, true);

  if (req.method === 'POST' && parsedUrl.pathname === '/pr/created') {
    onPrCreated(req, res);
  } else {
    res.statusCode = 404;
    res.end('Not Found');
  }
});

// Start the server
server.listen(process.env.PORT || 3978, () => {
  console.log(`Server started on port ${server.address().port}`);
});