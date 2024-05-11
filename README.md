### README.md

#### Script Manager

#### Description
This Go program, `Script Viewer`, lists files within a specified directory along with the last line of each file as it's description.

#### Usage
1. **Build from Source**:
   - **Install Go**: Ensure you have Go installed on your system. If not, you can download and install it from the [official Go website](https://golang.org/dl/).
   - **Download the Source Code**:
     - Clone this repository to your local machine:
       ```shell
       git clone https://github.com/Dekrazi/script_viewer.git
       ```
   - **Configure Scripts Folder**:
     - Open the `main.go` file in your preferred text editor.
     - Locate the `scriptsFolder` variable at the top of the file.
     - Change the value of `scriptsFolder` to the directory path where your scripts are stored.
   - **Build the Program**:
     - Navigate to the directory where you cloned the repository.
     - Build the program using the following command:
       ```shell
       go build -o script_viewer
       ```
2. **Use Provided Binary**:
   - If you prefer not to build from source, you can use the provided binary.
     - Download the binary from [this page](https://github.com/Dekrazi/script_viewer/blob/main/script_viewer).
     - Ensure the binary is executable:
       ```shell
       chmod +x script_viewer
       ```

3. **Run the Program**:
   - Execute the compiled binary with the following command:
     ```shell
     ./script_viewer
     ```
   - This will list the files within the specified directory along with their last lines, using the last line of each file as its description.

#### Example
Suppose you have script files in `/home/user/scripts/` and you want to see the list of them. You would set `scriptsFolder` in `main.go` to `/home/user/scripts/`, then build and run the program as described above.

Write short description at the last line of each script file to get brief info about it in Script Viewer.

#### Note
- Ensure that the directory specified in `scriptsFolder` is accessible and contains the files you want to list.
- This program assumes that the last line of each file contains meaningful information. If this is not the case for your files, you may need to modify the `readLastLineAsDesc` function to suit your needs.

#### License
This project is licensed under the [MIT License](LICENSE). Feel free to use and modify it according to your requirements.

For any issues or suggestions, please open an [issue](https://github.com/yourusername/yourrepository/issues) on GitHub. Contributions are welcome!
