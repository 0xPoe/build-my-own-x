pub fn is_palindrome(s: String) -> bool {
    let chars: Vec<char> = s.to_lowercase().chars().collect();
    let (mut i, mut j) = (0, s.len() - 1);
    while i < j {
        if !chars[i].is_ascii_alphanumeric() {
            i += 1;
            continue;
        }
        if !chars[j].is_ascii_alphanumeric() {
            j -= 1;
            continue;
        }
        if chars[i] != chars[j] {
            return false;
        }

        i += 1;
        j -= 1;
    }
    true
}

#[test]
fn test_is_palindrome() {
    assert!(is_palindrome("A man, a plan, a canal: Panama".to_string()));
    assert!(!is_palindrome("0P".to_string()));
    assert!(!is_palindrome("00PA".to_string()));
}
