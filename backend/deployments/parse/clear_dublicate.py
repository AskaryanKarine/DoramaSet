import sys

file_name = sys.argv[1]

content = []
with open(file_name, "r", encoding="utf-8") as f:
    lines = f.readlines()
    for line in lines:
        content.append(line.strip().split("|"))

# print(content)
tmp_dict = {}
for title in content:
    key = title[0]
    value = title[1]
    tmp_dict[key] = value

# content = []
# for k in tmp_dict:
#     s = str(k)+ "|" + str(tmp_dict[k]) + "\n"
#     content.append(s)

with open(file_name, "w", encoding="utf-8") as f:
    for k in tmp_dict:
        s = str(tmp_dict[k]) + "\n"
        # s = str(k) + "\n"
        f.write(s)

