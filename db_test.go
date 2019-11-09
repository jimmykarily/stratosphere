package main_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/jimmykarily/stratosphere"
)

var _ = Describe("Db", func() {
	var (
		path string
	)

	Describe("Creating a new DB object", func() {
		Context("using a path to an existing file", func() {
			BeforeEach(func() {
				path = "fixtures/test.sqlite3"
			})

			It("should not return an error", func() {
				_, err := NewDb(path)
				Expect(err).To(BeNil())
			})
		})

		Context("with a path to a non existent file", func() {
			BeforeEach(func() {
				path = "fixtures/not_a_file"
			})

			It("should return an error", func() {
				_, err := NewDb(path)
				Expect(err.Error()).To(Equal("Database file doesn't exist: fixtures/not_a_file"))
			})
		})
	})

	Describe("#Activities", func() {
		var db *Db

		BeforeEach(func() {
			path = "fixtures/test.sqlite3"
			db, _ = NewDb(path)
		})

		It("returns all Activities", func() {
			activities, err := db.Activities()
			Expect(err).To(BeNil())
			Expect(len(activities)).To(Equal(30))
		})

		It("returns Activity objects", func() {
			activities, err := db.Activities()
			Expect(err).To(BeNil())
			Expect(fmt.Sprintf("%T", activities[0])).To(Equal("*main.Activity"))
		})
	})
})
