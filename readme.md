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

To filter out pull requests where the TimeDifference is greater than 8 hours (in nanoseconds) using jq, you need to compare TimeDifference to the equivalent value of 8 hours in nanoseconds.
8 hours in nanoseconds:

    1 hour = 3,600,000,000,000 nanoseconds
    8 hours = 8 * 3,600,000,000,000 = 28,800,000,000,000 nanoseconds

```shell
jq '.[] | select(.TimeDifference > 28800000000000)' <myfile>.json
```

And to have the count:
```shell
 jq '[.[] | select(.TimeDifference > 28800000000000)] | length' <myFile>.json 
15
```
