# gistdirect

## Gistdirect is a simple URL redirection service. Imagine bitly.com or goo.gl and your hosts file combined

Create a Github Gist or just any plain file like with an alis or URL and provide the URL to this file to Gistdirect. If you hit any path aliased, it will redirect.

Sample "host" file:
```
gh https://github.com
```

It is a simple Go function which should be easy to run on any serverless platform. I got it running with Google Cloud Functions and Firebase hosting.