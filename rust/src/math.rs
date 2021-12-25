pub fn add(x: i32, y: i32) -> i32 {
    x + y
}

#[cfg(test)]
mod tests {
    use crate::math::add;
    #[cfg(test)]
    extern crate speculate;

    #[cfg(test)]
    use speculate::speculate;

    #[test]
    fn should_add_numbers() {
        assert_eq!(10, add(1, 4));
    }

    speculate! {
        describe "math" {
            before {
                let i = 0;
                println!("{}", i);
            }
            it "should add 1 + 2" {
                assert_eq!(add(1, 2), 4);
            }
        }
    }

    use rstest::rstest;

    #[rstest]
    #[case(0, 0)]
    #[case(1, 1)]
    fn param_test(#[case] a: i32, #[case] b: i32) {
        assert_eq!(1, add(a, b))
    }
}