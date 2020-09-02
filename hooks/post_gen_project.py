from __future__ import print_function
import os
from subprocess import Popen

# Get the root project directory
PROJECT_DIRECTORY = os.path.realpath(os.path.curdir)


def remove_file(filename):
    """
    generic remove file from project dir
    """
    fullpath = os.path.join(PROJECT_DIRECTORY, filename)
    if os.path.exists(fullpath):
        os.remove(fullpath)


def remove_echo_files():
    """
    Removes files needed for echo API Rest
    """
    os.remove(os.path.join(PROJECT_DIRECTORY, "pkg/users/interfaces/handlers/echo.go"))
    os.remove(
        os.path.join(PROJECT_DIRECTORY, "pkg/users/interfaces/handlers/echo_test.go")
    )


def remove_inmemory_files():
    """
    Removes files needed for inMemory storage
    """
    os.rmdir(
        os.path.join(PROJECT_DIRECTORY, "pkg/users/infrastructure/adapters/inmemory")
    )


def init_git():
    """
    Initialises git on the new project folder
    """
    GIT_COMMANDS = [
        ["git", "init"],
        ["git", "add", "."],
        ["git", "commit", "-a", "-m", "Initial Commit."],
    ]

    for command in GIT_COMMANDS:
        git = Popen(command, cwd=PROJECT_DIRECTORY)
        git.wait()


# 1. Remove echo API rest
if "{{ cookiecutter.use_echo_api }}".lower() != "y":
    remove_echo_files()

# 2. Remove inMemory storage
if "{{ cookiecutter.use_inmemory_storage }}".lower() != "y":
    remove_inmemory_files()

# 3. Initialize Git
if "{{ cookiecutter.use_git }}".lower() == "y":
    init_git()
else:
    remove_file(".gitignore")
