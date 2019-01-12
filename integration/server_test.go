package integration_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os/exec"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"github.com/henryaj/hangman-henry/server"
)

var _ = Describe("Server", func() {
	BeforeEach(func() {
		command := exec.Command(serverBinaryPath)
		_, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		time.Sleep(150 * time.Millisecond) // this is bad and you should feel bad
	})

	AfterEach(func() {
		gexec.KillAndWait()
	})

	Describe("GET /games", func() {
		BeforeEach(func() {
			resp, err := http.Post("http://0.0.0.0:8000/games", "", nil)
			Expect(err).NotTo(HaveOccurred())
			resp.Body.Close()
		})

		It("returns a list of all games", func() {
			resp, err := http.Get("http://0.0.0.0:8000/games")
			Expect(err).NotTo(HaveOccurred())
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)

			var games server.GamesMap
			err = json.Unmarshal(body, &games)
			Expect(err).NotTo(HaveOccurred())
			Expect(games).To(HaveLen(1))
		})
	})
})
