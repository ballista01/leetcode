PASSED=$(ls *.go)
if [$PASSED -eq ""];
then 
	PASSED=$(ls *.java)
fi
cp $PASSED ../passed/$PASSED
rm -v $PASSED
git stage ..
