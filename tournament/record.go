package tournament

import (
	"fmt"
	"strings"
)

const (
	whiteSpaceFill = 31
	header         = "Team                           | MP |  W |  D |  L |  P\n"
	recordTemplate = "%s%s|  %d |  %d |  %d |  %d |  %d\n"
)

type record struct {
	team                                 string
	matches, wins, losses, draws, points int
}

func (r record) String() string {
	var fillUp int
	if len(r.team) < whiteSpaceFill {
		fillUp = whiteSpaceFill - len(r.team)
	}
	return fmt.Sprintf(recordTemplate, r.team, strings.Repeat(" ", fillUp), r.matches, r.wins, r.draws, r.losses, r.points)
}

type records []*record

func (rr records) Len() int {
	return len(rr)
}

func (rr records) Less(i, j int) bool {
	if rr[i].points == rr[j].points {
		return rr[i].team < rr[j].team
	}
	return rr[i].points > rr[j].points
}

func (rr records) Swap(i, j int) {
	rr[i], rr[j] = rr[j], rr[i]
}

func (rr records) String() string {
	b := strings.Builder{}
	b.WriteString(header)
	for _, r := range rr {
		b.WriteString(r.String())
	}
	return b.String()
}
