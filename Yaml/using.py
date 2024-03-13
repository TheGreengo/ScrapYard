import yaml

file =  open('all_things.yaml', 'r')

docs = yaml.safe_load_all(file)
for i in docs:
    print(i)

file.close()