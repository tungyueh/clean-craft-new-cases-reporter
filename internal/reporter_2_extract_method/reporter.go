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

type NewCasesReporter struct {
	totalCases  int
	stateCounts map[string]int
	counties    []County
}

func (r *NewCasesReporter) MakeReport(countyCsv string) string {
	lines := strings.Split(countyCsv, "\n")
	r.calculateCounties(lines)

	report := r.makeHeader()
	report.WriteString(r.makeCountyDetails())
	report.WriteString("\n")
	report.WriteString(r.makeStateTotals())
	report.WriteString(fmt.Sprintf("Total Cases: %d\n", r.totalCases))
	return report.String()
}

func (r *NewCasesReporter) makeStateTotals() string {
	states := make([]string, 0, len(r.stateCounts))
	for state := range r.stateCounts {
		states = append(states, state)
	}
	sort.Strings(states)
	statesTotal := ""
	for _, state := range states {
		x := fmt.Sprintf("%s cases: %d\n", state, r.stateCounts[state])
		statesTotal += x
	}
	return statesTotal
}

func (r *NewCasesReporter) makeCountyDetails() string {
	details := ""
	for _, county := range r.counties {
		x := fmt.Sprintf("%-11s%-10s%.2f\n", county.County, county.State, county.RollingAverage)
		details += x
	}
	return details
}

func (*NewCasesReporter) makeHeader() *strings.Builder {
	report := strings.Builder{}
	report.WriteString("County     State     Avg New Cases\n")
	report.WriteString("======     =====     =============\n")
	return &report
}

func (r *NewCasesReporter) calculateCounties(lines []string) {
	r.stateCounts = make(map[string]int)
	for _, line := range lines {
		county := r.calcluateCounty(line)
		r.counties = append(r.counties, county)
	}
}

func (r *NewCasesReporter) calcluateCounty(line string) County {
	tokens := strings.Split(line, ",")
	county := County{
		County: strings.TrimSpace(tokens[0]),
		State:  strings.TrimSpace(tokens[1]),
	}

	county.RollingAverage = r.calculateRollingAverage(tokens)

	cases := r.calculateSumOfCases(tokens)
	r.totalCases += cases
	r.incrementStateCount(county.State, cases)
	return county
}

func (r *NewCasesReporter) incrementStateCount(state string, cases int) {
	r.stateCounts[state] = r.stateCounts[state] + cases
}

func (*NewCasesReporter) calculateSumOfCases(tokens []string) int {
	cases := 0
	for i := 2; i < len(tokens); i++ {
		val, _ := strconv.Atoi(strings.TrimSpace(tokens[i]))
		cases += val
	}
	return cases
}

func (*NewCasesReporter) calculateRollingAverage(tokens []string) float64 {
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
	rollingAverage := float64(sum) / n
	return rollingAverage
}
