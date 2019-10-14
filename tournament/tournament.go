// Package tournament provides functionality for keeping track of small football matches.
package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

const (
	win          = "win"
	loss         = "loss"
	draw         = "draw"
	hostIndex    = 0
	visitorIndex = 1
	resultIndex  = 2
	pointsPerWin = 3
)

// Tally sums up the results of football matches and creates a file with a table like the following:
//
//Team                           | MP |  W |  D |  L |  P
//Devastating Donkeys            |  3 |  2 |  1 |  0 |  7
//Allegoric Alaskans             |  3 |  2 |  0 |  1 |  6
//Blithering Badgers             |  3 |  1 |  0 |  2 |  3
//Courageous Californians        |  3 |  0 |  1 |  2 |  1
//
// Will skip comments (starts with #) and blank lines.
func Tally(reader io.Reader, writer io.Writer) error {
	teamToRecord := make(map[string]*record)
	sc := bufio.NewScanner(reader)
	for sc.Scan() {
		match := sc.Text()
		if isComment(match) || isEmptyRow(match) {
			continue
		}
		matchData := strings.Split(match, ";")

		if len(matchData) < 3 {
			return fmt.Errorf("invalid match: [%v]", matchData)
		}

		if err := evaluate(teamToRecord, matchData); err != nil {
			return err
		}
	}
	return writeTable(teamToRecord, writer)
}

func isComment(match string) bool {
	return strings.HasPrefix(match, "#")
}

func isEmptyRow(match string) bool {
	return len(strings.TrimSpace(match)) == 0
}

func evaluate(teamToRecord map[string]*record, matchData []string) error {
	host := getOrCreateRecord(matchData[hostIndex], teamToRecord)
	visitor := getOrCreateRecord(matchData[visitorIndex], teamToRecord)

	matchResult := matchData[resultIndex]
	switch matchResult {
	case win:
		host.wins++
		host.points += pointsPerWin
		visitor.losses++
	case draw:
		host.draws++
		host.points++
		visitor.draws++
		visitor.points++
	case loss:
		host.losses++
		visitor.wins++
		visitor.points += pointsPerWin
	default:
		return fmt.Errorf("invalid result for match: [%s]", matchResult)
	}
	host.matches++
	visitor.matches++
	return nil
}

func getOrCreateRecord(team string, recordMap map[string]*record) *record {
	r, ok := recordMap[team]
	if !ok {
		r = &record{team: team}
		recordMap[team] = r
	}
	return r
}

func writeTable(teamToRecord map[string]*record, writer io.Writer) error {
	var rr = records{}
	for _, r := range teamToRecord {
		rr = append(rr, r)
	}
	sort.Sort(rr)
	_, err := io.WriteString(writer, rr.String())
	return err
}
