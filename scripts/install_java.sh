#!/bin/bash

JDK=jdk-18_linux-x64_bin.tar.gz
JDK_PATH=jdk-18
PROFILE=$HOME/.bash_profile

echo "===> check java installed..."
java -version &>/dev/null
version=$?

if [ "$version" -ne 0 ]
then
	echo "===> not installed java"
	echo "===> install java tar: $JDK"
	wget https://download.oracle.com/java/18/latest/$JDK
	tar -xzf $JDK
	rm $JDK
	echo "export JAVA_HOME=$PWD/$JDK_PATH" >> $PROFILE
	echo "export PATH=\$PATH:$PWD/$JDK_PATH/bin" >> $PROFILE
	. $PROFILE
	echo "===> installed java"
else
	echo "===> already installed java"
fi