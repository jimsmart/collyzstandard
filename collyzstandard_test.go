package collyzstandard_test

import (
	"encoding/hex"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/jimsmart/collysqlite"
	"github.com/jimsmart/collyzstandard"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cache", func() {

	It("should Init and Destroy", func() {
		Context("with a vanilla name (no path)", func() {
			name := randomName("test-cache-db-")
			filename := name + ".sqlite"
			c := collysqlite.NewCache(name)
			w := collyzstandard.NewCompressor(c)
			Expect(w.Init()).To(BeNil())
			Expect(filename).To(BeAnExistingFile())
			Expect(c.Destroy()).To(BeNil())
			Expect(filename).NotTo(BeAnExistingFile())
		})
	})

	It("should Put, Get and Remove", func() {
		name := randomName("test-cache-db-")
		c := collysqlite.NewCache(name)
		w := collyzstandard.NewCompressor(c)
		w.Logging = true
		Expect(w.Init()).To(BeNil())
		defer c.Destroy()

		// Put.
		url := "http://example.org"
		data := []byte{0, 1, 2, 3, 4, 5, 6, 7}
		Expect(w.Put(url, data)).To(BeNil())
		// Get existing.
		got, err := w.Get(url)
		Expect(err).To(BeNil())
		Expect(got).To(Equal(data))
		// Remove.
		Expect(w.Remove(url)).To(BeNil())
		// Get non-existing.
		got, err = w.Get(url)
		Expect(err).To(BeNil())
		Expect(got).To(BeNil())
		// Remove non-existing.
		Expect(w.Remove(url)).To(BeNil())
	})

	It("should Put and Get some real webpages", func() {
		name := randomName("test-cache-db-")
		c := collysqlite.NewCache(name)
		w := collyzstandard.NewCompressor(c)
		w.Logging = true
		Expect(w.Init()).To(BeNil())
		defer c.Destroy()

		// Put.
		url := "https://google.com/"
		data, err := get(url)
		Expect(err).To(BeNil())
		Expect(w.Put(url, data)).To(BeNil())
		// Get.
		got, err := w.Get(url)
		Expect(err).To(BeNil())
		Expect(got).To(Equal(data))
		// Remove.
		// Expect(c.Remove(url)).To(BeNil())

		// Put.
		url = "https://facebook.com/"
		data, err = get(url)
		Expect(err).To(BeNil())
		Expect(w.Put(url, data)).To(BeNil())
		// Get.
		got, err = w.Get(url)
		Expect(err).To(BeNil())
		Expect(got).To(Equal(data))
		// Remove.
		// Expect(c.Remove(url)).To(BeNil())

		// Put.
		url = "https://twitter.com/"
		data, err = get(url)
		Expect(err).To(BeNil())
		Expect(w.Put(url, data)).To(BeNil())
		// Get.
		got, err = w.Get(url)
		Expect(err).To(BeNil())
		Expect(got).To(Equal(data))
		// Remove.
		Expect(c.Remove(url)).To(BeNil())
	})

})

func randomName(prefix string) string {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	h := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(h, b)
	return prefix + string(h)
}

func get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
