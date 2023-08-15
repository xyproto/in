extern crate clap;
extern crate env_logger;
extern crate log;

use clap::{App, Arg};
use glob::glob;
use log::{error, info};
use std::fs;
use std::path::{Path, PathBuf};
use std::process::Command;

fn run_command_and_cleanup(
    dir: &Path,
    command_parts: &[&str],
    created_dirs: &[PathBuf],
) -> Result<(), Box<dyn std::error::Error>> {
    run_command_in_dir(dir, command_parts)?;
    cleanup_empty_directories_if_created(dir, created_dirs)?;
    Ok(())
}

fn cleanup_empty_directories_if_created(
    dir: &Path,
    created_dirs: &[PathBuf],
) -> Result<(), Box<dyn std::error::Error>> {
    if created_dirs.contains(&dir.to_path_buf()) && is_directory_empty(dir)? {
        fs::remove_dir(dir)?;
        info!("Removed empty directory: {:?}", dir);
    }
    Ok(())
}

fn is_directory_empty(dir: &Path) -> Result<bool, Box<dyn std::error::Error>> {
    Ok(dir.read_dir()?.next().is_none())
}

fn ensure_directory_exists(path: &Path) -> Result<Vec<PathBuf>, Box<dyn std::error::Error>> {
    let mut created_dirs = Vec::new();
    let mut current_path = path.to_path_buf();

    if !current_path.exists() {
        created_dirs.push(current_path.clone());
        fs::create_dir_all(&current_path)?;
    }

    while let Some(parent) = current_path.parent() {
        if !parent.exists() {
            created_dirs.push(parent.to_path_buf());
        }
        current_path.pop();
    }

    Ok(created_dirs)
}

fn run_command_in_dir(
    dir: &Path,
    command_parts: &[&str],
) -> Result<(), Box<dyn std::error::Error>> {
    info!("Running command in directory: {:?}", dir);
    let status = Command::new(command_parts[0])
        .args(&command_parts[1..])
        .current_dir(dir)
        .status()?;

    if !status.success() {
        error!("Command exited with error.");
        return Err(Box::new(std::io::Error::new(
            std::io::ErrorKind::Other,
            "Command execution failed",
        )));
    }
    Ok(())
}

fn main() -> Result<(), Box<dyn std::error::Error>> {
    env_logger::init();

    let matches = App::new("in")
        .version("1.7.0")
        .author("Alexander F. Rødseth <xyproto@archlinux.org>")
        .about("Utility to execute commands in directories, and create directories if needed.")
        .arg(
            Arg::with_name("DIRECTORY_OR_PATTERN")
                .help("Target directory or pattern")
                .required(true),
        )
        .arg(
            Arg::with_name("COMMAND")
                .help("Command to run")
                .required(true)
                .multiple(true),
        )
        .get_matches();

    let path = matches.value_of("DIRECTORY_OR_PATTERN").unwrap();
    let command_parts: Vec<&str> = matches.values_of("COMMAND").unwrap().collect();

    if path.contains("**") {
        // Glob/wildcard mode
        for entry in glob(path)? {
            match entry {
                Ok(path_buf) => {
                    if path_buf.is_file() {
                        let dir = path_buf.parent().unwrap_or(Path::new("."));
                        run_command_in_dir(dir, &command_parts)?;
                    }
                }
                Err(e) => error!("Error with glob entry: {}", e),
            }
        }
    } else {
        // Standard mode
        let path = PathBuf::from(path);
        let created_dirs = ensure_directory_exists(&path)?;
        run_command_and_cleanup(&path, &command_parts, &created_dirs)?;
    }
    Ok(())
}
