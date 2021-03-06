const fs = require('fs');

const env = {
  "GClientId": "670077302427-0r21asrffhmuhkvfq10qa8kj86cslojn.apps.googleusercontent.com",
  "EntityTable": `portals-me-${process.argv[2]}-entities`,
  "SortIndex": `DataTable`,
  "FeedTable": `portals-me-${process.argv[2]}-feeds`,
  "TimelineTable": `portals-me-${process.argv[2]}-timeline`,
  "EntityStreamFanoutTopicArn": `arn:aws:sns:ap-northeast-1:941528793676:portals-me-${process.argv[2]}-entity-stream-hook`,
  "IdentityPoolId": process.env.COGNITO_IDP || fs.readFileSync('./token/cognito.key', 'utf8'),
  "JwtPrivate": process.env.JWT_PRIVATE ? process.env.JWT_PRIVATE.split('\\n').join('\n') : fs.readFileSync('./token/jwtES256.key', 'utf8'),
  "JwtPublic": process.env.JWT_PUBLIC ? process.env.JWT_PUBLIC.split('\\n').join('\n') : fs.readFileSync('./token/jwtES256.key.pub', 'utf8'),
  "TwitterKey": process.env.TWITTER_KEY || fs.readFileSync('./token/twitter.key', 'utf8'),
};

console.log(JSON.stringify(env));
