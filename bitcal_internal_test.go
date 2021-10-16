package bitcal

import (
	"fmt"
	"math"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("bitcal internal module", func() {
	Context("date2index()", func() {
		When("2009/1/1", func() {
			It("returns 0,0", func() {

				start_date := time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC)

				arr_ind, bit_ind := date2index(start_date)

				Expect(arr_ind).To(Equal(0))
				Expect(bit_ind).To(Equal(0))
			})
		})

		When("2009/12/31", func() {
			It("returns 5,44", func() {

				start_date := time.Date(2009, 12, 31, 0, 0, 0, 0, time.UTC)

				arr_ind, bit_ind := date2index(start_date)

				Expect(arr_ind).To(Equal(5))
				Expect(bit_ind).To(Equal(44))
			})
		})

	})

	Context("index2date()", func() {
		When("2009 0 0", func() {
			It("returns time.Time(2009, 1, 1, 0, 0, 0, 0, time.UTC)", func() {

				start_date := time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC)

				dt := index2date(2009, 0, 0)

				Expect(dt).To(Equal(start_date))
			})
		})

		When("2009 5 44", func() {
			It("returns time.Time(2009, 12, 31, 0, 0, 0, 0, time.UTC)", func() {

				start_date := time.Date(2009, 12, 31, 0, 0, 0, 0, time.UTC)

				dt := index2date(2009, 5, 44)

				Expect(dt).To(Equal(start_date))
			})
		})
	})

	Context("init_year()", func() {
		When("init_year 2009 with US workdays", func() {
			It("2009 exists and bits match up", func() {

				bc := Create(USWorkDays, Holidays{})

				bc.init_year(2009)

				y, y_ok := bc.bmap[2009]

				Expect(y_ok).To(Equal(true))
				Expect(y.year).To(Equal(2009))
				Expect(y.diy).To(Equal(int32(365)))

				t1_str := fmt.Sprintf("%064b", y.bits[0])
				Expect(t1_str).To(Equal("1100111110011111001111100111110011111001111100111110011111001111"))

				t1_str = fmt.Sprintf("%064b", y.bits[1])
				Expect(t1_str).To(Equal("1001111100111110011111001111100111110011111001111100111110011111"))

				t1_str = fmt.Sprintf("%064b", y.bits[2])
				Expect(t1_str).To(Equal("0011111001111100111110011111001111100111110011111001111100111110"))

				t1_str = fmt.Sprintf("%064b", y.bits[3])
				Expect(t1_str).To(Equal("0111110011111001111100111110011111001111100111110011111001111100"))

				t1_str = fmt.Sprintf("%064b", y.bits[4])
				Expect(t1_str).To(Equal("1111100111110011111001111100111110011111001111100111110011111001"))

				t1_str = fmt.Sprintf("%064b", y.bits[5])
				Expect(t1_str).To(Equal("1111001111100111110011111001111100111110011110000000000000000000"))

			})
		})
	})

	Context("clear_right()", func() {
		When("1111111111111111111111111111111111111111111111111111111111111111, 32", func() {
			It("1111111111111111111111111111111100000000000000000000000000000000", func() {

				var t1 uint64

				t1 = math.MaxUint64

				clear_right(&t1, 32)

				t1_str := fmt.Sprintf("%064b", t1)

				Expect(t1_str).To(Equal("1111111111111111111111111111111100000000000000000000000000000000"))
			})
		})

		When("1111111111111111111111111111111111111111111111111111111111111111, 0", func() {
			It("1111111111111111111111111111111111111111111111111111111111111111", func() {

				var t1 uint64

				t1 = math.MaxUint64

				clear_right(&t1, 0)

				t1_str := fmt.Sprintf("%064b", t1)

				Expect(t1_str).To(Equal("1111111111111111111111111111111111111111111111111111111111111111"))
			})
		})

		When("1111111111111111111111111111111111111111111111111111111111111111, 1", func() {
			It("1111111111111111111111111111111111111111111111111111111111111110", func() {

				var t1 uint64

				t1 = math.MaxUint64

				clear_right(&t1, 1)

				t1_str := fmt.Sprintf("%064b", t1)

				Expect(t1_str).To(Equal("1111111111111111111111111111111111111111111111111111111111111110"))
			})
		})

	})

	Context("clear_left()", func() {
		When("1111111111111111111111111111111111111111111111111111111111111111, 32", func() {
			It("0000000000000000000000000000000011111111111111111111111111111111", func() {

				var t1 uint64

				t1 = math.MaxUint64

				clear_left(&t1, 32)

				t1_str := fmt.Sprintf("%064b", t1)

				Expect(t1_str).To(Equal("0000000000000000000000000000000011111111111111111111111111111111"))
			})
		})

		When("1111111111111111111111111111111111111111111111111111111111111111, 0", func() {
			It("1111111111111111111111111111111111111111111111111111111111111111", func() {

				var t1 uint64

				t1 = math.MaxUint64

				clear_left(&t1, 0)

				t1_str := fmt.Sprintf("%064b", t1)

				Expect(t1_str).To(Equal("1111111111111111111111111111111111111111111111111111111111111111"))
			})
		})

		When("1111111111111111111111111111111111111111111111111111111111111111, 1", func() {
			It("0111111111111111111111111111111111111111111111111111111111111111", func() {

				var t1 uint64

				t1 = math.MaxUint64

				clear_left(&t1, 1)

				t1_str := fmt.Sprintf("%064b", t1)

				Expect(t1_str).To(Equal("0111111111111111111111111111111111111111111111111111111111111111"))
			})
		})

	})

})
