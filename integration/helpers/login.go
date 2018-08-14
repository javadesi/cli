package helpers

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

func SetAPI() (string, string) {
	apiURL := GetAPI()
	skipSSLValidation := skipSSLValidation()
	Eventually(CF("api", apiURL, skipSSLValidation)).Should(Exit(0))
	return apiURL, skipSSLValidation
}

func UnsetAPI() {
	Eventually(CF("api", "--unset")).Should(Exit(0))
}

func skipSSLValidation() string {
	if skip, err := strconv.ParseBool(os.Getenv("SKIP_SSL_VALIDATION")); err == nil && !skip {
		return ""
	}
	return "--skip-ssl-validation"
}

func GetAPI() string {
	apiURL := os.Getenv("CF_INT_API")
	if apiURL == "" {
		return "https://api.bosh-lite.com"
	}
	if !strings.HasPrefix(apiURL, "http") {
		apiURL = fmt.Sprintf("https://%s", apiURL)
	}

	return apiURL
}

func LoginCF() string {
	username, password := GetCredentials()
	env := map[string]string{
		"CF_USERNAME": username,
		"CF_PASSWORD": password,
	}
	Eventually(CFWithEnv(env, "auth")).Should(Exit(0))

	return username
}

func LoginCFWithClientCredentials() string {
	username, password := SkipIfClientCredentialsNotSet()
	env := map[string]string{
		"CF_USERNAME": username,
		"CF_PASSWORD": password,
	}
	Eventually(CFWithEnv(env, "auth", "--client-credentials")).Should(Exit(0))

	return username
}

// GetCredentials returns back the username and the password.
func GetCredentials() (string, string) {
	username := os.Getenv("CF_INT_USERNAME")
	if username == "" {
		username = "admin"
	}
	password := os.Getenv("CF_INT_PASSWORD")
	if password == "" {
		password = "admin"
	}
	return username, password
}

// GetOIDCCredentials returns back the username and the password for OIDC origin.
func GetOIDCCredentials() (string, string) {
	username := os.Getenv("CF_INT_OIDC_USERNAME")
	if username == "" {
		username = "admin_oidc"
	}
	password := os.Getenv("CF_INT_OIDC_PASSWORD")
	if password == "" {
		password = "admin"
	}
	return username, password
}

func LogoutCF() {
	Eventually(CF("logout")).Should(Exit(0))
}

func TargetOrgAndSpace(org string, space string) {
	Eventually(CF("target", "-o", org, "-s", space)).Should(Exit(0))
}

func TargetOrg(org string) {
	Eventually(CF("target", "-o", org)).Should(Exit(0))
}

func ClearTarget() {
	LogoutCF()
	LoginCF()
}

func SetupCF(org string, space string) {
	LoginCF()
	CreateOrgAndSpace(org, space)
	TargetOrgAndSpace(org, space)
}
