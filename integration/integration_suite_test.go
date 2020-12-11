package integration_test

import (
	"github.com/onsi/gomega/gexec"
	"io/ioutil"
	"net/http"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

var (
	serverSession *gexec.Session
)

var _ = BeforeSuite(func() {
	//Expect(os.Getenv("DATABASE_URI"), Not(BeEmpty()))

	var err error
	binaryPath, err := gexec.Build(filepath.Join("github.com", "aemengo", "investorbook-api"))
	Expect(err).NotTo(HaveOccurred())

	serverCommand := exec.Command(binaryPath)
	serverSession, err = gexec.Start(serverCommand, nil, nil)
	Expect(err).NotTo(HaveOccurred())

	// Wait for server to come up...
	time.Sleep(time.Second)
})

var _ = AfterSuite(func() {
	serverSession.Terminate()
	gexec.CleanupBuildArtifacts()
})

func output(resp *http.Response) string {
	contents, err := ioutil.ReadAll(resp.Body)
	Expect(err).NotTo(HaveOccurred())
	return string(contents)
}
