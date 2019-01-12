package hangman_henry_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHangmanHenry(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HangmanHenry Suite")
}
