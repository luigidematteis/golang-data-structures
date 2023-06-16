package main

import "fmt"
import "learning/arrays"

func main() {

	var current string = "reverseString3"

	if current == "part1" {
		fmt.Println("Hello, World!")

		// check if the words are palindromes
		letMeKnowIfPalindrome("Hello")
		letMeKnowIfPalindrome("racecar")
		letMeKnowIfPalindrome("abcdedcba")

		// check if in the given arr there is a pair of value that is the sum of the target
		arrays.TwoSum([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 112)

		arrays.FindSumOfTwoNumbersEqualToTargetGivenAnUnsortedArray([]int{3, 2, 4}, 6)

		// given two sorted integer arrays, return an array that combines both of them and is also sorted
		arrays.CombineArrays([]int{1, 4, 7, 20, 32}, []int{3, 5, 6})

		// given two strings s and t, return true if s is a subsequence of t, or false otherwise
		var isSubsequence = arrays.IsSubsequence("ace", "abcde")
		fmt.Println(isSubsequence)

		// reverse strings
		arrays.ReverseString([]byte("hello"))
		arrays.ReverseString2([]byte("Hannah"))

		//square of a sorted array
		arrays.SortedSquares([]int{-4, -1, 0, 3, 10})
		arrays.SortedSquares([]int{-1})
		arrays.SortedSquares([]int{-5, -3, -2, -1})
		arrays.SortedSquares([]int{-7, -3, 2, 3, 11, 11, 3})
		arrays.SortedSquares([]int{-4, -1, 0, 3, 10, 0, 0, 1032, 3, 3, 0, 3})
		arrays.SortedSquares([]int{
			-997, -997, -990, -987, -982, -982, -980, -980, -979, -978, -974, -973, -972, -972, -966, -965, -964,
			-963, -962, -959, -957, -957, -956, -955, -952, -952, -950, -948, -947, -944, -944, -943, -940, -932,
			-926, -925, -924, -919, -918, -918, -917, -917, -916, -915, -913, -913, -911, -909, -906, -905, -904,
			-904, -901, -899, -898, -898, -896, -892, -887, -886, -882, -875, -875, -874, -871, -871, -868, -864,
			-864, -852, -850, -850, -845, -844, -839, -834, -829, -828, -819, -819, -818, -818, -817, -817, -808,
			-806, -806, -806, -801, -797, -797, -794, -794, -794, -793, -793, -791, -790, -786, -785, -784, -784,
			-781, -775, -773, -771, -769, -763, -762, -757, -757, -749, -746, -741, -740, -739, -739, -739, -738,
			-734, -730, -729, -727, -727, -726, -726, -725, -721, -721, -719, -718, -718, -714, -710, -702, -702,
			-701, -701, -701, -700, -699, -698, -694, -693, -690, -690, -689, -689, -685, -685, -682, -682, -679,
			-679, -679, -674, -673, -672, -670, -669, -667, -667, -666, -663, -656, -655, -654, -653, -652, -651,
			-648, -642, -641, -638, -634, -632, -631, -630, -629, -629, -625, -624, -623, -614, -614, -611, -610,
			-606, -603, -603, -598, -597, -595, -594, -592, -590, -586, -586, -586, -585, -582, -579, -578, -577,
			-576, -575, -575, -574, -570, -570, -569, -568, -566, -562, -561, -557, -552, -551, -547, -546, -545,
			-543, -542, -540, -537, -536, -534, -529, -527, -526, -526, -522, -519, -518, -518, -515, -515, -515,
			-512, -511, -508, -507, -505, -504, -500, -498, -498, -497, -495, -493, -489, -488, -487, -482, -482,
			-480, -479, -477, -476, -476, -472, -472, -472, -471, -470, -470, -464, -464, -463, -463, -461, -460,
			-458, -447, -446, -443, -443, -442, -442, -441, -434, -431, -431, -430, -427, -427, -426, -424, -421,
			-420, -417, -416, -416, -416, -415, -413, -411, -411, -411, -411, -411, -409, -402, -399, -399, -397,
			-395, -394, -393, -389, -387, -384, -380, -380, -378, -377, -376, -373, -373, -365, -365, -363, -358,
			-351, -351, -345, -345, -342, -340, -338, -338, -336, -336, -335, -334, -331, -331, -331, -324, -323,
			-322, -320, -319, -317, -317, -315, -314, -312, -310, -310, -306, -305, -303, -296, -294, -292, -292,
			-291, -284, -283, -283, -282, -279, -276, -274, -271, -270, -267, -265, -258, -258, -258, -253, -251,
			-250, -246, -243, -242, -239, -238, -237, -236, -236, -235, -235, -230, -227, -226, -220, -220, -219,
			-219, -217, -213, -212, -212, -208, -206, -202, -201, -201, -200, -198, -197, -197, -195, -194, -192,
			-192, -192, -190, -186, -185, -180, -177, -176, -173, -172, -168, -163, -163, -162, -159, -156, -156,
			-154, -152, -151, -150, -147, -146, -140, -138, -136, -136, -134, -133, -133, -131, -130, -130, -127,
			-122, -118, -116, -116, -115, -113, -111, -111, -110, -107, -106, -105, -102, -100, -100, -100, -96,
			-94, -94, -90, -89, -84, -84, -82, -82, -81, -81, -78, -76, -75, -73, -70, -67, -61, -49, -48, -47, -47,
			-47, -42, -39, -37, -36, -34, -32, -28, -27, -24, -23, -20, -18, -13, -7, -7, -6, -5, -3, -2, -1, 0, 8, 11,
			12, 13, 14, 15, 16, 18, 19, 22, 27, 27, 27, 29, 31, 33, 34, 39, 40, 44, 45, 46, 46, 47, 48, 49, 49, 50, 52, 57,
			64, 66, 67, 69, 70, 71, 73, 73, 78, 80, 81, 81, 81, 91, 92, 93, 96, 97, 99, 104, 105, 106, 108, 108, 108, 109,
			112, 115, 115, 118, 118, 119, 120, 123, 124, 128, 131, 132, 133, 134, 135, 135, 136, 136, 142, 144, 145,
			145, 147, 150, 152, 153, 155, 159, 162, 166, 166, 171, 173, 174, 177, 178, 179, 180, 184, 185, 186, 187,
			189, 190, 192, 193, 193, 196, 201, 202, 202, 206, 207, 208, 210, 211, 213, 217, 218, 221, 223, 226, 227,
			230, 233, 233, 238, 241, 241, 246, 246, 246, 247, 247, 250, 251, 252, 253, 254, 255, 255, 258, 258, 261,
			261, 262, 266, 267, 272, 274, 275, 276, 278, 280, 285, 287, 287, 288, 299, 304, 304, 307, 308, 310, 313,
			315, 317, 318, 321, 323, 323, 325, 325, 328, 332, 332, 333, 335, 336, 336, 336, 338, 339, 340, 341, 343,
			346, 349, 351, 353, 356, 357, 357, 359, 367, 370, 374, 376, 380, 380, 380, 383, 386, 389, 389, 390, 391,
			392, 393, 394, 395, 396, 401, 407, 415, 418, 422, 422, 424, 426, 427, 429, 431, 433, 434, 436, 439, 443,
			446, 447, 453, 458, 462, 467, 470, 470, 472, 473, 475, 476, 476, 478, 479, 480, 483, 484, 484, 486, 487,
			489, 490, 492, 497, 499, 502, 503, 507, 509, 510, 510, 513, 514, 515, 516, 517, 520, 528, 529, 531, 532,
			533, 536, 539, 540, 542, 549, 552, 558, 559, 563, 569, 569, 571, 574, 575, 576, 578, 578, 579, 582, 585,
			588, 588, 589, 589, 591, 594, 595, 595, 596, 600, 601, 602, 603, 606, 607, 609, 610, 613, 615, 616, 616,
			616, 617, 617, 618, 619, 621, 623, 625, 625, 628, 630, 636, 639, 643, 643, 647, 647, 649, 654, 656, 659,
			661, 661, 662, 662, 665, 667, 667, 669, 670, 671, 671, 672, 683, 683, 683, 684, 685, 687, 687, 694, 697,
			698, 699, 704, 710, 712, 715, 715, 715, 717, 718, 719, 719, 721, 723, 726, 726, 727, 727, 728, 728, 729,
			730, 730, 731, 732, 734, 734, 736, 736, 739, 739, 741, 742, 745, 745, 747, 748, 749, 749, 760, 763, 764,
			775, 776, 779, 783, 784, 786, 794, 798, 799, 801, 803, 805, 808, 810, 811, 814, 816, 819, 820, 822, 826,
			828, 838, 839, 839, 848, 849, 850, 850, 851, 851, 852, 852, 855, 865, 867, 868, 869, 870, 872, 874, 878,
			881, 884, 884, 886, 887, 887, 890, 890, 893, 893, 894, 894, 895, 895, 896, 900, 908, 911, 915, 915, 917,
			919, 923, 923, 924, 926, 926, 926, 930, 933, 937, 937, 938, 941, 941, 944, 944, 945, 948, 948, 948, 949,
			950, 952, 960, 960, 961, 968, 970, 971, 973, 973, 977, 979, 980, 983, 986, 992, 992, 992, 993, 995, 996,
			997})

		// sliding windows
		arrays.FindLengthOfSubarrayWithSumLessThanK([]int{3, 1, 2, 7, 4, 2, 1, 1, 5}, 8)
		arrays.FindNumOfSubarrayWithProductLessThanK([]int{10, 5, 2, 6}, 100)

		// maximum average subarray
		arrays.MaximumAverageSubarray([]int{5}, 1)
		arrays.MaximumAverageSubarray([]int{1, 4, 2, 3, 2}, 5)
		arrays.MaximumAverageSubarray([]int{1, 12, -5, -6, 50, 3}, 4)
		arrays.MaximumAverageSubarray([]int{-1}, 1)
		arrays.MaximumAverageSubarray([]int{4, 0, 4, 3, 3}, 5)
		arrays.MaximumAverageSubarray([]int{8860, -853, 6534, 4477, -4589, 8646, -6155, -5577, -1656, -5779, -2619, -8604, -1358, -8009,
			4983, 7063, 3104, -1560, 4080, 2763, 5616, -2375, 2848, 1394, -7173, -5225, -8244, -809, 8025, -4072, -4391, -9579, 1407, 6700,
			2421, -6685, 5481, -1732, -8892, -6645, 3077, 3287, -4149, 8701, -4393, -9070, -1777, 2237, -3253, -506, -4931, -7366, -8132,
			5406, -6300, -275, -1908, 67, 3569, 1433, -7262, -437, 8303, 4498, -379, 3054, -6285, 4203, 6908, 4433, 3077, 2288, 9733, -8067,
			3007, 9725, 9669, 1362, -2561, -4225, 5442, -9006, -429, 160, -9234, -4444, 3586, -5711, -9506, -79, -4418, -4348, -5891}, 93)

		arrays.MaxAverageSubarray_2([]int{5}, 1)
		arrays.MaxAverageSubarray_2([]int{1, 4, 2, 3, 2}, 5)
		arrays.MaxAverageSubarray_2([]int{1, 12, -5, -6, 50, 3}, 4)
		arrays.MaxAverageSubarray_2([]int{-1}, 1)
		arrays.MaxAverageSubarray_2([]int{4, 0, 4, 3, 3}, 5)
		arrays.MaxAverageSubarray_2([]int{8860, -853, 6534, 4477, -4589, 8646, -6155, -5577, -1656, -5779, -2619, -8604, -1358, -8009,
			4983, 7063, 3104, -1560, 4080, 2763, 5616, -2375, 2848, 1394, -7173, -5225, -8244, -809, 8025, -4072, -4391, -9579, 1407, 6700,
			2421, -6685, 5481, -1732, -8892, -6645, 3077, 3287, -4149, 8701, -4393, -9070, -1777, 2237, -3253, -506, -4931, -7366, -8132,
			5406, -6300, -275, -1908, 67, 3569, 1433, -7262, -437, 8303, 4498, -379, 3054, -6285, 4203, 6908, 4433, 3077, 2288, 9733, -8067,
			3007, 9725, 9669, 1362, -2561, -4225, 5442, -9006, -429, 160, -9234, -4444, 3586, -5711, -9506, -79, -4418, -4348, -5891}, 93)
		arrays.MaxAverageSubarray_2([]int{6, 8, 6, 8, 0, 4, 1, 2, 10, 10, 9, 9}, 2)
	}

	if current == "part2" {
		// maximum of consecutive ones in an array representing a binary number
		arrays.MaxConsecutiveOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2)

		// answer to queries that are true if the sum of the subarray from x to y is less than limit
		arrays.AnswerQueries([]int{1, 6, 3, 2, 7, 2}, [][]int{{0, 3}, {2, 5}, {2, 4}}, 13)

		// ways to split an array
		arrays.WaysToSplitArray([]int{10, 4, -8, 7})
		arrays.WaysToSplitArray2([]int{10, 4, -8, 7})

		// running sum
		arrays.RunningSum([]int{1, 2, 3, 4})
	}

	if current == "prefix_sum" {
		// minimum value to get positive step by step sum
		arrays.MinStartValue([]int{-3, 2, -3, 4, 2})
		arrays.MinStartValue([]int{1, -2, -3})
		arrays.MinStartValue([]int{-5, -2, 4, 4, -2})
		arrays.MinStartValue([]int{2, 3, 5, -5, -1})
		arrays.MinStartValue([]int{1, 2})
		arrays.MinStartValue([]int{4, 2, -1})
	}

	if current == "kradius" {
		// k-radius subarray average
		arrays.GetAverage([]int{7, 4, 3, 9, 1, 8, 5, 2, 6}, 3)
	}

	if current == "reverseString3" {
		//arrays.ReverseWords("Let's take LeetCode contest")
		arrays.ReverseWords2("Let's take LeetCode contest")
	}

}

func letMeKnowIfPalindrome(word string) {
	if arrays.TwoPointers(word) {
		fmt.Println("The word", word, "is a palindrome")
	} else {
		fmt.Println("The word", word, "is not a palindrome")
	}
}
