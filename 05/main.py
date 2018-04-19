import pickle

d = pickle.load(open("./banner.p", "rb"))

for l in d:
	for x in l:
		print(x[0] * x[1], end='')
	print()

