package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"testing"
	"time"
	"os/exec"
)

var pathToServerBinary string
var serverSession *gexec.Session

var _ = BeforeSuite(func() {
	var err error
	pathToServerBinary, err = gexec.Build("github.com/bborbe/ip/cmd/ip-client")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

func TestIpClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IpClient Suite")
}

var _ = Describe("ip-client", func() {
	var err error
	It("returns with exitcode != 0 if no parameters have been given", func() {
		serverSession, err = gexec.Start(exec.Command(pathToServerBinary), GinkgoWriter, GinkgoWriter)
		Expect(err).To(BeNil())
		serverSession.Wait(time.Second)
		Expect(serverSession.ExitCode()).To(Equal(1))
	})
})
