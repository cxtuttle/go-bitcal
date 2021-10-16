//  Package bitcal provides functions for doing date math with business calendars.  It uses bit vectors to track working days from non-working days.
//
package bitcal

import (
//	"log"
	"math"
	"math/bits"
	"time"
)

type bityear struct {
	year int
	diy  int32     // days in year
	bits [6]uint64 // Enough bits for 366 days
}

type bitcal_map map[int]bityear

type BitCal struct {
	wd   Workdays
	hd   Holidays
	bmap bitcal_map
}

// Workdays type is used to assign whether a given day is a workday. Assign true to a day to indicate it is a workday.
type Workdays struct {
	Mon bool
	Tue bool
	Wed bool
	Thu bool
	Fri bool
	Sat bool
	Sun bool
}

// TODO: The Holidays type will eventually add holiday support. The dates here will be omitted from a work calendar.  Expect type to change.
type Holidays struct {
	dates []time.Time
}

// Predefined Workdays type for the US work week. Mon-Fri are marked true.
var USWorkDays = Workdays{true, true, true, true, true, false, false}

// Create takes a Workday and Holidays struct and returns an initialized struct in return that can be used for calls against that Calendar.
func Create(wd Workdays, hd Holidays) (c *BitCal) {

	c = &BitCal{}

	c.wd = wd
	c.hd = hd

	c.bmap = make(map[int]bityear)

	return
}

func convertWeekday(wd time.Weekday) (d uint64) {

	if wd == 0 { // Sunday
		d = 6
	} else {
		d = uint64(wd - 1)
	}

	return
}

func isLeapYear(year int) bool {
	return year%400 == 0 || year%4 == 0 && year%100 != 0
}

func (c *BitCal) get_wd_bits() (wd_bits uint64) {

	if c.wd.Mon {
		wd_bits |= (1 << 7)
	}

	if c.wd.Tue {
		wd_bits |= (1 << 6)
	}

	if c.wd.Wed {
		wd_bits |= (1 << 5)
	}

	if c.wd.Thu {
		wd_bits |= (1 << 4)
	}

	if c.wd.Fri {
		wd_bits |= (1 << 3)
	}

	if c.wd.Sat {
		wd_bits |= (1 << 2)
	}

	if c.wd.Sun {
		wd_bits |= (1 << 1)
	}

	/// 1st bit is 0

	return
}

func (c *BitCal) get_nonwd_bits() (wd_bits uint64) {

	if !c.wd.Mon {
		wd_bits |= (1 << 7)
	}

	if !c.wd.Tue {
		wd_bits |= (1 << 6)
	}

	if !c.wd.Wed {
		wd_bits |= (1 << 5)
	}

	if !c.wd.Thu {
		wd_bits |= (1 << 4)
	}

	if !c.wd.Fri {
		wd_bits |= (1 << 3)
	}

	if !c.wd.Sat {
		wd_bits |= (1 << 2)
	}

	if !c.wd.Sun {
		wd_bits |= (1 << 1)
	}

	/// 1st bit is 0

	return
}

func (c *BitCal) init_year(year int) {

	_, y_ok := c.bmap[year]

	if y_ok {
		// already cached ... do nothing
	} else {
		var by bityear

		by.year = year

		if isLeapYear(year) {
			by.diy = 366
		} else {
			by.diy = 365
		}

		for i := 0; i < 6; i++ {
			by.bits[i] = math.MaxUint64 //  Set all bits to 1
		}

		nwd_bits := c.get_nonwd_bits()

		d1 := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

		wd := convertWeekday(d1.Weekday())

		//  Now clear all non work days

		for i := 0; i < 6; i++ {
			if i > 0 {
				wd++ //  Increment starting weekday each time

				if wd > 6 {
					wd = 0
				}
			}

			by.bits[i] &^= (nwd_bits << (56 + wd))
			by.bits[i] &^= (nwd_bits << (49 + wd))
			by.bits[i] &^= (nwd_bits << (42 + wd))
			by.bits[i] &^= (nwd_bits << (35 + wd))
			by.bits[i] &^= (nwd_bits << (28 + wd))
			by.bits[i] &^= (nwd_bits << (21 + wd))
			by.bits[i] &^= (nwd_bits << (14 + wd))
			by.bits[i] &^= (nwd_bits << (7 + wd))
			by.bits[i] &^= (nwd_bits << (0 + wd))
			by.bits[i] &^= (nwd_bits >> (7 - wd))
		}

		// Remove bits after year end
		if isLeapYear(year) {
			clear_right(&by.bits[5], (64 - 46))
		} else {
			clear_right(&by.bits[5], (64 - 45))
		}

		// Now remove holidays  --  TO DO

		// Set value
		c.bmap[year] = by

	}
}

func date2index(date time.Time) (arr_ind int, bit_ind int) {

	yd := date.YearDay()

	yd--

	arr_ind = int(yd / 64)
	bit_ind = yd - (arr_ind * 64)

	return
}

func index2date(year int, arr_ind int, bit_ind int) (ret_date time.Time) {

	ret_date = time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	add_days := (64 * arr_ind) + bit_ind

	ret_date = ret_date.AddDate(0, 0, add_days)

	return
}

func (c *BitCal) index2work(year int, arr_ind int, bit_ind int) (work bool) {

	work = (c.bmap[year].bits[arr_ind] & (1 << bit_ind)) > 0

	return
}

func clear_left(workb *uint64, bit_ind int) {
	*workb &= (math.MaxUint64 >> bit_ind)
}

func clear_right(workb *uint64, bit_ind int) {
	var bm uint64

	bm = (math.MaxUint64 >> (64 - bit_ind))
	bm = ^bm

	*workb &= bm
}

func (c *BitCal) get_forward_bits(year int, arr_ind int, bit_ind int, include bool) (ret_bits uint64, ret_year int, ret_arr_ind int, ret_bit_ind int) {

	if bit_ind > 62 || (arr_ind == 5 && ((isLeapYear(year) && bit_ind > 45) || (!isLeapYear(year) && bit_ind > 44))) {
		if arr_ind > 4 {
			ret_bit_ind = 63
			ret_year = year + 1
			ret_arr_ind = 0

			c.init_year(ret_year)
		} else {
			ret_bit_ind = 63
			ret_year = year
			ret_arr_ind = arr_ind + 1
		}

		ret_bits = c.bmap[ret_year].bits[ret_arr_ind]
	} else {
		ret_bit_ind = 63
		ret_year = year
		ret_arr_ind = arr_ind

		_, y_ok := c.bmap[ret_year]

		if !y_ok {
			c.init_year(ret_year)
		}

		ret_bits = c.bmap[ret_year].bits[ret_arr_ind]

		clear_left(&ret_bits, bit_ind) // Clear everything to the left of the bit we're starting at

		lz := bits.LeadingZeros64(ret_bits)

//		log.Printf("FB YEAR: %d ARR_IND: %d BIT_INT: %d  LZ: %d",year, arr_ind, bit_ind, lz)

		if include {
			if lz > bit_ind {
//				log.Printf("FB INC1 YEAR: %d ARR_IND: %d BIT_INT: %d  LZ: %d",year, arr_ind, bit_ind, lz)
				clear_left(&ret_bits, lz) // Date on non workday so clear next workday
			} else {
//				log.Printf("FB INC2 YEAR: %d ARR_IND: %d BIT_INT: %d  LZ: %d",year, arr_ind, bit_ind, lz)
				clear_left(&ret_bits, bit_ind+1) // clear the workday
			}
		} else {
			if lz > bit_ind {
//				log.Printf("FB NOINC1 YEAR: %d ARR_IND: %d BIT_INT: %d  LZ: %d",year, arr_ind, bit_ind, lz)
				clear_left(&ret_bits, lz+1) // Date on non workday so clear next workday
			} else {
//				log.Printf("FB NOINC2 YEAR: %d ARR_IND: %d BIT_INT: %d  LZ: %d",year, arr_ind, bit_ind, lz)
				clear_left(&ret_bits, bit_ind+1) // clear the workday
			}
		}
	}

	return
}

func (c *BitCal) get_backword_bits(year int, arr_ind int, bit_ind int, include bool) (ret_bits uint64, ret_year int, ret_arr_ind int, ret_bit_ind int) {

//	log.Printf("YEAR: %d ARR_IND: %d BIT_INT: %d",year, arr_ind, bit_ind)

	if bit_ind < 1 {
		if arr_ind < 1 {
			ret_bit_ind = 0
			ret_year = year - 1
			ret_arr_ind = 5

			c.init_year(ret_year)
		} else {
			ret_bit_ind = 0
			ret_year = year
			ret_arr_ind = arr_ind - 1
		}

		ret_bits = c.bmap[ret_year].bits[ret_arr_ind]
	} else {
		ret_bit_ind = 0
		ret_year = year
		ret_arr_ind = arr_ind

		_, y_ok := c.bmap[ret_year]

		if !y_ok {
			c.init_year(ret_year)
		}

		ret_bits = c.bmap[ret_year].bits[ret_arr_ind]

		clear_right(&ret_bits, (64 - (bit_ind + 1))) // Clear everything to the right of the bit we're starting at

		tz := bits.TrailingZeros64(ret_bits)

//		log.Printf("BB YEAR: %d ARR_IND: %d BIT_INT: %d  TZ: %d",year, arr_ind, bit_ind, tz)

		if include {
			if tz > bit_ind {
//				log.Printf("BB INC1 YEAR: %d ARR_IND: %d BIT_INT: %d  TZ: %d",year, arr_ind, bit_ind, tz)
				clear_right(&ret_bits, tz) // Date on non workday so clear next workday
			} else {
//				log.Printf("BB INC2 YEAR: %d ARR_IND: %d BIT_INT: %d  TZ: %d",year, arr_ind, bit_ind, tz)
				clear_right(&ret_bits, 64-bit_ind) // clear the workday
			}
		} else {
			if tz > bit_ind {
//				log.Printf("BB NOINC1 YEAR: %d ARR_IND: %d BIT_INT: %d  TZ: %d",year, arr_ind, bit_ind, tz)
				clear_right(&ret_bits, tz+1) // Date on non workday so clear next workday
			} else {
//				log.Printf("BB NOINC2 YEAR: %d ARR_IND: %d BIT_INT: %d  TZ: %d",year, arr_ind, bit_ind, tz)
				clear_right(&ret_bits, 64-bit_ind) // clear the workday
			}
		}
	}

	return
}

func clear_right_still_greater(workb *uint64, days int, count int, include bool) {

	trim_ratio := []int{2, 4, 8}

	for _, ratio := range trim_ratio {
		tw := *workb
		trail := bits.TrailingZeros64(*workb)

		bits_to_clear := int((64 - trail) / ratio)

		if bits_to_clear < 2 {
			clear_right(&tw, trail+1)
		} else {
			clear_right(&tw, int((64-trail)/ratio)+trail)
		}

		tw_cnt := bits.OnesCount64(tw)

		if count+tw_cnt >= days {
			*workb = tw
			return
		}
	}

	// Still here just clear 1 bit
	trail := bits.TrailingZeros64(*workb)
	clear_right(workb, trail+1)
}

func clear_left_still_greater(workb *uint64, days int, count int, include bool) {

	trim_ratio := []int{2, 4, 8}

	for _, ratio := range trim_ratio {
		tw := *workb
		lead := bits.LeadingZeros64(*workb)

		bits_to_clear := int((64 - lead) / ratio)

		if bits_to_clear < 2 {
			clear_left(&tw, lead+1)
		} else {
			clear_left(&tw, int((64-lead)/ratio)+lead)
		}

		tw_cnt := bits.OnesCount64(tw)

		if count+tw_cnt >= -days {
			*workb = tw
			return
		}
	}

	// Still here just clear 1 bit
	lead := bits.LeadingZeros64(*workb)
	clear_left(workb, lead+1)
}

// AddWorkDays take a starting time.Time and the number of work days( positive or negative ) to add to the date. When include is true,  the current date is included and if it is a non work-day it will the not skip the forwards/backwords work-day.  A time.Time is returned.
func (c *BitCal) AddWorkDays(date time.Time, days int, include bool) (ret_date time.Time) {

	var (
		day_found bool
		count     int
	)

	if days == 0 {
		ret_date = date
		return
	}

	c.init_year(date.Year())

	arr_ind, bit_ind := date2index(date)

//	log.Printf("Year: %d ARR_IND: %d BIT_IND: %d",date.Year(), arr_ind, bit_ind)

	if days > 0 {
		workb, cyear, carr_ind, cbit_ind := c.get_forward_bits(date.Year(), arr_ind, bit_ind, include)

		for !day_found {
			cc := bits.OnesCount64(workb)

			prev_count := count
			count += cc

			if count >= days {
				//  Day in this segment
				if count == days {
					// Get the last work day
					trail := bits.TrailingZeros64(workb)

					// log.Printf("Year: %d Carr_ind: %d Cbit_ind: %d Trail: %d", cyear, carr_ind, cbit_ind, trail)

					if trail > 0 {
						cbit_ind -= trail
					}

					ret_date = index2date(cyear, carr_ind, cbit_ind)
					day_found = true

				} else {
					count = prev_count
					clear_right_still_greater(&workb, days, count, include)
				}
			} else {
				workb, cyear, carr_ind, cbit_ind = c.get_forward_bits(cyear, carr_ind, cbit_ind, include)
			}
		}
	} else {
		workb, cyear, carr_ind, cbit_ind := c.get_backword_bits(date.Year(), arr_ind, bit_ind, include)

		for !day_found {
//			log.Printf("Year: %d Carr_ind: %d Cbit_ind: %d", cyear, carr_ind, cbit_ind)

			cc := bits.OnesCount64(workb)

			prev_count := count
			count += cc

			if count >= -days {
				if count == -days {
					// Get the last work day
					lead := bits.LeadingZeros64(workb)

//					log.Printf("Year: %d Carr_ind: %d Cbit_ind: %d Lead: %d", cyear, carr_ind, cbit_ind, lead)

					if lead > 0 {
						cbit_ind += lead
					}

					ret_date = index2date(cyear, carr_ind, cbit_ind)
					day_found = true

				} else {
					count = prev_count
					clear_left_still_greater(&workb, days, count, include)
				}
			} else {
				workb, cyear, carr_ind, cbit_ind = c.get_backword_bits(cyear, carr_ind, cbit_ind, include)
			}
		}
	}

	return
}

// IsWorkDay returns true if the date given is a work day
func (c *BitCal) IsWorkDay(date time.Time) (ret bool) {

	c.init_year(date.Year())

	arr_ind, bit_ind := date2index(date)

	return c.index2work(date.Year(), arr_ind, bit_ind)
}


// WorkDayPrevIncl take a starting time.Time and returns the Previous Workday Including the date given.  A time.Time is returned.
func (c *BitCal) WorkDayPrevIncl(date time.Time) (ret_date time.Time) {

	if c.IsWorkDay(date) {
		ret_date = date
	} else {
		ret_date = c.AddWorkDays(date, -1, true)
	}

	return
}

// WorkDayNextIncl take a starting time.Time and returns the Next Workday Including the date given.  A time.Time is returned.
func (c *BitCal) WorkDayNextIncl(date time.Time) (ret_date time.Time) {

	if c.IsWorkDay(date) {
		ret_date = date
	} else {
		ret_date = c.AddWorkDays(date, 1, true)
	}

	return
}

// DeltaWorkDays takes two time.Time's and returns the number of work days between them.
func (c *BitCal) DeltaWorkDays(date1 time.Time, date2 time.Time) (days int) {

	var (
		neg_sign bool
		ldt      time.Time
		gdt      time.Time
		lyr      int
		gyr      int
	)

	if date1.After(date2) {
		neg_sign = true
		ldt = date2
		gdt = date1
	} else {
		neg_sign = false
		ldt = date1
		gdt = date2
	}

	lyr = ldt.Year()
	gyr = gdt.Year()

	for y := lyr; y <= gyr; y++ {
		c.init_year(y)
	}

	arr_ind1, bit_ind1 := date2index(ldt)
	arr_ind2, bit_ind2 := date2index(gdt)

	for y := lyr; y <= gyr; y++ {
		for ar := 0; ar < 6; ar++ {
			if y == gyr && ar > arr_ind2 {
				break
			}

			if y == lyr && ar == 0 {
				ar = arr_ind1
			}

			bt := c.bmap[y].bits[ar]

			if y == lyr && ar == arr_ind1 {
				clear_left(&bt, bit_ind1+1)
			}

			if y == gyr && ar == arr_ind2 {
				clear_right(&bt, 64-bit_ind2)
			}

			days += bits.OnesCount64(bt)
		}
	}

	if neg_sign {
		days = -days
	}

	return
}
