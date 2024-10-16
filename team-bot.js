const { TeamsActivityHandler } = require('botbuilder');

class TeamsBot extends TeamsActivityHandler {
    constructor(conversationState, userState) {
        super();
        this.conversationState = conversationState;
        this.userState = userState;
    }

    async handleTeamsMessagingExtensionSubmitAction(context, action) {
        // Handle webhook request from Azure DevOps
        const webhookData = action.data;
        const message = `New Pull Request Created: ${webhookData.resource.pullRequestId} by ${webhookData.resource.createdBy.displayName}`;
        
        // Send message to chat group
        await context.sendActivity(message);
    }
}

exports.TeamsBot = TeamsBot;