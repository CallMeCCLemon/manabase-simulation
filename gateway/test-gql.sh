curl -g "http://localhost:8888/graphql" -d '
{
    echo(message: "Hello World!") {
        message
    }
}'