use std::{cell::RefCell, collections::VecDeque, rc::Rc};

use crate::utils::TreeNode;

pub fn level_order(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<Vec<i32>> {
    if root.is_none() {
        return vec![];
    }

    let mut result = vec![];
    let mut queue = VecDeque::new();
    queue.push_back(root);
    while !queue.is_empty() {
        let q_len = queue.len();
        let mut level = vec![];
        for _ in 0..q_len {
            let node = queue.pop_front().unwrap();
            if let Some(node) = node {
                level.push(node.borrow().val);
                queue.push_back(node.borrow().left.clone());
                queue.push_back(node.borrow().right.clone());
            }
        }
        if !level.is_empty() {
            result.push(level);
        }
    }

    result
}

#[test]
fn test_level_order() {
    use crate::tree;

    assert_eq!(
        vec![vec![3], vec![9, 20], vec![15, 7]],
        level_order(tree!(3, 9, 20, null, null, 15, 7))
    )
}
