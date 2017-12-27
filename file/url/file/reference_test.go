package file_test

import (
	"errors"
	"io/ioutil"

	"github.com/dpb587/metalink/file"
	. "github.com/dpb587/metalink/file/url/file"

	boshsysfakes "github.com/cloudfoundry/bosh-utils/system/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("File", func() {
	var subject file.Reference
	var fs *boshsysfakes.FakeFileSystem

	BeforeEach(func() {
		fs = boshsysfakes.NewFakeFileSystem()

		fs.WriteFileString("/somewhere/useful", "something useful")
	})

	Describe("NewReference", func() {
		It("expands paths", func() {
			fs.ExpandPathExpanded = "/root/somewhere"

			subject := NewReference(fs, "~/somewhere")

			Expect(subject).ToNot(BeNil())
			Expect(subject.ReaderURI()).To(Equal("file:///root/somewhere"))
		})
	})

	Describe("Name", func() {
		It("gives the base name", func() {
			subject = NewReference(fs, "/somewhere/useful")

			value, err := subject.Name()

			Expect(err).ToNot(HaveOccurred())
			Expect(value).To(Equal("useful"))
		})
	})

	Describe("Size", func() {
		It("gives the size", func() {
			subject = NewReference(fs, "/somewhere/useful")

			value, err := subject.Size()

			Expect(err).ToNot(HaveOccurred())
			Expect(value).To(Equal(uint64(16)))
		})

		It("errors gracefully", func() {
			fs.RegisterOpenFile("/somewhere/useful", &boshsysfakes.FakeFile{
				StatErr: errors.New("fake-err"),
			})

			subject = NewReference(fs, "/somewhere/useful")

			_, err := subject.Size()

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Checking file size"))
		})
	})

	Describe("Reader", func() {
		It("opens for reading", func() {
			subject = NewReference(fs, "/somewhere/useful")

			reader, err := subject.Reader()

			Expect(err).ToNot(HaveOccurred())

			readerString, _ := ioutil.ReadAll(reader)

			Expect(string(readerString)).To(Equal("something useful"))
		})

		It("errors gracefully", func() {
			fs.OpenFileErr = errors.New("fake-err")

			subject = NewReference(fs, "/somewhere/useful")

			_, err := subject.Reader()

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Opening file"))
			Expect(err.Error()).To(ContainSubstring("fake-err"))
		})
	})
})
