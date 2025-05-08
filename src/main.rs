use clap::{App, Arg};
use glob::glob;
use log::{error, info};
use std::fs;
use std::process::Command;
use std::{
    error::Error,
    path::{Path, PathBuf},
};

fn ensure_directory_exists(path: &Path) -> Result<Vec<PathBuf>, Box<dyn std::error::Error>> {
    let mut created_dirs = Vec::new();
    let mut current_path = PathBuf::new();

    for component in path.components() {
        current_path.push(component);
        if !current_path.exists() {
            fs::create_dir(&current_path)?;
            created_dirs.push(current_path.clone());
        }
    }

    Ok(created_dirs)
}

fn cleanup_empty_directories_if_created(
    created_dirs: &[PathBuf],
) -> Result<(), Box<dyn std::error::Error>> {
    for dir in created_dirs.iter().rev() {
        if is_directory_empty(dir)? {
            fs::remove_dir(dir)?;
            info!("Removed empty directory: {:?}", dir);
        }
    }
    Ok(())
}

fn is_directory_empty(dir: &Path) -> Result<bool, Box<dyn std::error::Error>> {
    Ok(dir.read_dir()?.next().is_none())
}

fn run_command_in_dir(
    dir: &Path,
    command_parts: &[&str],
) -> Result<(), Box<dyn std::error::Error>> {
    info!("Running command in directory: {:?}", dir);

    let dir_str = dir.to_str().unwrap_or_default();
    let status = Command::new(command_parts[0])
        .args(&command_parts[1..])
        .current_dir(dir)
        .env("PWD", dir_str)
        .env("INDIR", dir_str)
        .status()?;

    if !status.success() {
        error!("Command exited with error.");
        return Err(
            std::io::Error::new(std::io::ErrorKind::Other, "Command execution failed").into(),
        );
    }
    Ok(())
}

fn main() -> Result<(), Box<dyn Error>> {
    let matches = App::new("in")
        .version("1.7.4")
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

    let pattern = matches.value_of("DIRECTORY_OR_PATTERN").unwrap();
    let command_parts: Vec<&str> = matches.values_of("COMMAND").unwrap().collect();

    // glob mode
    if pattern.contains("**") {
        for entry in glob(pattern)? {
            match entry {
                Ok(path_buf) if path_buf.is_file() => {
                    let dir = path_buf.parent().unwrap_or_else(|| Path::new("."));
                    run_command_in_dir(dir, &command_parts)?;
                }
                Ok(_) => { /* skip non-files */ }
                Err(e) => error!("Error with glob entry: {}", e),
            }
        }
        return Ok(());
    }

    let path = PathBuf::from(pattern);
    let created_dirs = ensure_directory_exists(&path)?;
    run_command_in_dir(&path, &command_parts)?;
    cleanup_empty_directories_if_created(&created_dirs)?;

    Ok(())
}
