use std::fs;

fn main() {
    let x = fs::read_to_string("input.txt").expect("could not finish");
    println!("{}", x);
}
