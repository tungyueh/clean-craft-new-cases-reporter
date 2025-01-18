package reporter

import (
    "testing"
)

const DELTA = 0.0001

var reporter *NewCasesReporter

func setUp() {
    reporter = &NewCasesReporter{}
}

func TestCountyReport(t *testing.T) {
    setUp()
    report := reporter.MakeReport("" +
        "c1, s1, 1, 1, 1, 1, 1, 1, 1, 7\n" +
        "c2, s2, 2, 2, 2, 2, 2, 2, 2, 7")
    expected := "" +
        "County     State     Avg New Cases\n" +
        "======     =====     =============\n" +
        "c1         s1        1.86\n" +
        "c2         s2        2.71\n\n" +
        "s1 cases: 14\n" +
        "s2 cases: 21\n" +
        "Total Cases: 35\n"
    if report != expected {
        t.Errorf("Expected:\n%s\nGot:\n%s", expected, report)
    }
}

func TestStateWithTwoCounties(t *testing.T) {
    setUp()
    report := reporter.MakeReport("" +
        "c1, s1, 1, 1, 1, 1, 1, 1, 1, 7\n" +
        "c2, s1, 2, 2, 2, 2, 2, 2, 2, 7")
    expected := "" +
        "County     State     Avg New Cases\n" +
        "======     =====     =============\n" +
        "c1         s1        1.86\n" +
        "c2         s1        2.71\n\n" +
        "s1 cases: 35\n" +
        "Total Cases: 35\n"
    if report != expected {
        t.Errorf("Expected:\n%s\nGot:\n%s", expected, report)
    }
}

func TestStatesWithShortLines(t *testing.T) {
    setUp()
    report := reporter.MakeReport("" +
        "c1, s1, 1, 1, 1, 1, 7\n" +
        "c2, s2, 7")
    expected := "" +
        "County     State     Avg New Cases\n" +
        "======     =====     =============\n" +
        "c1         s1        2.20\n" +
        "c2         s2        7.00\n\n" +
        "s1 cases: 11\n" +
        "s2 cases: 7\n" +
        "Total Cases: 18\n"
    if report != expected {
        t.Errorf("Expected:\n%s\nGot:\n%s", expected, report)
    }
}
