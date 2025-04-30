How to run Unit Tests
```Bash
cmake -S . -B build
cmake --build build
./build/runTests
```
1. `cmake -S . -B build`
This command configures your project using CMake.

`-S .`: Specifies the source directory, `.` refers to the current directory where your `CMakeLists.txt` is located.

`-B build`: Specifies the build directory where CMake will generate the necessary build files (like Makefiles or Xcode project files) inside the build folder.

Effectively, this command sets up the environment so that you can later build the project.

2. `cmake --build build`
This command actually builds the project.

`--build build`: Tells CMake to build the project using the generated build files in the build directory. CMake will invoke the default build system (e.g., make or ninja, depending on what’s available on your system).

This step compiles the source files, links them, and generates your executable (e.g., `runTests`).

3. `./build/runTests`
This runs the compiled executable.

`./build/runTests`: Executes the program runTests that was created in the build directory. This is where your test cases are executed, and you see the output (pass/fail results).

---

Sed Command Explained
Used for swapping out the target CPP file used for testing.

The basic structure
sed (stream editor) is a command-line tool that processes text, allowing you to perform operations like search and replace.
```Bash
sed -i '' "s/add_executable(runTests [^)]*)/add_executable(runTests $REPLACEMENT)/" CMakeLists.txt
```

Breaking down each part:

1. `sed` - The command itself.

2. `-i ''` - The `-i` flag means "in-place editing" (modify the file directly rather than outputting to standard output).
  - The empty quotes `''` are specific to macOS. On Linux, you would just use `-i` alone.
  - This tells sed to edit the file in place without creating a backup file.

3. `"s/add_executable(runTests [^)]*)/add_executable(runTests $REPLACEMENT)/"` - This is the substitution command:
  - `s/` - Starts the substitution command.
  - `add_executable(runTests [^)]*)` - The pattern to search for:
    - `add_executable(runTests`  - Matches this exact text.
    - `[^)]*` - A character class that matches any character EXCEPT `)`, and `*` means "zero or more of these characters". This effectively matches everything between runTests and the closing parenthesis.
    - `)` - Matches the closing parenthesis.
  - `/` - Separator between what to find and what to replace it with.
  - `add_executable(runTests $REPLACEMENT)` - The replacement text:
    - We keep the beginning part identical.
    - `$REPLACEMENT` - The bash variable containing your command-line argument.
    - `)` - The closing parenthesis.
  - `/` - Marks the end of the substitution command.

4. `CMakeLists.txt` - The input file to process.

---

Sample Code
```cpp
#include <gtest/gtest.h>

TEST(SampleTest, BasicAssertion) {
    EXPECT_EQ(1 + 1, 2);
}

int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
```
