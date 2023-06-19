exe = findMacs
all : $(exe)

$(exe): $(exe).org
	bash ../scripts/org2nw $(exe).org | notangle -R$(exe).sh > $(exe)
	chmod +x $(exe)
test: $(exe) bedtools.sh
	bash test.sh
bedtools.sh: $(exe).org
	bash ../scripts/org2nw $(exe).org | notangle -Rbedtools.sh > bedtools.sh
clean:
	rm -f $(exe) bedtools.sh
