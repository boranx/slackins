# slackins [![Build Status](https://travis-ci.org/boranx/slackins.svg?branch=master)](https://travis-ci.org/boranx/slackins) [![Coverage Status](https://coveralls.io/repos/github/boranx/slackins/badge.svg)](https://coveralls.io/github/boranx/slackins)
Slacking is a stand alone service that integrates slack and jenkins. No plug-in required.
You can execute any job via slackbot.

## Configuration
Create a bot from https://yourteam.slack.com/apps/manage/custom-integrations and save the token somewhere.
Configure Jenkins Jobs with Parameterized Build and enable remote triggering.

### Set Environment Variable :
- SLACK_BOT_API_TOKEN=xxxxx
- URI=http://localjenkins:8080
- TOKEN=xxxx

### Usage
- @chatbot execute deployapi version=1.1.1,transform=true
