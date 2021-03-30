use clap::Clap;
use std::fs::File;
use std::io::BufRead;
use std::io::BufReader;
use std::io::stdin;

#[derive(Clap, Debug)]
#[clap(
    name = "My RPN Program",
    version = "1.0.0",
    author = "hsmtkk",
    about = "Supre awesome sample RPN calculator"
)]

struct Opts{
    #[clap(short, long)]
    verbose: bool,
    #[clap(name = "FILE")]
    formula_file: Option<String>,
}

fn main(){
    let opts = Opts::parse();

    if let Some(path) = opts.formula_file {
        let f = File::open(path).unwrap();
        let reader = BufReader::new(f);
        read_file(reader, opts.verbose);
    }else{
        let stdin = stdin();
        let reader = stdin.lock();
        read_file(reader, opts.verbose);
    }
}

fn read_file<R:BufRead>(reader:R, verbose:bool){
    for line in reader.lines(){
        let line = line.unwrap();
        println!("{}", line);
    }
}