use std::fs::read_to_string;


fn main() {
    
    let result: Vec<String> = read_lines("data.txt");
    
    let mut calories: Vec<i32> = calculate_calories(&result);
    
    let max = find_max(&calories);
    let top3 = sum_top3_max(&mut calories);
    
    println!("max = {max}");
    println!("top 3 max = {top3}");

}

fn read_lines(filename: &str) -> Vec<String> {
    let mut result = Vec::new();

    for line in read_to_string(filename).unwrap().lines() {
        result.push(line.to_string());
    }
    
    return result;
}

fn calculate_calories(vec: &Vec<String>) -> Vec<i32> {
    
    let mut result = Vec::new();
    let mut sum: i32 = 0;
    for row in vec {
        if !row.is_empty() {
            let value: i32 = row.trim().parse().expect("could not read string");
            sum = sum + value; 
        }
        if row.is_empty() {
            result.push(sum);
            sum = 0;
        }
    }
    return result;
}

fn find_max(vec: &Vec<i32>) -> i32 {
    let mut max: i32 = 0;
    for row in vec {
        if max < *row {
            max = *row
        }
    }
    return max;
}

fn sum_top3_max(vec: &mut Vec<i32>) -> i32 {
    
    vec.sort_by(|a,b| b.cmp(a)); // uses quicksort so O(n * log n); 

    let result: i32 = vec[0] + vec[1] + vec[2];

    return result;
}
    
