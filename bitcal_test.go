package bitcal_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cxtuttle/go-bitcal"
)

var _ = Describe("bitcal module", func() {
	Context("bitcal.AddWorkDays()", func() {
		When("Add 2 US Work Days to 2009/1/1", func() {
			It("returns 2009/1/5", func() {
				bc := bitcal.Create(bitcal.USWorkDays, bitcal.Holidays{})

				start_date := time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC)

				nd := bc.AddWorkDays(start_date, 2, false)

				Expect(nd.Year()).To(Equal(2009))
				Expect(nd.Month()).To(Equal(time.Month(1)))
				Expect(nd.Day()).To(Equal(5))
			})
		})
	})

	Context("bitcal.AddWorkDays()", func() {
		When("Add 1 US Work Days to 2009/1/1 Inclusive", func() {
			It("returns 2009/1/2", func() {
				bc := bitcal.Create(bitcal.USWorkDays, bitcal.Holidays{})

				start_date := time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC)

				nd := bc.AddWorkDays(start_date, 1, true)

				Expect(nd.Year()).To(Equal(2009))
				Expect(nd.Month()).To(Equal(time.Month(1)))
				Expect(nd.Day()).To(Equal(2))
			})
		})
	})

	Context("bitcal.AddWorkDays()", func() {
		When("Add 1 US Work Days to 2009/1/3 Inclusive", func() {
			It("returns 2009/1/5", func() {
				bc := bitcal.Create(bitcal.USWorkDays, bitcal.Holidays{})

				start_date := time.Date(2009, 1, 3, 0, 0, 0, 0, time.UTC)

				nd := bc.AddWorkDays(start_date, 1, true)

				Expect(nd.Year()).To(Equal(2009))
				Expect(nd.Month()).To(Equal(time.Month(1)))
				Expect(nd.Day()).To(Equal(5))
			})
		})
	})

	Context("bitcal.AddWorkDays()", func() {
		When("Add 1 US Work Days to 2009/2/28 Inclusive", func() {
			It("returns 2009/3/2", func() {
				bc := bitcal.Create(bitcal.USWorkDays, bitcal.Holidays{})

				start_date := time.Date(2009, 2, 28, 0, 0, 0, 0, time.UTC)

				nd := bc.AddWorkDays(start_date, 1, true)

				Expect(nd.Year()).To(Equal(2009))
				Expect(nd.Month()).To(Equal(time.Month(3)))
				Expect(nd.Day()).To(Equal(2))
			})
		})
	})

	Context("bitcal.AddWorkDays()", func() {
		When("Add 22 US Work Days to 2009/9/1", func() {
			It("returns 2009/10/1", func() {
				bc := bitcal.Create(bitcal.USWorkDays, bitcal.Holidays{})

				start_date := time.Date(2009, 9, 1, 0, 0, 0, 0, time.UTC)

				nd := bc.AddWorkDays(start_date, 22, false)

				Expect(nd.Year()).To(Equal(2009))
				Expect(nd.Month()).To(Equal(time.Month(10)))
				Expect(nd.Day()).To(Equal(1))
			})
		})
	})

	Context("bitcal.AddWorkDays()", func() {
		When("Add 1 US Work Days to 2009/11/01", func() {
			It("returns 2010/11/03", func() {
				bc := bitcal.Create(bitcal.USWorkDays, bitcal.Holidays{})

				start_date := time.Date(2009, 11, 1, 0, 0, 0, 0, time.UTC)

				nd := bc.AddWorkDays(start_date, 1, false)

				Expect(nd.Year()).To(Equal(2009))
				Expect(nd.Month()).To(Equal(time.Month(11)))
				Expect(nd.Day()).To(Equal(3))
			})
		})
	})

	Context("bitcal.AddWorkDays()", func() {
		When("Add -1 US Work Days to 2015/11/16", func() {
			It("returns 2010/11/13", func() {
				bc := bitcal.Create(bitcal.USWorkDays, bitcal.Holidays{})

				start_date := time.Date(2015, 11, 16, 0, 0, 0, 0, time.UTC)

				nd := bc.AddWorkDays(start_date, -1, false)

				Expect(nd.Year()).To(Equal(2015))
				Expect(nd.Month()).To(Equal(time.Month(11)))
				Expect(nd.Day()).To(Equal(13))
			})
		})
	})

	Context("bitcal.AddWorkDays()", func() {
		When("Add -1 US Work Days to 2015/11/16 Inclusive", func() {
			It("returns 2010/11/13", func() {
				bc := bitcal.Create(bitcal.USWorkDays, bitcal.Holidays{})

				start_date := time.Date(2015, 11, 16, 0, 0, 0, 0, time.UTC)

				nd := bc.AddWorkDays(start_date, -1, true)

				Expect(nd.Year()).To(Equal(2015))
				Expect(nd.Month()).To(Equal(time.Month(11)))
				Expect(nd.Day()).To(Equal(13))
			})
		})
	})

	Context("bitcal.AddWorkDays()", func() {
		When("Add -1 US Work Days to 2015/11/15 Inclusive", func() {
			It("returns 2010/11/13", func() {
				bc := bitcal.Create(bitcal.USWorkDays, bitcal.Holidays{})

				start_date := time.Date(2015, 11, 15, 0, 0, 0, 0, time.UTC)

				nd := bc.AddWorkDays(start_date, -1, true)

				Expect(nd.Year()).To(Equal(2015))
				Expect(nd.Month()).To(Equal(time.Month(11)))
				Expect(nd.Day()).To(Equal(13))
			})
		})
	})

	Context("bitcal.AddWorkDays()", func() {
		When("Add -1 US Work Days to 2015/11/14 Inclusive", func() {
			It("returns 2010/11/13", func() {
				bc := bitcal.Create(bitcal.USWorkDays, bitcal.Holidays{})

				start_date := time.Date(2015, 11, 14, 0, 0, 0, 0, time.UTC)

				nd := bc.AddWorkDays(start_date, -1, true)

				Expect(nd.Year()).To(Equal(2015))
				Expect(nd.Month()).To(Equal(time.Month(11)))
				Expect(nd.Day()).To(Equal(13))
			})
		})
	})

	Context("bitcal.AddWorkDays()", func() {
		When("Add -1 US Work Days to 2015/11/14 Inclusive", func() {
			It("returns 2010/11/13", func() {
				bc := bitcal.Create(bitcal.USWorkDays, bitcal.Holidays{})

				start_date := time.Date(2015, 11, 14, 0, 0, 0, 0, time.UTC)

				nd := bc.AddWorkDays(start_date, -1, true)

				Expect(nd.Year()).To(Equal(2015))
				Expect(nd.Month()).To(Equal(time.Month(11)))
				Expect(nd.Day()).To(Equal(13))
			})
		})
	})

	Context("bitcal.AddWorkDays()", func() {
		When("Add -1 US Work Days to 2015/11/21 Inclusive", func() {
			It("returns 2010/11/20", func() {
				bc := bitcal.Create(bitcal.USWorkDays, bitcal.Holidays{})

				start_date := time.Date(2015, 11, 21, 0, 0, 0, 0, time.UTC)

				nd := bc.AddWorkDays(start_date, -1, true)

				Expect(nd.Year()).To(Equal(2015))
				Expect(nd.Month()).To(Equal(time.Month(11)))
				Expect(nd.Day()).To(Equal(20))
			})
		})
	})

	Context("bitcal.AddWorkDays()", func() {
		When("Add -1 US Work Days to 2015/11/22 Inclusive", func() {
			It("returns 2010/11/20", func() {
				bc := bitcal.Create(bitcal.USWorkDays, bitcal.Holidays{})

				start_date := time.Date(2015, 11, 22, 0, 0, 0, 0, time.UTC)

				nd := bc.AddWorkDays(start_date, -1, true)

				Expect(nd.Year()).To(Equal(2015))
				Expect(nd.Month()).To(Equal(time.Month(11)))
				Expect(nd.Day()).To(Equal(20))
			})
		})
	})


	
	Context("bitcal.AddWorkDays()", func() {
		When("Add 1 US Work Days to 2009/11/01 inclusive", func() {
			It("returns 2010/11/02", func() {
				bc := bitcal.Create(bitcal.USWorkDays, bitcal.Holidays{})

				start_date := time.Date(2009, 11, 1, 0, 0, 0, 0, time.UTC)

				nd := bc.AddWorkDays(start_date, 1, true)

				Expect(nd.Year()).To(Equal(2009))
				Expect(nd.Month()).To(Equal(time.Month(11)))
				Expect(nd.Day()).To(Equal(2))
			})
		})
	})

	Context("bitcal.AddWorkDays()", func() {
		When("Add 75 US Work Days to 2009/11/01", func() {
			It("returns 2010/02/15", func() {
				bc := bitcal.Create(bitcal.USWorkDays, bitcal.Holidays{})

				start_date := time.Date(2009, 11, 1, 0, 0, 0, 0, time.UTC)

				nd := bc.AddWorkDays(start_date, 75, false)

				Expect(nd.Year()).To(Equal(2010))
				Expect(nd.Month()).To(Equal(time.Month(2)))
				Expect(nd.Day()).To(Equal(15))
			})
		})
	})

	Context("bitcal.AddWorkDays()", func() {
		When("Add 75 US Work Days to 2009/11/01 inclusive", func() {
			It("returns 2010/02/12", func() {
				bc := bitcal.Create(bitcal.USWorkDays, bitcal.Holidays{})

				start_date := time.Date(2009, 11, 1, 0, 0, 0, 0, time.UTC)

				nd := bc.AddWorkDays(start_date, 75, true)

				Expect(nd.Year()).To(Equal(2010))
				Expect(nd.Month()).To(Equal(time.Month(2)))
				Expect(nd.Day()).To(Equal(12))
			})
		})
	})


	Context("bitcal.AddWorkDays()", func() {
		When("Add -1 US Work Days to 2009/11/06", func() {
			It("returns 2009/11/05", func() {
				bc := bitcal.Create(bitcal.USWorkDays, bitcal.Holidays{})

				start_date := time.Date(2009, 11, 6, 0, 0, 0, 0, time.UTC)

				nd := bc.AddWorkDays(start_date, -1, false)

				Expect(nd.Year()).To(Equal(2009))
				Expect(nd.Month()).To(Equal(time.Month(11)))
				Expect(nd.Day()).To(Equal(5))
			})
		})
	})

	Context("bitcal.AddWorkDays()", func() {
		When("Add -1 US Work Days to 2009/11/06 inclusive", func() {
			It("returns 2009/11/05", func() {
				bc := bitcal.Create(bitcal.USWorkDays, bitcal.Holidays{})

				start_date := time.Date(2009, 11, 6, 0, 0, 0, 0, time.UTC)

				nd := bc.AddWorkDays(start_date, -1, true)

				Expect(nd.Year()).To(Equal(2009))
				Expect(nd.Month()).To(Equal(time.Month(11)))
				Expect(nd.Day()).To(Equal(5))
			})
		})
	})

	Context("bitcal.AddWorkDays()", func() {
		When("Add -65 US Work Days to 2009/11/06", func() {
			It("returns 2009/08/07", func() {
				bc := bitcal.Create(bitcal.USWorkDays, bitcal.Holidays{})

				start_date := time.Date(2009, 11, 6, 0, 0, 0, 0, time.UTC)

				nd := bc.AddWorkDays(start_date, -65, false)

				Expect(nd.Year()).To(Equal(2009))
				Expect(nd.Month()).To(Equal(time.Month(8)))
				Expect(nd.Day()).To(Equal(7))
			})
		})
	})

	Context("bitcal.AddWorkDays()", func() {
		When("Add -1 US Work Days to 2009/01/04", func() {
			It("returns 2009/01/01", func() {
				bc := bitcal.Create(bitcal.USWorkDays, bitcal.Holidays{})

				start_date := time.Date(2009, 1, 4, 0, 0, 0, 0, time.UTC)

				nd := bc.AddWorkDays(start_date, -1, false)

				Expect(nd.Year()).To(Equal(2009))
				Expect(nd.Month()).To(Equal(time.Month(1)))
				Expect(nd.Day()).To(Equal(1))
			})
		})
	})

	Context("bitcal.AddWorkDays()", func() {
		When("Add -1 US Work Days to 2009/01/04 inclusive", func() {
			It("returns 2009/01/02", func() {
				bc := bitcal.Create(bitcal.USWorkDays, bitcal.Holidays{})

				start_date := time.Date(2009, 1, 4, 0, 0, 0, 0, time.UTC)

				nd := bc.AddWorkDays(start_date, -1, true)

				Expect(nd.Year()).To(Equal(2009))
				Expect(nd.Month()).To(Equal(time.Month(1)))
				Expect(nd.Day()).To(Equal(2))
			})
		})
	})

	Context("bitcal.AddWorkDays()", func() {
		When("Add -55 US Work Days to 2009/01/04", func() {
			It("returns 2008/10/17", func() {
				bc := bitcal.Create(bitcal.USWorkDays, bitcal.Holidays{})

				start_date := time.Date(2009, 1, 4, 0, 0, 0, 0, time.UTC)

				nd := bc.AddWorkDays(start_date, -55, false)

				Expect(nd.Year()).To(Equal(2008))
				Expect(nd.Month()).To(Equal(time.Month(10)))
				Expect(nd.Day()).To(Equal(17))
			})
		})
	})

	Context("bitcal.AddWorkDays()", func() {
		When("Add -55 US Work Days to 2009/01/04 inclusive", func() {
			It("returns 2008/10/20", func() {
				bc := bitcal.Create(bitcal.USWorkDays, bitcal.Holidays{})

				start_date := time.Date(2009, 1, 4, 0, 0, 0, 0, time.UTC)

				nd := bc.AddWorkDays(start_date, -55, true)

				Expect(nd.Year()).To(Equal(2008))
				Expect(nd.Month()).To(Equal(time.Month(10)))
				Expect(nd.Day()).To(Equal(20))
			})
		})
	})

	Context("bitcal.AddWorkDays()", func() {
		When("Add 255 US Work Days to 2020/02/04", func() {
			It("returns 2021/01/26", func() {
				bc := bitcal.Create(bitcal.USWorkDays, bitcal.Holidays{})

				start_date := time.Date(2020, 2, 4, 0, 0, 0, 0, time.UTC)

				nd := bc.AddWorkDays(start_date, 255, false)

				Expect(nd.Year()).To(Equal(2021))
				Expect(nd.Month()).To(Equal(time.Month(1)))
				Expect(nd.Day()).To(Equal(26))
			})
		})
	})

	Context("bitcal.DeltaWorkDays()", func() {
		When("2009/01/04 - 2009/01/13", func() {
			It("10", func() {
				bc := bitcal.Create(bitcal.USWorkDays, bitcal.Holidays{})

				start_date := time.Date(2009, 1, 4, 0, 0, 0, 0, time.UTC)
				end_date := time.Date(2009, 1, 13, 0, 0, 0, 0, time.UTC)

				delta := bc.DeltaWorkDays(start_date, end_date)

				Expect(delta).To(Equal(6))
			})
		})
	})

	// Cross year
	Context("bitcal.DeltaWorkDays()", func() {
		When("2009/8/04 - 2010/01/13", func() {
			It("10", func() {
				bc := bitcal.Create(bitcal.USWorkDays, bitcal.Holidays{})

				start_date := time.Date(2009, 8, 4, 0, 0, 0, 0, time.UTC)
				end_date := time.Date(2010, 1, 13, 0, 0, 0, 0, time.UTC)

				delta := bc.DeltaWorkDays(start_date, end_date)

				Expect(delta).To(Equal(115))
			})
		})
	})

	// Negative
	Context("bitcal.DeltaWorkDays()", func() {
		When("2009/8/04 - 2008/01/13", func() {
			It("10", func() {
				bc := bitcal.Create(bitcal.USWorkDays, bitcal.Holidays{})

				start_date := time.Date(2009, 8, 4, 0, 0, 0, 0, time.UTC)
				end_date := time.Date(2008, 1, 13, 0, 0, 0, 0, time.UTC)

				delta := bc.DeltaWorkDays(start_date, end_date)

				Expect(delta).To(Equal(-406))
			})
		})
	})

	// Leap year
	Context("bitcal.DeltaWorkDays()", func() {
		When("2019/8/04 - 2021/01/13", func() {
			It("10", func() {
				bc := bitcal.Create(bitcal.USWorkDays, bitcal.Holidays{})

				start_date := time.Date(2019, 8, 4, 0, 0, 0, 0, time.UTC)
				end_date := time.Date(2021, 1, 13, 0, 0, 0, 0, time.UTC)

				delta := bc.DeltaWorkDays(start_date, end_date)

				Expect(delta).To(Equal(377))
			})
		})
	})

})
