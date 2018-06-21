package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"testing"
	"time"
	"os/exec"
	"github.com/onsi/gomega/gbytes"
)

var pathToBinary string
var session *gexec.Session

var _ = BeforeSuite(func() {
	var err error
	pathToBinary, err = gexec.Build("github.com/bborbe/ip/cmd/ip-client")
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
	Context("without parameters", func() {
		It("returns with exitcode != 0", func() {
			session, err = gexec.Start(exec.Command(pathToBinary), GinkgoWriter, GinkgoWriter)
			Expect(err).To(BeNil())
			session.Wait(time.Second)
			Expect(session.ExitCode()).To(Equal(1))
		})
		It("returns with exitcode != 0", func() {
			session, err = gexec.Start(exec.Command(pathToBinary), GinkgoWriter, GinkgoWriter)
			Expect(err).To(BeNil())
			session.Wait(time.Second)
			Expect(session.Err).To(gbytes.Say("parameter url missing"))
		})
	})
})
