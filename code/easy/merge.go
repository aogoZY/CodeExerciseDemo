package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < m; {
		for j := 0; j < n; {
			if nums1[i]<=nums2[j]{
				i++
			}else{
				nums1[i],nums1[i+1]=nums1[i+1], nums1[i]
				nums1[i]=nums2[j]
				j++
			}
		}
	}
}

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}

