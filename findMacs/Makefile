exe = findMacs
all : $(exe)

$(exe): $(exe).org
	bash ../scripts/org2nw $(exe).org | notangle -R$(exe).sh > $(exe)
	chmod +x $(exe)
clean:
	rm -f $(exe)
test:
	bash test.sh
