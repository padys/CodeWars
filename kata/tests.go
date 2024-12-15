package kata

import "math"

var _ = Describe("Tests", func() {
	It("Sample tests", func() {
		Expect(Int32ToIp(2154959208)).To(Equal("128.114.17.104"))
		Expect(Int32ToIp(2149583361)).To(Equal("128.32.10.1"))
		Expect(Int32ToIp(0)).To(Equal("0.0.0.0"))
		Expect(Int32ToIp(math.MaxUint32)).To(Equal("255.255.255.255"))
	})
})
