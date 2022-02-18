# scplanner

golang API for http://scplanner.net

## Details

We use the `jwt` value of the cookie for authentication. You can set that by running the following with the cookie that you get from inspecting the dev console:

```
./scripts/usecookie.sh <cookie>
```

See `cli/main.go` for commands. The use case is searching for contracts of a certain type and then placing bids for a given URL, e.g. *placing bids for all "Rock" contracts for the URL https://soundcloud.com/someArtist/someTrack*:

```
./scripts/run.sh SearchAndBid --genre "Rock" --url https://soundcloud.com/someArtist/someTrack
```