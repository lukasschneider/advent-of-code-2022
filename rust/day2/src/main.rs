use std::fs;

fn main() {
    

    let mut final_score: i32 = 0;
    let mut final_score_predict: i32 = 0;

    for line in fs::read_to_string("data.txt").unwrap().lines() {
        let result = line.to_string();
        let mut value = result.chars();
        let pair: (char, char) = (value.next().unwrap(), value.last().unwrap());
        final_score += calc_score(pair);
        final_score_predict += calc_score_on_predict(pair.clone());
    }

    println!("final_score: {}", final_score);
    println!("final_score_predict: {}", final_score_predict);

}

fn calc_score(pair: (char, char)) -> i32 {
    let mut score_per_round: i32 = 0;

    /*
     * Rock X: 1
     * Paper Y: 2
     * Scissor Z: 3
     */

    match pair.0 {
        'A' => match pair.1 { // Rock
            'X' => score_per_round = 3 + 1, // draw
            'Y' => score_per_round = 6 + 2, // win
            'Z' => score_per_round = 0 + 3, // lose
            _ => eprintln!("error"),
        },
        'B' => match pair.1 { // Paper
            'X' => score_per_round = 0 + 1, // lose 
            'Y' => score_per_round = 3 + 2, // draw
            'Z' => score_per_round = 6 + 3, // win
            _ => eprintln!("error"),
        },
        'C' => match pair.1 { // Scissor
            'X' => score_per_round = 6 + 1, // win
            'Y' => score_per_round = 0 + 2, // lose
            'Z' => score_per_round = 3 + 3, // draw 
            _ => eprintln!("error"),
        }
        _ => eprintln!("error"),
    };

    score_per_round
}

fn calc_score_on_predict(pair: (char, char)) -> i32 {
    let mut score_per_round: i32 = 0;
    
    /*
     * Rock X: 1
     * Paper Y: 2
     * Scissor Z: 3
     */

    match pair.0 {
        'A' => match pair.1 { // Rock
            'X' => score_per_round = 0 + 3, // lose
            'Y' => score_per_round = 3 + 1, // draw
            'Z' => score_per_round = 6 + 2, // win
            _ => eprintln!("error"),
        },
        'B' => match pair.1 { // Paper
            'X' => score_per_round = 0 + 1, // lose
            'Y' => score_per_round = 3 + 2, // draw
            'Z' => score_per_round = 6 + 3, // win
            _ => eprintln!("error"),
        },
        'C' => match pair.1 { // Scissor
            'X' => score_per_round = 0 + 2, // lose 
            'Y' => score_per_round = 3 + 3, // draw
            'Z' => score_per_round = 6 + 1, // win 
            _ => eprintln!("error"),
        }
        _ => eprintln!("error"),
    };

    score_per_round
}























