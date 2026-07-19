use std::cmp::{max, min};

pub fn max_area(height: Vec<i32>) -> i32 {
    if height.is_empty() {
        return 0;
    }

    let mut result = 0;
    let (mut i, mut j) = (0, height.len() - 1);
    while i < j {
        let area = (j - i) * (min(height[i], height[j]) as usize);
        result = max(result, area as i32);
        if height[i] < height[j] {
            i += 1;
        } else {
            j -= 1;
        }
    }
    result
}

#[test]
fn test_max_area() {
    assert_eq!(49, max_area(vec![1, 8, 6, 2, 5, 4, 8, 3, 7]))
}
