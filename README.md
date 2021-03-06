# Blog CLI

This Go CLI lets users store Markdown blog post files in a remote database. It's still a work in progress.

## Running the CLI

To simply run the CLI, invoke it with the path to a Markdown file:

```bash
go run ./ my-markdown-file.md
```

You can pass a path to the file as well:

```bash
go run ./ ../some-file.md
```

## Running tests

```bash
go test -v
``` 

## Running a local DynamoDB instance

**NOTE: This part is probably deprecated; we may not end up using DynamoDB to store posts**

Open up a terminal and run the following command to start a DynamoDB image and expose it on your local port 8000:


```bash
docker run --rm -p 8000:8000 amazon/dynamodb-local -jar DynamoDBLocal.jar -sharedDb -inMemory
```

If this is your first time running the command, Docker will take a bit longer to download the image before starting it. You'll see output like this, which means it's running successfully and waiting to take requests:

```
Initializing DynamoDB Local with the following configuration:
Port:	8000
InMemory:	true
DbPath:	null
SharedDb:	false
shouldDelayTransientStatuses:	false
CorsParams:	*

2021-01-01 18:48:08.463:INFO::main: Logging initialized @963ms to org.eclipse.jetty.util.log.StdErrLog
```

Leave that running and open a new terminal. We're going to use the AWS CLI to create the table our CLI uses to store create and read blog posts.

```bash
aws dynamodb create-table --table-name blog_posts --attribute-definitions AttributeName=slug,AttributeType=S --key-schema AttributeName=slug,KeyType=HASH --billing-mode PAY_PER_REQUEST --endpoint-url http://localhost:8000
```