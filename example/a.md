this file will be our final file. everything will be included in it. so let's
include another file in the bottom of it.

and here it starts >>>>>  "[ pasta : b file ](./b.md)"

there shold be 2 lines about blueberries and canberries. we put just the file
about blueberries `b.md` it had another pasta in it. so it works recursively.

all these files were embedded because they had `[pasta: ]` prefix in their
titles. keep in mind that pasta links should be the only thing in the line.
regular links with no pasta prefix will stay as they are. example:

[b file](./b.md)

pasta pathes are relative to the executable, so cd to your final document's
folder before using pasta.
