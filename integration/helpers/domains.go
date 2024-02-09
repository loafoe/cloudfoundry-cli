package helpers

import (
	"fmt"
	"regexp"
	"strings"

	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

func DomainName(prefix ...string) string {
	if len(prefix) > 0 {
		return fmt.Sprintf("integration-%s.com", PrefixedRandomName(prefix[0]))
	}
	return fmt.Sprintf("integration%s.com", PrefixedRandomName(""))
}

type Domain struct {
	Org  string
	Name string
}

func NewDomain(org string, name string) Domain {
	return Domain{
		Org:  org,
		Name: name,
	}
}

// globally cached
var foundDefaultDomain string

func DefaultSharedDomain() string {
	if foundDefaultDomain == "" {
		session := CF("domains")
		Eventually(session).Should(Exit(0))

		regex := regexp.MustCompile(`(.+?)\s+shared`)

		output := strings.Split(string(session.Out.Contents()), "\n")
		for _, line := range output {
			if line != "" && !strings.HasPrefix(line, "integration-") {
				matches := regex.FindStringSubmatch(line)
				if len(matches) == 2 {
					foundDefaultDomain = matches[1]
					break
				}
			}
		}
		Expect(foundDefaultDomain).ToNot(BeEmpty())
	}
	return foundDefaultDomain
}

func (d Domain) Create() {
	Eventually(CF("create-domain", d.Org, d.Name)).Should(Exit(0))
	Eventually(CF("domains")).Should(And(Exit(0), Say(d.Name)))
}

func (d Domain) CreatePrivate() {
	Eventually(CF("create-private-domain", d.Org, d.Name)).Should(Exit(0))
}

func (d Domain) CreateShared() {
	Eventually(CF("create-shared-domain", d.Name)).Should(Exit(0))
}

func (d Domain) CreateWithRouterGroup(routerGroup string) {
	Eventually(CF("create-shared-domain", d.Name, "--router-group", routerGroup)).Should(Exit(0))
}

func (d Domain) CreateInternal() {
	Eventually(CF("create-shared-domain", d.Name, "--internal")).Should(Exit(0))
}

func (d Domain) Share() {
	Eventually(CF("share-private-domain", d.Org, d.Name)).Should(Exit(0))
}

func (d Domain) V7Share(orgName string) {
	Eventually(CF("share-private-domain", orgName, d.Name)).Should(Exit(0))
}

func (d Domain) Delete() {
	Eventually(CF("delete-domain", d.Name, "-f")).Should(Exit(0))
}

func (d Domain) DeleteShared() {
	Eventually(CF("delete-shared-domain", d.Name, "-f")).Should(Exit(0))
}
