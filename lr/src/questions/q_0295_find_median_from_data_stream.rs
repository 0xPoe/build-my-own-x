use std::{cmp::Reverse, collections::BinaryHeap};

struct MedianFinder {
    small_heap: BinaryHeap<i32>,
    large_heap: BinaryHeap<Reverse<i32>>,
}

impl MedianFinder {
    fn new() -> Self {
        MedianFinder {
            small_heap: BinaryHeap::new(),
            large_heap: BinaryHeap::new(),
        }
    }

    fn add_num(&mut self, num: i32) {
        self.small_heap.push(num);
        if !self.small_heap.is_empty()
            && !self.large_heap.is_empty()
            && self.small_heap.peek().unwrap() > &self.large_heap.peek().unwrap().0
        {
            let num = self.small_heap.pop().unwrap();
            self.large_heap.push(Reverse(num));
        }
        if self.small_heap.len() > self.large_heap.len() + 1 {
            let num = self.small_heap.pop().unwrap();
            self.large_heap.push(Reverse(num));
        }
        if self.large_heap.len() > self.small_heap.len() + 1 {
            let num = self.large_heap.pop().unwrap().0;
            self.small_heap.push(num);
        }
    }

    fn find_median(&self) -> f64 {
        let small_len = self.small_heap.len();
        let large_len = self.large_heap.len();

        if small_len > large_len {
            *self.small_heap.peek().unwrap() as f64
        } else if small_len < large_len {
            self.large_heap.peek().unwrap().0 as f64
        } else {
            (*self.small_heap.peek().unwrap() + self.large_heap.peek().unwrap().0) as f64 / 2.0
        }
    }
}

#[test]
fn test_find_median() {
    let mut finder = MedianFinder::new();
    finder.add_num(1);
    finder.add_num(2);
    assert_eq!(1.5, finder.find_median())
}
