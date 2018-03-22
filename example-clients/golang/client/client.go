package nba

// NOTE: This file was generated automatically, do not edit directly!

import "net/http"

const UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.101 Safari/537.36"

const Referer = "http://stats.nba.com/scores/"

const Origin = "http://stats.nba.com"

func PlayerProfile(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/playerprofilev2", params)
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

func BoxScoreSummary(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/boxscoresummaryv2", params)
}

func BoxScore(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/boxscoretraditionalv2", params)
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

func HomepageV2(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/homepagev2", params)
}

func AssistTracker(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/assisttracker", params)
}

func PlayerStats(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/leaguedashplayerstats", params)
}

func PlayerClutch(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/leaguedashplayerclutch", params)
}

func TeamClutch(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/leaguedashteamclutch", params)
}

func PlayerShooting(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/leaguedashplayerptshot", params)
}

func TeamShooting(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/leaguedashteamptshot", params)
}

func LeagueGameLog(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/leaguegamelog", params)
}

func LeagueLeaders(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/leagueLeaders", params)
}

func LeagueStandings(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/leaguestandings", params)
}

func PlayerHustleLeaders(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/leaguehustlestatsplayerleaders", params)
}

func TeamHustleLeaders(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/leaguehustlestatsteamleaders", params)
}

func PlayerHustle(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/leaguehustlestatsplayer", params)
}

func TeamHustle(params map[string]string) (resp *http.Response, err error) {
    return apiGet("http://stats.nba.com/stats/leaguehustlestatsteam", params)
}