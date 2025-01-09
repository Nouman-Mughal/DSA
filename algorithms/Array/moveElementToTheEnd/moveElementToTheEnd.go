package moveElementsToTheEnd

// moving zero to the end of array in the specific examaple
func moveZeroes(nums []int) []int {
    // using two pointer approach

    for left, right:= 0,0 ; right < len(nums); right++ {
        if nums[right] != 0 {
            nums[left], nums[right] = nums[right], nums[left]
            left++
        }
    } 
    return nums
}