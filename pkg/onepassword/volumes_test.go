package onepassword

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
)

func TestAreVolmesUsingSecrets(t *testing.T) {
	secretNamesToSearch := map[string]*corev1.Secret{
		"onepassword-database-secret": {},
		"onepassword-api-key":         {},
	}

	volumeSecretNames := []string{
		"onepassword-database-secret",
		"onepassword-api-key",
		"some_other_key",
	}

	volumes := generateVolumes(volumeSecretNames)

	if !AreVolumesUsingSecrets(volumes, secretNamesToSearch) {
		t.Errorf("Expected that volumes were using secrets but they were not detected.")
	}
}

func TestAreVolumesNotUsingSecrets(t *testing.T) {
	secretNamesToSearch := map[string]*corev1.Secret{
		"onepassword-database-secret": {},
		"onepassword-api-key":         {},
	}

	volumeSecretNames := []string{
		"some_other_key",
	}

	volumes := generateVolumes(volumeSecretNames)

	if AreVolumesUsingSecrets(volumes, secretNamesToSearch) {
		t.Errorf("Expected that volumes were not using secrets but they were detected.")
	}
}
