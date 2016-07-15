package main

import (
    "os"
    "fmt"
    "io/ioutil"
    
    "github.com/nickb1080/nba/client"
)

func check (e error) {
    if (e != nil) {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
}

func main () {
    params := map[string]string{
        "GraphEndSeason": "2015-16",
        "PlayerID": "201939",
        "Season": "2015-16",
        "GraphStat": "PTS",
        "GraphStartSeason": "2015-16",
        "LeagueID": "00",
        "SeasonType": "Regular Season",
    }
    
    resp, e := nba.PlayerProfile(params)
    check(e)

    body, e := ioutil.ReadAll(resp.Body)
    check(e)

    fmt.Print(string(body))
}