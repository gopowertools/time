package ttime_test

import (
	"time"
	"ttime"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Provider", func() {
	var (
		err              error
		t, ft, ftPlusSec time.Time
		sut              *ttime.Provider
	)

	BeforeEach(func() {
		ft, err = time.Parse(time.RFC3339, "2022-04-24T00:00:00Z")
		if err != nil {
			Fail("Invalid future test time" + err.Error())
		}

		ftPlusSec, err = time.Parse(time.RFC3339, "2022-04-24T00:00:01Z")
		if err != nil {
			Fail("Invalid future test time" + err.Error())
		}

		sut = ttime.NewProvider()
	})

	Describe("Now", func() {
		Context("FixNow() not previously called", func() {
			It("returns the system time", func() {
				t = sut.Now()
				Expect(t).Should(BeTemporally("~", time.Now(), time.Second))
			})
		})
		Context("FixNow() previously called", func() {
			BeforeEach(func() {
				sut.FixNow(ft)
			})
			It("returns the fixed time", func() {
				t = sut.Now()
				Expect(t).Should(BeTemporally("==", ft))
			})
		})
	})
	Describe("Since", func() {
		Context("The given time is one second before 'now'", func() {
			BeforeEach(func() {
				sut.FixNow(ftPlusSec)
			})
			It("returns 1 second", func() {
				d := sut.Since(ft)
				Expect(d).To(Equal(time.Second))
			})
		})
		Context("The given time is one second after 'now'", func() {
			BeforeEach(func() {
				sut.FixNow(ft)
			})
			It("returns -1 second", func() {
				d := sut.Since(ftPlusSec)
				Expect(d).To(Equal(-time.Second))
			})
		})
	})
	Describe("Until", func() {
		Context("The given time is one second after 'now'", func() {
			BeforeEach(func() {
				sut.FixNow(ft)
			})
			It("returns 1 second", func() {
				d := sut.Until(ftPlusSec)
				Expect(d).To(Equal(time.Second))
			})
		})
		Context("The given time is one second before 'now'", func() {
			BeforeEach(func() {
				sut.FixNow(ftPlusSec)
			})
			It("returns -1 second", func() {
				d := sut.Until(ft)
				Expect(d).To(Equal(-time.Second))
			})
		})
	})
})
