PASSED=$(ls *.go)
cp $PASSED ../passed/$PASSED
rm -v $PASSED
git stage ..
