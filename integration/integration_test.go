package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
)

var _ = Describe("Integration", func() {
	Describe("GET /investors/:investorId/connections", func() {
		It("returns connections", func() {
			resp, err := http.Get("http://localhost:8000/investors/90234/connections")
			Expect(err).NotTo(HaveOccurred())
			Expect(output(resp)).To(Equal("Dennis Chambers"))
		})
	})

	Describe("GET /investors/:investorId/mutual/:otherInvestorId", func() {
		It("returns mutual connections", func() {
			resp, err := http.Get("http://localhost:8000/investors/1/mutual/2")
			Expect(err).NotTo(HaveOccurred())
			Expect(output(resp)).To(Equal(`Antonia Brandt
Louanne Richard
Debbie Byrd
Gustav Jørgensen
Océane Rey
Victor Scott
Muharrem Jaeger
Olive Johnson
Sophie Lavoie
Constance Simon
Ianis Almeida
Amelia Roche
Marinice Costa
Martha Byrd
Felix Rasmussen
Morgane Aubert`))
		})
	})
})
