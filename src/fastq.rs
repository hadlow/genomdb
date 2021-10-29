use std::fmt;
use std::fmt::Debug;
use std::fs::File;
use std::io::{self, prelude::*, BufReader};
use std::path::PathBuf;

pub struct Fastq
{
    path: PathBuf,
}

impl Fastq
{
    pub fn new(path: &PathBuf) -> Self
    {
        Self
        {
            path: PathBuf::from(path),
        }
    }

    pub fn tokenize(&self) -> io::Result<()>
    {
        let file = File::open(self.path.to_str().unwrap())?;
        let reader = BufReader::new(file);

        for line in reader.lines()
        {
            println!("{}", line?);
        }

        Ok(())
    }
}

impl Debug for Fastq
{
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result
    {
        write!(f, "{}", self.path.to_str().unwrap())
    }
}
