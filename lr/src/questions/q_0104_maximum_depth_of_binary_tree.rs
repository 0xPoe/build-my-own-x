use std::{cell::RefCell, rc::Rc};

use crate::utils::TreeNode;

pub fn max_depth(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
    if root.is_none() {
        return 0;
    }
    let left = 1 + max_depth(root.as_ref().unwrap().borrow().left.clone());
    let right = 1 + max_depth(root.as_ref().unwrap().borrow().right.clone());
    if left > right {
        return left;
    }
    right
}

#[test]
fn test_max_depth() {
    use crate::tree;

    let tree = tree!(3, 9, 20, null, null, 15, 7);
    assert_eq!(3, max_depth(tree))
}
