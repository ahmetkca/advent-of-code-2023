package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetOverlaps1(t *testing.T) {
	r1 := &Range{
		10,
		7,
	}

	r2 := &Range{
		19,
		3,
	}

	overlap, nonoverlap, err := r1.GetOverlaps(r2)
	if err != nil {
		t.Fatalf("Error while processing overlaps: %s\n", err)
	}

	if overlap != nil {
		t.Errorf("r1.GetOverlaps(%s), got %s, expected %v\n", r2, overlap, nil)
	}

	if len(nonoverlap) != 1 && nonoverlap[0].Start != 10 && nonoverlap[0].Length != 7 {
		t.Errorf("r1.GetOverlaps(%s), got %s, expected %s\n", r2, nonoverlap, r1)
	}
}

// 11    (this)    15
// |---|---|---|---|  (other)
//
//	12  13  14  |----------|
//
// Ʌ           Ʌ   15         21
// |	       |
// ( nonoverlap)
func TestGetOverlaps2(t *testing.T) {
	r1 := &Range{
		11,
		5,
	}

	r2 := &Range{
		15,
		7,
	}

	overlap, nonoverlap, err := r1.GetOverlaps(r2)
	if err != nil {
		t.Fatalf("Error while processing overlaps: %s\n", err)
	}

	expectedOverlap := &Range{
		15,
		1,
	}

	if overlap != nil && !cmp.Equal(overlap, expectedOverlap) {
		t.Fatalf("r1.GetOverlaps(%s): got %s, expected %s\n", r2, overlap, expectedOverlap)
	}

	if len(nonoverlap) != 1 {
		t.Fatalf("Expected 1 nonoverlap ranges, got %d\n", len(nonoverlap))
	}

	nonoverlap1 := nonoverlap[0]

	expectedNonoverlap1 := &Range{
		11,
		4,
	}

	if !cmp.Equal(nonoverlap1, expectedNonoverlap1) {
		t.Fatalf(
			"Expected first nonoverlap range %s, got %s\n",
			expectedNonoverlap1,
			nonoverlap1,
		)
	}
}

// 9   (r1)   21
// |----------|
//
//	|---------------|
//	13     (r2)     29
func TestGetOverlaps3(t *testing.T) {
	r1 := &Range{
		9,
		13,
	}

	r2 := &Range{
		13,
		17,
	}

	expectedOverlap := &Range{
		13,
		9,
	}

	expectedNonoverlap := []*Range{
		{
			9,
			4,
		},
	}

	overlap, nonoverlap, err := r1.GetOverlaps(r2)
	if err != nil {
		t.Fatalf("Error while processing overlaps: %s\n", err)
	}

	if overlap != nil && !cmp.Equal(overlap, expectedOverlap) {
		t.Errorf("r1.GetOverlaps(%s): got %s, expected %s\n", r2, overlap, expectedOverlap)
	}

	if len(nonoverlap) != 1 {
		t.Errorf("Expected 1 nonoverlap ranges, got %d\n", len(nonoverlap))
	}

	if !cmp.Equal(nonoverlap[0], expectedNonoverlap[0]) {
		t.Errorf(
			"Expected first nonoverlap range %s, got %s\n",
			expectedNonoverlap[0],
			nonoverlap[0],
		)
	}
}

//	(this)
//
// |---------------|
//
//	|---------|
//	  (other)
func TestGetOverlaps4(t *testing.T) {
	r1 := &Range{
		9,
		7,
	}

	r2 := &Range{
		11,
		5,
	}

	expectedOverlap := &Range{
		11,
		5,
	}

	expectedNonoverlap := []*Range{
		{
			9,
			2,
		},
	}

	overlap, nonoverlap, err := r1.GetOverlaps(r2)
	if err != nil {
		t.Fatalf("Error while processing overlaps: %s\n", err)
	}

	if overlap != nil && !cmp.Equal(overlap, expectedOverlap) {
		t.Errorf("r1.GetOverlaps(%s): got %s, expected %s\n", r2, overlap, expectedOverlap)
	}

	if len(nonoverlap) != 1 {
		t.Errorf("Expected 1 nonoverlap ranges, got %d\n", len(nonoverlap))
	}

	if !cmp.Equal(nonoverlap[0], expectedNonoverlap[0]) {
		t.Errorf(
			"Expected first nonoverlap range %s, got %s\n",
			expectedNonoverlap[0],
			nonoverlap[0],
		)
	}
}

//	      (this)
//	|--------------|
//
// |-------------------------|
//
//	(other)
func TestGetOverlaps5(t *testing.T) {
	//        1   2   3   4   5   6   7
	// r1 = [11, 12, 13, 14, 15, 16, 17]
	r1 := &Range{
		11,
		7,
	}

	//       1  2  3   4   5   6   7   8   9  10  11  12  13
	// r2 = [7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19]
	r2 := &Range{
		7,
		13,
	}

	// expectedOverlap = [11, 12, 13, 14, 15, 16, 17]
	expectedOverlap := &Range{
		11,
		7,
	}

	overlap, nonoverlap, err := r1.GetOverlaps(r2)
	if err != nil {
		t.Fatalf("Error while processing overlaps: %s\n", err)
	}

	if overlap != nil && !cmp.Equal(overlap, expectedOverlap) {
		t.Errorf("r1.GetOverlaps(%s): got %s, expected %s\n", r2, overlap, expectedOverlap)
	}

	if len(nonoverlap) != 0 {
		t.Errorf("Expected 0 nonoverlap ranges, got %d\n", len(nonoverlap))
	}
}

//	  3       (this)        11
//	  |---------------------|
//	        |---------|
//			   5 (other) 6
func TestGetOverlaps6(t *testing.T) {
	// r1 = [3, 4, 5, 6, 7, 8, 9, 10, 11]
	r1 := &Range{
		3,
		9,
	}

	// r2 = [5, 6]
	r2 := &Range{
		5,
		2,
	}

	expectedOverlap := &Range{
		5,
		2,
	}

	expectedNonoverlap := []*Range{
		{
			3,
			2,
		},
		{
			7,
			5,
		},
	}

	overlap, nonoverlap, err := r1.GetOverlaps(r2)
	if err != nil {
		t.Errorf("Error while processing overlaps: %s\n", err)
	}

	if overlap != nil && !cmp.Equal(overlap, expectedOverlap) {
		t.Errorf("r1.GetOverlaps(%s): got %s, expected %s\n", r2, overlap, expectedOverlap)
	}

	if len(nonoverlap) != 2 {
		t.Errorf("Expected 2 nonoverlap ranges, got %d\n", len(nonoverlap))
	}

	if !cmp.Equal(nonoverlap[0], expectedNonoverlap[0]) {
		t.Errorf(
			"Expected first nonoverlap range %s, got %s\n",
			expectedNonoverlap[0],
			nonoverlap[0],
		)
	}

	if !cmp.Equal(nonoverlap[1], expectedNonoverlap[1]) {
		t.Errorf(
			"Expected second nonoverlap range %s, got %s\n",
			expectedNonoverlap[1],
			nonoverlap[1],
		)
	}
}

//	      (r1)
//	|-------------|
//
// |-------------------------|
//
//	(r2)
func TestGetOverlaps7(t *testing.T) {
	//        1   2   3   4   5
	// r1 = [11, 12, 13, 14, 15]
	r1 := &Range{
		11,
		5,
	}

	//		 1  2  3  4  5  6  7   8   9  10  11  12  13
	// r2 = [3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15]
	r2 := &Range{
		3,
		13,
	}

	expectedOverlap := &Range{
		11,
		5,
	}

	overlap, nonoverlap, err := r1.GetOverlaps(r2)
	if err != nil {
		t.Errorf("Error while processing overlaps: %s\n", err)
	}

	if overlap != nil && !cmp.Equal(overlap, expectedOverlap) {
		t.Errorf("r1.GetOverlaps(%s): got %s, expected %s\n", r2, overlap, expectedOverlap)
	}

	if len(nonoverlap) != 0 {
		t.Errorf("Expected 0 nonoverlap ranges, got %d\n", len(nonoverlap))
	}
}

//	         (this)
//	|---------------------|
//
// |---------------------|Ʌ             Ʌ
//
//	        (other         |             |
//		                      |             |
//		                      |             |
//		                      (nonoverlap #1)
func TestGetOverlaps8(t *testing.T) {
	// 	      1   2   3   4   5   6   7
	// r1 = [21, 22, 23, 24, 25, 26, 27]
	r1 := &Range{21, 7}

	// 		 1    2   3   4   5   6   7   8   9
	// r2 = [17, 18, 19, 20, 21, 22, 23, 24, 25]
	r2 := &Range{17, 9}

	// expectedOverlap = [21, 22, 23, 24, 25]
	expectedOverlap := &Range{21, 5}

	expectedNonoverlap := []*Range{
		{26, 2},
	}

	overlap, nonoverlap, err := r1.GetOverlaps(r2)
	if err != nil {
		t.Errorf("Error while processing overlaps: %s\n", err)
	}

	if overlap != nil && !cmp.Equal(overlap, expectedOverlap) {
		t.Errorf("r1.GetOverlaps(%s): got %s, expected %s\n", r2, overlap, expectedOverlap)
	}

	if len(nonoverlap) != 1 {
		t.Errorf("Expected 1 nonoverlap ranges, got %d\n", len(nonoverlap))
	}

	if !cmp.Equal(nonoverlap[0], expectedNonoverlap[0]) {
		t.Errorf(
			"Expected first nonoverlap range %s, got %s\n",
			expectedNonoverlap[0],
			nonoverlap[0],
		)
	}
}

//		                   (this)
//		                 11  12  13
//			             |---|---|
//	  7   8   9   10  11
//		 |---|---|---|---|
//		     (other)
func TestGetOverlaps9(t *testing.T) {
	//        1   2   3
	// r1 = [11, 12, 13]
	r1 := &Range{11, 3}

	//       1  2  3   4   5
	// r2 = [7, 8, 9, 10, 11]
	r2 := &Range{7, 5}

	expectedOverlap := &Range{11, 1}

	expectedNonoverlap := []*Range{{12, 2}}

	overlap, nonoverlap, err := r1.GetOverlaps(r2)
	if err != nil {
		t.Errorf("Error while processing overlaps: %s\n", err)
	}

	if overlap != nil && !cmp.Equal(overlap, expectedOverlap) {
		t.Errorf("r1.GetOverlaps(%s): got %s, expected %s\n", r2, overlap, expectedOverlap)
	}

	if len(nonoverlap) != 1 {
		t.Errorf("Expected 1 nonoverlap ranges, got %d\n", len(nonoverlap))
	}

	if !cmp.Equal(nonoverlap[0], expectedNonoverlap[0]) {
		t.Errorf(
			"Expected first nonoverlap range %s, got %s\n",
			expectedNonoverlap[0],
			nonoverlap[0],
		)
	}
}

func TestGetOverlaps10(t *testing.T) {
	r1 := &Range{33, 7}

	r2 := &Range{19, 3}

	expectedNonoverlap := []*Range{{33, 7}}

	overlap, nonoverlap, err := r1.GetOverlaps(r2)

	if err != nil {
		t.Errorf("Error while processing overlaps: %s\n", err)
	}

	if overlap != nil {
		t.Errorf("r1.GetOverlaps(%s): got %s, expected %v\n", r2, overlap, nil)
	}

	if len(nonoverlap) != 1 {
		t.Errorf("Expected 1 nonoverlap ranges, got %d\n", len(nonoverlap))
	}

	if !cmp.Equal(nonoverlap[0], expectedNonoverlap[0]) {
		t.Errorf(
			"Expected first nonoverlap range %s, got %s\n",
			expectedNonoverlap[0],
			nonoverlap[0],
		)
	}
}

func TestGetOverlaps11(t *testing.T) {
	r1 := &Range{10, 3}
	r2 := &Range{10, 7}

	expectedOverlap := &Range{10, 3}

	overlap, nonoverlap, err := r1.GetOverlaps(r2)
	if err != nil {
		t.Errorf("Error while processing overlaps: %s\n", err)
	}

	if overlap != nil && !cmp.Equal(overlap, expectedOverlap) {
		t.Errorf("r1.GetOverlaps(%s): got %s, expected %s\n", r2, overlap, expectedOverlap)
	}

	if len(nonoverlap) != 0 {
		t.Errorf("Expected 0 nonoverlap ranges, got %d\n", len(nonoverlap))
	}
}

func TestGetOverlaps12(t *testing.T) {
	r1 := &Range{10, 7}
	r2 := &Range{10, 3}

	expectedOverlap := &Range{10, 3}

	expectedNonoverlap := []*Range{{13, 4}}

	overlap, nonoverlap, err := r1.GetOverlaps(r2)
	if err != nil {
		t.Errorf("Error while processing overlaps: %s\n", err)
	}

	if overlap != nil && !cmp.Equal(overlap, expectedOverlap) {
		t.Errorf("r1.GetOverlaps(%s): got %s, expected %s\n", r2, overlap, expectedOverlap)
	}

	if len(nonoverlap) != 1 {
		t.Errorf("Expected 1 nonoverlap ranges, got %d\n", len(nonoverlap))
	}

	if !cmp.Equal(nonoverlap[0], expectedNonoverlap[0]) {
		t.Errorf(
			"Expected first nonoverlap range %s, got %s\n",
			expectedNonoverlap[0],
			nonoverlap[0],
		)
	}
}

func TestEdgeCase1(t *testing.T) {
	r1 := &Range{Start: 33679712, Length: 18514855}
	r2 := &Range{Start: 1227445490, Length: 16145987}

	expectedNonoverlap := []*Range{{Start: 33679712, Length: 18514855}}

	overlap, nonoverlap, err := r1.GetOverlaps(r2)
	if err != nil {
		t.Errorf("Error while processing overlaps: %s\n", err)
	}

	if overlap != nil {
		t.Errorf("r1.GetOverlaps(%s): got %s, expected %v\n", r2, overlap, nil)
	}

	if len(nonoverlap) != 1 {
		t.Errorf("Expected 1 nonoverlap ranges, got %d\n", len(nonoverlap))
	}

	if !cmp.Equal(nonoverlap[0], expectedNonoverlap[0]) {
		t.Errorf(
			"Expected first nonoverlap range %s, got %s\n",
			expectedNonoverlap[0],
			nonoverlap[0],
		)
	}

}

func TestGetOverlaps13(t *testing.T) {
	r1 := &Range{47, 17}
	r2 := &Range{47, 17}

	expectedOverlap := &Range{47, 17}

	overlap, nonoverlap, err := r1.GetOverlaps(r2)
	if err != nil {
		t.Errorf("Error while processing overlaps: %s\n", err)
	}

	if overlap != nil && !cmp.Equal(overlap, expectedOverlap) {
		t.Errorf("r1.GetOverlaps(%s): got %s, expected %s\n", r2, overlap, expectedOverlap)
	}

	if len(nonoverlap) != 0 {
		t.Errorf("Expected 0 nonoverlap ranges, got %d\n", len(nonoverlap))
	}
}
