# emailambda

## Motivation

Like [emaild](https://github.com/4thel00z/emaild) but deployable to netlify.

## Deploy to netlify

<a href="https://app.netlify.com/start/deploy?repository=https://github.com/4thel00z/emailambda">
    <img src="https://www.netlify.com/img/deploy/button.svg" title="Deploy to Netlify">
</a>

## Configuration

There are three main env vars that configure emailambda:

| Name         | Description                                                                                                                                                         |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| BASIC_AUTH   | The value that the http authorization header is expected to have. This value is very important because it secures your lambda from being called by evil häggerbois. |
| GMAIL_TOKEN  | The gmail token generated via gmail-token (find it in the 4thel00z/emaild github repository) encoded as a base64 string                                             |
| GMAIL_CONFIG | The credentials.json file that you have generate followed my smartboi tutorial (also in the 4thel00z/emaild github repository) encoded as a base64 string           |


## License

This project is licensed under the GPL-3 license.
