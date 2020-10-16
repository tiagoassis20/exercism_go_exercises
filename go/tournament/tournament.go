package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Scores struct {
	Team                             string
	Matches, Win, Draw, Loss, Points int
}

func (s Scores) String() string {
	return fmt.Sprintf("%-30s | %2d | %2d | %2d | %2d | %2d", s.Team, s.Matches, s.Win, s.Draw, s.Loss, s.Points)
}

func Tally(in io.Reader, out io.Writer) error {
	// parses the input
	board, err := parseScores(in)
	if err != nil {
		return err
	}

	// sort slice by points and by name
	sort.Slice(board, func(i, j int) bool {
		return board[i].Points > board[j].Points || (board[i].Points == board[j].Points && board[i].Team < board[j].Team)
	})

	// write board on the output
	fmt.Fprintln(out, fmt.Sprintf("%-30s | %2s | %2s | %2s | %2s | %2s", "Team", "MP", "W", "D", "L", "P"))
	for _, score := range board {
		fmt.Fprintln(out, score)
	}
	return nil
}

func parseScores(in io.Reader) ([]Scores, error) {
	var teams = make(map[string]Scores)

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		if !(strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "") {
			parsed := strings.Split(line, ";")
			if len(parsed) != 3 {
				return nil, fmt.Errorf("bad format")
			}
			if err := updateTeam(teams, parsed[0], parsed[2], true); err != nil {
				return nil, err
			}
			if err := updateTeam(teams, parsed[1], parsed[2], false); err != nil {
				return nil, err
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return mapToSlice(teams), nil
}

func updateTeam(teams map[string]Scores, teamName, result string, homeTeam bool) error {
	team := teams[teamName]
	team.Team = teamName
	team.Matches++

	switch result {
	case "win":
		if homeTeam {
			team.Win++
			team.Points += 3
		} else {
			team.Loss++
		}
	case "draw":
		team.Draw++
		team.Points++
	case "loss":
		if !homeTeam {
			team.Win++
			team.Points += 3
		} else {
			team.Loss++
		}
	default:
		return fmt.Errorf("Result not recognized")
	}
	teams[teamName] = team
	return nil
}

func mapToSlice(m map[string]Scores) []Scores {
	slice := []Scores{}
	for _, value := range m {
		slice = append(slice, value)
	}
	return slice
}
