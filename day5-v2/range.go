package main

import (
	"fmt"
)

// Range represents a range of integers
// Start is the first integer in the range
// Length is the number of integers in the range
// Example: Range{Start: 10, Length: 5} represents the range [10, 11, 12, 13, 14]
type Range struct {
	Start  int
	Length int
}

func (r *Range) String() string {
	return fmt.Sprintf("Range{Start: %d, Length: %d}", r.Start, r.Length)
}

func (r *Range) End() int {

	return r.Start + r.Length - 1
}

func (r *Range) Contains(num int) bool {
	return num >= r.Start && num <= r.End()
}

func (r *Range) ContainsRange(other *Range) bool {
	return r.Contains(other.Start) && r.Contains(other.End())
}

func (r Range) Equal(other Range) bool {
	isItEqual := false
	for i := 0; i < r.Length; i++ {
		if r.Start+i == other.Start+i {
			isItEqual = true
		} else {
			isItEqual = false
			break
		}
	}
	return isItEqual && r.Length == other.Length && r.Start == other.Start && r.End() == other.End()
}

func (this *Range) GetOverlaps(other *Range) (overlap *Range, nonoverlap []*Range, err error) {
	// TODO: change with slog
	// log.Printf("this: %s, other: %s\n", this, other)
	switch {
	//           End
	// Start     |
	// |         |
	// v  (this) V
	// |---------|    (other)
	//              |---------|
	case this.End() < other.Start:
		// TODO: change with slog
		// log.Println("Case 1")
		return nil, []*Range{{this.Start, this.Length}}, nil

		//   (this)
		// |--------|  (other)
		//          |----------|
	case this.Start < other.Start && this.End() == other.Start:
		// TODO: change with slog
		// log.Println("Case 2")
		return &Range{
				this.End(),
				1,
			}, []*Range{
				{this.Start, (other.Start) - this.Start},
			}, nil

	//      (this)
	//   |----------|  (other)
	//          |--------------|
	case this.Start < other.Start && this.Start+this.Length > other.Start && this.Start+this.Length < other.Start+other.Length:
		// TODO: change with slog
		// log.Println("Case 3")
		return &Range{
				other.Start,
				(this.Start + this.Length) - other.Start,
			}, []*Range{
				{this.Start, (other.Start) - this.Start},
			}, nil

	//        (this)
	//    |---------------|
	//          |---------|
	//            (other)
	case this.Start < other.Start && this.Start+this.Length > other.Start && this.Start+this.Length == other.Start+other.Length:
		// TODO: change with slog
		// log.Println("Case 4")
		return &Range{
				other.Start,
				other.Length,
			}, []*Range{
				{this.Start, (other.Start) - this.Start},
			}, nil

	//              (this)
	//        |--------------|
	//  |-------------------------|
	//       (other)
	case this.Start > other.Start && this.Start+this.Length < other.Start+other.Length:
		// TODO: change with slog
		// log.Println("Case 5")
		return &Range{
				this.Start,
				this.Length,
			},
			[]*Range{},
			nil

	//           (this)
	//   |---------------------|
	//         |---------|
	//			  (other)
	//
	case this.Start < other.Start && this.Start+this.Length > other.Start+other.Length:
		// TODO: change with slog
		// log.Println("Case 6")
		return &Range{
				other.Start,
				other.Length,
			}, []*Range{
				{this.Start, (other.Start) - this.Start},
				{
					other.End() + 1,
					this.End() - other.End(),
				},
			}, nil

	//                    (this)
	//               |-------------|
	//   |-------------------------|
	//        (other)
	case this.Start > other.Start && this.Start+this.Length == other.Start+other.Length:
		// TODO: change with slog
		// log.Println("Case 7")
		return &Range{
			this.Start,
			this.Length,
		}, []*Range{}, nil

	//                          (this)
	//                 |---------------------|
	//  |---------------------|Ʌ             Ʌ
	//     (other)             |             |
	//                         |             |
	//                         |             |
	//                         (nonoverlap #1)
	case this.Start > other.Start && this.Start < other.Start+other.Length && this.Start+this.Length > other.Start+other.Length:
		// TODO: change with slog
		// log.Println("Case 8")
		return &Range{
				this.Start,
				(other.End()) - this.Start + 1, // +1 because we want to include the last element
			}, []*Range{
				{
					(other.End()) + 1,
					this.End() - (other.End()),
				},
			}, nil

	//
	//                     (this)
	//                |-------------|
	//  |-------------|
	//      (other)
	//
	case this.Start > other.Start && this.Start == other.Start+other.Length && this.Start+this.Length > other.Start+other.Length:
		// TODO: change with slog
		// log.Println("Case 9")
		return &Range{
				this.Start,
				1,
			}, []*Range{
				{this.Start + 1, this.End()},
			}, nil

	//
	//						    (this)
	//                     |-------------|
	//    |------------|
	//        (other)
	case other.Start < this.Start && this.Start > other.Start+other.Length && this.Start+this.Length > other.Start+other.Length:
		// TODO: change with slog
		// log.Println("Case 10")
		return nil, []*Range{{this.Start, this.Length}}, nil

	//        (this)
	//  |----------------|
	//  |--------------------------|
	//             (other)
	case this.Start == other.Start && this.End() < other.End():
		return &Range{
			this.Start,
			this.Length,
		}, []*Range{}, nil

	//            (this)
	// |---------------------------|
	// |----------------|
	//      (other)
	case this.Start == other.Start && this.End() > other.End():
		return &Range{
				this.Start,
				other.Length,
			}, []*Range{
				{
					other.End() + 1,
					this.End() - other.End(),
				},
			}, nil

	//            (this)
	// |---------------------------|
	// |---------------------------|
	//             (other)
	case this.Start == other.Start && this.End() == other.End():
		return &Range{
			this.Start,
			this.Length,
		}, []*Range{}, nil

	default:
		return nil, []*Range{}, fmt.Errorf(
			"%s\n",
			"Error while processing overlaps with given ranges, unknown overlap and nonoverlap",
		)
	}
}

type SingleRangeMap struct {
	SourceRange      *Range
	DestinationRange *Range
}

func (rm *SingleRangeMap) String() string {
	return fmt.Sprintf(
		"SourceRange: %s, DestinationRange: %s",
		rm.SourceRange,
		rm.DestinationRange,
	)
}

func (rm *SingleRangeMap) Map(r *Range) (*Range, error) {
	// if r is not in the source range, return nil
	if !rm.SourceRange.ContainsRange(r) {
		return nil, fmt.Errorf("Error: %s is not in the source range %s", r, rm.SourceRange)
	}

	srcStartDiff := r.Start - rm.SourceRange.Start
	destStart := rm.DestinationRange.Start + srcStartDiff

	// if r is in the source range, return the corresponding range in the destination range
	return &Range{
		destStart,
		r.Length,
	}, nil
}
