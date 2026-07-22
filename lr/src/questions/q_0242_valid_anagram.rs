use std::collections::HashMap;

pub fn is_anagram(s: String, t: String) -> bool {
    let mut s_map: HashMap<char, i32> = HashMap::new();
    for c in s.chars() {
        *s_map.entry(c).or_insert(0) += 1;
    }
    for c in t.chars() {
        *s_map.entry(c).or_insert(0) -= 1;
    }

    s_map.values().all(|x| *x == 0)
}

#[test]
fn test_is_anagram() {
    assert!(is_anagram("anagram".to_string(), "nagaram".to_string()));
    assert!(!is_anagram("car".to_string(), "arr".to_string()))
}
