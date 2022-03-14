# Generates the README file based on the comments in my
# go files. It generates and populates a table with the 
# following headings:

# | # | NAME | CONCEPTS | STATS |

# To tag your files use the following comments anywhere in
# your file:

# // Name: asd
# // Num: asdf
# // Tags: asdf
# // Stats: asdf

import os

readme = open("README.md", "w")
readme.write("These are a collection of some of my solutions to LeetCode problems in `golang`. The file `tagger.py` automatically parses the details in the comments of each .go file to render the table below. \n\n")

readme.write("| # | Name | Concepts | Stats | \n")
readme.write("|:--- | :--- | :--- | :--- | \n")

currDir = os.getcwd()
goFiles = []
for filename in os.listdir(currDir):
    if filename.endswith(".go"):
        goFiles.append(filename)
        f = open(filename, "r")
        lines = f.readlines()

        stats = ""
        tags = ""
        num =  filename.replace(".go", "")
        if num == "template":
            continue
        name = ""
        for line in lines:
            if line.startswith("//"):
                words = line.split()[1:]
                # print(words)
                word = words[0]

                if word == "Num:" or word == "Number:":
                    num = words[-1]
                    
                if word == "Name:":
                    name = line[2:].replace("Name:", "").replace("\n", "")

                if word == "Tags:":
                    tags = line[2:].replace("Tags:", "").replace("\n", "")

                if word == "Stats:":
                    stats = line[2:].replace("Stats:", "").replace("\n", "")

        filestats = "| " + num + " | " + name + " | " + tags + " | " + stats + " | \n"
        readme.write(filestats)
        f.close()

readme.close()
readme = open("README.md", "r")
filesInReadme = set()
for line in readme.readlines():
    if line[0] == "|" and (line[2] != "-" or line[2] != "#") :
        filesInReadme.add(line.split("|")[1].split(" ")[1]+".go")

if set(goFiles).remove("template.go") is filesInReadme:
    print("Uh oh! These files don't appear in the README:\n")
    print(set(goFiles).difference(filesInReadme))
else:
    print("all good!")
