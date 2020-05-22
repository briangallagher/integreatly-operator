package common

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

// The running of this test is determined by the "UPGRADE_TESTING" environment variable
// This is (will be) set in the upgrade pipeline:
// https://gitlab.cee.redhat.com/integreatly-qe/ci-cd/blob/master/pipelines/rhmi-upgrade.groovy
//
// The logic behind this test is specified in: https://issues.redhat.com/browse/INTLY-7394
// In summary:
// - before upgrade current version is n -1 and target version is empty
// - during upgrade current version in n-1 and target version is n
// - once stage goes to completed the current version should be n and target version empty
//
// However, at the time of writing, the upgrade pipeline deploys 2.2.0 of rhmi which does not
// contain the CR fields in question i.e. toVersion and Version. So, the logic has to be altered
// to account for this. Hence, the testing logic is:
// - during upgrade the target version is n
// - After upgrade the current version should be n and target version empty
//
// TODO: When the upgrade pipeline is updated to install 2.2.1 update this test to
// run the former logic. Also, add testing around the prometheus metric: rhmi_version
func TestUpgradeVersions(t *testing.T, ctx *TestingContext) {

	var (
		preUpgrade                          = os.Getenv("PRE_UPGRADE")
		postUpgrade                         = os.Getenv("POST_UPGRADE")
		inProgressUpgradeCurrentRhmiVersion = os.Getenv("IN_PROGRESS_UPGRADE_CURRENT_RHMI_VERSION")
		inProgressUpgradeTargetRhmiVersion  = os.Getenv("IN_PROGRESS_UPGRADE_TARGET_RHMI_VERSION")
	)

	t.Logf("Test rhmi upgrades, preUpgrade: %s, postUpgrade: %s, currentRhmiVersion: %s, targetRhmiVersion: %s", preUpgrade, postUpgrade, currentRhmiVersion, targetRhmiVersion)

	if postUpgrade == "true" {

		// Verify the RHMI versions during the upgrade

		// Convert SEMVER to int
		current := semverToInt(t, inProgressUpgradeCurrentRhmiVersion)
		target := semverToInt(t, inProgressUpgradeTargetRhmiVersion)

		if target <= current {
			t.Fatal("The target version during an upgrade is less than or equal the current, %s, %s", inProgressUpgradeTargetRhmiVersion, inProgressUpgradeCurrentRhmiVersion)
		}

		// Verify the RHMI versions now

		// get console master url
		rhmi, err := getRHMI(ctx.Client)
		if err != nil {
			t.Fatalf("error getting RHMI CR: %v", err)
		}

		if rhmi.Status.ToVersion != "" {
			t.Fatal("ToVersion should be empty post upgrade")
		}

		if rhmi.Status.Version != inProgressUpgradeTargetRhmiVersion {
			t.Fatal("Post Upgrade, the current version should be the same as the target version during the upgrade")
		}
	}
}

func semverToInt(t *testing.T, semver string) int {
	num := strings.Replace(semver, ".", "", -1)
	val, err := strconv.Atoi(num)
	if err != nil {
		t.Fatal("Failed to convert string version to int")
	}
	return val
}
