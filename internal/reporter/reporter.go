package reporter

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type County struct {
	County         string
	State          string
	RollingAverage float64
}

type NewCasesReporter struct{}

func (r *NewCasesReporter) MakeReport(countyCsv string) string {
	totalCases := 0
	stateCounts := make(map[string]int)
	var counties []County
	lines := strings.Split(countyCsv, "\n")

	for _, line := range lines {
		tokens := strings.Split(line, ",")
		county := County{
			County: strings.TrimSpace(tokens[0]),
			State:  strings.TrimSpace(tokens[1]),
		}

		// compute rolling average
		lastDay := len(tokens) - 1
		firstDay := lastDay - 7 + 1
		if firstDay < 2 {
			firstDay = 2
		}
		n := float64(lastDay - firstDay + 1)
		sum := 0
		for day := firstDay; day <= lastDay; day++ {
			val, _ := strconv.Atoi(strings.TrimSpace(tokens[day]))
			sum += val
		}
		county.RollingAverage = float64(sum) / n

		// compute sum of cases
		cases := 0
		for i := 2; i < len(tokens); i++ {
			val, _ := strconv.Atoi(strings.TrimSpace(tokens[i]))
			cases += val
		}
		totalCases += cases
		stateCounts[county.State] = stateCounts[county.State] + cases
		counties = append(counties, county)
	}

	report := strings.Builder{}
	report.WriteString("County     State     Avg New Cases\n")
	report.WriteString("======     =====     =============\n")
	for _, county := range counties {
		report.WriteString(fmt.Sprintf("%-11s%-10s%.2f\n", county.County, county.State, county.RollingAverage))
	}
	report.WriteString("\n")
	states := make([]string, 0, len(stateCounts))
	for state := range stateCounts {
		states = append(states, state)
	}
	sort.Strings(states)
	for _, state := range states {
		report.WriteString(fmt.Sprintf("%s cases: %d\n", state, stateCounts[state]))
	}
	report.WriteString(fmt.Sprintf("Total Cases: %d\n", totalCases))

	return report.String()
}
