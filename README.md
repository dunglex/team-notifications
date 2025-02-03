# Team Notifications 🔔

A repository with configuration and samples for team notifications using Azure DevOps and Microsoft Teams Adaptive Cards.

## 🌳 Repository Structure

This repo contains 3 branches:
- `master` - Documentation, tests, curl samples, diagrams, and JSON schemas
- `go` - Application in golang
- `node` - Previous application in nodejs (having issue out of memory)

## 📁 Files

- `sample_adaptive_card.json` - PR notification card template 
- `sample_parse_json.txt` - JSON schema and parsing config
- `test-curl.sh` - Test script for sending notifications
- `test-request.sh` - Sample request payload
- `adaptive-card.json` - Full card schema spec

## 🧪 Usage

### Testing Notifications

Run the test scripts to simulate notifications:

```bash
# Test with sample payload
./test-request.sh

# Send to Teams endpoint
./test-curl.sh