Submission for [Bristol-Myers Squibb's Molecular Translation Competition](https://www.kaggle.com/c/bms-molecular-translation/)
============
### Language
[Golang](https://golang.org/) (Go)
### Local Development Setup (macOS)
+ Repeats [this guide](https://gocv.io/getting-started/macos/) for installing gocv, then installs the kaggle cli to fetch the competition data
+ Assumes you have a Kaggle account and are registered for [Bristol-Myers Squibb's Molecular Translation Competition](https://www.kaggle.com/c/bms-molecular-translation/) to have access to the competition data
+ After the contest ends on 2021-05-26, contrived competition data may be provided in the repository. If it's not found, create 2 directories at the top level of this repository named **train** and **test**, and within each directory, create 3 nested directories named 0, then place any line diagrams of any chemicals that you like, in .png file format. For example, your files would be named train/0/0/0/0d3aef24.png and test/0/0/0/9c8fb2ed.png . 
<pre>
# Assumes you have installed Go
$ cd bms-molecular-translation

# Install opencv (10 minutes)
$ brew install opencv

# Install pkgconfig (1 minute)
$ brew install pkgconfig

# Install the kaggle cli (5 minutes)
$ pip install kaggle

# Set up your kaggle credentials at https://github.com/Kaggle/kaggle-api#api-credentials
# Then download the competition data (8GB, 30 minutes)

$ chmod 600 ~/.kaggle/kaggle.json
$ kaggle competitions download -c bms-molecular-translation

# Build the script
$ go install translate

# Run the executable
$ bin/translate
</pre>

This change should not trigger a github action

This change, coupled with the comment in another file, should trigger
the github action. 

Subsequent changes should trigger the github action, ignoring paths-ignore
