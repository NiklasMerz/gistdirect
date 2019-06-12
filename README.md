# gistdirect

## Gistdirect is a simple URL redirection service / URL shortener. 

> Imagine bitly.com or goo.gl and your hosts file combined

[![Run on Google Cloud](https://storage.googleapis.com/cloudrun/button.svg)](https://console.cloud.google.com/cloudshell/editor?shellonly=true&cloudshell_image=gcr.io/cloudrun/button&cloudshell_git_repo=https://github.com/niklasmerz/gistdirect)

Create a Github Gist or just any plain file hosted somewhere. Gistdirect will grap that file and redirect the alias to the full url.

Sample "host" file:
```
gh https://github.com
```

[Test File](https://gist.github.com/NiklasMerz/a9b5905f742b5863197a0af0465a39f6)

The URL to the file is provided as an environment variable: `GIST_URL`

It is a simple Go function which is easy to run on any serverless platform. You can also run the binary somewhere and it will listen on port 8080. Executables are avaible from [releases](https://github.com/NiklasMerz/gistdirect/releases).

I got it running with Google Cloud Functions and Firebase hosting. You can deploy the function from "function.go" via the web interface or gcloud CLI. Firebase can be used to assign a [custom domain](https://firebase.google.com/docs/hosting/custom-domain) to the function. Firebase CLI and the "firebase.json" file help you to assing the function to the root of the domain.

## Using the Dockerfile

Gistdirect is on Docker Hub so you can just run:

```
docker run -p 8080:8080 -e GIST_URL='https://gist.github.com/NiklasMerz/a9b5905f742b5863197a0af0465a39f6/raw/' niklasmerz/gistdirect
```


