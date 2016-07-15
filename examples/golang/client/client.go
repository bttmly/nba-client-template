package nba

// NOTE: This file was generated automatically, do not edit directly!

import "net/http"

func PlayerProfile(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/playerprofile", params)
}

func PlayerInfo(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/commonplayerinfo", params)
}

func PlayersInfo(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/commonallplayers", params)
}

func TeamStats(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/leaguedashteamstats", params)
}

func TeamSplits(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/teamdashboardbygeneralsplits", params)
}

func TeamYears(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/commonteamyears", params)
}

func PlayerSplits(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/playerdashboardbygeneralsplits", params)
}

func Shots(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/shotchartdetail", params)
}

func Scoreboard(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/scoreboard", params)
}

func PlayByPlay(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/playbyplay", params)
}

func TeamHistoricalLeaders(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/teamhistoricalleaders", params)
}

func TeamInfoCommon(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/teaminfocommon", params)
}

func CommonTeamRoster(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/commonteamroster", params)
}

func TeamPlayerDashboard(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/teamplayerdashboard", params)
}

func Lineups(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/leaguedashlineups", params)
}

func PlayerTracking(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/leaguedashptstats", params)
}