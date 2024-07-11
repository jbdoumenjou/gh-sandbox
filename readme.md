# Context
This project is a sandbox to play with the Github API.
One of the purpose is to extract some metric from the code lifecycle inspired by DORA metrics.

# Usage

First, you must create a token that allow to read the repository data.
Then define an the following env variable.

```shell
export GITHUB_TOKEN=XXXXXXXXXX
```

Then run the main:

```shell
go run ./... --owner=yourOrgOrUsername --repo=yourRepo --count=20
```

It will display the result on the standard output.