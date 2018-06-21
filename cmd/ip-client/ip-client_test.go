package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"testing"
	"time"
	"os/exec"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/ghttp"
	"net/http"
)

var pathToBinary string
var session *gexec.Session
var server *ghttp.Server

var _ = BeforeSuite(func() {
	var err error
	pathToBinary, err = gexec.Build("github.com/bborbe/ip/cmd/ip-client")
	Expect(err).NotTo(HaveOccurred())
})

var _ = BeforeEach(func() {
	server = ghttp.NewServer()
	server.RouteToHandler(http.MethodGet, "/", ghttp.RespondWith(http.StatusOK, "1.2.3.4"))
})

var _ = AfterEach(func() {
	session.Interrupt()
	Eventually(session).Should(gexec.Exit())
	server.Close()
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
	var args []string
	Context("without parameters", func() {
		BeforeEach(func() {
			args = []string{}
		})
		It("returns with exitcode != 0", func() {
			session, err = gexec.Start(exec.Command(pathToBinary, args...), GinkgoWriter, GinkgoWriter)
			Expect(err).To(BeNil())
			session.Wait(time.Second)
			Expect(session.ExitCode()).To(Equal(1))
		})
		It("say parameter url missing", func() {
			session, err = gexec.Start(exec.Command(pathToBinary, args...), GinkgoWriter, GinkgoWriter)
			Expect(err).To(BeNil())
			session.Wait(time.Second)
			Expect(session.Err).To(gbytes.Say("parameter url missing"))
		})
	})
	Context("with invalid url parameter", func() {
		BeforeEach(func() {
			args = []string{"--url", "foobar"}
		})
		It("returns with exitcode != 0", func() {
			session, err = gexec.Start(exec.Command(pathToBinary, args...), GinkgoWriter, GinkgoWriter)
			Expect(err).To(BeNil())
			session.Wait(time.Second)
			Expect(session.ExitCode()).To(Equal(1))
		})
		It("say parameter url invalid", func() {
			session, err = gexec.Start(exec.Command(pathToBinary, args...), GinkgoWriter, GinkgoWriter)
			Expect(err).To(BeNil())
			session.Wait(time.Second)
			Expect(session.Err).To(gbytes.Say("parameter url invalid"))
		})
	})
	Context("with not existing ip-server url", func() {
		BeforeEach(func() {
			args = []string{"--url", "http://localhost:1234"}
		})
		It("returns with exitcode = 0", func() {
			session, err = gexec.Start(exec.Command(pathToBinary, args...), GinkgoWriter, GinkgoWriter)
			Expect(err).To(BeNil())
			session.Wait(time.Second)
			Expect(session.ExitCode()).To(Equal(1))
		})
		It("says get ip failed", func() {
			session, err = gexec.Start(exec.Command(pathToBinary, args...), GinkgoWriter, GinkgoWriter)
			Expect(err).To(BeNil())
			session.Wait(time.Second)
			Expect(session.Err).To(gbytes.Say("get ip failed"))
		})
	})
	Context("with valid ip-server url", func() {
		BeforeEach(func() {
			args = []string{"--url", server.URL()}
		})
		It("returns with exitcode = 0", func() {
			session, err = gexec.Start(exec.Command(pathToBinary, args...), GinkgoWriter, GinkgoWriter)
			Expect(err).To(BeNil())
			session.Wait(time.Second)
			Expect(session.ExitCode()).To(Equal(0))
		})
		It("says get ip failed", func() {
			session, err = gexec.Start(exec.Command(pathToBinary, args...), GinkgoWriter, GinkgoWriter)
			Expect(err).To(BeNil())
			session.Wait(time.Second)
			Expect(session.Out).To(gbytes.Say("1.2.3.4"))
		})
	})
})
